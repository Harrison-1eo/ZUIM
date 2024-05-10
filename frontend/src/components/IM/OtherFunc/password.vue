<template>
  <div>
    <h1>修改密码</h1>
    <form @submit.prevent="changePassword">
      <label for="oldPassword">旧密码：</label>
      <input type="password" id="oldPassword" v-model="oldPassword"><br>
      <label for="newPassword">新密码：</label>
      <input type="password" id="newPassword" v-model="newPassword"><br>
      <button type="submit">确认修改</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Password',
  data() {
    return {
      oldPassword: '',
      newPassword: ''
    };
  },
  methods: {
    changePassword() {
      // 在这里可以添加修改密码的逻辑

      console.log('旧密码：', this.oldPassword);
      console.log('新密码：', this.newPassword);
      // 调用修改密码的接口等操作
    }
  }
};
</script>

<style scoped>
/* 这里可以添加组件的样式 */
</style>
