<!-- src/components/IMLeftMenu.vue -->

<template>
  <el-row class="im-left-menu">
    <el-col style="width: min-content;">
      <el-menu
          default-active="me"
          class="el-menu-vertical-demo"
          style="height: 100%; background: none;"
          @select="handleSelect"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-menu-item class="el-menu-item-button" index="home">
          <template #title>
            <el-icon><House /></el-icon>
            <span>首页</span>
          </template>
        </el-menu-item>
        <el-sub-menu class="el-menu-item-button" index="me" style="background: none">
          <template #title>
            <el-icon><Pointer /></el-icon>
            <span>关于我</span>
          </template>
          <el-menu-item class="el-menu-item-button" index="info">个人信息</el-menu-item>
          <el-menu-item class="el-menu-item-button" index="updateInfo">修改资料</el-menu-item>
          <el-menu-item class="el-menu-item-button" index="password">修改密码</el-menu-item>
        </el-sub-menu>
        <el-menu-item class="el-menu-item-button" index="friend">
          <template #title>
            <el-icon><User /></el-icon>
<!--              <span>好友</span>-->
              <span>终端</span>
          </template>
        </el-menu-item>
        <el-menu-item class="el-menu-item-button" index="chat">
          <template #title>
            <el-icon><ChatRound /></el-icon>
            <span>通信</span>
          </template>
        </el-menu-item>
        <el-menu-item class="el-menu-item-button" index="setting">
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
  created() {
    this.$emit('selectRoute', this.activeIndex)
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
    height: 100%; /* 使导航占满整个左侧部分 */
    text-align: center;
    background: linear-gradient(180deg, #F0EDF7 .03%, #ebeaf5 32.19%, #e8e8f3 68.86%, #e4eaf7 99.12%);
    border-right: 1px solid #dee2eb;
}

.el-menu-vertical-demo {
    border-right: none;
    background: none;
}

.el-menu-item-button {
    background: none;
    border-bottom: #333333 1px;
    /*border-radius: 10px;*/
}

.el-menu-item-button:hover {
    background: #ffffff;
}
</style>
