<template>
  <div class="message-item" :class="{ 'message-right': align === 'right', 'message-left': align === 'left' }">
    <div v-if="message.type === 'text'" class="text-message">
      <div v-if="align === 'left'" class="avatar">
        <img src="avatar" alt="avatar" />
      </div>
      <span class="sender-name">{{ message.sender_name }}</span>
      <span class="send-time">{{ message.send_time }}</span>
      <p>{{ message.content }}</p>
    </div>
    <el-image v-else-if="message.type === 'pic'" :src="message.content" class="pic-message"></el-image>
    <el-link v-else-if="message.type === 'file'" :href="message.content" target="_blank" class="file-message">{{ message.content }}</el-link>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: Object,
      required: true
    }
  },
  avatar: {
      type: "../../assets/avatar/dick1.jpg", // 头像的URL
      default: "../../assets/avatar/dick1.jpg" // 默认为空
    }
};
</script>

<style scoped>
.message-item {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
}

.message-right {
  align-self: flex-end;
}

.message-left {
  align-self: flex-start;
}

.text-message {
  padding: 10px;
  border-radius: 5px;
  background-color: #f0f0f0;
}

.sender-name {
  font-weight: bold;
}

.send-time {
  font-size: 12px;
  color: #888;
}

.pic-message {
  max-width: 200px;
  max-height: 200px;
}

.file-message {
  padding: 5px 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}
</style>