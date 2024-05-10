<template>
  <div class="chat-box" id="chat-box">
    <div class="chat-messages" ref="messageContainer" id="message-box">
<!--      <p>消息内容将在这里显示 {{ roomID }} </p>-->
      <el-scrollbar ref="scrollbar" height="500px">
        <div ref="inner">
          <el-button type="text" text> 加载更多历史消息 </el-button>
          <MessageItem v-for="(message, index) in messages" :key="index" :message="message" class="message"/>
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
import loading from "@element-plus/icons/lib/Loading";

export default {
  name: 'RoomChat',
  computed: {
    loading() {
      return loading
    }
  },
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
    this.getHistoryMessages(0); // 组件创建时调用获取历史消息函数
    this.initWebSocket(); // 初始化 WebSocket
  },
  watch: {
    roomID() {
      this.messages = []; // 切换房间时清空消息列表
      this.getHistoryMessages(0); // 监听 roomID 变化时重新获取历史消息
    }
  },
  methods: {
    sendMessageToParent(message) {
      // 接收来自子组件的消息并处理
      console.log('Sending message to parent:', message);
      // 在这里可以进行进一步的处理，比如发送给服务器等操作
    },
    async getHistoryMessages(lastMessageId) {
      try {
        const response = await axios.post(
            'http://localhost:8000/api/message/list',
            {
              "room_id": this.roomID,
              "last_message_id": lastMessageId,
              "limit": 20
            },
            {
              headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` }
            });
        this.messages = [...this.messages, ...response.data.data];
        this.$nextTick(() => {
          this.scrollChatToBottom(); // 获取历史消息后滚动到底部
        });
      } catch (error) {
        console.error('Failed to fetch messages:', error);
      }
    },
    sendMessage() {
      // 发送消息逻辑（使用 WebSocket）
      if (!this.newMessage.trim()) return;
      this.ws.send(JSON.stringify({
        type: 'text',
        content: this.newMessage
      }));
      this.newMessage = ''; // 发送后清空输入框
    },
    initWebSocket() {
      this.ws = new WebSocket('ws://localhost:8000/api/message/ws');
      this.ws.onopen = () => {
        console.log('WebSocket connected');
      };
      this.ws.onmessage = (event) => {
        const message = JSON.parse(event.data);
        this.messages.push(message); // 接收到新消息时添加到消息列表
        this.$nextTick(() => {
          this.scrollChatToBottom(); // 滚动到底部
        });
      };
      this.ws.onerror = (error) => {
        console.error('WebSocket error:', error);
      };
      this.ws.onclose = () => {
        console.log('WebSocket closed');
      };
    },
    scrollChatToBottom() {
      const messageContainer = this.$refs.messageContainer;
      messageContainer.scrollTop = messageContainer.scrollHeight;
    }
  }
};
</script>

<style scoped>
.chat-box {
  height: 100%;
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
  //flex: 0;
  padding: 10px;
  height: 50px;
}

#message-box {
  height: calc(100% - 72px - 42px); /* 减去头部和输入框的高度 */
}
</style>

