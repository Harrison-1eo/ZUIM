<template>
    <div class="friend-container">
        <div class="friend-list">
            <h2>好友列表</h2>
            <el-menu class="el-menu-vertical-demo">
                <el-menu-item v-for="friend in friends" :key="friend.ID" :index="friend.ID"
                    @click="fetchFriendInfo(friend.ID)">
                    <i class="el-icon-user-solid"></i>
                    {{ friend.username }}
                </el-menu-item>
            </el-menu>

        </div>

        <div class="friend-details">
            <div class="friend-header">
                <p class="friend-title"> {{ activeFriend === null ? '请选择好友' : activeFriend.username }} </p>
                <el-icon v-if="activeFriendId" @click="drawer = true" class="more-icon">
                    <More />
                </el-icon>
            </div>
            <FriendChat v-if="activeFriendId" :friendID="activeFriendId" :friendName="activeFriend.username" />
            <FriendDrawer v-if="activeFriendId" v-model="drawer" v-model:ifFetch="ifFetch" :friendID="activateFriendId"
                :friend="activeFriend" />
        </div>
    </div>

</template>

<script>
import axios from "@/axios-config";
import FriendDrawer from "@/components/IM/OtherFunc/FriendDrawer.vue";
import FriendChat from "@/components/IM/OtherFunc/FriendChat.vue";
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
                const response = await axios.get('/api/user/friends');
                this.friends = response.data.data;
                console.log(this.friends);
            } catch (error) {
                console.error('Failed to fetch friends', error);
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
    display: flex;
    height: 100%;
    overflow-y: hidden;
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
