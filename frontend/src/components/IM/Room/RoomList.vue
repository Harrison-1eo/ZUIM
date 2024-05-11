<template>
    <div class="chatroom-container">
        <div class="chat-list">
            <el-button type="primary" class="new-chat-button" @click="createRoomBox">创建新房间</el-button>
            <el-menu class="el-menu-vertical-demo">
                <el-menu-item v-for="room in rooms" :key="room.ID" :index="room.ID" >
                    {{ room.name }}
                    <!-- 一个按钮，点击后获取该房间的信息 -->
                    <el-button type="text" @click="fetchRoomInfo(room.ID)">进入</el-button>
                    <!-- 一个按钮，删除该房间的按钮 -->
                    <el-button type="text" @click="deleteRoom(room.ID)">删除</el-button>
                </el-menu-item>
            </el-menu>
        </div>
        <div class="chat-details">
            <div class="chat-header">
              <p class="room-title"> {{ activeRoom===null ? '请选择聊天室' : activeRoom.name }} </p>
              <el-icon v-if="activeRoomId" @click="drawer=true" class="more-icon"><More /></el-icon>
            </div>
            <RoomChat v-if="activeRoomId" :roomID="activeRoomId" />
            <RoomDrawer v-if="activeRoomId" v-model="drawer" :roomID="activeRoomId" :room="activeRoom"/>
        </div>
    </div>
</template>

<script>
import axios from '@/axios-config';
import RoomChat from "@/components/IM/Room/RoomChat.vue";
import RoomDrawer from "@/components/IM/Room/RoomDrawer.vue";
import { ElMessage, ElMessageBox } from "element-plus";

export default {
    data() {
        return {
            rooms: [], // 存储聊天室列表
            activeRoom: null, // 当前激活的聊天室详情
            newMessage: '', // 绑定输入框的内容
            activeRoomId: null, // 当前激活的聊天室ID
            drawer: false
        };
    },
    components: {
        RoomChat,
        RoomDrawer
    },
    created() {
        this.fetchRooms(); // 获取聊天室列表
    },
    methods: {
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
                const response = await axios.post('http://localhost:8000/api/room/create',
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
        deleteRoomById(id) {
            // 执行删除房间
            axios.delete('http://localhost:8000/api/room/delete?room_id=' + id,
                { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } });
            // 删除成功后刷新房间列表


            // 示例：在控制台输出删除的房间ID
            console.log('删除房间，房间ID:', id);
        },
        async deleteRoom(id) {
            try {
                const response = await axios.delete('http://localhost:8000/api/room/delete?room_id=' + id,
                { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } });
                console.log('Delete room:', response.data);
                if (response.data.code !== 0) {
                  alert('删除失败');
                    return false;
                    
                }
                await this.fetchRooms(); // 删除成功后刷新房间列表
            } catch (error) {
                console.error('Failed to delete room:', error);
                return false;
            }
        },
        async fetchRooms() {
            try {
                // const response = await axios.get(
                //     'http://localhost:5000/api/room/list',
                //     { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } }
                // );
                const response = await axios.get(
                    'http://localhost:8000/api/room/list',
                );
                this.rooms = response.data.data; // 假设返回的数据是聊天室数组
                console.log('Rooms:', this.rooms);
            } catch (error) {
                console.error('Failed to fetch rooms:', error);
            }
        },
        async fetchRoomInfo(id) {
            try {
                // const response = await axios.get('http://localhost:8000/api/room/members?room_id=' + id,
                //     { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` } });
                this.activeRoom = this.rooms.find(room => room.ID === id); // 假设返回的数据是聊天室详情
                this.activeRoomId = id; // 更新当前激活的聊天室ID
            } catch (error) {
                console.error('Failed to fetch room info:', error);
            }
        },


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
</style>
