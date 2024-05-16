<!-- /src/components/IM/Room/RoomList.vue -->

<template>
    <div class="chatroom-container">
        <div class="chat-list">
            <el-button type="primary" class="new-chat-button" @click="createRoomBox">创建新房间</el-button>
            <el-menu class="el-menu-vertical-demo">
                <el-menu-item v-for="room in rooms" :key="room.ID" :index="room.ID" @click="fetchRoomInfo(room.ID)">
                    {{ room.name }}
                </el-menu-item>
            </el-menu>
        </div>
        <div class="chat-details">
            <div class="chat-header">
                <p class="room-title"> {{ activeRoom===null ? '请选择聊天室' : activeRoom.name }} </p>
                <el-icon v-if="activeRoomId" @click="drawer=true" class="more-icon">
                    <More />
                </el-icon>
            </div>
            <RoomChat v-if="activeRoomId" :roomID="activeRoomId" />
            <RoomDrawer v-if="activeRoomId" v-model="drawer" v-model:ifFetch="ifFetch" :roomID="activeRoomId" :room="activeRoom" />
        </div>
    </div>
</template>

<script>
import axios from '@/axios-config';
import RoomChat from "@/components/IM/Room/RoomChat.vue";
import RoomDrawer from "@/components/IM/Room/RoomDrawer.vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { mapGetters } from 'vuex';
import store from '@/store/store';

export default {
    data() {
        return {
            rooms: [], // 存储聊天室列表
            activeRoom: null, // 当前激活的聊天室详情
            activeRoomId: null, // 当前激活的聊天室ID
            drawer: false,
            ifFetch: false
        };
    },
    components: {
        RoomChat,
        RoomDrawer
    },
    computed: {
        // ...mapGetters(['activeRoomId']),
        activeRoomId() {
            return store.state.activeRoom.ID
        },
        activeRoom() {
            return store.state.activeRoom; // 从 Vuex store 中获取激活的房间ID
        }
    },
    watch: {
        ifFetch() {
            this.selectNoRoom();
            this.drawer = false;
            this.ifFetch = false;
        },
        activeRoom(newValue, oldValue) {
            // 在这里监听 activeRoomId 的变化，并做出相应的响应
            this.activeRoom = newValue;
            console.log('activeRoom changed:', newValue);
        }

    },
    created() {
        this.fetchRooms(); // 获取聊天室列表
    },
    methods: {
        selectNoRoom() {
            this.activeRoom = null;
            this.activeRoomId = null;
            this.fetchRooms();
        },
        createRoomBox() {
            ElMessageBox.prompt('请输入房间名称', '创建新房间').then(({ value }) => {
                if (!value) {
                    ElMessage.error('房间名称不能为空');
                    return;
                }
                // 创建新房间
                if (this.createRoom(value)) {
                    ElMessage.success('创建成功');
                } else {
                    ElMessage.error('创建失败');
                }
            }).catch(() => {
                ElMessage.info('取消创建');
            });
        },
        async createRoom(name) {
            try {
                const response = await axios.post('/api/room/create',
                    {
                        'name': name,
                        'description': ''
                    },
                    { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } }
                );
                console.log('Create room:', response.data);
                if (response.data.code !== 0) {
                    return false;
                }
                await this.fetchRooms(); // 创建成功后刷新房间列表
            } catch (error) {
                console.error('Failed to create room:', error);
                return false;
            }
            return true;
        },
        async fetchRooms() {
            try {
                const response = await axios.get(
                    '/api/room/list',
                );
                this.rooms = response.data.data; // 假设返回的数据是聊天室数组
                console.log('Rooms:', this.rooms);
            } catch (error) {
                console.error('Failed to fetch rooms:', error);
            }
        },
        async fetchRoomInfo(id) {
            this.activeRoom = this.rooms.find(room => room.ID === id); // 假设返回的数据是聊天室详情
            this.activeRoomId = id; // 更新当前激活的聊天室ID
            store.commit('setActiveRoom', this.activeRoom);
            console.log('Active room in fetchroominfo func:', this.activeRoom);
        },
        // async fetchRoomInfo(id) {
        //     // 通过 Vuex store 分发 fetchRoomInfo 方法
        //     await this.$store.dispatch('fetchRoomInfo', id);
        // }


    }
};
</script>

<style scoped>
/* CSS样式与之前相同 */
.chatroom-container {
    display: flex;
    height: 100%;
    overflow-y: hidden;
}

.chat-list {
    width: 200px;
    display: flex;

    flex-direction: column;
    border-right: 1px solid #ebeef5;
}

.chat-details {
    flex: 1;
    display: flex;
    flex-direction: column;
    height: 100vh;
}

.chat-header {
    padding: 10px;
    background-color: #f5f7fa;
    text-align: center;
    height: 2em;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.new-chat-button {
    margin: 10px;
    height: 40px;
}

.el-menu-vertical-demo {
    border: none;
    overflow-y: auto;
}

.more-icon {
    cursor: pointer;
}
</style>
