<template>
  <h2>
    Welcome to Our IM System
  </h2>
  <el-tabs v-model="activeTab" class="demo-tabs" stretch @tab-click="handleClick">
    <el-tab-pane label="登录" name="loginTab"></el-tab-pane>
    <el-tab-pane label="注册" name="registerTab"></el-tab-pane>

    <div class="tab-content">
      <keep-alive>
        <component :is="curTabComponents[activeTab]"></component>
      </keep-alive>
    </div>

  </el-tabs>




</template>

<script>
import UserLogin from '../components/Login/UserLogin.vue'
import UserRegister from '../components/Login/UserRegister.vue'

export default {
  name: 'UserLogin',
  data() {
    return {
      activeTab: 'loginTab',
      curTabComponents: {
        loginTab: UserLogin,
        registerTab: UserRegister
      }
    }
  },
  components: {
    UserLogin,
    UserRegister
  },
  methods: {
    handleClick(tab) {
      this.activeTab = tab.name;
    }
  }
}
</script>


<style>

</style>