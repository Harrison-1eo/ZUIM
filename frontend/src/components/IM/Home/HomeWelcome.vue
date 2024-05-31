<template>
    <div class="title-box">

        <div class="title-box-content">
            <h1>{{ beforeAvatarText }}</h1>
        </div>

        <div class="user-avatar">
            <el-avatar v-if="avatarVisible" :size="60" :src="getAvatar(userInfo.avatar)" />
        </div>

        <div class="title-box-content">
            <h1>{{ afterAvatarText }}</h1>
        </div>

    </div>
</template>

<script>
import {backendBaseUrl} from "@/utils/base-url-setting";

export default {
    name: 'HomeWelcome',
    props: {
        userInfo: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            fullTextBeforeAvatar: `${this.getGreeting()}，`,
            fullTextAfterAvatar: '', 
            beforeAvatarText: '',
            afterAvatarText: '',
            avatarVisible: false,
            textIndex: 0,
            isDeleting: false,
            typingSpeed: 200,
            deletingSpeed: 100,
            pauseTime: 1500,
        };
    },
    methods: {
        getGreeting() {
            const now = new Date();
            const hours = now.getHours();
            let greeting;

            if (hours >= 6 && hours < 12) {
                greeting = "早上好";
            } else if (hours >= 12 && hours < 14) {
                greeting = "中午好";
            } else if (hours >= 14 && hours < 18) {
                greeting = "下午好";
            } else {
                greeting = "晚上好";
            }

            return greeting;
        },
        getAvatar(url) {
            if (url) {
                return backendBaseUrl + url;
            }
            return backendBaseUrl + '/static/avatars/nopic.png';
        },
        typeText() {
            if (this.isDeleting) {
                if (this.afterAvatarText.length > 0) {
                    this.afterAvatarText = this.afterAvatarText.substring(0, this.afterAvatarText.length - 1);
                } else if (this.avatarVisible) {
                    this.avatarVisible = false;
                    this.textIndex = this.fullTextBeforeAvatar.length;  
                } else if (this.beforeAvatarText.length > 0) {
                    this.beforeAvatarText = this.beforeAvatarText.substring(0, this.beforeAvatarText.length - 1);
                    this.textIndex--;
                } else {
                    this.isDeleting = false;  // 所有文本删除完毕，重置删除标志
                    setTimeout(this.typeText, this.pauseTime / 3);  // 等待一段时间后重新开始打字
                    return;  // 退出当前递归调用
                }
            } else {
                if (this.textIndex < this.fullTextBeforeAvatar.length) {
                    this.beforeAvatarText += this.fullTextBeforeAvatar.charAt(this.textIndex);
                } else if (this.textIndex === this.fullTextBeforeAvatar.length) {
                    this.avatarVisible = true;
                } else if (this.textIndex > this.fullTextBeforeAvatar.length && this.textIndex <= this.fullTextBeforeAvatar.length + this.fullTextAfterAvatar.length) {
                    this.afterAvatarText += this.fullTextAfterAvatar.charAt(this.textIndex - this.fullTextBeforeAvatar.length - 1);
                }
                this.textIndex++;

                if (this.textIndex > this.fullTextBeforeAvatar.length + this.fullTextAfterAvatar.length) {
                    this.isDeleting = true;
                    setTimeout(this.typeText, this.pauseTime);  // 等待一段时间后开始删除文字
                    return;  // 退出当前递归调用
                }
            }

            setTimeout(this.typeText, this.isDeleting ? this.deletingSpeed : this.typingSpeed);
        },
    },
    mounted() {
        this.fullTextAfterAvatar = ` ${JSON.parse(localStorage.getItem('user'))}。`;
        this.typeText();
    }
};
</script>

<style scoped>

h1 {
    font-size: 3em;
    font-weight: bolder;
    margin: 0;
}

.title-box {
    display: flex;
    align-items: center;
    height: 64px;
}

.user-avatar {
    margin: 0 10px;
}
</style>
