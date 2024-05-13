<template>
  <div class="message-item" :class="{ 'message-from-user': isSenderUser }">
    <el-avatar :src="message.sender_avatar || 'http://localhost:8000/static/avatars/nopic.png'"></el-avatar>
    <div class="message-content">
      <div :class="{'sender-name-from-user': isSenderUser, 'sender-name': !isSenderUser}">
        {{ isSenderUser ? '我' : message.sender_name || 'No sender-name'}}
      </div>
      <div v-if="message.type === 'text'" class="text-message">
        {{ message.content }}
      </div>
      <div v-else-if="message.type === 'image'" class="image-message">
        <el-image :src="message.content" fit="contain" style="max-width: 100%; height: auto;"></el-image>
      </div>
      <div v-else-if="message.type === 'file'" class="file-message">
        <el-icon><Link /></el-icon>
        <el-link :href="getFileUrl(message.content)" target="_blank">下载文件：{{ this.getFileName(message.content) }}</el-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ElAvatar, ElImage, ElLink } from 'element-plus';
import {Link} from "@element-plus/icons";

export default {
  components: {
    Link,
    ElAvatar,
    ElImage,
    ElLink
  },
  props: {
    message: {
      type: Object,
      required: true
    }
  },
  computed: {
    isSenderUser() {
      const userId = localStorage.getItem('userId');
      return userId && parseInt(userId, 10) === this.message.sender_id;
    }
  },
  methods: {
    getFileUrl(content) {
      const regex = /you can download the file on (http:\/\/localhost:8000\/static\/files\/\S+)/;
      const match = content.match(regex);
      if (match) {
        console.log(match[1]); // 输出匹配到的URL
        return match[1]; // 返回匹配到的完整URL
      }
      return ""; // 如果没有匹配到，返回null
    },
    getFileName(content) {
      // content中的数据内容是：response.data.data.file_name + ' you can download the file on http://localhost:8000' + response.data.data.file_url
      // 使用正则表达式匹配出文件的名称
      const regex = /(.+) you can download the file on http:\/\/localhost:8000\/\S+/g;
      return content.replace(regex, '$1');
    }
  }
}
</script>

<style scoped>
.message-item {
  display: flex;
  align-items: flex-start;
  margin: 10px;
}
.message-from-user {
  flex-direction: row-reverse;
}
.message-content {
  margin-left: 10px;
  margin-right: 10px;
}
.sender-name {
  font-weight: bold;
}
.sender-name-from-user {
  font-weight: bold;
  text-align: right;
}

.text-message, .image-message, .file-message {
  background: #F2F2F2;
  border: 1px solid #D9D9D9;
  border-radius: 12px;
  margin: 5px;
  padding: 5px 10px 5px 10px;
}

.file-message {
  display: flex;
  align-items: center;
}

</style>
