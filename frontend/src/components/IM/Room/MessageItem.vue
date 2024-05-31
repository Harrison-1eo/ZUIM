<template>

    <!--  消息显示框：如果是用户发送的消息，则显示在右侧，否则显示在左侧  -->
    <div class="message-item"
         :class="{ 'message-from-user': isSenderUser }">
        <!--  头像显示  -->
        <el-avatar shape="square" :src="getAvatarUrl(message.sender_avatar)" style="margin-top: 0px"></el-avatar>
        <!--  消息内容显示：包含发送者名称、消息内容  -->
        <div class="message-content">
            <!--  用户名称显示  -->
            <div v-if="isSenderUser===false" class="sender-name, sender-name-from-other">
                {{ message.sender_name }}
            </div>

            <!--  消息内容显示  -->
            <div @mouseenter="showMoreButton = true"
                 @mouseleave="showMoreButton = false">
                <div :class="{ 'bubble-and-more-left': !isSenderUser, 'bubble-and-more-right': isSenderUser }">
                    <div class="chat-bubble chat-bubble-top"
                         :class="isSenderUser ? 'chat-bubble-from-right' : 'chat-bubble-from-left'"
                         style="--chat-bubble-background-color: #ffffff;
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
                                <Link/>
                            </el-icon>
                            <el-link :href="getFileUrl(message.content)" target="_blank">
                                下载文件：{{ this.getFileName(message.content) }}
                            </el-link>
                        </div>
                    </div>
                    <el-icon v-if="showMoreButton" style="margin: 5px 10px; cursor: pointer;" @click="showDialog = true">
                        <More/>
                    </el-icon>
                </div>
            </div>
        </div>
    </div>
    <el-dialog v-model="showDialog" title="加密信息" width="800">
        <template #header>
            <div>加密信息</div>
        </template>
        <!-- 展示message.encryptInfo中的信息 -->
        <div class="message-encrypt-info">
            <div class="message-encrypt-info-key">加密方式： </div>
            <div class="message-encrypt-info-value">ZUC-256 加速算法</div>
        </div>

        <div class="message-encrypt-info">
            <div class="message-encrypt-info-key">消息长度：</div>
            <div class="message-encrypt-info-value">{{ message.encryptInfo.length }}</div>
        </div>

        <div class="message-encrypt-info">
            <div class="message-encrypt-info-key">起始位置：</div>
            <div class="message-encrypt-info-value">{{ message.encryptInfo.position }}</div>
        </div>

        <div class="message-encrypt-info">
            <div class="message-encrypt-info-key">加密密钥：</div>
            <div class="message-encrypt-info-value">{{ message.encryptInfo.key }}</div>
        </div>

        <div class="message-encrypt-info">
            <div class="message-encrypt-info-key">加密结果（经过Base64编码）：</div>
            <div class="message-encrypt-info-value">{{ message.encryptInfo.data }}</div>
        </div>
    </el-dialog>
</template>

<script>
import {ElAvatar, ElImage, ElLink} from 'element-plus';
import {Link} from "@element-plus/icons";
import {backendBaseUrl} from "@/utils/base-url-setting";
import "@/assets/chatbubble.css";

export default {
    data() {
        return {
            showMoreButton: false,
            showDialog: false,
        }
    },
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
            if (!avatar) {
                // 如果 user.avatar 为 null，返回默认的 URL
                return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
            } else if (typeof avatar === 'string') {
                return backendBaseUrl + avatar;
            }
            // 如果头像文件不存在或user.avatar格式不正确，则返回默认的 URL
            return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
        },
        getPicUrl(url) {
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
    margin: 10px 20px;
}

.message-from-user {
    flex-direction: row-reverse;
}

.message-content {
    display: flex;
    flex-direction: column;

}

.sender-name {
    /* font-weight: lighter; */
    font-size: 14px;
    font-stretch: semi-condensed;
    color: #808080;
}


.sender-name-from-other {
    margin-left: 10px;
    text-align: left;
}

.sender-name-from-user {
    margin-right: 10px;
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

.bubble-and-more-left {
    display: flex;
    flex-direction: row;
    align-items: center;
}

.bubble-and-more-right {
    display: flex;
    flex-direction: row-reverse;
    align-items: center;
}

.message-encrypt-info {
    display: flex;
    flex-direction: column;
    margin: 15px;
}

.message-encrypt-info-key {
    font-weight: bold;
    font-size: large;
}
</style>
