<template>
    <div class="chat-input">
        <el-input type="textarea" placeholder="输入消息..." v-model="newMessage" class="message-input" :rows="4" />

        <div class="send-buttons">
            <!-- 发送文件 -->
            <div class="Filecontainer">
                <!-- 文件上传按钮 -->
                <el-button type="primary" class="send-button" @click="openFileChooser" style="margin-bottom: 10px">发送文件</el-button>
                <!-- 隐藏的文件输入 -->
                <input type="file" ref="fileInput" style="display: none;" @change="handleFileUpload">
            </div>
            <!-- <el-button type="primary" class="send-button" @click="changeVisible" style="margin-bottom: 10px">添加附件</el-button> -->

            <el-button type="success" class="send-button" @click="sendMessage('text')">发送消息</el-button>
            <!-- 开启视频 -->
            <el-button type="success" class="send-button" @click="startVideo">开启视频</el-button>
        </div>

        <!-- 对话框 -->
        <el-dialog title="上传文件" v-model="dialogVisible">
            <!-- 上传组件 -->

        </el-dialog>
    </div>
</template>

<script>
import axios from "@/axios-config";
import {ElMessage} from "element-plus";
export default {
    props: {
        roomID: {
            type: Number,
            required: true
        }
    },
    data() {
        return {
            newMessage: '', // 绑定输入框的内容
            dialogVisible: false,
            roomId: this.roomID
        };
    },
    methods: {
        changeVisible() {
            this.dialogVisible = !this.dialogVisible;
            console.log("changeVisible", this.dialogVisible);
        },
        sendMessage(type) {
            if (!this.newMessage) {
                ElMessage.warning('消息不能为空');
                return;
            }
            // 向父组件发送输入的消息
            this.$emit('send', type, this.newMessage);
            this.newMessage = ''; // 发送后清空输入框
        },
        // 打开文件选择器
        openFileChooser() {
            this.$refs.fileInput.click();
        },

        // 处理文件上传
        handleFileUpload(event) {
            const file = event.target.files[0]; // 获取文件
            if (!file) {
                return;
            }

            // 在这里添加文件传输的逻辑
            console.log('文件选择成功：', file.name);
            // 假设有一个函数 uploadFile 来处理文件上传
            this.uploadFile(file);
        },

        // 示例上传函数（需要自己实现具体上传逻辑）
        async uploadFile(file) {
            // 使用 FormData 来包装文件
            // const formData = new FormData();
            // formData.append('file', file);

            // 发送文件到服务器
            try {
                const formData = new FormData();
                formData.append('room_id', this.roomId);
                formData.append('file', file);

                const response = await axios.post('/api/message/upload', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                });
                console.log(response);

                if (response.data.code === 0) {
                    // 上传成功，返回文件信息
                    console.log('File uploaded:', response.data.data);
                    const msg = {
                        room_id: this.roomId,
                        type: 'file',
                        content: response.data.data.file_name
                    }
                    console.log("msg", msg);
                    console.log(response.data.data.file_name);
                    this.$emit('send', 'file', response.data.data.file_name + ' you can download the file on http://localhost:8000' + response.data.data.file_url);
                    // this.sendMessage(msg); // 发送文件消息
                    // try {
                    //     const res = await axios.post('/api/message/send', {
                    //         room_id: this.roomId,
                    //         type: 'file', // 假设消息类型为文本，可以根据实际需要修改
                    //         content: response.data.data.file_name // 将子组件传递过来的消息内容发送到服务器
                    //     });
                    //     if (res.data.code === 0) {
                    //         // 更新消息列表，这里假设服务器返回的消息格式与历史消息格式一致
                    //         // this.messages.push(response.data.data);// 将服务器返回的消息添加到消息列表
                    //         console.log("response.data.data", res.data.data);
                    //         this.newMessage = ''; // 发送成功后清空输入框
                    //         this.messages = [];
                    //         await this.getHistoryMessages(0, 10);
                    //     } else {
                    //         console.error('Failed to send message:', res.data.msg);
                    //     }
                    // } catch (error) {
                    //     console.error('Failed to send message:', error);
                    // }

                    return response.data.data;
                } else {
                    // 上传失败，抛出错误信息
                    throw new Error(response.data.msg || '上传文件失败');
                }
            } catch (error) {
                console.error('Failed to upload file:', error);
                throw error;
            }
        },




        // uploadFile(file) {
        //     // 使用 FormData 来包装文件
        //     const formData = new FormData();
        //     formData.append('file', file);

        //     // 发送文件到服务器

        // },
        startVideo() {
            // 开启视频
            console.log('Starting video...');
            // 开启摄像头
            // 创建本地视频流
            navigator.mediaDevices.getUserMedia({ video: true, audio: false })
                .then(function (stream) {
                    var localVideo = document.getElementById("localVideo");
                    localVideo.srcObject = stream;

                    // 初始化 PeerJS
                    var peer = new Peer();

                    // 当 Peer 连接到 PeerJS 服务器时执行
                    peer.on('open', function (id) {
                        console.log('My peer ID is: ' + id);

                        // 当有来自其他 Peer 的连接请求时执行
                        peer.on('call', function (call) {
                            // 接听来自其他 Peer 的视频通话请求
                            call.answer(stream);

                            // 在远程视频元素中显示对方视频流
                            var remoteVideo = document.getElementById('remoteVideo');
                            call.on('stream', function (remoteStream) {
                                remoteVideo.srcObject = remoteStream;
                            });
                        });
                    });
                })
                .catch(function (err) {
                    console.error('getUserMedia error:', err);
                });

        }
    }
};
</script>

<style scoped>
.chat-input {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    background-color: #f4f4f4;
}

.message-input {
    flex: 1; /* 输入框占据剩余空间 */
    margin-right: 10px; /* 输入框和按钮之间的间隔 */
}

.send-button {
    width: 80px; /* 设置发送按钮的宽度 */
    height: 40px; /* 设置按钮的高度，与输入框高度一致 */
    margin-right: 10px;
    margin-left: 10px;
    margin-top: 10px;
    margin-block-end: 10px;
}
</style>
