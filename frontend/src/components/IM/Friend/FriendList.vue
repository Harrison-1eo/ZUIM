<template>
    <div class="friend-container">
        <div class="chat-list">
            <h2 style="padding-left: 20px">终端列表</h2>
            <el-menu class="el-menu-rooms">
                <el-scrollbar wrap-class="scrollbar-wrapper">
                    <el-menu-item
                        class="el-menu-item-button"
                        v-for="friend in friends"
                        :key="friend.ID"
                        :index="friend.ID"
                        @click="fetchFriendInfo(friend.ID)">
<!--                        <i class="el-icon-user-solid"></i>-->
                        {{ friend.username }}
                    </el-menu-item>
                </el-scrollbar>
            </el-menu>
            
        </div>

        <div class="friend-details">
            <div class="friend-header">
                <p class="friend-title"> {{ activeFriend === null ? '请选择终端' : activeFriend.username }} </p>
                <el-icon v-if="activeFriendId" @click="drawer = true" class="more-icon">
                    <More />
                </el-icon>
            </div>
            <FriendChat v-if="activeFriendId" :friendID="activeFriendId" :friendName="activeFriend.username" />
            <FriendDrawer v-if="activeFriendId" v-model="drawer" v-model:ifFetch="ifFetch" :friendID="activateFriendId" :friend="activeFriend" />
        </div>
    </div>
    
</template>

<script>
import axios_config from "@/utils/axios-config";
import FriendDrawer from "@/components/IM/Friend/FriendDrawer.vue";
import FriendChat from "@/components/IM/Friend/FriendChat.vue";
export default {
    data() {
        return {
            friends: [],
            activateFriendId: null,
            drawer: false,
            ifFetch: false,
            activeFriend: null,


        };
    },
    components: {
        FriendDrawer,
        FriendChat
    },
    watch: {
        ifFetch() {
            this.selectNoFriend();
            this.drawer = false;
            this.ifFetch = false;
        }
    },
    created() {
        this.fetchFriends();
    },
    methods: {
        selectNoFriend() {
            this.activeFriend = null;
            this.activeFriendId = null;
            this.fetchFriends();
        },
        createdRoomBoxwithChosenFriend() {
            const description = "chat with" + this.activeFriend.username;
            this.$emit('createRoom', description);
        },
        async fetchFriends() {
            try {
                const response = await axios_config.get('/api/user/friends');
                this.friends = response.data.data;
                // 排除自己
                console.log('localStorage.getItem(userId): ', localStorage.getItem('userId'));
                this.friends = this.friends.filter(friend => friend.ID.toString() !== localStorage.getItem('userId').toString());
                console.log('friends: ', this.friends);
            } catch (error) {
                console.error('Failed to fetch friends' ,error);
            }
        },
        fetchFriendInfo(id) {
            this.activeFriendId = id;
            console.log('fetchFriendInfo', id);
            this.activeFriend = this.friends.find(friend => friend.ID === id);

        },

    }
};
</script>

<style>
.friend-container {
    margin: 0;
    padding: 0;
    height: 100%;
    overflow-y: hidden;
}

.chat-list {
    padding: 20px;
    width: 200px;
    display: flex;
    flex-direction: column;
    border-right: 1px solid #ffffff;
    background: linear-gradient(180deg,#f0edf7 .03%,#ebeaf5 32.19%,#e8e8f3 68.86%,#e4eaf7 99.12%);
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




.friend-list {
    padding: 20px;
    width: 200px;
    display: flex;

    flex-direction: column;
    border-right: 1px solid #ebeef5;
}

.friend-details {
    flex: 1;
    display: flex;
    flex-direction: column;
    height: 100vh;
}
</style>
