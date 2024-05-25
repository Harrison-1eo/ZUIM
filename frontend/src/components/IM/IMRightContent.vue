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
    background: linear-gradient(180deg,#f5f4f6,#e6ebf7);
    padding: 0;
}
</style>
