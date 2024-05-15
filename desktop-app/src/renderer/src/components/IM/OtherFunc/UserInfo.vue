<template>
    <el-card class="user-info-card" shadow="hover">
        <div class="user-info-header">
            <el-avatar :src="getAvatarUrl(userInfo.avatar)" size="400" class="user-avatar" shape="circle"></el-avatar>
            <div class="user-info-main">
                <h2>{{ userInfo.username || 'No username provided' }}</h2>
                <p>{{ userInfo.email || 'No email provided' }}</p>
            </div>
        </div>
        <el-divider></el-divider>
        <el-descriptions border column="1">
            <el-descriptions-item label="User ID">{{ userInfo.user_id }}</el-descriptions-item>
            <el-descriptions-item label="Nickname">{{
                userInfo.nike_name || 'N/A'
            }}</el-descriptions-item>
            <el-descriptions-item label="Country">{{ userInfo.country || 'N/A' }}</el-descriptions-item>
            <el-descriptions-item label="Province">{{ userInfo.province || 'N/A' }}</el-descriptions-item>
            <el-descriptions-item label="City">{{ userInfo.city || 'N/A' }}</el-descriptions-item>
            <el-descriptions-item label="Account Created">{{
                formatDate(userInfo.CreatedAt)
                }}</el-descriptions-item>
        </el-descriptions>
    </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
    ElCard,
    ElAvatar,
    ElDescriptions,
    ElDescriptionsItem,
    ElDivider,
    ElMessage
} from 'element-plus'
import axios from '@/axios-config'

const userInfo = ref({})

onMounted(async () => {
    try {
        const response = await axios.get('/api/user/my')
        if (response.data.code === 0) {
            userInfo.value = response.data.data
        } else {
            console.error('Failed to fetch user info:', response.data.msg)
            ElMessage.error('获取用户信息失败')
        }
    } catch (error) {
        console.error('Error fetching user information:', error)
        ElMessage.error('获取用户信息失败')
    }
})

function formatDate(dateString) {
    const date = new Date(dateString)
    return date.toLocaleDateString()
}

function getAvatarUrl(url) {
    if (!url) {
        return 'http://localhost:8000/static/avatars/nopic.png'
    } else {
        return 'http://localhost:8000' + url
    }
}
</script>

<style scoped>
.user-info-card {
    max-width: 500px;
    margin: 40px auto;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.user-info-header {
    display: flex;
    align-items: center;
    gap: 20px;
}

.user-info-main h2 {
    margin: 0;
    font-size: 24px;
    font-weight: bold;
    color: #333;
}

.user-info-main p {
    margin: 4px 0 0;
    color: #666;
    font-size: 16px;
}

.user-avatar {
    border: 2px solid #eee;
}
</style>
