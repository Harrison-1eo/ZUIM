<template>
    <div>
      <h1>Register</h1>
      <form @submit.prevent="register">
        <input v-model="username" placeholder="Username" required>
        <input type="password" v-model="password" placeholder="Password" required>
        <button type="submit">Register</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',
        password: ''
      };
    },
    methods: {
      async register() {
        const res = await axios.post('/api/register', {
          username: this.username,
          password: this.password
        });
        if (res.data.success) {
          this.$router.push('/login');
        } else {
          alert(res.data.message);
        }
      }
    }
  };
  </script>