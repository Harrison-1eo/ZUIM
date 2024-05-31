<template>
    <div class="tab-content-2">
        <div class="logincontainer">
            <div class="loginorreg">
                <p>账号</p>
                <el-input v-model="loginUsername" placeholder="请输入账号" type="text" clearable />
            </div>
            <div class="loginorreg">
                <p>密码</p>
                <el-input v-model="loginPassword" placeholder="请输入密码" type="password" style="text-align: center, width:80%" show-password />
            </div>
            <!-- div居中 -->

        </div>
        <div class="login-or-regsiter" style="text-align: center">
            <el-button type="primary" @click="apiLogin">登录</el-button>
        </div>
    </div>
</template>

<script>
import axios_origin from "@/utils/axios-origin";
import {
    userCipherFrontend,
    userCipherBackend,
    userCipherWebsocketFrontend,
    userCipherWebsocketBackend
} from "@/utils/encrypt";
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
                const res = await axios_origin.post('/login', {
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

                        let salts = ['backend', 'frontend', 'websocketBackend', 'websocketFrontend']
                        let keys = []
                        for (let i = 0; i < salts.length; i++) {
                            const sha256Password = require('js-sha256').sha256(this.loginPassword);
                            const k = res.data.data.token + sha256Password + res.data.data.user.ID + salts[i];
                            const kk = require('js-sha256').sha256(k);
                            localStorage.setItem(salts[i] + 'Password', kk);
                            keys.push(kk);
                        }

                        userCipherBackend.init(keys[0]);
                        userCipherFrontend.init(keys[1]);
                        userCipherWebsocketBackend.init(keys[2]);
                        userCipherWebsocketFrontend.init(keys[3]);


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
/* .parent {
display: grid;
grid-template-columns: repeat(16, 1fr);
grid-template-rows: repeat(20, 1fr);
grid-column-gap: 0px;
grid-row-gap: 0px;
} */
.tab-content-2 {
    display: grid;
    height: 350px;
    grid-template-columns: repeat(16, 1fr);
    grid-template-rows: repeat(20, 1fr);
    grid-column-gap: 0px;
    grid-row-gap: 0px;
}
.logincontainer {
    grid-area: 1 / 3 / 15 / 15;
    /* 子元素并排 */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: 0;
}
.login-or-regsiter {
    grid-area: 16 / 8 / 18 / 10;
}
.loginorreg {
    width: 80%;
    height: 80px;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0;
}
.loginorreg p {
    width: 20%;
    text-align: center;
    margin-right: 10px;
}
.loginorreg .el-input {
    margin-bottom: 15px;
    width: 70%; /* 只占整个宽度的一半 */
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
}
.loginorreg .el-button {
    height: 44px;
    width: 120px;
}
</style>