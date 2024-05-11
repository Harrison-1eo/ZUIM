<template>
    <div class="chat-input">
        <el-input type="textarea" placeholder="输入消息..." v-model="newMessage" class="message-input" :rows="4" />

        <div class="send-buttons">
            <!-- 发送文件 -->
            <div class="Filecontainer">
                <!-- 文件上传按钮 -->
                <el-button type="primary" class="send-button" @click="openFileChooser" style="margin-bottom: 10px">发送文件</el-button>
                <!-- 隐藏的文件输入 -->
                <input type="file" ref="fileInput" style="display: none;" @change="handleFileUpload">
            </div> 
            <!-- <el-button type="primary" class="send-button" @click="changeVisible" style="margin-bottom: 10px">添加附件</el-button> -->
            
            <el-button type="success" class="send-button" @click="sendMessage">发送消息</el-button>
            <!-- 开启视频 -->
            <el-button type="success" class="send-button" @click="startVideo">开启视频</el-button>
        </div>

        <!-- 对话框 -->
        <el-dialog title="上传文件" v-model="dialogVisible">
            <!-- 上传组件 -->

        </el-dialog>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newMessage: '', // 绑定输入框的内容
            dialogVisible: false
        };
    },
    methods: {
        changeVisible() {
            this.dialogVisible = !this.dialogVisible;
            console.log("changeVisible", this.dialogVisible);
        },
        sendMessage() {
            // 向父组件发送输入的消息
            this.$emit('send', 'text', this.newMessage);
            this.newMessage = ''; // 发送后清空输入框
        },
        // 打开文件选择器
        openFileChooser() {
            this.$refs.fileInput.click();
        },

        // 处理文件上传
        handleFileUpload(event) {
            const file = event.target.files[0]; // 获取文件
            if (!file) {
                return;
            }

            // 在这里添加文件传输的逻辑
            console.log('文件选择成功：', file.name);
            // 假设有一个函数 uploadFile 来处理文件上传
            // this.uploadFile(file);
        },

        // 示例上传函数（需要自己实现具体上传逻辑）
        uploadFile(file) {
            // 使用 FormData 来包装文件
            const formData = new FormData();
            formData.append('file', file);

            // 发送文件到服务器
            fetch('你的上传 API 地址', {
                method: 'POST',
                body: formData,
            })
                .then(response => response.json())
                .then(data => console.log('文件上传成功:', data))
                .catch(error => console.error('文件上传失败:', error));
        },
        startVideo() {
            // 开启视频
            console.log('Starting video...');
        }
    }
};
</script>

<style scoped>
.chat-input {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    background-color: #f4f4f4;
}

.message-input {
    flex: 1; /* 输入框占据剩余空间 */
    margin-right: 10px; /* 输入框和按钮之间的间隔 */
}

.send-button {
    width: 80px; /* 设置发送按钮的宽度 */
    height: 40px; /* 设置按钮的高度，与输入框高度一致 */
    margin-right: 10px;
    margin-left: 10px;
    margin-top: 10px;
    margin-block-end: 10px;

}


</style>
