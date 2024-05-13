<template>
    <div class="chat-input">
        <el-input type="textarea"
                  placeholder="输入消息..."
                  v-model="newMessage"
                  class="message-input"
                  :rows="4"
                  @keyup.enter="sendMessage()"
        />

        <div class="send-buttons">

          <div class="upload-icon-box">
            <!-- 发送图片 -->
            <div class="upload-icon">
              <el-tooltip
                  class="box-item"
                  effect="dark"
                  content="发送图片"
                  placement="top"
              >
                <el-icon size="30px" @click="openImageChooser"><Picture /></el-icon>
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
                <el-icon  size="30px" @click="openFileChooser"><FolderAdd /></el-icon>
              </el-tooltip>
              <!-- 隐藏的文件输入 -->
              <input type="file" ref="fileInput" style="display: none;" @change="handleFileUpload">
            </div>

          </div>

            <el-button type="success" class="send-button" @click="sendMessage()">发送消息</el-button>
            <!-- 开启视频 -->
<!--            <el-button type="success" class="send-button" @click="startVideo">开启视频</el-button>-->
        </div>

    </div>
</template>

<script>
import axios from "@/axios-config";
import {ElMessage} from "element-plus";
import {FolderAdd, Picture} from "@element-plus/icons";

export default {
  components: {
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

        // 示例上传函数（需要自己实现具体上传逻辑）
        async uploadFile(file, type) {
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
                    if (type === 'file'){
                      this.$emit('send', 'file', response.data.data.file_name + ' you can download the file on http://localhost:8000' + response.data.data.file_url);
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
    margin-top: 8px;
    margin-block-end: 10px;
}

.upload-icon-box {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 5px;
}

.upload-icon {
    margin-right: 10px;
    margin-left: 10px;
    cursor: pointer;
}
</style>
