<template>

  <el-drawer show-close>
    <template #header>
      <h2>{{ room.name }}</h2>
    </template>


    <el-card style="max-width: 480px">
      <template #header>
        <div class="card-description">
          <span>房间描述</span>
        </div>
      </template>
      <p>{{ room.description }}</p>
    </el-card>

    <p></p>

    <!-- 用于显示房间成员列表 -->
    <el-card style="max-width: 480px">
      <template #header>
        <div class="card-header">
          <span>房间成员</span>
        </div>
      </template>
      <p v-for="user in roomUsers" :key="user.id">{{ user.username }}
      </p>
      <template #footer>
        <el-button type="primary" >邀请成员（未绑定事件）</el-button>
      </template>

    </el-card>

    <!-- 用于显示房间文件列表 -->

    <template #footer>
      <el-button type="danger" >退出房间（未绑定事件）</el-button>
    </template>


  </el-drawer>

</template>


<script>
import axios from "@/axios-config";
import {ElMessage} from "element-plus";

export default {
  name: 'RoomDrawer',
  props: {
    roomID: {
      type: Number,
      required: true
    },
    room: {
      type: Object,
      required: true
    },
    drawer: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return{
      roomUsers: [],
    }
  },
  watch: {
    roomID() {
      this.fetchRoomUsers();
    }
  },
  created() {
    this.fetchRoomUsers();
  },
  methods: {
    async fetchRoomUsers() {
      // 获取房间成员列表
      try {
        const response = await axios.get('http://localhost:8000/api/room/members?room_id=' + this.roomID);

        if (response.data.code === 0) {
          this.roomUsers = response.data.data;
        } else {
          ElMessage.error(response.data.msg);
        }
      } catch (error) {
        console.error('获取房间成员列表失败', error);
        ElMessage.error('获取房间成员列表失败');
      }


    }
  }
};
</script>
