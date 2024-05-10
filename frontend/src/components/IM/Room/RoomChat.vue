<template>
  <div class="chat-box" id="chat-box">
    <div class="chat-messages" ref="messageContainer" id="message-box">
<!--      <p>消息内容将在这里显示 {{ roomID }} </p>-->
      <el-scrollbar ref="scrollbar">
        <div ref="inner" class="message-inner-list">
          <MessageItem v-for="(message, index) in messages" :key="index" :message="message" class="message"/>
          <el-button type="text" link @click="getMoreHistoryMessages" style="margin-bottom: 9px"> 加载更多历史消息 </el-button>
        </div>
      </el-scrollbar>
    </div>
    <div class="input-box" id="input-box">
      <ChatInput @send="sendMessageToParent"/>
    </div>
  </div>


</template>

<script>
import axios from "@/axios-config";
import MessageItem from "@/components/IM/Room/MessageItem.vue";
import ChatInput from "@/components/IM/Room/ChatInput.vue";
import {ElMessage} from "element-plus";

export default {
  name: 'RoomChat',
  props: {
    roomID: {
      type: Number,
      required: true
    }
  },
  components: {
    MessageItem,
    ChatInput
  },
  data() {
    return {
      messages: [], // 用于存储消息列表
      newMessage: '', // 用于绑定输入框的内容
      ws: null, // WebSocket 对象
    };
  },
  created() {
    this.getHistoryMessages(0,10); // 组件创建时调用获取历史消息函数
  },
  watch: {
    roomID() {
      this.messages = []; // 切换房间时清空消息列表
      this.getHistoryMessages(0, 10); // 监听 roomID 变化时重新获取历史消息
    }
  },
  methods: {
    sendMessageToParent(message) {
      // 接收来自子组件的消息并处理
      console.log('Sending message to parent:', message);
      // 在这里可以进行进一步的处理，比如发送给服务器等操作
    },
    async getHistoryMessages(lastMessageId, limit) {
      try {
        const response = await axios.post(
            'http://localhost:8000/api/message/list',
            {
              "room_id": this.roomID,
              "last_message_id": lastMessageId,
              "limit": limit
            });
        if (response.data.data.length === 0) {
          ElMessage.info('没有更多历史消息了')
          return;
        }
        this.messages = [...this.messages, ...response.data.data];
      } catch (error) {
        console.error('Failed to fetch messages:', error);
      }
    },
    getMoreHistoryMessages() {
      console.log(this.messages);
      if (this.messages[this.messages.length - 1].ID <= 1) {
        ElMessage.info('没有更多历史消息了')
        return;
      }
      const lastMessageId = this.messages.length > 0 ? this.messages[this.messages.length - 1].ID : 0;
      console.log('Called with lastMessageId:', lastMessageId);
      this.getHistoryMessages(lastMessageId, 10);
    },
  }
};
</script>

<style scoped>
.chat-box {
  height: calc(100vh - 7em);
  display: flex;
  flex-direction: column;
  position: relative; /* 添加相对定位 */
}

.chat-messages {
  flex: 1;
}

.message {
  margin-bottom: 10px;
}
.input-box{
  padding: 0;
}

#message-box {
  height: calc(100% - 72px - 42px); /* 减去头部和输入框的高度 */
}

.message-inner-list{
  display: flex;
  flex-direction: column;
}
</style>

