<template>

    <div class="logincontainer">
        <div class="loginorreg">
            <p>账号</p>
            <el-input v-model="loginUsername" placeholder="请输入账号" type="text" clearable />
        </div>
        <div class="loginorreg">
            <p>密码</p>
            <el-input v-model="loginPassword" placeholder="请输入密码" type="password" style="text-align: center" show-password />
        </div>
        <!-- div居中 -->

        <div style="text-align: center">
            <el-button type="primary" @click="apiLogin">登录</el-button>
        </div>
    </div>

</template>

<script>
import axios_config from "@/utils/axios-config";
import { ElMessage } from "element-plus";

export default {
    name: 'UserLogin',
    data() {
        return {
            activeTab: 'loginTab',
            loginUsername: '',
            loginPassword: '',
        }
    },
    methods: {
        checkLogin() {
            if (this.loginUsername === '' || this.loginPassword === '') {
                ElMessage.error('用户名或密码不能为空');
                return false;
            } else {
                return true;
            }
        },
        async apiLogin() {
            if (this.checkLogin() === false) {
                return;
            }
            try {
                const res = await axios.post('/login', {
                    'username': this.loginUsername,
                    'password': this.loginPassword
                });
                if (res.status === 200) {
                    if (res.data.code === 0) {
                        // Save token and user information to local storage or Vuex
                        localStorage.setItem('token', res.data.data.token);
                        localStorage.setItem('user', JSON.stringify(res.data.data.user.username));
                        localStorage.setItem('userId', JSON.stringify(res.data.data.user.ID));
                        this.$router.push('/im');
                        ElMessage.success('登录成功');
                    } else {
                        ElMessage.error(res.data.msg);
                    }
                } else {
                    // Handle non-200 response status (server connection failure)
                    ElMessage.error('登录请求失败');
                }
            } catch (error) {
                console.error('An error occurred:', error);
                ElMessage.error('连接服务器请求失败');
            }
        },
    }
}
</script>


<style>
.logincontainer {
    width: 100%;
    max-width: 600px; /* 根据需要调整最大宽度 */
    margin: 0 auto;
    padding: 20px;
    display: flex;
    flex-direction: column; /* 设置 flex 方向为垂直 */
    align-items: center; /* 水平居中 */
}
/* loginorreg宽占container宽的一半 */
.loginorreg {
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
}
.loginorreg p {
    width: 30%;
    text-align: center;
    margin-right: 10px;
}
.loginorreg .el-input {
    margin-bottom: 15px;
    width: 50%; /* 只占整个宽度的一半 */
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
}
</style>