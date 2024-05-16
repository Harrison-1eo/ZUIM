<template>
    <!-- frienddrawer的功能是点击展开，有功能如下：创建我和他之间的聊天室，即必须得调用RoomList.vue中的createRoomBox和RoomDrawer中的addUserBox函数 -->
    <!-- createRoomBox函数接口： -->
    <el-drawer show-close drawe>
        <template #header>
            <h2>{{friend.username}}</h2>
            <h2>{{friend.ID}}</h2>
        </template>

        <el-card style="max-width: 480px">
            <!-- <template #header>
            <div class="card-description">
                <span>创建聊天室</span>
            </div>
        </template> -->
            <el-button type="primary" class="changedescription" @click="createRoomChatWithThisGuy(friend.ID, friend.username)">创建聊天室</el-button>

        </el-card>
    </el-drawer>
</template>

<script>
import axios from "@/axios-config";
import { descriptionProps, ElMessage, ElMessageBox } from "element-plus";
import store from "@/store/store";
export default {
    name: 'FriendDrawer',
    props: {
        friend: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            roomUsers: []
        };
    },
    watch: {
        friend() {
            // this.fetchRoomUsers();
        }
    },
    created() {
        // this.fetchRoomUsers();
    },
    computed: {
        activeRoom() {
            return store.state.activeRoom; // 从 Vuex store 中获取激活的房间ID
        }
    },
    methods: {
        // createRoomBox() {
        //     const description = 'chat with ' + this.friend.name;
        //     this.$emit('createRoomBox', description);
        // },
        async createRoomChatWithThisGuy(ID, username) {
            // ID不能是我自己的ID
            if (ID.toString() === localStorage.getItem('userId')) {
                ElMessage.error('不能和自己聊天');
                console.log('不能和自己聊天');
                return false;
            }
            // 创建我和他之间的聊天室，跳转到/im/chat界面
            const description = "chat with " + username;
            const roomname = "chat with " + ID;
            try {
                const response = await axios.post('/api/room/create', {
                    'name': roomname,
                    'description': description,
                },
                    { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } }
                );
                console.log('Create room:', response.data);
                if (response.data.code !== 0) {
                    console.error('创建聊天室失败', response.data.msg);
                    ElMessage.error(response.data.msg);
                    return false;
                }
                this.roomID = response.data.data.ID;
                // 把他加入到聊天室中
                const response2 = await axios.post('/api/room/add_user', {
                    'room_id': response.data.data.ID,
                    'user_name': username,
                    'role': 'member'
                });
                console.log('Add user:', response2.data);

                // 创建成功，跳转到聊天室界面
                this.$router.push({ path: '/im/chat' });
                await this.fetchRoomUsers();
                // await this.getHistoryMessages(0, 10);
                console.log('before commit');
                store.commit('setActiveRoom', response.data.data);
                console.log('after commit');


            } catch (error) {
                console.error('创建聊天室失败', error);
                ElMessage.error('创建聊天室失败');
            }

        },

        async fetchRoomUsers() {
            // 获取房间成员列表
            try {
                const response = await axios.get('/api/room/members?room_id=' + this.roomID);

                if (response.data.code === 0) {
                    this.roomUsers = response.data.data;
                } else {
                    ElMessage.error(response.data.msg);
                }
            } catch (error) {
                console.error('获取房间成员列表失败', error);
                ElMessage.error('获取房间成员列表失败');
            }


        },
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
    }
};
</script>