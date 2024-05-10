<template>
  <div>
    <h1>个人信息</h1>
    <p>用户名：{{ username }}</p>
    <p>邮箱：{{ email }}</p>
  </div>
</template>

<script>
export default {
  name: 'Info',
  data() {
    return {
      username: '用户A',
      email: 'user@example.com'
    };
  }
};
</script>

<style scoped>
/* 这里可以添加组件的样式 */
</style>
