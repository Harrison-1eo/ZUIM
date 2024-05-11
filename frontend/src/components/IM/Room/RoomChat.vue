<template>
    <!-- 邀请别人进入聊天室 -->
    <div class="invite-frineds" id="invite-friends" style="text-align: center; margin-bottom: 30px">
        <!-- <div style="margin-bottom: 10px; height:60">
            <el-input v-model="inviteUsername" placeholder="请输入好友用户名" style="width:300px; height=60px"></el-input>
        </div> -->
        <el-button class="float-right-button" type="primary" @click="inviteFriend">邀请好友进入聊天室</el-button>
    </div>
    <div class="chat-box" id="chat-box">
        <div class="chat-messages" ref="messageContainer" id="message-box">
            <!-- 一个按钮清空聊天记录 -->
            <div style="text-align: center; margin-bottom: 10px; height: 50px">
                <el-button class="float-right-button"type="primary" @click="clearMessages">清空聊天记录</el-button>
            </div>
            <!-- <p>消息内容将在这里显示 {{ roomID }} </p> -->
            <el-scrollbar ref="scrollbar">
                <!-- <div class="avatar">
                    <img src="" alt="avatar" />
                    <span>{{item.from_user}}</span>“聊天姓名”
                </div> -->
                <div class="message-container">
                    <div ref="inner" class="message-inner-list">
                        <!-- 设置头像 -->
                        <!-- <div class="avatar">
                            <img src="../../assets/avatar/dick1.png" alt="avatar" />
                        </div> -->
                        <MessageItem v-for="(message, index) in messages" :align="message.sender_id === currentUser.user_id ? 'right' : 'left'" :avatar="message.sender_id === currentUser.user_id ? currentUser.avatar : 'frontend/src/assets/avatar/dick1.jpg'":key="index" :message="message" class="message" />
                        <el-button type="text" link @click="getMoreHistoryMessages" style="margin-bottom: 9px"> 加载更多历史消息 </el-button>
                    </div>
                </div>
            </el-scrollbar>
        </div>
        <div class="input-box" id="input-box">
            <ChatInput @send="sendMessageToParent" />
        </div>
    </div>

</template>

<script>
import axios from "@/axios-config";
import MessageItem from "@/components/IM/Room/MessageItem.vue";
import ChatInput from "@/components/IM/Room/ChatInput.vue";
import { ElMessage } from "element-plus";

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
        this.getHistoryMessages(0, 10); // 组件创建时调用获取历史消息函数
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
            this.messages = []; // 切换房间时清空消息列表
            this.getHistoryMessages(0, 10); // 监听 roomID 变化时重新获取历史消息
        },

    },
    methods: {
        // 情况消息暂时不用
        clearMessages() {
            this.messages = [];
            // 向后端发送清空消息请求
            const response = axios.post('/api/message/clear', {
                room_id: this.roomID
            });
            if (response.data.code === 0) {
                ElMessage.success('清空成功');
            } else {
                ElMessage.error('清空失败');
            }
        },
        sendMessageToParent(message) {
            // 接收来自子组件的消息并处理
            console.log('Sending message to parent:', message);
            // 在这里可以进行进一步的处理，比如发送给服务器等操作
            this.sendMessageToServer(message);
            // this.getHistoryMessages(0, 10);
        },
        async sendMessageToServer(message) {
            try {
                const response = await axios.post('/api/message/send', {
                    room_id: this.roomID,
                    type: 'text', // 假设消息类型为文本，可以根据实际需要修改
                    content: message // 将子组件传递过来的消息内容发送到服务器
                });

                // 检查服务器是否成功接收并处理消息
                if (response.data.code === 0) {
                    // 更新消息列表，这里假设服务器返回的消息格式与历史消息格式一致
                    // this.messages.push(response.data.data);// 将服务器返回的消息添加到消息列表
                    this.newMessage = ''; // 发送成功后清空输入框

                    this.getHistoryMessages(0, 10);
                } else {
                    console.error('Failed to send message:', response.data.msg);
                }
            } catch (error) {
                console.error('Failed to send message:', error);
            }
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
                console.log('Fetched history messages:', response.data.data);
                this.messages = [...response.data.data];

                // console.log('Fetched history messages:', response.data.data[0]);
                // this.messages = [...response.data.data];
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
            
        }
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
.message-container {
    height: calc(100vh - 400px); /* 指定固定高度 */
    overflow: auto; /* 添加滚动条 */
}
.message {
    margin-bottom: 10px;
}
.input-box {
    height: 200px;
    padding: 0;
}
.float-right-button {
  float: right;
  min-width: 66px;
  min-height: 48px;
  max-width: 188px;
    max-height: 68px;
}
#message-box {
    height: calc(100% - 72px - 42px); /* 减去头部和输入框的高度 */
}

.message-inner-list {
    display: flex;
    flex-direction: column;
}
</style>

