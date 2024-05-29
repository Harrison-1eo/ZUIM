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
        </div>
        <el-dialog v-model="recieveVideoVisible" title="视频聊天室" width="640" height="480" :visible.sync="recieveVideoVisible" :close-on-click-modal="false" :close-on-press-escape="false">
            <div class="video-container" width="640" height="500" v-if="recieveVideoVisible">
                <!-- <LiveStream :roomID="roomID" :VideoChunks="VideoChunks" /> -->
                
                <LiveStream :roomID="roomID" :VideoChunk="VideoChunk" />

            </div>
        </el-dialog>

        <div class="input-box" id="input-box">
            <ChatInput :roomID="roomID" @send="sendMessageToParent" />
        </div>
        <!-- 隐藏的视频 -->

    </div>

</template>

<script>
import axios_config from "@/utils/axios-config";
import MessageItem from "@/components/IM/Room/MessageItem.vue";
import ChatInput from "@/components/IM/Room/ChatInput.vue";
import { ElMessage } from "element-plus";
import WebsocketClass from "@/utils/websocket";
// import { createFFmpeg, fetchFile, FFmpeg } from "@ffmpeg/ffmpeg";
import LiveStream from "@/components/IM/Room/LiveStream.vue"
import LoginView from '../../../../../desktop-app/src/renderer/src/views/LoginView.vue';
import {wsBaseUrl} from "@/utils/base-url-setting";

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
        ChatInput,
        LiveStream
    },
    data() {
        return {
            messages: [], // 用于存储消息列表
            newMessage: '', // 用于绑定输入框的内容
            ws: new WebsocketClass(
                wsBaseUrl,
                localStorage.getItem('token'),
                this.addNewMessage,
                true,
            ), // WebSocket 对象
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
            this.$router.push('/login');
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
        },
        // VideoChunks: {
        //     deep: true,
        //     handler() {
        //         console.log('VideoChunks changed in RoomChat');
        //     }
        // }

    },
    mounted() {
        this.scrollToBottom();
        // this.setupVideoPlayer();
        // this.initializeLiveStreaming();


    },
    beforeUnmount() {
        this.ws.close();
    },
    methods: {

        addNewMessage(message) {
            // 接收来自服务器的消息并处理
            // console.log('Received new message:', message);
            if (message.code === 0) {
                if (message.data.room_id !== this.roomID) {
                    return;
                }
                // message.data.sender_avatar = this.currentUser.avatar;
                console.log('Received new message:', message);
                // 如果不是video类型的数据，直接添加到消息列表
                if (message.data.type === 'video') {
                    // 视频流的处理方式，将视频展示出来
                    // 先解码
                    console.log('丢雷老母');
                    // console.log('Received video message:', message.data.content);
                    this.recieveVideoVisible = true;
                    console.log('recieveVideoVisible changed to true');
                    // console.log('Received video message:', message.data);
                    this.ReceiveVideoChunk_Base64(message.data.content);
                    console.log('god damn it');

                } else {
                    console.log('in else   Received text message:', message.data);
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



        ReceiveVideoChunk_Base64(VideoChunk_Base64) {
            // console.log('old videochunks:', this.VideoChunks);
            this.VideoChunks.push(VideoChunk_Base64);
            // console.log('new videochunks:', this.VideoChunks);
            console.log('is videochunk will change?', this.VideoChunk === VideoChunk_Base64);
            this.VideoChunk = VideoChunk_Base64;
            // console.log('VideoChunk change in room chat:', this.VideoChunk);
            // 直接向子组件LiveStream传递VideoChunks得了
            // this.$emit('update:VideoChunks', this.VideoChunks);

        },
        // handleReceivedVideoChunk_Base64(VideoChunk_Base64) {
        //     if (!this.sourceBuffer || this.mediaSource.readyState !== 'open') {
        //         console.log('MediaSource not ready or SourceBuffer not initialized.');
        //         return;
        //     }
        //     // 创建一个新的video对象用来播放视频
        //     this.videoElement.autoplay = true;
        //     this.videoElement.muted = true;
        //     // 对VideoChunk_Base64进行处理，在videoElement中播放
        //     const videoBlob = this.base64ToBlob(VideoChunk_Base64, 'video/webm');
        //     // 将视频块添加到SourceBuffer
        //     const fileReader = new FileReader();
        //     fileReader.onload = () => {
        //         this.VideoChunks.push(fileReader.result);
        //         this.sourceBuffer.push(new Uint8Array(fileReader.result));
        //     };
        //     fileReader.readAsArrayBuffer(videoBlob);
        //     // 播放视频
        //     this.videoElement.src = URL.createObjectURL(videoBlob);
        //     this.videoElement.play();


        // },
        // base64ToBlob(base64, mimeType) {
        //     const bytes = atob(base64);
        //     const len = bytes.length;
        //     const buffer = new ArrayBuffer(len);
        //     const view = new Uint8Array(buffer);
        //     for (let i = 0; i < len; i++) {
        //         view[i] = bytes.charCodeAt(i);
        //     }
        //     return new Blob([buffer], { type: mimeType });
        // },



        // 接受来自 ChatInput 组件的消息并发送
        sendMessageToParent(type, message) {
            // 接收来自子组件的消息并处理
            // console.log('Sending message to parent:', message);
            // 在这里可以进行进一步的处理，比如发送给服务器等操作
            // this.sendMessageToServer(message);
            console.log('type:', type);
            try {
                this.ws.send({
                    room_id: this.roomID,
                    type: type,
                    content: message
                });
                console.log('send send!!!!!');

            } catch (error) {
                console.error('Failed to send message:', error);
            }

        },

        // 用于获取历史消息
        async getHistoryMessages(lastMessageId, limit) {
            try {
                const response = await axios.post(
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
                { direction: 'bottom', target: this.$refs.scrollbar.$el.querySelector('.el-scrollbar__wrap') }
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

