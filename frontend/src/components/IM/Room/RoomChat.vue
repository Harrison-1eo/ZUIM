<template>
    <div class="chat-box" id="chat-box">
        <div class="chat-messages" ref="messageContainer" id="message-box">
            <!-- <p>消息内容将在这里显示 {{ roomID }} </p> -->
            <el-scrollbar ref="scrollbar">
                <div ref="inner" class="message-inner-list">
                    <el-button type="text" link @click="getMoreHistoryMessages" style="margin: 9px"> 加载更多历史消息 </el-button>
                    <MessageItem v-for="(message, index) in messages" :key="index" :message="message" class="message" />
                    <!--                  <el-backtop :visibility-height="1" > </el-backtop>-->
                </div>
            </el-scrollbar>
    <div class="chat-details">
        <div class="chat-header">
            <p class="room-title"> {{ roomInfo.name }} </p>
            <el-icon v-if="roomID" @click="drawer=true" class="more-icon">
                <More/>
            </el-icon>
        </div>
        <RoomDrawer v-if="roomID"
                    v-model="drawer"
                    v-model:ifUpdate="ifUpdate"
                    v-model:roomInfo="roomInfo"
                    :roomID="roomID"/>

        <el-dialog v-model="recieveVideoVisible" title="视频聊天室" width="640" height="480" :visible.sync="recieveVideoVisible" :close-on-click-modal="false" :close-on-press-escape="false">
            <div class="video-container" width="640" height="500" v-if="recieveVideoVisible">
                <!-- <LiveStream :roomID="roomID" :VideoChunks="VideoChunks" /> -->

                <LiveStream :roomID="roomID" :VideoChunk="VideoChunk" @closeCamera="closeCamera" @ForceforceDeleteVideoBeTrue="ForceforceDeleteVideoBeTrue" />

        <div class="chat-box" id="chat-box">
            <div class="chat-messages" ref="messageContainer" id="message-box">
                <!-- <p>消息内容将在这里显示 {{ roomID }} </p> -->
                <el-scrollbar ref="scrollbar">
                    <div ref="inner" class="message-inner-list">
                        <el-button type="text" link @click="getMoreHistoryMessages" style="margin: 9px">
                            加载更多历史消息
                        </el-button>
                        <MessageItem v-for="(message, index) in messages" :key="index" :message="message"
                                     class="message"/>
                    </div>
                </el-scrollbar>
            </div>

        <div class="input-box" id="input-box">
            <ChatInput ref="ChatInput" :roomID="roomID" @send="sendMessageToParent" @ForceforceDeleteVideoBeFalse="ForceforceDeleteVideoBeFalse" />
            <el-dialog v-model="recieveVideoVisible" title="视频聊天室" width="640" height="480"
                       v-model:visible="recieveVideoVisible" :close-on-click-modal="false" :close-on-press-escape="false">
                <div class="video-container" width="640" height="500" v-if="recieveVideoVisible">
                    <LiveStream :roomID="roomID" :VideoChunk="VideoChunk"/>
                </div>
            </el-dialog>

            <div class="input-box" id="input-box">
                <ChatInput :roomID="roomID" @send="sendMessageToParent"/>
            </div>
        </div>
    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import MessageItem from "@/components/IM/Room/MessageItem.vue";
import ChatInput from "@/components/IM/Room/ChatInput.vue";
import {ElMessage} from "element-plus";
import WebsocketClass from "@/utils/websocket";
// import { createFFmpeg, fetchFile, FFmpeg } from "@ffmpeg/ffmpeg";
import LiveStream from "@/components/IM/Room/LiveStream.vue"
import { wsBaseUrl } from "@/utils/base-url-setting";
import {wsBaseUrl} from "@/utils/base-url-setting";
import RoomDrawer from "@/components/IM/Room/RoomDrawer.vue";

export default {
    name: 'RoomChat',
    props: {
        roomID: {
            type: Number,
            required: true
        },
        ifUpdateRoomList2Chat: {
            type: Boolean,
        },
    },
    components: {
        RoomDrawer,
        MessageItem,
        ChatInput,
        LiveStream
    },
    data() {
        return {
            roomInfo: {},
            messages: [], // 用于存储消息列表
            newMessage: '', // 用于绑定输入框的内容
            // WebSocket 对象
            ws: new WebsocketClass(
                    wsBaseUrl,
                    localStorage.getItem('token'),
                    this.addNewMessage,
                    true,
            ),
            drawer: false,
            ifUpdate: false,
            currentUser: null, // 当前用户信息
            mediaSource: new MediaSource(),
            sourceBuffer: [],
            recieveVideoVisible: false,
            VideoChunks: [],
            VideoChunk: null,
            videoElement: null,
            videoDuration: 0,
            ffmpeg: null,
            videoTime: 10,
            forceDeleteVideo: false,
        };
    },
    created() {
        this.getHistoryMessages(0, 10); // 组件创建时调用获取历史消息函数
        this.fetchRoomInfo();
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
            this.$router.push('/login');
        }
    },
    watch: {
        roomID() {
            this.messages = [];                           // 切换房间时清空消息列表
            this.getHistoryMessages(0, 10); // 监听 roomID 变化时重新获取历史消息
            this.fetchRoomInfo();
        },
        messages: {
            deep: true,
            handler() {
                this.handleNewMessage();
            }
        },
        ifUpdate() {
            if (this.ifUpdate) {
                this.$emit('update:ifUpdateRoomList2Chat', true);
                this.ifUpdate = false;
            }
        },
    },
    mounted() {
        this.scrollToBottom();
    },
    beforeUnmount() {
        this.ws.close();
    },
    methods: {
        // 对新消息进行处理
        addNewMessage(message) {
            // 接收来自服务器的消息并处理
            // console.log('Received new message:', message);
        fetchRoomInfo() {
            axios_config.get(`/api/room/info?room_id=${this.roomID}`).then(response => {
                if (response.data.code === 0) {
                    this.roomInfo = response.data.data;
                    console.log('获取房间信息成功:', this.roomInfo);
                } else {
                    console.error('获取房间信息失败:', response.data.msg);
                }
            }).catch(error => {
                console.error('获取房间信息失败:', error);
            });
        },
        // 对新消息进行处理
        addNewMessage(message) {
            // 接收来自服务器的消息并处理
            console.log('Received new message:', message);
            if (message.code === 0) {
                if (message.data.room_id !== this.roomID) {
                    return;
                }
                // message.data.sender_avatar = this.currentUser.avatar;
                // console.log('Received new message:', message);
                message.data.encryptInfo = message.en_data;
                message.data.encryptInfo.key = localStorage.getItem('websocketBackendPassword');
                if (message.data.type === 'video') {
                    if (this.forceDeleteVideo === false) {
                        console.log('forceDeleteVideo is false');
                        this.recieveVideoVisible = true;
                        this.ReceiveVideoChunk_Base64(message.data.content);
                    } else {
                        console.log('forceDeleteVideo is true');
                    }

                console.log('Received new message:', message);
                message.data.encryptInfo = message.en_data;
                message.data.encryptInfo.key = localStorage.getItem('websocketBackendPassword');
                if (message.data.type === 'video') {
                    this.recieveVideoVisible = true;
                    this.ReceiveVideoChunk_Base64(message.data.content);
                } else {
                    this.messages.push(message.data);
                }
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
            // console.log('Sending message to parent:', message);
            console.log('Sending message to parent:', message);
            // 在这里可以进行进一步的处理，比如发送给服务器等操作
            // this.sendMessageToServer(message);
            this.ws.send({
                room_id: this.roomID,
                type: type,
                content: message
            });
        },
        ReceiveVideoChunk_Base64(VideoChunk_Base64) {
            this.VideoChunks.push(VideoChunk_Base64);
            this.VideoChunk = VideoChunk_Base64;
        },
        // 用于获取历史消息
        async getHistoryMessages(lastMessageId, limit) {
            try {
                const response = await axios_config.post(
                    '/api/message/list',
                    {
                        "room_id": this.roomID,
                        "last_message_id": lastMessageId,
                        "limit": limit
                    });
                //   console.log('roomID:', this.roomID);
                //   console.log('lastMessageId:', lastMessageId);
                //     console.log('limit:', limit);
                //   console.log('Fetched history messages111:', response.data.data);
                if (response.data.data.length === 0) {
                    ElMessage.info('没有更多历史消息了')
                    return;
                }
                response.data.data.forEach(item => {
                    item.encryptInfo = response.data.en_data;
                    item.encryptInfo.key = localStorage.getItem('backendPassword');
                });
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
            console.log('minMessageId:', minMessageId);
            this.getHistoryMessages(minMessageId, 10);
        },
        // 滚动消息列表到底部
        scrollToBottom() {
            this.$nextTick(() => {
                const container = this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap');
                container.scrollTop = container.scrollHeight;
            });
            this.$refs.scrollbar.handleScroll(
                    {direction: 'bottom', target: this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap')}
            )
        },
        //判断消息列表是否在底部
        isAtBottom() {
            const container = this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap');
            // 判断精度，差值小于10px则认为在底部
            return container.scrollTop + container.clientHeight + 10 >= container.scrollHeight;
            // return container.scrollTop + container.clientHeight === container.scrollHeight;
        },
        // 当有新消息时，并且消息列表在底部时，自动滚动到底部，保持最新消息可见
        handleNewMessage() {
            if (this.isAtBottom()) {
                console.log('scroll to bottom');
                this.scrollToBottom();
            }
        },
        closeCamera() {
            this.$refs.ChatInput.closeCamera();
            this.recieveVideoVisible = false;
        },
        ForceforceDeleteVideoBeFalse() {
            this.forceDeleteVideo = false;
        },
        ForceforceDeleteVideoBeTrue() {
            this.forceDeleteVideo = true;
        }
    }
};
</script>

<style scoped>

.chat-header {
    padding: 10px 30px;
    background-color: #f5f7fa;
    text-align: center;
    height: 3em;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.more-icon {
    cursor: pointer;
}

.room-title {
    font-size: 20px;
    font-weight: bold;
}

.chat-box {
    margin: 0 60px;
    height: calc(100vh - 8em);
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
    height: calc(100% - 60px - 52px - 221px - 1em); /* 减去头部和输入框的高度 */
}

.message-inner-list {
    display: flex;
    flex-direction: column;
}
</style>

