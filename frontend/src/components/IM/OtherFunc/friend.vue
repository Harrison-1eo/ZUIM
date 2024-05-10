<template>
  <div class="friend-list">
    <h2>好友列表</h2>
    <el-divider></el-divider>
    <el-scrollbar style="height: 300px;">
      <el-menu
        default-active="1"
        class="el-menu-vertical-demo"
        background-color="#ffffff"
        text-color="#000000"
        active-text-color="#ffffff"
        @select="handleSelect"
      >
        <el-menu-item-group title="在线好友">
          <el-menu-item v-for="friend in onlineFriends" :key="friend.id" :index="friend.id">
            <i class="el-icon-user-solid"></i>
            {{ friend.name }}
          </el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="离线好友">
          <el-menu-item v-for="friend in offlineFriends" :key="friend.id" :index="friend.id">
            <i class="el-icon-user"></i>
            {{ friend.name }}
          </el-menu-item>
        </el-menu-item-group>
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      onlineFriends: [
        { id: 1, name: '张三' },
        { id: 2, name: '李四' },
        { id: 3, name: '王五' }
      ],
      offlineFriends: [
        { id: 4, name: '赵六' },
        { id: 5, name: '钱七' }
      ]
    };
  },
  methods: {
    handleSelect(index) {
      console.log('选中的好友ID:', index);
      // 在这里可以进行选中好友后的操作，比如打开聊天窗口等
    }
  }
};
</script>

<style>
.friend-list {
  padding: 20px;
}
</style>
