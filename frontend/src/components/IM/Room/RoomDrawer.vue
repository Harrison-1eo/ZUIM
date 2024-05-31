<!-- /src/components/IM/Room/RoomDrawer.vue -->

<template>
    <el-drawer show-close drawer size="450px">
        <template #header>
            <div class="room-drawer-header">
                <h2>{{ roomInfo.name }}</h2>
                <el-icon style="margin: 20px; cursor: pointer;" @click="updateRoomNameBox"><Edit /></el-icon>
            </div>
        </template>

        <el-card style="max-width: 480px">
            <template #header>
                <div class="card-description">
                    <span>房间描述</span>
                </div>
            </template>
            <p>{{ roomInfo.description }}</p>
            <template #footer>
                <el-button type="primary" class="changedescription" @click="updateRoomInfoBox">修改房间信息</el-button>
            </template>
        </el-card>

        <p></p>

        <el-card style="max-width: 480px">
            <template #header>
                <div class="card-header">
                    <span>房间成员</span>
                </div>
            </template>
            <!--      <p v-for="user in roomUsers" :key="user.id">{{ user.username }}-->
            <!--      </p>-->
            <el-row>
                <el-col v-for="user in roomUsers" :key="user.id" :span="12" style="display:flex; text-align: center; margin: 10px 0;">
                    <el-avatar :src="avatarUrl(user.avatar)" :size="40"></el-avatar>
                    <span style="margin: auto auto auto 10px;">{{ user.username }}</span>
                </el-col>
            </el-row>
            <template #footer>
                <el-button type="primary" @click="addUserBox">邀请成员</el-button>
            </template>

        </el-card>

        <template #footer>
            <el-button type="danger" @click="deleteRoom">退出房间</el-button>
        </template>

    </el-drawer>

</template>


<script>
import axios_config from "@/utils/axios-config";
import { ElMessage, ElMessageBox } from "element-plus";
import {backendBaseUrl} from "@/utils/base-url-setting";

export default {
    name: 'RoomDrawer',
    props: {
        roomID: {
            type: Number,
            required: true
        },
        roomInfo: {
            type: Object,
            required: true
        },
    },
    data() {
        return {
            roomUsers: [],
        }
    },
    watch: {
        roomID() {
            this.fetchRoomUsers();
        }
    },
    created() {
        this.fetchRoomUsers();
    },
    methods: {
        // ElMessageBox.prompt 的参数设置参考：https://blog.csdn.net/weixin_47872288/article/details/139044826
        updateRoomNameBox() {
            ElMessageBox.prompt(
                    '请输入新的房间名称',
                    '修改房间名称',
                    {
                        inputValue: this.roomInfo.name,
                        confirmButtonText: '修改',
                        cancelButtonText: '取消',
                        inputPattern: /\S/,
                        inputErrorMessage: '房间名称不能为空'
                    }
            ).then(({ value }) => {
                if (!value) {
                    ElMessage.error('房间名称不能为空');
                    return;
                }
                // 修改房间名称
                const newRoomInfo = this.roomInfo;
                newRoomInfo.name = value;
                this.updateRoomInfo(newRoomInfo);
            }).catch(() => {
                ElMessage.info('取消修改');
            });
        },
        updateRoomInfoBox() {
            ElMessageBox.prompt(
                    '请输入新的房间描述',
                    '修改房间信息',
                    {
                        inputValue: this.roomInfo.description,
                        confirmButtonText: '修改',
                        cancelButtonText: '取消',
                        inputType: 'textarea',
                        inputPattern: /\S/,
                        inputErrorMessage: '房间描述不能为空'
                    }
            ).then(({ value }) => {
                if (!value) {
                    ElMessage.error('房间描述不能为空');
                    return;
                }
                // 修改房间描述
                const newRoomInfo = this.roomInfo;
                newRoomInfo.description = value;
                this.updateRoomInfo(newRoomInfo);
            }).catch(() => {
                ElMessage.info('取消修改');
            });
        },
        addUserBox() {
            ElMessageBox.prompt(
                    '请输入邀请的用户名称',
                    '邀请朋友',
                    {
                        confirmButtonText: '邀请',
                        cancelButtonText: '取消',
                        inputType: 'textarea',
                        inputErrorMessage: '用户名称不能为空'
                    }
            ).then(({ value }) => {
                if (!value) {
                    ElMessage.error('用户名称不能为空');
                    return;
                }
                // 创建新房间
                this.addUser(value);
            }).catch(() => {
                ElMessage.info('取消邀请');
            });
        },
        async updateRoomInfo(newRoomInfo) {
            try {
                const response = await axios_config.post('/api/room/update',
                    {
                        'room_id': this.roomID,
                        'name': newRoomInfo.name,
                        'description': newRoomInfo.description
                    },
                );
                if (response.data.code !== 0) {
                    ElMessage.error(response.data.msg);
                    return false;
                }
                ElMessage.success('修改成功');
                this.$emit('update:roomInfo', newRoomInfo);
                this.$emit('update:ifUpdate', true);
            } catch (error) {
                console.error('Failed to update room info:', error);
                ElMessage.error('服务器错误');
                return false;
            }
            return true;
        },
        async addUser(name) {
            try {
                const response = await axios_config.post('/api/room/add_user',
                    {
                        'room_id': this.roomID,
                        'user_name': name,
                        'role': 'member'
                    },
                );
                if (response.data.code !== 0) {
                    ElMessage.error(response.data.msg);
                    return false;
                }
                await this.fetchRoomUsers();
            } catch (error) {
                console.error('Failed to add user:', error);
                ElMessage.error('服务器错误');
                return false;
            }
            ElMessage.success('邀请成功');
            return true;
        },
        async deleteRoom() {
            try {
                const response = await axios_config.delete(
                    '/api/room/delete?room_id=' + this.roomID
                );
                if (response.data.code !== 0) {
                    ElMessage.error(response.data.msg);
                    return false;
                }
                // 将v-model传递的ifUpdate设置为true，触发父组件的watch
                this.$emit('update:ifUpdate', true);
                ElMessage.success('退出房间成功');
            } catch (error) {
                console.error('Failed to delete room:', error);
                return false;
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
                console.error('获取房间成员列表失败', error);
                ElMessage.error('获取房间成员列表失败');
            }


        },
        avatarUrl(url) {
            if (url) {
                return backendBaseUrl + url;
            } else {
                return backendBaseUrl + '/static/avatars/nopic.png';
            }
        }


    }
};
</script>

<style scoped>

.room-drawer-header {
    display: flex;
    justify-content: flex-start;
    align-items: center;
}

</style>
