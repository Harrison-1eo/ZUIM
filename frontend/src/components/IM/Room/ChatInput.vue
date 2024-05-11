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

            <el-button type="success" class="send-button" @click="sendMessage">发送消息</el-button>
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
        sendMessage() {
            // 向父组件发送输入的消息
            this.$emit('send', 'text', this.newMessage);
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
        // startVideo() {
        //     // 开启视频
        //     console.log('Starting video...');
        // }
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
