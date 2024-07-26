<template>
    <el-card class="user-info-card" shadow="hover">
        <h2>修改个人资料</h2>
        <el-divider></el-divider>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">修改头像</div>
            </el-col>
            <el-col :span="16">
                <div class="avatar-uploader-box">
                    <el-upload class="avatar-uploader" action="http://localhost:8000/api/user/upload_avatar"
                               :headers="{ Authorization: 'Bearer ' + this.token }" :show-file-list="false" :on-success="handleAvatarSuccess"
                               :on-error="handleAvatarError" :before-upload="beforeAvatarUpload">
                        <img v-if="avatarUrl" :src="avatarUrl" class="avatar">
                        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                </div>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">用户 ID</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.ID" style="width: 240px" disabled></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">用户名称 (唯一)</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.username" style="width: 240px" placeholder="输入用户名称"></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">邮箱</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.email" style="width: 240px" placeholder="输入邮箱"></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">用户昵称</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.nike_name" style="width: 240px" placeholder="输入昵称" clearable></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">国家</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.country" style="width: 240px" placeholder="输入国家" clearable></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="20" justify="space-between" align="middle">
            <el-col :span="8" class="hint-text">
                <div style="display: flex;justify-content: end">创建于</div>
            </el-col>
            <el-col :span="16">
                <el-input v-model="this.userInfo.CreatedAt" style="width: 240px" disabled></el-input>
            </el-col>
        </el-row>

        <div style="display: flex;justify-content: center;margin: 20px 40px 20px 20px;">
            <el-button type="primary" @click="updateUserInfo">更新</el-button>
            <el-button @click="getUserInfo">重置</el-button>
        </div>


    </el-card>
</template>


<script>
import axios_config from "@/utils/axios-config";
import { backendBaseUrl } from "@/utils/base-url-setting";

import {ElCard, ElDivider, ElMessage, ElRow, ElCol} from "element-plus";

export default {
    components: {ElCard, ElDivider, ElRow, ElCol},
    data() {
        return {
            uploadUrl: backendBaseUrl + '/api/user/upload_avatar',
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
        async updateUserInfo() {
            // 更新用户信息
            // UserID    uint   `json:"user_id"   gorm:"index"`
            // Username  string `json:"username" gorm:"unique"`
            // NikeName  string `json:"nike_name"`
            // Avatar    string `json:"avatar"`
            // Sexuality string `json:"sexuality"`
            // Year      uint   `json:"year"`
            // Month     uint   `json:"month"`
            // Day       uint   `json:"day"`
            // Country   string `json:"country"`
            // Province  string `json:"province"`
            // City      string `json:"city"`
            // Email     string `json:"email" gorm:"unique,index"`
            try {
                const response = await axios_config.post('/api/user/update', this.userInfo);
                if (response.data.code === 0) {
                    ElMessage.success('更新用户信息成功');
                    localStorage.setItem('user', this.userInfo.username);
                } else {
                    ElMessage.error(response.data.msg);
                }
            } catch (error) {
                ElMessage.error('更新用户信息失败');
            }
        },
        async getUserInfo() {
            try {
                const response = await axios_config.get('/api/user/my');
                if (response.data.code === 0) {
                    this.userInfo = response.data.data;
                } else {
                    ElMessage.error('获取用户信息失败');
                }
            } catch (error) {
                ElMessage.error('获取用户信息失败');
            }
            this.avatarUrl = this.userInfo.avatar ? backendBaseUrl + this.userInfo.avatar : backendBaseUrl + '/static/avatars/nopic.png';
        },
        handleAvatarError(err, file, fileList) {
            ElMessage.error('上传头像失败');
        },
        handleAvatarSuccess(res) {
            if (res.code === 0) {
                ElMessage.success('上传头像成功');
                this.userInfo = res.data;
                this.avatarUrl = this.userInfo.avatar ? backendBaseUrl + this.userInfo.avatar : '';
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
            } else {
                ElMessage.error('上传头像失败');
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
.el-row {
    margin-bottom: 20px;
}

.user-info-card {
    max-width: 700px;
    margin: 40px auto;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.avatar-uploader-box .avatar-uploader .avatar{
    width: 100px;
    height: 100px;
    display: block;
}

.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: border-color 0.3s linear;
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

.hint-text {
    color: #606266;
    font-size: 14px;
}
</style>