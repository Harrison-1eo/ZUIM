<template>
    <h2>
        Welcome to Our IM System
    </h2>
    <el-tabs v-model="activeTab" class="demo-tabs" stretch @tab-click="handleClick">
        <el-tab-pane label="登录" name="loginTab"></el-tab-pane>
        <el-tab-pane label="注册" name="registerTab"></el-tab-pane>

        <div class="tab-content">
            <keep-alive>
                <component :is="curTabComponents[activeTab]"></component>
            </keep-alive>
        </div>

    </el-tabs>




</template>

<script>
import UserLogin from '../components/Login/UserLogin.vue'
import UserRegister from '../components/Login/UserRegister.vue'

export default {
    name: 'UserLogin',
    data() {
        return {
            activeTab: 'loginTab',
            curTabComponents: {
                loginTab: UserLogin,
                registerTab: UserRegister
            }
        }
    },
    components: {
        UserLogin,
        UserRegister
    },
    methods: {
        handleClick(tab) {
            this.activeTab = tab.name;
        }
    }
}
</script>


<style>
.demo-tabs {
    width: 100%;
}

.tab-content {
    margin-top: 20px;
}

h2 {
    font-size: 24px;
    color: #333;
    margin-bottom: 20px;
}

.el-tab-pane {
    padding: 20px;
}

.el-tabs__item {
    font-size: 16px;
    color: #333;
}

.el-tabs__active-bar {
    height: 2px;
    background-color: #1890ff;
}

.tab-content {
    background-color: #f5f5f5;
    padding: 20px;
}

.form-container {
    max-width: 400px;
    margin: 0 auto;
}

.form-group {
    margin-bottom: 20px;
}

.form-label {
    font-size: 16px;
    color: #333;
}

.form-input {
    width: 100%;
    height: 40px;
    padding: 8px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.form-submit {
    width: 100%;
    height: 40px;
    background-color: #1890ff;
    color: #fff;
    font-size: 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.form-submit:hover {
    background-color: #1476d7;
}

.error-message {
    color: #f5222d;
}

.success-message {
    color: #52c41a;
}
</style>