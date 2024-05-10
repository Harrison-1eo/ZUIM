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
/* 样式与之前相同 */
</style>
