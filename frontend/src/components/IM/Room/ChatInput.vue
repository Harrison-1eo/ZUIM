<template>
  <div class="chat-input">
    <el-input type="textarea" placeholder="输入消息..." v-model="newMessage" class="message-input" :rows="4"/>

    <div class="send-buttons">
      <!-- 发送文件 -->
      <el-button type="primary" class="send-button" @click="sendFile" style="margin-bottom: 10px">发送文件</el-button>
      <el-button type="primary" class="send-button" @click="changeVisible" style="margin-bottom: 10px">添加附件</el-button>
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
export default {
  data() {
    return {
      newMessage: '', // 绑定输入框的内容
      dialogVisible: false
    };
  },
  methods: {
    changeVisible() {
      this.dialogVisible = !this.dialogVisible;
      console.log("changeVisible", this.dialogVisible);
    },
    sendMessage() {
      // 向父组件发送输入的消息
      this.$emit('send', this.newMessage);
      this.newMessage = ''; // 发送后清空输入框
    },
    sendFile() {
      // 发送文件

      console.log('Sending file...');
    },
    startVideo() {
      // 开启视频
      console.log('Starting video...');
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
}

.send-buttons {
  display: flex;
  align-items: center;
  flex-direction: column;
}
</style>
