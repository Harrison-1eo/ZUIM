<template>
    <div class="video-container" width="640" height="480">
        <canvas ref="canvasPlayer" class="canvas-player" width="640" height="480"></canvas>
        <button @click="StopLive">关闭</button>
    </div>
</template>

<script>

export default {
    name: 'LiveStream',
    props: {
        roomID: {
            type: Number,
            required: true
        },
        VideoChunk: {
            type: String,
            required: true
        },

    },
    mounted() {
        this.initializeLiveStreaming();
    },
    data() {
        return {
            isInitialized: false,
            canvasElement: null,
            mediaSource: null,
            sourceBuffer: null,
            queue: [],

        }
    },

    watch: {
        VideoChunk: {
            deep: true,
            handler(newVal, oldVal) {
                if (!this.isInitialized) {
                    this.initializeLiveStreaming();
                }
                this.handleReceivedVideoChunk_Base64(newVal);
                // this.processQueue();
            }
        },
        'queue.length'(newLength) {
            if (newLength > 0 && this.isInitialized) {
                this.processQueue(); // 当队列非空时，开始处理队列
            }
        }


    },
    methods: {
        initializeLiveStreaming() {
            this.canvasElement = this.$refs.canvasPlayer;
            // this.canvasElement = this.$refs.video;
            if (!this.canvasElement) {
                ElMessage.error('Video element is not initialized.');
                return;
            }
            this.isInitialized = true;
        },
        handleReceivedVideoChunk_Base64(VideoChunk_Base64) {
            this.queue.push(VideoChunk_Base64); // 将视频块添加到队列中
            // this.processQueue(); // 尝试处理队列
        },
        processQueue() {

            try {
                while (this.queue.length > 0) {
                    const VideoChunk_Base64 = this.queue.shift(); // 从队列中移除一个元素
                    // 把经base64编码的图片转为canvas，显示在canvasElement上，且左右反转
                    const img = new Image();
                    img.src = VideoChunk_Base64;
                    img.onload = () => {
                        const canvas = this.canvasElement;
                        const ctx = canvas.getContext('2d');
                        ctx.drawImage(img, 0, 0, canvas.width, canvas.height);
                    };
                }
            } catch (error) {
                ElMessage.error('Error processing queue:', error);
            }

        },

        StopLive() {
            // 调用chatinput的closecamara函数，关闭摄像头，关闭直播
            this.$emit('closeCamera');
            // 消除视频流，消除this.queue,消除this.VideoChunk，消除this.canvasElement，消除video-container
            this.queue = [];
            this.canvasElement = null;
            this.isInitialized = false;
            this.$el.remove();
            this.$emit('ForceforceDeleteVideoBeTrue');

        }

    }
}
</script>

<style scoped>
.video-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 640px;
    background-color: #000;
}
.video-container button {
    /* 右上角 */
    position: absolute;
    top: 0;
    right: 0;
    z-index: 999;
    background-color: #f56c6c;
    color: white;
    border: none;
    border-radius: 5px;
    padding: 5px 10px;
    cursor: pointer;
    font-size: 16px;
    margin: 10px;
}
</style>
