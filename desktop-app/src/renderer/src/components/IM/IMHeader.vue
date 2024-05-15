<!-- src/components/IMHeader.vue -->

<template>
    <el-header class="header-container">
        <div class="logo-container">
            <img src="../../assets/logo.png" alt="Logo" height="50px" />
            <span>IM SYSTEM</span>
        </div>
        <el-button @click="logout" class="logout-button">登出</el-button>
    </el-header>
</template>

<script>
import { ElMessage } from "element-plus";

export default {
    name: 'IMHeader',
    methods: {
        logout() {
            ElMessage.success('登出成功');
            localStorage.removeItem('token')
            localStorage.removeItem('user')
            localStorage.removeItem('userId')
            this.$router.push('/login')
        }
    }
}

</script>
<style scoped>
.header-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    /* 确保左侧和右侧内容分布在两端 */
    padding: 0 20px;
    /* 添加一些内边距 */
}

.logo-container {
    display: flex;
    align-items: center;
    /* 图标和标题上下居中 */
}

.logout-button {
    /* 根据需要可以添加更多样式 */
}
</style>