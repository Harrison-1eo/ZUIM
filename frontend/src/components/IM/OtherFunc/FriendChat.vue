<template>
    <div class="friend-container">
        <!-- 和这个friend共同的chatroom -->
        <div class="friend-chat">
            <el-card style="max-width: 480px">
                <template #header>
                    <div class="card-description">
                        <!-- <span>和{{friendID}}的聊天</span> -->
                        <span>和{{ friendName }}的聊天</span>
                    </div>
                </template>
                <el-scrollbar style="height: 300px;">
                    <el-menu class="el-menu-vertical-demo">
                        <el-menu-item v-for="room in rooms" :key="room.ID" :index="room.ID" @click="fetchRoomInfo(room.ID)">
                            {{ room.name }}
                        </el-menu-item>
                    </el-menu>
                </el-scrollbar>
            </el-card>
        </div>
        <div class="friend-details">
            <div class="friend-header">
                <p class="friend-title"> {{ activeRoom===null ? '请选择聊天室' : activeRoom.name }} </p>
                <el-icon v-if="activeRoomId" @click="drawer=true" class="more-icon"><More /></el-icon>
            </div>
            <!-- <RoomDrawer v-if="activeRoomId" v-model="drawer" v-model:ifFetch="ifFetch" :roomID="activeRoomId" :room="activeRoom"/> -->
        </div>
    </div>
</template>

<script>
import axios from '@/axios-config';
import RoomDrawer from "@/components/IM/Room/RoomDrawer.vue";
import { ElMessage, ElMessageBox } from "element-plus";

export default {
    name: 'FriendChat',
    props: {
        friendID: {
            type: Number,
            required: true
        },
        friendName: {
            type: String,
            required: true
        }
    },
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
        RoomDrawer
    },
    watch: {
        ifFetch() {
            this.selectNoRoom();
            this.drawer = false;
            this.ifFetch = false;
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
        fetchRooms() {
            axios.get('/api/room/list').then(res => {
                this.rooms = res.data;
            }).catch(err => {
                ElMessage.error('获取聊天室列表失败');
            });
        },
        fetchRoomInfo(roomID) {
            axios.get('/api/room/info', {
                params: {
                    roomID: roomID
                }
            }).then(res => {
                this.activeRoom = res.data;
                this.activeRoomId = roomID;
            }).catch(err => {
                ElMessage.error('获取聊天室信息失败');
            });
        }
    }
};
</script>

<style>
.friend-container {
    display: flex;
    justify-content: space-between;
    padding: 20px;
}
.friend-chat {
    width: 30%;
}
.friend-details {
    width: 68%;
}
.friend-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid #f4f4f4;
}
.friend-title {
    font-size: 20px;
}
.more-icon {
    cursor: pointer;
}
.card-description {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>
