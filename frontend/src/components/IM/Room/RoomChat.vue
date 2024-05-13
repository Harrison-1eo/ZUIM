<template>
    <div class="chat-box" id="chat-box">
        <div class="chat-messages" ref="messageContainer" id="message-box">
            <!-- <p>消息内容将在这里显示 {{ roomID }} </p> -->
            <el-scrollbar ref="scrollbar">
                <div ref="inner" class="message-inner-list">
                  <el-button type="text" link @click="getMoreHistoryMessages" style="margin: 9px"> 加载更多历史消息 </el-button>
                  <MessageItem v-for="(message, index) in messages" :key="index" :message="message" class="message" />
                </div>
            </el-scrollbar>
        </div>
        <div class="input-box" id="input-box">
            <ChatInput :roomID="roomID" @send="sendMessageToParent" />
        </div>
    </div>

</template>

<script>
import axios from "@/axios-config";
import MessageItem from "@/components/IM/Room/MessageItem.vue";
import ChatInput from "@/components/IM/Room/ChatInput.vue";
import { ElMessage } from "element-plus";
import WebsocketClass from "@/utils/websocket";

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
            ws: new WebsocketClass(
                `ws://localhost:8000/ws/join`,
                localStorage.getItem('token'),
                this.addNewMessage
            ), // WebSocket 对象
            currentUser: null, // 当前用户信息
        };
    },
    created() {
        this.getHistoryMessages(0, 10); // 组件创建时调用获取历史消息函数
        this.ws.connect();
         // 创建 WebSocket 实例
        // 定义当前用户信息
        const userId = localStorage.getItem('userId');
        const username = localStorage.getItem('user');
        if (userId && username) {
            this.currentUser = {
                user_id: JSON.parse(userId),
                username: JSON.parse(username),
                avatar: "../../../assets/avatar/dick1.jpg", // 这里需要根据实际情况获取用户头像
            };
        } else {
            // 处理未登录情况，比如跳转到登录页面
        }
    },
    watch: {
        roomID() {
            this.messages = [];                           // 切换房间时清空消息列表
            this.getHistoryMessages(0, 10); // 监听 roomID 变化时重新获取历史消息
        },
        messages: {
            deep: true,
            handler() {
              this.handleNewMessage();
            }
        }
    },
    mounted() {
      this.scrollMaxHeight = this.$refs.scrollbar.$el.clientHeight - 380;
    },
    beforeUnmount() {
      this.ws.close();
    },
    methods: {
      // 对新消息进行处理
      addNewMessage(message) {
          // 接收来自服务器的消息并处理
          console.log('Received new message:', message);
          if (message.code === 0) {
            if (message.data.room_id !== this.roomID) {
              return;
            }
            this.messages.push(message.data);
          } else if (message.code === 1) {
              console.error('收到错误消息:', message.msg);
          } else if (message.code === 2) {
              // 忽略心跳包
          } else {
              console.error('收到未知消息:', message);
          }
      },
      // 接受来自 ChatInput 组件的消息并发送
      sendMessageToParent(type, message) {
          // 接收来自子组件的消息并处理
          console.log('Sending message to parent:', message);
          // 在这里可以进行进一步的处理，比如发送给服务器等操作
          // this.sendMessageToServer(message);
          this.ws.send({
              room_id: this.roomID,
              type: type,
              content: message
          });
      },
      // 用于获取历史消息
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
              console.log('Fetched history messages:', response.data.data);

              // response.data.data 获取的消息列表ID是从大到小的，需要将其反转
              this.messages = [...response.data.data.reverse(), ...this.messages];
          } catch (error) {
              console.error('Failed to fetch messages:', error);
          }
      },
      // 用于点击“加载更多历史消息”按钮时调用
      getMoreHistoryMessages() {
          console.log(this.messages);
          if (this.messages.length === 0) {
              ElMessage.info('没有更多历史消息了')
              return;
          }
          if (this.messages[this.messages.length - 1].ID <= 1) {
              ElMessage.info('没有更多历史消息了')
              return;
          }
          const minMessageId = Math.min(...this.messages.map(item => item.ID));
          console.log('Called with lastMessageId:', minMessageId);
          this.getHistoryMessages(minMessageId, 10);
      },
      // 滚动消息列表到底部
      scrollToBottom() {
        this.$nextTick(() => {
          const container = this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap');
          container.scrollTop = container.scrollHeight;
        });
        this.$refs.scrollbar.handleScroll(
          { direction: 'bottom', target: this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap') }
        )
      },
      // 判断消息列表是否在底部
      isAtBottom() {
        const container = this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap');
        return container.scrollTop + container.clientHeight === container.scrollHeight;
      },
      // 当有新消息时，并且消息列表在底部时，自动滚动到底部，保持最新消息可见
      handleNewMessage() {
        if (this.isAtBottom()) {
          console.log('scroll to bottom');
          this.scrollToBottom();
        }
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
.input-box {
    padding: 0;
}

#message-box {
    height: calc(100% - 72px - 42px); /* 减去头部和输入框的高度 */
}

.message-inner-list {
    display: flex;
    flex-direction: column;
}
</style>

