<template>
    <!--  消息显示框：如果是用户发送的消息，则显示在右侧，否则显示在左侧  -->
    <div class="message-item" :class="{ 'message-from-user': isSenderUser }">
        <!--  头像显示  -->
        <el-avatar :src="getAvatarUrl(message.sender_avatar)" style="margin-top: 20px"></el-avatar>
        <!--  消息内容显示：包含发送者名称、消息内容  -->
        <div class="message-content">
            <!--  用户名称显示  -->
            <div :class="{'sender-name-from-user': isSenderUser, 'sender-name': !isSenderUser}">
                {{ isSenderUser ? '我' : message.sender_name || 'No sender-name'}}
            </div>

            <!--  消息内容显示  -->
            <div class="chat-bubble chat-bubble-top"
                 :class="isSenderUser ? 'chat-bubble-from-right' : 'chat-bubble-from-left'"
                 style="--chat-bubble-background-color: #95EC69;
                    --chat-bubble-border-color: rgba(0, 0, 0, 0);
                    --chat-bubble-text-color: #000">

                <!--      文本显示 -->
                <!--            <div v-if="message.type === 'text'" class="text-message">-->
                <div v-if="message.type === 'text'">
                    {{ message.content }}
                </div>
                <!--      图片显示 -->
                <!--            <div v-else-if="message.type === 'image'" class="image-message">-->
                <div v-else-if="message.type === 'image'">
                    <el-image :src="getPicUrl(message.content)"
                              fit="contain"
                              :previewSrcList="[getPicUrl(message.content)]"
                              style="height: auto; display: block;"
                    ></el-image>
                </div>
                <!--      文件显示 -->
                <!--            <div v-else-if="message.type === 'file'" class="file-message file-message-link">-->
                <div v-else-if="message.type === 'file'" class="file-message-link">
                    <el-icon>
                        <Link />
                    </el-icon>
                    <el-link :href="getFileUrl(message.content)" target="_blank">
                        下载文件：{{ this.getFileName(message.content) }}
                    </el-link>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
import { ElAvatar, ElImage, ElLink } from 'element-plus';
import { Link } from "@element-plus/icons";
import backendBaseUrl from "@/utils/base-url-setting";
import "@/assets/chatbubble.css";

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
            const userId = localStorage.getItem('userId');
            return userId && parseInt(userId, 10) === this.message.sender_id;
        }
    },
    methods: {
        getFileUrl(content) {
            const regex = /you can download the file on (\/static\/files\/\S+)/;
            const match = content.match(regex);
            if (match) {
                console.log(match[1]); // 输出匹配到的URL
                return backendBaseUrl + match[1]; // 返回匹配到的完整URL
            }
            return ""; // 如果没有匹配到，返回null
        },
        getFileName(content) {
            // content中的数据内容是：response.data.data.file_name + ' you can download the file on ' + response.data.data.file_url
            // 使用正则表达式匹配出文件的名称
            const regex = /(.+) you can download the file on \/\S+/g;
            return content.replace(regex, '$1');
        },
        getAvatarUrl(avatar) {
            // console.log(user);
            console.log(avatar);
            if (!avatar) {
                // 如果 user.avatar 为 null，返回默认的 URL
                console.log('1');
                return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
            } else if (typeof avatar === 'string') {
                return backendBaseUrl + avatar;
            }
            // 如果头像文件不存在或user.avatar格式不正确，则返回默认的 URL
            console.log('4');
            return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
        },
        getPicUrl(url) {
            // console.log(url);
            if (url) {
                return backendBaseUrl + url;
                // return this.avatarUrl(url);
            }
            return backendBaseUrl + '/static/avatars/nopic.png';
        }
    }
}
</script>

<style scoped>
.message-item {
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    margin: 10px;
}
.message-from-user {
    flex-direction: row-reverse;
}
.message-content {
    display: flex;
    flex-direction: column;

}
.sender-name {
    margin-left: 10px;
    font-weight: bold;
}
.sender-name-from-user {
    margin-right: 10px;
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

.file-message-link {
    display: flex;
    align-items: center;
}
</style>
