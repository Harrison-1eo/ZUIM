<!-- /src/components/IM/Room/RoomList.vue -->

<template>
    <div class="chatroom-container">
        <div class="chat-list">
            <h2 style="padding-left: 20px">通讯列表</h2>

            <el-menu class="el-menu-rooms">
                <el-scrollbar wrap-class="scrollbar-wrapper">
                    <el-menu-item
                            :default-active="activeRoomId"
                            class="el-menu-item-button"
                            v-for="room in rooms"
                            :key="room.ID"
                            :index="room.ID"
                            @click="activeOneRoom(room.ID)">
                        {{ room.name }}
                    </el-menu-item>
                </el-scrollbar>
            </el-menu>
            <el-button
                    type="primary"
                    class="new-chat-button"
                    @click="createRoomBox">
                创建新通讯
            </el-button>
        </div>
        <div class="chat-details">
            <RoomChat v-if="activeRoomId" :roomID="activeRoomId" v-model:ifUpdateRoomList2Chat="ifUpdateRoomList2Chat"/>
        </div>
    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import RoomChat from "@/components/IM/Room/RoomChat.vue";
import {ElMessage, ElMessageBox} from "element-plus";
// import store from '@/store/store';

export default {
    data() {
        return {
            rooms: [], // 存储聊天室列表
            activeRoom: null, // 当前激活的聊天室详情
            activeRoomId: null, // 当前激活的聊天室ID
            ifUpdateRoomList2Chat: false,
        };
    },
    components: {
        RoomChat,
    },
    watch: {
        ifFetch() {
            this.selectNoRoom();
            this.drawer = false;
            this.ifFetch = false;
        },
        activeRoom(newValue, oldValue) {
            this.activeRoom = newValue;
            console.log('activeRoom changed:', newValue, oldValue);
        },
        ifUpdateRoomList2Chat() {
            this.fetchRooms();
        },
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
            ElMessageBox.prompt('请输入通讯名称', '创建新通讯').then(({value}) => {
                if (!value) {
                    ElMessage.error('通讯名称不能为空');
                    return;
                }
                // 创建新通讯
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
                const response = await axios_config.post('/api/room/create',
                        {
                            'name': name,
                            'description': ''
                        },
                );
                console.log('Create room:', response.data);
                if (response.data.code !== 0) {
                    return false;
                }
                await this.fetchRooms(); // 创建成功后刷新通讯列表
                // 创建成功后激活新通讯：无法实现，element-plus的el-menu-item无法通过代码激活
                // await this.activeOneRoom(response.data.data.ID);
            } catch (error) {
                console.error('Failed to create room:', error);
                return false;
            }
            return true;
        },
        async fetchRooms() {
            try {
                const response = await axios_config.get(
                        '/api/room/list',
                );
                this.rooms = response.data.data; // 假设返回的数据是聊天室数组
                if (this.ifUpdateRoomList2Chat) {
                    this.ifUpdateRoomList2Chat = false;
                    // 如果当前选中的聊天室仍在列表中，则重新获取聊天室信息，并选定当前聊天室不变
                    if (this.rooms.find(room => room.ID === this.activeRoomId)) {
                        console.log('activeRoomId 没变:', this.activeRoomId, this.rooms);
                        this.activeOneRoom(this.activeRoomId);
                    } else {
                        console.log('activeRoomId 变了:', this.activeRoomId);
                        this.activeRoom = null;
                        this.activeRoomId = null;
                    }
                }
                console.log('Rooms:', this.rooms);
            } catch (error) {
                console.error('Failed to fetch rooms:', error);
            }
        },
        async activeOneRoom(id) {
            this.activeRoom = this.rooms.find(room => room.ID === id);      // 假设返回的数据是聊天室详情
            this.activeRoomId = id;                                         // 更新当前激活的聊天室ID
            // store.commit('setActiveRoom', this.activeRoom);
            // console.log('Active room in fetchroominfo func:', this.activeRoom);
        },
    }
};
</script>

<style scoped>
.chatroom-container {
    display: flex;
    height: 100%;
    overflow-y: hidden;
}

.chat-list {
    padding: 20px;
    width: 200px;
    display: flex;
    flex-direction: column;
    border-right: 1px solid #ffffff;
    background: linear-gradient(180deg, #f0edf7 .03%, #ebeaf5 32.19%, #e8e8f3 68.86%, #e4eaf7 99.12%);
}

.chat-details {
    flex: 1;
    display: flex;
    flex-direction: column;
    height: 100vh;
}

.new-chat-button {
    margin: 20px;
    height: 40px;
    background: linear-gradient(135deg, #897dfb, #3471f6);
    border: none;
}

.el-menu-rooms {
    border: none;
    overflow-y: auto;
    background: none;
}

.el-menu-item-button {
    background: none;
    border-bottom: #333333 1px;
    border-radius: 10px;
}

.el-menu-item-button:hover {
    background: #ffffff;
}


</style>
