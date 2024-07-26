<template>
    <el-drawer show-close drawe>
        <template #header>
            <h2>{{friend.username}}</h2>
        </template>

        <el-card style="max-width: 480px">
            <el-button type="primary" class="changedescription" @click="createRoomChatWithThisGuyBox(friend.ID, friend.username)">创建聊天室</el-button>
        </el-card>
        <template #footer>
            <el-button type="danger">删除终端</el-button>
        </template>
    </el-drawer>
</template>

<script>
import axios_config from "@/utils/axios-config";
import {ElMessage, ElMessageBox} from "element-plus";
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
        createRoomChatWithThisGuyBox(ID, username) {
            ElMessageBox.prompt(
                    '请输入新的房间名称',
                    '创建与' + username + '的聊天室',
                    {
                        inputValue: 'chat with ' + username,
                        confirmButtonText: '创建',
                        cancelButtonText: '取消',
                        inputPattern: /\S/,
                        inputErrorMessage: '房间名称不能为空'
                    }
            ).then(({ value }) => {
                if (!value) {
                    ElMessage.error('房间名称不能为空');
                    return;
                }
                this.createRoomChatWithThisGuy(ID, username, value);
            }).catch(() => {
                ElMessage.info('创建');
            });
        },
        async createRoomChatWithThisGuy(ID, username, roomname) {
            // ID不能是我自己的ID
            if (ID.toString() === localStorage.getItem('userId')) {
                ElMessage.error('不能和自己聊天');
                return false;
            }
            // 创建我和他之间的聊天室，跳转到/im/chat界面
            const description = "自动创建" + roomname
            try {
                const response = await axios_config.post('/api/room/create', {
                    'name': roomname,
                    'description': description,
                });
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
                this.$store.commit('setActiveRoom', response.data.data);
                this.$router.push({ path: '/im/chat' });

                ElMessage.success('创建聊天室成功, 跳转到聊天室界面');
                // await this.fetchRoomUsers();

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