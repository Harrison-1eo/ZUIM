<template>
    <div class="backgroud">
        <div class="title">
            <h2>
                Welcome to Our IM System
            </h2>
        </div>
        <!-- <div class="login-box">
            <form action="#" class="login-form">
                <el-tabs v-model="activeTab" class="demo-tabs"  stretch @tab-click="handleClick" style="height: 200px">
                    <el-tab-pane label="登录" :label-style="{ fontSize: '80px' }" name="loginTab" class="large-text"></el-tab-pane>
                    <el-tab-pane label="注册" :label-style="{ fontSize: '80px' }" name="registerTab" class="large-text"></el-tab-pane>

                    <div class="tab-content">
                        <keep-alive>
                            <component :is="curTabComponents[activeTab]"></component>
                        </keep-alive>
                    </div>
                </el-tabs>
            </form>
        </div> -->
        <div class="login-container">
            <div class="login-form">
                <div class="tabs">
                    <div class="damn" @click="handleClick">
                        <div id="loginTab" class="loginTab" style="color: #f60;">账号密码登录</div>
                        <div class="registerTab">注册</div>
                    </div>
                    <div class="tab-content">
                        <keep-alive>
                            <component :is="curTabComponents[activeTab]"></component>
                        </keep-alive>
                    </div>
                </div>

            </div>
        </div>
        <!-- 新增的文字说明部分 -->
        <div class="description">
            <div class="about">
                <p>
                    一个vue项目，后端使用Go语言构建，通信采用websocket，加解密算法使用C语言编写。
                    一个基于软件优化后的ZUC-256国密算法的即时多人聊天系统，可以文字聊天，可以发文件，开启多人视频会议，传输的内容都经过了我们小组封装好的协议的加密保护，确保了保密性，软件优化后的ZUC-256算法高效且多密钥状态下加解密速度快，达到38Gbps，使得加解密速度不再成为系统性能瓶颈。

                </p>
                <div @click="goPage('https://github.com/Harrison-1eo/ZUIM/')">项目源码</div>

            </div>
        </div>
    </div>

</template>

<script>
import UserLogin from '../components/Login/UserLogin.vue'
import UserRegister from '../components/Login/UserRegister.vue'
export default {
    name: 'UserLogin',
    data() {
        return {
            activeTab: 'loginTab',
            // activeTab显示为橙色, 其他为黑色


            curTabComponents: {
                loginTab: UserLogin,
                registerTab: UserRegister
            }
        }
    },
    created() {
        // activeTab显示为橙色, 其他为黑色
        var loginTab = document.querySelector('#loginTab');
        // 确保元素存在
        if (loginTab) {
            // 修改样式
            loginTab.style.color = '#f60';
        } else {
            console.error('Element with id "loginTab" not found');
        }

    },
    components: {
        UserLogin,
        UserRegister
    },
    methods: {
        goToLogin() {
            this.activeTab = 'loginTab'
        },
        goToRegister() {
            this.activeTab = 'registerTab'
        },
        handleClick(e) {
            if (e.target.innerText === '账号密码登录') {
                this.goToLogin();
                // activeTab显示为橙色, 其他为黑色
                e.target.style.color = '#f60';
                e.target.nextElementSibling.style.color = '#000';
            } else {
                this.goToRegister();
                e.target.style.color = '#f60';
                e.target.previousElementSibling.style.color = '#000';
            }
        },
        goPage(url) {
            window.location.href = url
        }
    }
}
</script>


<style>
.damn {
    display: flex;
    justify-content: space-around;
    /* align-items: center; */
}
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    /* background-color: #f9f9f9; */
}
.login-form {
    width: 400px;
    padding: 20px;
    background-color: #fff;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 10px;
    padding: 30px;
}
.tabs {
    display: flex;
    justify-content: space-around;
    margin-bottom: 20px;
    flex-direction: column;
}
.tab {
    padding: 10px;
    cursor: pointer;
}
.tab.active {
    color: #f60;
    border-bottom: 2px solid #f60;
}
.input-group {
    margin-bottom: 15px;
}
.input-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 5px;
}
.login-button {
    width: 100%;
    padding: 10px;
    background-color: #f60;
    border: none;
    border-radius: 5px;
    color: #fff;
    font-size: 16px;
    cursor: pointer;
}
.login-button:hover {
    background-color: #e55d00;
}
.links {
    display: flex;
    justify-content: space-between;
    margin-top: 10px;
}
.links a {
    color: #999;
    text-decoration: none;
}
.other-login {
    margin-top: 20px;
    text-align: center;
}
.other-login .icons {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-top: 10px;
}
.title {
    text-align: center;
    /* 字体变大 */
    font-size: 24px;
    margin-top: 30px;
}
.title h2 {
    font-size: 54px;
    color: #333;
    /* margin: 0 0 20px; */
    margin-top: 0;
}
.backgroud {
    background-image: url("../assets/beach.png");
    background-size: cover;
    height: 100%;
    display: flex;
    align-items: center;
    flex-direction: column;

    justify-content: space-around;
}

.login-box {
    /* 透明 */
    /* background-color: rgba(255, 255, 255, 0.8); */
    max-width: 600px;
    height: 460px;
    margin: 0 auto;
    padding: 20px;
    /* background-color: #fff; */
    /* border-radius: 4px; */
    /* box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); */
    opacity: 0.9;
}

/* 实现毛玻璃效果 */
.login-box::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    /* background: url("../assets/backgroud.gif") no-repeat; */
    background-size: cover;
    filter: blur(10px);
    opacity: 0.9;
    z-index: -1;
}

.login-form {
    max-width: 600px;
    margin: 0 auto;
}

.demo-tabs {
    width: 100%;
}

.tab-content {
    margin-top: 20px;
    height: 340px;
    border-radius: 10px;
}

h2 {
    font-size: 24px;
    color: #333;
    margin-bottom: 20px;
}

/* .el-tab-pane {
    font-size: 36px;
    color: #333;

} */
.large-text {
    font-size: 44px;
}

.el-tabs__item {
    font-size: 16px;
    color: #333;
}

.el-tabs__active-bar {
    height: 2px;
    background-color: #27c9b7;
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

.description {
    width: 60%;
    text-align: center;
    margin-top: 20px;
    /* 实现毛玻璃效果 */
    /* background-color: rgba(255, 255, 255, 0.8); */
    /* border-radius: 4px; */
    /* box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); */
    opacity: 0.7;
}
.about {
    /* 实现毛玻璃效果 */
    background-color: rgba(106, 166, 198, 0.8);
    /* border-radius: 4px; */
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    opacity: 0.7;
    padding: 20px;
    border-radius: 10px;
}
.about p {
    font-size: 16px;
    color: #333;
    margin: 0 auto;
    text-align: justify;
    line-height: 0.8cm;
}
</style>