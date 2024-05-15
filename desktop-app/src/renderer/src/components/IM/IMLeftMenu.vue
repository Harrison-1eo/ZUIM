<!-- src/components/IMLeftMenu.vue -->

<template>
  <el-row class="im-left-menu">
    <el-col style="width: min-content;">
      <el-menu
          default-active="me"
          class="el-menu-vertical-demo"
          style="height: 100%"
          @select="handleSelect"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-menu-item index="home">
          <template #title>
            <el-icon><House /></el-icon>
            <span>首页</span>
          </template>
        </el-menu-item>
        <el-sub-menu index="me">
          <template #title>
            <el-icon><Pointer /></el-icon>
            <span>关于我</span>
          </template>
          <el-menu-item index="info">个人信息</el-menu-item>
          <el-menu-item index="updateInfo">修改资料</el-menu-item>
          <el-menu-item index="password">修改密码</el-menu-item>
        </el-sub-menu>
        <el-menu-item index="friend">
          <template #title>
            <el-icon><User /></el-icon>
            <span>好友</span>
          </template>
        </el-menu-item>
        <el-menu-item index="chat">
          <template #title>
            <el-icon><ChatRound /></el-icon>
            <span>聊天室</span>
          </template>
        </el-menu-item>
        <el-menu-item index="setting">
          <el-icon><setting /></el-icon>
          <span>设置</span>
        </el-menu-item>
      </el-menu>
    </el-col>
  </el-row>
</template>

<script >

export default {
  name: 'IMLeftMenu',
  data() {
    return {
      activeIndex: 'home'
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath)
      this.$emit('selectRoute', key)
    },
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },
  }
}
</script>

<style scoped>
.im-left-menu {
  text-align: center;
}

.el-menu-vertical-demo {
  border-right: 1px solid #ebeef5; /* 添加右侧边框 */
}
</style>
