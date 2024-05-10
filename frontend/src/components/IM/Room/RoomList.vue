<template>
    <div class="chatroom-container">
        <div class="chat-list">
            <el-button type="primary" class="new-chat-button" @click="createRoomBox">创建新房间</el-button>
            <div class="menu-wrapper">
                <el-menu :default-active="activeRoomId" class="el-menu-vertical-demo" @select="fetchRoomInfo">
                    <el-menu-item v-for="room in rooms" :key="room.ID" :index="room.ID">
                        {{ room.name }}
                    </el-menu-item>
                </el-menu>

                <!-- <el-menu :default-active="activeRoomId" class="el-menu-vertical-demo" @select="fetchRoomInfo">
        <el-menu-item v-for="room in rooms" :key="room.ID" :index="room.ID">
          {{ room.name }}
        </el-menu-item>
      </el-menu> -->

            </div>
            <!-- <div class="chat-details">
      <div class="chat-header">
        {{ activeRoom || '选择一个聊天室' }}
      </div>
      <RoomChat v-if="activeRoomId" :roomID="activeRoomId"/>
    </div> -->
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import RoomChat from "@/components/IM/Room/RoomChat.vue";
import { ElMessage, ElMessageBox } from "element-plus";

export default {
    data() {
        return {
            rooms: [], // 存储聊天室列表
            activeRoom: null, // 当前激活的聊天室详情
            newMessage: '', // 绑定输入框的内容
            activeRoomId: null // 当前激活的聊天室ID
        };
    },
    components: {
        RoomChat
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
        async fetchRooms() {
            try {
                // const response = await axios.get(
                //     'http://localhost:5000/api/room/list',
                //     { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } }
                // );
                const response = await axios.get(
                    'http://localhost:8000/api/room/list',
                    // authorization: Bearer+token
                    { headers: { Authorization: `Bearer ${localStorage.getItem("token")}` } }
                    // { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` }}
                );
                this.rooms = response.data.data; // 假设返回的数据是聊天室数组
                console.log('Rooms:', this.rooms);
            } catch (error) {
                console.error('Failed to fetch rooms:', error);
            }
        },
        async fetchRoomInfo(id) {
            try {
                const response = await axios.get('http://localhost:8000/api/room/members?room_id=' + id,
                    // { headers: { Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzOTg4MjYsInVzZXJfaWQiOjJ9.u1coQZoetjzqKRVjQqWdLWC1Jr5ymGoqcUCqLc0eLFY` }});
                    { headers: { Authorization: `Bearer ${localStorage.getItem("token")}` } });
                this.activeRoom = response.data.data[0].username; // 设置当前激活的聊天室详情

            } catch (error) {
                console.error('Failed to fetch room info:', error);
            }
        }
    }
};
</script>

<style scoped>
/* CSS样式与之前相同 */
.chatroom-container {
    height: 50%;
    flex-direction: column;
    align-items: center;
}

/* .chat-list {
    display: grid;
    grid-template-columns: repeat(5, 1fr); 
    gap: 10px; 
    justify-items: center; 
} */

.menu-wrapper {
  overflow: hidden; /* 清除浮动 */
}

.el-menu-item {
  width: 20%; /* 每个元素占据一行的 20% 宽度 */
  float: left; /* 设置为浮动布局 */
  box-sizing: border-box;
}
/* .chat-details {
    flex: 1;
    display: flex;
    flex-direction: column;
} */

.new-chat-button {
    margin-bottom: 20px; /* 按钮下方留出一些空间 */
}

.chat-header {
    padding: 10px;
    background-color: #f5f7fa;
    text-align: center;
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
