<template>
    <div class="avatar-uploader-box">
        <el-upload class="avatar-uploader" action="http://localhost:8000/api/user/upload_avatar"
            :headers="{ Authorization: 'Bearer ' + this.token }" :show-file-list="false" :on-success="handleAvatarSuccess"
            :on-error="handleAvatarError" :before-upload="beforeAvatarUpload">
            <img v-if="avatarUrl" :src="avatarUrl" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
    </div>
</template>


<script>
import axios from "@/axios-config";
import { ElMessage } from "element-plus";

export default {
    data() {
        return {
            userInfo: {},
            token: localStorage.getItem('token'),
            avatarUrl: ''
        };
    },
    created() {
        // 获取用户信息
        this.getUserInfo();
    },
    methods: {
        async getUserInfo() {
            try {
                const response = await axios.get('/api/user/my');
                if (response.data.code === 0) {
                    this.userInfo = response.data.data;
                } else {
                    console.error('Failed to fetch user info:', response.data.msg);
                    ElMessage.error('获取用户信息失败');
                }
            } catch (error) {
                console.error('Error fetching user information:', error);
                ElMessage.error('获取用户信息失败');
            }
            this.avatarUrl = this.userInfo.avatar ? 'http://localhost:8000' + this.userInfo.avatar : '';
        },
        handleAvatarError(err, file, fileList) {
            console.error('上传头像失败:', err, file, fileList);
            ElMessage.error('上传头像失败');
        },
        handleAvatarSuccess(res) {
            if (res.code === 0) {
                ElMessage.success('上传头像成功');
                this.userInfo = res.data;
                this.avatarUrl = this.userInfo.avatar ? 'http://localhost:8000' + this.userInfo.avatar : '';
                var temp = {
                    'id': this.userInfo.ID,
                    'avatar': this.avatarUrl

                }
                var sender_id_avatar_temp = JSON.parse(localStorage.getItem('sender_id_avatar'));
                if (typeof sender_id_avatar_temp === 'list') {
                    sender_id_avatar_temp.push(temp);
                } else {
                    sender_id_avatar_temp = [temp];
                }
                localStorage.setItem('sender_id_avatar', JSON.stringify(sender_id_avatar_temp));
                console.log('本地头像更改');
            } else {
                ElMessage.error('上传头像失败');
                console.error('Failed to upload avatar:', res.data.msg);
            }
        },
        beforeAvatarUpload(file) {
            const isJPGorPng = file.type === 'image/jpeg' || file.type === 'image/png';
            const isLt2M = file.size / 1024 / 1024 < 2;

            if (!isJPGorPng) {
                this.$message.error('上传头像图片只能是 JPG 或者 PNG 格式!');
            }
            if (!isLt2M) {
                this.$message.error('上传头像图片大小不能超过 2MB!');
            }
            return isJPGorPng && isLt2M;
        }
    }
}
</script>

<style>
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
}

.avatar-uploader .el-upload:hover {
    border-color: #409eff;
}

.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
}

.avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>