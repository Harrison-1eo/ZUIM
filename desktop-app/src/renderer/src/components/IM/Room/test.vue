<template>
  <div class="chat-input">
    <el-input type="textarea" placeholder="输入消息..." v-model="newMessage" class="message-input" :rows="4" />

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
            <el-icon size="30px"><Picture /></el-icon>
          </el-tooltip>
          <!-- 图片文件输入 -->
          <input type="file" ref="imageInput" style="display: none;" accept="image/*" @change="handleImageUpload">
          <el-icon size="30px" @click="openImageChooser"><Picture /></el-icon>
        </div>

        <!-- 发送文件 -->
        <div class="upload-icon">
          <el-tooltip
              class="box-item"
              effect="dark"
              content="发送文件"
              placement="top"
          >
            <el-icon  size="30px" @click="openFileChooser"><FolderAdd /></el-icon>
          </el-tooltip>
          <!-- 隐藏的文件输入 -->
          <input type="file" ref="fileInput" style="display: none;" @change="handleFileUpload">
        </div>
      </div>

      <el-button type="success" class="send-button" @click="sendMessage('text')">发送消息</el-button>
    </div>

  </div>
</template>

<script>
import axios from "@/axios-config";
import {ElMessage} from "element-plus";
import {FolderAdd, Picture} from "@element-plus/icons-vue";

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
    sendMessage(type) {
      if (!this.newMessage) {
        ElMessage.warning('消息不能为空');
        return;
      }
      this.$emit('send', type, this.newMessage);
      this.newMessage = '';
    },
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
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.uploadFile(file, 'file');
      }
    },
    async uploadFile(file, type) {
      try {
        const formData = new FormData();
        formData.append('room_id', this.roomId);
        formData.append('file', file);

        const response = await axios.post(`/api/message/upload/${type}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        if (response.data.code === 0) {
          const fileUrl = `http://localhost:8000${response.data.data.file_url}`;
          const message = `${response.data.data.file_name} you can download/view the ${type} on ${fileUrl}`;
          this.$emit('send', type, message);
        } else {
          throw new Error(response.data.msg || `上传${type}失败`);
        }
      } catch (error) {
        console.error(`Failed to upload ${type}:`, error);
        ElMessage.error(`上传${type}失败`);
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
.message-input
