<!-- src/components/IMRightContent.vue -->
<template>
  <el-main class="im-right-content">
    <router-view></router-view>
  </el-main>
</template>

<script>
export default {
  name: 'IMRightContent'
}
</script>

<style scoped>
.im-right-content {
  padding: 20px; /* 添加内部空间 */
  background-color: #fff; /* 背景颜色 */
  overflow-y: auto; /* 添加滚动条 */
}
</style>
