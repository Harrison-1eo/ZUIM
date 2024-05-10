<template>

  <div >
    <p>账号</p>
    <el-input v-model="registerUsername" placeholder="请输入账号" type="text" clearable/>

    <p>密码</p>
    <el-input v-model="registerPassword" placeholder="请输入密码" type="password" show-password/>

    <p>确认密码</p>
    <el-input v-model="registerPassword2" placeholder="请再次输入密码" type="password" show-password/>
    <el-button type="primary" @click="apiRegister">注册</el-button>
  </div>


</template>

<script>
import axios from "axios";
import {ElMessage} from "element-plus";

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
        const res = await axios.post('/register', {
          username: this.username,
          password: this.password
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

</style>