<template>

    <div>
        <div class="Logincontainer">
            <div class="loginorreg">
                <p>账号</p>
                <el-input v-model="registerUsername" placeholder="请输入账号" type="text" clearable />
            </div>
            <div class="loginorreg">
                <p>密码</p>
                <el-input v-model="registerPassword" placeholder="请输入密码" type="password" show-password />
            </div>
            <div class="loginorreg">
                <p>确认密码</p>
                <el-input v-model="registerPassword2" placeholder="请再次输入密码" type="password" show-password />
            </div>
            <div style="text-align: center">
                <el-button type="primary" @click="apiRegister">注册</el-button>
            </div>
        </div>

    </div>

</template>

<script>
import axios_origin from "@/utils/axios-origin";
import { ElMessage } from "element-plus";

export default {
    name: 'UserLogin',
    data() {
        return {
            activeTab: 'loginTab',
            registerUsername: '',
            registerPassword: '',
            registerPassword2: ''
        }
    },
    methods: {
        checkRegister() {
            if (this.registerUsername === '' || this.registerPassword === '' || this.registerPassword2 === '') {
                ElMessage.error('用户名或密码不能为空');
                return false;
            }
            if (this.registerPassword !== this.registerPassword2) {
                ElMessage.error('两次输入密码不一致');
                return false;
            }
            return true;
        },
        async apiRegister() {
            if (this.checkRegister() === false) {
                return;
            }
            try {
                const res = await axios_origin.post('/register', {
                    'username': this.registerUsername,
                    'password': this.registerPassword
                });
                if (res.status === 200) {
                    if (res.data.code === 0) {
                        ElMessage.success('注册成功');
                        this.activeTab = 'loginTab';
                        this.password = '';
                        this.password2 = '';
                    } else {
                        ElMessage.error(res.data.msg);
                    }
                } else {
                    ElMessage.error('注册请求失败');
                }
            } catch (error) {
                console.error('An error occurred:', error);
                ElMessage.error('连接服务器请求失败');
            }
        }
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