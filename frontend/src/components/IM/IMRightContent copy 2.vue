<template>
  <div>
    <el-main class="im-right-content">
      <router-view></router-view>
    </el-main>
    <!-- 聊天框 -->
    <div class="chat-room">
      <h2>Chat Room</h2>
      <div class="chat-container">
        <div class="chat-box">
          <div class="chat-header">
            <span>聊天</span>
          </div>
          <div class="chat-messages" ref="chatBox">
            <div class="message" v-for="message in messages" :key="message.id">
              <p>{{ message.content }}</p>
              <small>{{ message.time }}</small>
            </div>
          </div>
        </div>
        <div class="chat-input">
          <input v-model="message" placeholder="Enter message" />
          <button @click="sendMessage">Send</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IMRightContent'
}
</script>

<style scoped>
.im-right-content {
  background-color: #d8d8d8;
  padding: 0px;
}

.chat-room {
  display: grid;
  grid-template-rows: auto 1fr auto; /* 设置网格布局，第一行高度自适应，第二行占据剩余空间，第三行高度自适应 */
  height: calc(100vh - 60px); /* 60px 是可能的其他头部和底部的高度 */
}

.chat-box {
  border: 1px solid #ccc; /* 添加边框 */
  border-radius: 5px; /* 圆角边框 */
  overflow-y: auto; /* 如果内容超出框的高度，显示滚动条 */
}

.chat-input {
  display: grid;
  grid-template-columns: 1fr auto; /* 设置网格布局，第一列占据剩余空间，第二列根据内容调整宽度 */
  align-items: center; /* 垂直居中 */
  padding: 10px; /* 添加一些内边距 */
}

.chat-header {
  background-color: #f0f0f0; /* 设置头部的背景颜色 */
  padding: 10px; /* 设置头部的内边距 */
  border-bottom: 1px solid #ccc; /* 添加底部边框 */
}

.chat-header span {
  font-weight: bold; /* 设置标题文本为粗体 */
}

.chat-messages {
  width: 100%; /* 设置聊天消息的宽度为整个框的宽度 */
}
</style>