<template>
  <div class="message-item">
    <div v-if="message.type === 'text'" class="text-message">
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
  }
};
</script>

<style scoped>
.message-item {
  margin-bottom: 10px;
}

.text-message {
  padding: 10px;
  background-color: #f0f0f0;
}

.sender-name {
  font-weight: bold;
}

.send-time {
  color: #666;
}

.pic-message {
  width: 200px;
}

.file-message {
  color: #3366cc;
}
</style>
