<template>
    <div class="message-item" :class="{ 'message-from-user': isSenderUser }">
        <el-avatar :src="getAvatarUrl(message.sender_avatar)"></el-avatar>
        <div class="message-content">
            <div :class="{ 'sender-name-from-user': isSenderUser, 'sender-name': !isSenderUser }">
                {{ isSenderUser ? '我' : message.sender_name || 'No sender-name' }}
            </div>
            <!-- 文本显示 -->
            <div v-if="message.type === 'text'" class="text-message">
                {{ message.content }}
            </div>
            <!-- 图片显示 -->
            <div v-else-if="message.type === 'image'" class="image-message">
                <el-image :src="getPicUrl(message.content)" fit="cover" :previewSrcList="[getPicUrl(message.content)]"
                    style="width: 200px; height: auto; display: block"></el-image>
            </div>
            <!-- 文件显示 -->
            <div v-else-if="message.type === 'file'" class="file-message">
                <el-icon>
                    <Link />
                </el-icon>
                <el-link :href="getFileUrl(message.content)" target="_blank">
                    下载文件：{{ this.getFileName(message.content) }}
                </el-link>
            </div>
        </div>
    </div>
</template>

<script>
import { ElAvatar, ElImage, ElLink } from 'element-plus'
import { Link } from '@element-plus/icons-vue'

export default {
    components: {
        Link,
        ElAvatar,
        ElImage,
        ElLink
    },
    props: {
        message: {
            type: Object,
            required: true
        }
    },
    computed: {
        isSenderUser() {
            const userId = localStorage.getItem('userId')
            return userId && parseInt(userId, 10) === this.message.sender_id
        }
    },
    methods: {
        getFileUrl(content) {
            const regex = /you can download the file on (http:\/\/localhost:8000\/static\/files\/\S+)/
            const match = content.match(regex)
            if (match) {
                console.log(match[1]) // 输出匹配到的URL
                return match[1] // 返回匹配到的完整URL
            }
            return '' // 如果没有匹配到，返回null
        },
        getFileName(content) {
            // content中的数据内容是：response.data.data.file_name + ' you can download the file on http://localhost:8000' + response.data.data.file_url
            // 使用正则表达式匹配出文件的名称
            const regex = /(.+) you can download the file on http:\/\/localhost:8000\/\S+/g
            return content.replace(regex, '$1')
        },
        getAvatarUrl(url) {
            if (!url) {
                return 'http://localhost:8000/static/avatars/nopic.png'
            } else {
                return 'http://localhost:8000' + url
            }
        },
        getPicUrl(url) {
            // console.log(url);
            if (url) {
                return 'http://localhost:8000' + url
                // return this.avatarUrl(url);
            }
            return 'http://localhost:8000/static/avatars/nopic.png'
        }
    }
}
</script>

<style scoped>
.message-item {
    display: flex;
    align-items: flex-start;
    margin: 10px;
}

.message-from-user {
    flex-direction: row-reverse;
}

.message-content {
    margin-left: 10px;
    margin-right: 10px;
}

.sender-name {
    font-weight: bold;
}

.sender-name-from-user {
    font-weight: bold;
    text-align: right;
}

.text-message,
.image-message,
.file-message {
    background: #f2f2f2;
    border: 1px solid #d9d9d9;
    border-radius: 12px;
    margin: 5px;
    padding: 5px 10px 5px 10px;
}

.file-message {
    display: flex;
    align-items: center;
}
</style>
