<!-- /src/components/IM/Room/RoomDrawer.vue -->

<template>
    <el-drawer show-close drawer>
        <template #header>
            <h2>{{ room.name }}</h2>
        </template>

        <el-card style="max-width: 480px">
            <template #header>
                <div class="card-description">
                    <span>房间描述</span>
                </div>
            </template>
            <p>{{ room.description }}</p>
            <template #footer>
                <el-button type="primary" class="changedescription" @click="changedescrpition">修改房间描述（未设置方法）</el-button>
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
                    <el-avatar :src="avatarUrl(user)" :size="40"></el-avatar>
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
import axios from "@/axios-config";
import { ElMessage, ElMessageBox } from "element-plus";
import dick1 from '../../../assets/avatar/dick1.png';
import { isNavigationFailure } from 'vue-router';
export default {
    name: 'RoomDrawer',
    props: {
        roomID: {
            type: Number,
            required: true
        },
        room: {
            type: Object,
            required: true
        },
    },
    data() {
        return {
            dick1: dick1,
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
    computed: {

    },
    methods: {
        changedescrpition() {
            ElMessage.info('修改房间描述功能未实现');
        },
        addUserBox() {
            ElMessageBox.prompt('请输入邀请的用户名称', '邀请朋友').then(({ value }) => {
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
        async addUser(name) {
            try {
                const response = await axios.post('http://localhost:8000/api/room/add_user',
                    {
                        'room_id': this.roomID,
                        'user_name': name,
                        'role': 'member'
                    },
                );
                console.log('Add user:', response.data);
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
                const response = await axios.delete(
                    'http://localhost:8000/api/room/delete?room_id=' + this.roomID
                );
                console.log('Delete room:', response.data);
                if (response.data.code !== 0) {
                    ElMessage.error(response.data.msg);
                    return false;
                }
                // 将v-model传递的ifFetch设置为true，触发父组件的watch
                this.$emit('update:ifFetch', true);
                ElMessage.success('退出房间成功');
            } catch (error) {
                console.error('Failed to delete room:', error);
                return false;
            }
        },
        async fetchRoomUsers() {
            // 获取房间成员列表
            try {
                const response = await axios.get('http://localhost:8000/api/room/members?room_id=' + this.roomID);

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
        avatarUrl(user) {
            console.log(user);
            console.log(user.avatar);
            if (!user.avatar) {
                // 如果 user.avatar 为 null，返回默认的 URL
                console.log('1');
                return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
            } else if (typeof user.avatar === 'string') {
                // 如果 user.avatar 是字符串，则尝试查找对应的头像文件
                const formats = ['png', 'jpg', 'jpeg'];
                console.log('2');
                for (const format of formats) {
                    const url = require(`@/assets/avatar/${user.avatar}.${format}`);
                    if (url) {
                        console.log('3');
                        return url;
                    }
                }
            }
            // 如果头像文件不存在或user.avatar格式不正确，则返回默认的 URL
            console.log('4');
            return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
        }


    }
};
</script>
