<template>
<!-- frienddrawer的功能是点击展开，有功能如下：创建我和他之间的聊天室，即必须得调用RoomList.vue中的createRoomBox和RoomDrawer中的addUserBox函数 -->
<!-- createRoomBox函数接口： -->
<el-drawer show-close drawe>
    <template #header>
        <h2>{{friend.name}}</h2>
    </template>

    <el-card style="max-width: 480px">
        <template #header>
            <div class="card-description">
                <span>创建聊天室</span>
            </div>
        </template>
        <el-button type="primary" class="changedescription" @click="createRoomBox">创建聊天室</el-button>
    </el-card>
</el-drawer>
</template>

<script>
import axios from "@/axios-config";
import { ElMessage, ElMessageBox } from "element-plus";

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
        friend() {
            this.fetchRoomUsers();
        }
    },
    created() {
        this.fetchRoomUsers();
    },
    methods: {
        createRoomBox() {
            const description = 'chat with ' + this.friend.name;
            this.$emit('createRoomBox', description);
        }
    }
};
</script>