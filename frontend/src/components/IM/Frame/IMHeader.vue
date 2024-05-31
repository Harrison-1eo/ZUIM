<!-- src/components/IMHeader.vue -->
<template>
    <el-header class="header-container">
        <div class="logo-container">
            <!--      <img src="../../assets/logo.png" alt="Logo" height="50px"/>-->
            <FluentCommunicationShield24Filled />
            <!--      <span>IM SYSTEM</span>-->
            <!--        <span class="system-title">IoT Messaging and Domestic Security Base (IMDSB) System</span>-->
            <span class="system-title">IMDSB System</span>
        </div>
        <el-button @click="logout" class="logout-button">登出</el-button>
    </el-header>
</template>

<script>
import { ElMessage } from "element-plus";
import FluentCommunicationShield24Filled from "@/components/IM/Frame/IMIcon.vue";


export default {
    name: 'IMHeader',
    components: {
        FluentCommunicationShield24Filled
    },
    methods: {
        logout() {
            ElMessage.success('登出成功');
            localStorage.clear();
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


.system-title {
    font-size: 20px;
    font-weight: bold;
    margin-left: 10px;
}
</style>