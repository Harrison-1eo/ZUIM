<template>
    <div class="backgroud">
        <div class="title">
            <h2>
                Welcome to Our IM System
            </h2>
        </div>
        <div class="login-box">
            <form action="#" class="login-form">
                <el-tabs v-model="activeTab" class="demo-tabs" stretch @tab-click="handleClick" style="height: 200px">
                    <el-tab-pane label="登录" name="loginTab" class="large-text"></el-tab-pane>
                    <el-tab-pane label="注册" name="registerTab" class="large-text"></el-tab-pane>

                    <div class="tab-content">
                        <keep-alive>
                            <component :is="curTabComponents[activeTab]"></component>
                        </keep-alive>
                    </div>
                </el-tabs>
            </form>
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
            console.log(tab);
            this.activeTab = tab.props.name;
        },
        goPage(url) {
            window.location.href = url
        }
    }
}
</script>


<style>
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