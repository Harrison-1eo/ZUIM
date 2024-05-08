<template>
  <div class="home-page">
    <!-- 个人信息部分 -->
    <div class="profile">
      <h2>Profile</h2>
      <p>Username: {{ userProfile.username }}</p>
      <p>Email: {{ userProfile.email }}</p>
      <!-- <button @click="edit">Edit</button> -->
    </div>

    <!-- 聊天历史部分 -->
    <div class="chat-history">
      <h2>Chat History</h2>
      <ul>
        <li v-for="history in chatHistories" :key="history.id">
          {{ history.room }} - Last message at: {{ history.lastMessage }}
        </li>
      </ul>
    </div>

    <!-- 聊天房间操作部分 -->
    <div class="chat-room-operations">
      <h2>Chat Rooms</h2>
      <button @click="createRoom">Create Chat Room</button>
      <input v-model="newRoomName" placeholder="Enter room name" />

      <h3>Join a Room</h3>
      <ul>
        <li v-for="room in rooms" :key="room.id">
          {{ room.name }} - <button @click="joinRoom(room.id)">Join</button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userProfile: {
        username: 'JohnDoe',
        email: 'john@example.com'
      },
      chatHistories: [
        { id: 1, room: 'VueJS Fans', lastMessage: '2023-04-03 14:22' },
        { id: 2, room: 'Family', lastMessage: '2023-04-03 09:15' }
      ],
      rooms: [
        { id: 1, name: 'VueJS Fans' },
        { id: 2, name: 'Family' }
      ],
      newRoomName: ''
    };
  },
  methods: {
    createRoom() {
      if (this.newRoomName !== '') {
        this.rooms.push({ id: this.rooms.length + 1, name: this.newRoomName });
        this.newRoomName = ''; // 清空输入框
      }
    },
    joinRoom(roomId) {
      // 在这里加入逻辑，例如可以是路由导航或者调用API
      console.log(`Joining room: ${roomId}`);
    }
  }
};
</script>

<style scoped>
.home-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}
.profile, .chat-history, .chat-room-operations {
  margin: 10px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}
</style>
