<template>
    <div class="video-container" width="640" height="480">
        <!-- <video ref="video" controls autoplay></video> -->
        <!-- <video id="video" ref="video" controls autoplay></video> -->
        <!-- <video-player class="video-player" ref="videoPlayer" :options="playerOptions"></video-player> -->
        <canvas ref="canvasPlayer" class="canvas-player" width="640" height="480"></canvas>
        <!-- <button @click="isInitialized=false">Stop</button> -->
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
                console.log('VideoChunk changed');
                if (!this.isInitialized) {
                    this.initializeLiveStreaming();
                }
                this.handleReceivedVideoChunk_Base64(newVal);
                // this.processQueue();
            }
        },
        'queue.length'(newLength) {
            if (newLength > 0) {
                this.processQueue(); // 当队列非空时，开始处理队列
            }
        }


    },
    methods: {

        initializeLiveStreaming() {
            console.log('Initializing live streaming');
            this.canvasElement = this.$refs.canvasPlayer;
            // this.canvasElement = this.$refs.video;
            if (!this.canvasElement) {
                console.error('Video element is not initialized.');
                return;
            }


        },
        handleReceivedVideoChunk_Base64(VideoChunk_Base64) {
            this.queue.push(VideoChunk_Base64); // 将视频块添加到队列中
            console.log('Video chunk added to queue');
            // this.processQueue(); // 尝试处理队列
        },
        processQueue() {

            console.log('Processing queue');
            // console.log(this.sourceBuffer);
            console.log(this.queue.length);
            try {
                while (this.queue.length > 0) {
                    const VideoChunk_Base64 = this.queue.shift(); // 从队列中移除一个元素
                    console.log('Appending video chunk to source buffer');
                    // 把经base64编码的图片转为canvas，显示在canvasElement上，且左右反转
                    const img = new Image();
                    img.src = VideoChunk_Base64;
                    // console.log('the VideoChunk_Base64 in live stream:', VideoChunk_Base64);
                    img.onload = () => {
                        const canvas = this.canvasElement;
                        const ctx = canvas.getContext('2d');
                        ctx.drawImage(img, 0, 0, canvas.width, canvas.height);


                    };

                }
            } catch (error) {
                console.error('Error processing queue:', error);
            }

        },

        // 尝试用接收的是视频直接解析的url
        // processQueue() {
        //     const video = document.createElement('video');
        //     video.controls = true;
        //     document.body.appendChild(video);

        //     const mediaSource = new MediaSource();
        //     video.src = URL.createObjectURL(mediaSource);
        //     mediaSource.addEventListener('sourceopen', () => {
        //         const sourceBuffer = mediaSource.addSourceBuffer('video/webm; codecs="vp8"');
        //         sourceBuffer.appendBuffer(this.queue.shift());
        //     });
        //     video.play();
        // },

    }
}
</script>

<style scoped>
.video-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 100%;
    background-color: #000;
}
</style>
