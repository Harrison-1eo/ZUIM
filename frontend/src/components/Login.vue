<template>
    <div>
        <h1>Login</h1>
        <form @submit.prevent="login">
            <input v-model="username" placeholder="Username" required>
            <input type="password" v-model="password" placeholder="Password" required>
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<script>
import axios from '@/axios-config';

export default {
    data() {
        return {
            username: '',
            password: ''
        };
    },
    methods: {
        async login() {
            const res = await axios.post('/login', {
                username: this.username,
                password: this.password
            });
            if (res.data.code === 0) {
                // 保存token和用户信息到本地存储或Vuex
                localStorage.setItem('token', res.data.data.token);
                localStorage.setItem('user', JSON.stringify(res.data.data.user));
                // this.$router.push('/');
                alert(res.data.msg);
            } else {
                alert(res.data.msg);
            }

        }
    }
};
</script>