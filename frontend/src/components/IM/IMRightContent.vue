<template>
    <div class="chat-room">
        <h2>Chat Room</h2>
        <el-container class="chat-container">
            <el-main class="chat-box">
                
                    <!-- <div slot="header" class="chat-header">
                        <span>聊天</span>
                    </div> -->
                    <div class="chat-messages">
                        <div class="message" v-for="message in messages" :key="message.id" :class="{ 'own-message': message.own }">
                            <div class="avatar">
                                <img :src="message.avatar" alt="avatar">
                            </div>
                            <div class="message-content">
                                <p>{{ message.content }}</p>
                                <small>{{ message.time }}</small>
                            </div>
                        </div>
                    </div>
                
            </el-main>
            <el-footer class="chat-input">
                <el-input v-model="message" placeholder="Enter message" />
                <el-button type="primary" @click="sendMessage">Send</el-button>
            </el-footer>
        </el-container>
    </div>
</template>

<script>
export default {
    name: 'ChatRoom',
    data() {
        return {
            messages: [],
            message: ''
        };
    },
    methods: {
        sendMessage() {
            if (this.message.trim() === '') {
                this.$message.error("消息内容不能为空");
                return;
            }
            // 模拟发送消息，实际中需要调用后端接口发送消息
            const newMessage = {
                id: this.messages.length + 1,
                content: this.message,
                time: new Date().toLocaleString(),
                own: true, // 标记消息是自己发送的
                avatar: "../../assets/logo.png" // 自己的头像地址
            };
            this.messages.push(newMessage);
            this.message = ''; // 清空输入框

            this.$nextTick(() => {
                const chatBox = this.$refs.chatBox;
                if (chatBox) {
                    chatBox.scrollTop = chatBox.scrollHeight;
                }
            });
        }

    }
};
</script>

<style scoped>
.chat-room {
    padding: 20px;
}

.chat-container {
    height: calc(50vh - 150px); /* 适配屏幕高度 */
}

.chat-box {
    overflow-y: auto;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.chat-header {
    background-color: #f0f0f0;
    padding: 10px;
}

.chat-messages {
    padding: 10px;
}

.chat-input {
    padding-top: 10px;
    display: flex;
    align-items: center;
}
</style>
