<template>
  <div class="chat-input">
    <el-input type="text" placeholder="输入消息..." v-model="newMessage"/>
    <el-button type="primary" @click="sendMessage">发送</el-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newMessage: '' // 用于绑定输入框的内容
    };
  },
  methods: {
    sendMessage() {
      // 向父组件发送输入的消息
      this.$emit('send', this.newMessage);
      this.newMessage = ''; // 发送后清空输入框
    }
  }
};
</script>

<style scoped>
.chat-input {
  display: flex;
  align-items: center;
  height: 50px; /* 设置高度为固定值 */
}

.chat-input el-input {
  flex: 1;
  margin-right: 10px; /* 调整输入框与按钮之间的间距 */
}

.chat-input el-button {
  width: 80px; /* 调整按钮的宽度 */
}
</style>
