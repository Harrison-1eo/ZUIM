<template>
    <div class="message-item" :class="{ 'message-from-user': isSenderUser }">
        <el-avatar :src="getAvatarUrl(message)"></el-avatar>
        <div class="message-content">
            <div :class="{'sender-name-from-user': isSenderUser, 'sender-name': !isSenderUser}">
                {{ isSenderUser ? '我' : message.sender_name || 'No sender-name'}}
            </div>
            <!--      文本显示 -->
            <div v-if="message.type === 'text'" class="text-message">
                {{ message.content }}
            </div>
            <!--      图片显示 -->
            <div v-else-if="message.type === 'image'" class="image-message">
                <el-image :src="getPicUrl(message.content)" fit="cover" :previewSrcList="[getPicUrl(message.content)]" style="width: 200px; height: auto" loading="lazy"></el-image>
            </div>
            <!--      文件显示 -->
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
import { ElAvatar, ElImage, ElLink } from 'element-plus';
import { Link } from "@element-plus/icons";
import axios from '@/axios-config';
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
            const regex = /you can download the file on (http:\/\/localhost:8000\/static\/files\/\S+)/;
            const match = content.match(regex);
            if (match) {
                console.log(match[1]); // 输出匹配到的URL
                return match[1]; // 返回匹配到的完整URL
            }
            return ""; // 如果没有匹配到，返回null
        },
        getFileName(content) {
            // content中的数据内容是：response.data.data.file_name + ' you can download the file on http://localhost:8000' + response.data.data.file_url
            // 使用正则表达式匹配出文件的名称
            const regex = /(.+) you can download the file on http:\/\/localhost:8000\/\S+/g;
            return content.replace(regex, '$1');
        },
        getAvatarUrl(message) {
            // console.log(user);
            console.log('the whole message', message);
            console.log('sender avatar', message.sender_avatar);
            // localStorage.getItem('sender_id_avatar')是一个list，里面每个元素都是json格式文件
            // 如果能在localStorage中找到对应的sender_id_avatar，就直接返回对应的avatar
            // var temp = JSON.parse(localStorage.getItem('sender_id_avatar'));
            // console.log('temp');
            // console.log(temp);
            if (localStorage.getItem('sender_id_avatar')) {
                console.log('sender_id_avatar', localStorage.getItem('sender_id_avatar'));
                var temp = JSON.parse(localStorage.getItem('sender_id_avatar'));
                console.log('temp');
                console.log(temp);
                // temp里面有东西
                for (let i = 0; i < temp.length; i++) {
                    if (temp[i].id === message.sender_id) {
                        console.log('从localStorage中获取头像');
                        return temp[i].avatar;
                    }
                }
            }
            // 如果找不到对应的sender_id_avatar，就从message中获取sender_avatar
            if (!message.sender_avatar) {
                // 如果 user.avatar 为 null，返回默认的 URL
                console.log('1');
                return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
            } else if (typeof message.sender_avatar === 'string') {
                // ？？？不能从前端获取用户头像，总不能访问网页时将所有用户头像都下载下来吧
                // 如果 user.avatar 是字符串，则尝试查找对应的头像文件
                // const formats = ['png', 'jpg', 'jpeg'];
                // console.log('2');
                // for (const format of formats) {
                //     const url = require(`@/assets/avatar/${avatar}.${format}`);
                //     if (url) {
                //         console.log('3');
                //         return url;
                //     }
                // }
                // 如果头像文件存在，则返回头像的 URL，并且保存对应的 sender_id 和 avatar到localStorage中
                const sender_id_avatar_temp = {
                    id: message.sender_id,
                    avatar: 'http://localhost:8000' + message.sender_avatar
                };
                // sender_id_avatar = sender_id_avatar + sender_id_avatar_temp
                if (typeof temp === 'list') {
                    temp.push(sender_id_avatar_temp);
                } else {
                    temp = [sender_id_avatar_temp];
                }
                localStorage.setItem('sender_id_avatar', JSON.stringify(temp));

                return 'http://localhost:8000' + message.sender_avatar;
            }
            // 如果头像文件不存在或user.avatar格式不正确，则返回默认的 URL
            console.log('4');
            return 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png';
        },
        getPicUrl(url) {
            if (url) {
                return 'http://localhost:8000' + url;
                // return this.avatarUrl(url);
            }
            return 'http://localhost:8000/static/avatars/nopic.png';
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
