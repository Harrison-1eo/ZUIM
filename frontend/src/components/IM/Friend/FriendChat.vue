<template>
    <div class="friend-chat-container">
        <!-- 显示这个friend的基本资料    -->
        <div class="friend-chat-info-container">
            <div class="friend-avatar">
                <el-avatar :src="avatarUrl()" :size="80" class="user-avatar" shape="square"></el-avatar>
            </div>
            <div class="friend-username">
                {{ userInfo.username || 'No username provided' }}
            </div>
            <div class="friend-details">
                <p>邮箱: {{ userInfo.email || 'No email provided' }}</p>
                <p>地区: {{ userInfo.country || '中国' }}</p>
            </div>
            <div class="friend-more">
                <el-icon @click="drawer = true" style="cursor: pointer">
                    <More />
                </el-icon>
                <FriendDrawer v-model="drawer" v-model:ifFetch="ifFetch" :friendID="activeFriend.ID" :friend="activeFriend" />
            </div>
        </div>

        <el-divider style="margin: 20px 0"/>

        <!-- 和这个friend共同的chatroom -->
        <div class="friend-chat-list-container">

            <el-menu class="el-menu-rooms friend-chat-list-menu"  v-if="roomsBetween.length>0">
                <h2 style="padding-left: 20px;margin: 0 0 10px;">共同通讯室</h2>
                <el-scrollbar wrap-class="scrollbar-wrapper">
                    <el-menu-item
                            class="el-menu-item-button"
                            style="font-size: large"
                            v-for="room in roomsBetween"
                            :key="room.ID"
                            :index="room.ID.toString()"
                            @click="fetchRoomInfo(room.ID)">
                        {{ room.name }}
                    </el-menu-item>
                </el-scrollbar>
            </el-menu>

            <el-empty v-else description="暂无共同通讯室"></el-empty>
        </div>
    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import {ElAvatar, ElDivider, ElMessage} from "element-plus";
import {backendBaseUrl} from "@/utils/base-url-setting";
import FriendDrawer from "@/components/IM/Friend/FriendDrawer.vue";

export default {
    name: 'FriendChat',
    components: {FriendDrawer, ElAvatar, ElDivider},
    props: {
        activeFriend: {
            type: Object,
            required: true,
        },
    },
    data() {
        return {
            roomsBetween: [], // 存储聊天室列表，两者共同的聊天室
            userInfo: {},
            ifFetch: false,
            drawer: false,
        };
    },
    watch: {
        activeFriend() {
            console.log('friendInfo change', this.activeFriend);
            this.fetchRoomsWithFriend();
            this.fetchFriendInfo();
        },
        ifFetch() {
            this.drawer = false;
            this.ifFetch = false;
        }
    },
    created() {
        this.fetchRoomsWithFriend(); // 获取聊天室列表
        this.fetchFriendInfo();
    },
    methods: {
        avatarUrl() {
            return backendBaseUrl + this.userInfo.avatar;
        },
        async fetchRoomsWithFriend() {
            try {
                const response = await axios_config.get('/api/user/common_rooms', {
                    params: {
                        friend_id: this.activeFriend.ID
                    }
                })
                if (response.data.code === 0) {
                    console.log('response.data.data', response.data.data)
                    this.roomsBetween = response.data.data;
                    console.log('roomsBetween', this.roomsBetween);
                } else {
                    ElMessage.error('获取聊天室列表失败: ', response.data.msg);
                }
            } catch (error) {
                console.error('Failed to fetch rooms', error);
            }
        },
        async fetchFriendInfo() {
            try {
                const response = await axios_config.get('/api/user/info', {
                    params: {
                        user_id: this.activeFriend.ID
                    }
                });
                if (response.data.code === 0) {
                    this.userInfo = response.data.data;
                    console.log('userInfo', this.userInfo);
                } else {
                    ElMessage.error('获取用户信息失败: ', response.data.msg);
                }
            } catch (error) {
                console.error('Failed to fetch friend information:', error);
            }
        },
    }
};
</script>

<style>
.friend-chat-container {
    width: calc(100% - 40px);
    height: calc(100% - 60px - 40px);
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    padding: 20px;
    margin: 0;
}

.friend-chat-info-container {
    display: grid;
    grid-template-columns: repeat(20, 1fr);
    grid-template-rows: repeat(5, 30px);
    grid-column-gap: 0;
    grid-row-gap: 0;
}

.friend-avatar {
    grid-area: 2 / 2 / 3 / 3;
}

.friend-username {
    grid-area: 2 / 4 / 4 / 10;
    font-size: 24px;
    font-weight: bold;
    color: #000000;
}

.friend-details {
    margin-top: 10px;
    grid-area: 3 / 4 / 6 / 19;
    color: #8b8a8a;
}

.friend-details p {
    margin: 0 0 5px;
}

.friend-more {
    grid-area: 2 / 20 / 3 / 21;
}

.friend-chat-list-container {
    width: 100%;
    margin: 20px;
}

.friend-chat-list-menu {
    width: 100%;
}


</style>
