<template>
    <el-drawer show-close drawe>
        <template #header>
            <h2>{{friend.username}}</h2>
        </template>

        <el-card style="max-width: 480px">
            <el-button type="primary" class="changedescription" @click="createRoomChatWithThisGuy(friend.ID, friend.username)">创建聊天室</el-button>
        </el-card>
        <template #footer>
            <el-button type="danger">删除终端</el-button>
        </template>
    </el-drawer>
</template>

<script>
import axios_config from "@/utils/axios-config";
import { ElMessage } from "element-plus";
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
    },
    created() {
    },
    computed: {
        activeRoom() {
            return store.state.activeRoom; // 从 Vuex store 中获取激活的房间ID
        }
    },
    methods: {
        async createRoomChatWithThisGuy(ID, username) {
            // ID不能是我自己的ID
            if (ID.toString() === localStorage.getItem('userId')) {
                ElMessage.error('不能和自己聊天');
                return false;
            }
            // 创建我和他之间的聊天室，跳转到/im/chat界面
            const description = "chat with " + username;
            const roomname = "chat with " + ID;
            try {
                const response = await axios_config.post('/api/room/create', {
                    'name': roomname,
                    'description': description,
                },
                    { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } }
                );
                if (response.data.code !== 0) {
                    ElMessage.error(response.data.msg);
                    return false;
                }
                this.roomID = response.data.data.ID;
                // 把他加入到聊天室中
                const response2 = await axios_config.post('/api/room/add_user', {
                    'room_id': response.data.data.ID,
                    'user_name': username,
                    'role': 'member'
                });

                // 创建成功，跳转到聊天室界面
                this.$router.push({ path: '/im/chat' });
                await this.fetchRoomUsers();
                store.commit('setActiveRoom', response.data.data);

            } catch (error) {
                ElMessage.error('创建聊天室失败');
            }

        },

        async fetchRoomUsers() {
            // 获取房间成员列表
            try {
                const response = await axios_config.get('/api/room/members?room_id=' + this.roomID);

                if (response.data.code === 0) {
                    this.roomUsers = response.data.data;
                } else {
                    ElMessage.error(response.data.msg);
                }
            } catch (error) {
                ElMessage.error('获取房间成员列表失败');
            }


        },
        async getHistoryMessages(lastMessageId, limit) {
            try {
                const response = await axios_config.post(
                    '/api/message/list',
                    {
                        "room_id": this.roomID,
                        "last_message_id": lastMessageId,
                        "limit": limit
                    });
                if (response.data.data.length === 0) {
                    ElMessage.info('没有更多历史消息了')
                    return;
                }
                // response.data.data 获取的消息列表ID是从大到小的，需要将其反转
                this.messages = [...response.data.data.reverse(), ...this.messages];
            } catch (error) {
                ElMessage.error('Failed to fetch messages:', error);
            }
        },
    }
};
</script>