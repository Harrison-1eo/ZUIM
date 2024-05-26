<template>
    <div class="chat-input">

        <div class="chat-input-icon-box">
            <!-- 发送图片 -->
            <div class="upload-icon">
                <el-tooltip
                        class="box-item"
                        effect="dark"
                        content="发送图片"
                        placement="top"
                >
                    <el-icon size="20px" @click="openImageChooser"><Picture /></el-icon>
                </el-tooltip>
                <input type="file" ref="imageInput" style="display: none;" accept="image/*" @change="handleImageUpload">
            </div>
            <!-- 发送文件 -->
            <div class="upload-icon">
                <el-tooltip
                        class="box-item"
                        effect="dark"
                        content="发送文件"
                        placement="top"
                >
                    <!-- 文件上传按钮 -->
                    <!--        <el-button type="primary" class="send-button" @click="openFileChooser" style="margin-bottom: 10px">发送文件</el-button>-->
                    <el-icon  size="20px" @click="openFileChooser"><FolderAdd /></el-icon>
                </el-tooltip>
                <!-- 隐藏的文件输入 -->
                <input type="file" ref="fileInput" style="display: none;" @change="handleFileUpload">
            </div>
        </div>

        <el-divider style="margin: 0 5px"/>

        <div class="chat-input-input-box">
<!--            <el-input type="textarea"-->
<!--                      placeholder="输入消息..."-->
<!--                      v-model="newMessage"-->
<!--                      resize="none"-->
<!--                      style="padding: 10px; border-radius: 10px;"-->
<!--                      :rows="4"-->
<!--                      @keyup.enter="sendMessage()"-->
<!--            />-->
            <textarea
                    rows="3"
                    v-model="newMessage"
                    placeholder="输入消息..."
                    class="message-input"
                    @keyup.enter="sendMessage()"
            />
        </div>

        <div class="chat-input-send-box">
            <el-button
                type="success"
                class="send-button"
                @click="sendMessage()">
                <!-- 发送消息 -->
                <el-icon size="20px"><Promotion /></el-icon>
            </el-button>
        </div>

    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import {ElMessage} from "element-plus";
import {FolderAdd, Picture} from "@element-plus/icons";
import {Bottom} from "@element-plus/icons-vue";

export default {
  components: {
      Bottom,
    Picture,
    FolderAdd,
  },
    props: {
        roomID: {
            type: Number,
            required: true
        }
    },
    data() {
        return {
            newMessage: '', // 绑定输入框的内容
            roomId: this.roomID
        };
    },
    methods: {
        sendMessage() {
            // 如果消息以回车结尾，去掉回车
            this.newMessage = this.newMessage.trim();
            // 如果消息为空，不发送
            if (!this.newMessage) {
              ElMessage.warning('消息不能为空');
              return;
            }
            // 向父组件发送输入的消息
            this.$emit('send', 'text', this.newMessage);
            this.newMessage = ''; // 发送后清空输入框
        },
        // 打开文件选择器
        openFileChooser() {
            this.$refs.fileInput.click();
        },
        openImageChooser() {
          this.$refs.imageInput.click();
        },
        handleImageUpload(event) {
          const file = event.target.files[0];
          if (file) {
            this.uploadFile(file, 'image');
          }
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
            this.uploadFile(file, 'file');
        },

        async uploadFile(file, type) {
            // 发送文件到服务器
            try {
                const formData = new FormData();
                formData.append('room_id', this.roomId);
                formData.append('file', file);

                const response = await axios_config.post('/api/message/upload', formData, {
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
                    if (type === 'file'){
                      this.$emit('send', 'file', response.data.data.file_name + ' you can download the file on ' + response.data.data.file_url);
                    } else if (type === 'image'){
                      this.$emit('send', 'image', response.data.data.file_url);
                    } else {
                      throw new Error('上传文件失败，无效的文件类型');
                    }
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
    }
};
</script>

<style scoped>
.chat-input {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    margin: 10px 0px;
    padding: 5px 10px;
    border-radius: 8px;
    border: 1px solid #ffffff;
    background: #ffffff;
    box-shadow: 0 16px 20px 0 rgba(174,167,223,.2);
    transition: border 0.5s linear;
}

.chat-input:hover {
    border: 1px solid #4955f5;
}

.chat-input-icon-box {
    padding: 5px 10px;
    display: flex;
    flex-direction: row;
    justify-content: start;
}

.upload-icon {
    margin-right: 10px;
    cursor: pointer;
}

.chat-input-input-box {
    padding: 5px 10px;
    caret-color: transparent;
}

.message-input {
    resize: none;
    flex: 1; /* 输入框占据剩余空间 */
    border: none;
    background: none;
    width: 100%;
    font-size: medium;
    font-family: "微软雅黑","黑体",serif;
    caret-color: #000000;
}

.message-input:focus {
    outline: none;
}



.chat-input-send-box {
    display: flex;
    flex-direction: row-reverse;
    justify-content: end;
}

.send-button {
    width: 60px; /* 设置发送按钮的宽度 */
    height: 35px; /* 设置按钮的高度，与输入框高度一致 */
    margin-right: 10px;
    margin-left: 10px;
    margin-top: 8px;
    margin-block-end: 10px;
    background: linear-gradient(322deg,#1f6ff5 7.93%,#9a7ffc 92.22%);
    border: none;
    border-radius: 20px;
    transition: background 0.5s linear;
}

.send-button:hover {
    background: linear-gradient(322deg, #628eed 7.93%, #a1a4f8 92.22%);
}


</style>
