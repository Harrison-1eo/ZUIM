<template>
    <div class="info-main-box">
        <div class="encryption-info-box">
            <div class="encryption-info-title div1">已使用ZUC256加速算法加密</div>
            <div class="encryption-info div2">HTTP数据传输密钥：<br>{{httpKey}}</div>
            <div class="encryption-info div3">WebSocket数据传输密钥：<br>{{wsKey}}</div>
            <div class="div4">
                <div class="encryption-info-icon ">
                    <el-icon><Promotion /></el-icon>
                </div>
            </div>
        </div>

        <div class="device-info-box">
            <div class="device-info div5">终端通信正常</div>
            <div class="device-info div6">与服务器的网络连接正常</div>
            <div class="device-info div7">文件系统访问正常</div>
            <div class="device-info div8">摄像头检测正常</div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'InfoMainBox',
    data() {
        return {
            httpKey: localStorage.getItem('backendPassword'),
            wsKey: localStorage.getItem('websocketBackendPassword')
        }
    },
    // method {
    //     ShowData () {

    //     }
    // }
};
</script>

<style scoped>
.info-main-box {
    height: 420px;
    width: 100%;
    margin: 15px 30px 20px 30px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.encryption-info-box, .device-info-box {
    height: 100%;
}

.encryption-info-box {
    padding: 20px;
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    background: #f9f9f9;
    border-radius: 10px;
    margin-bottom: 20px;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-template-rows: repeat(3, auto);
    row-gap: 20px;
    border: #ffffff 1px solid;
    transition: border 0.5s linear;
    word-wrap: break-word;
}

.encryption-info-box:hover {
    border: #4CAF50 1px solid;
}

.encryption-info-icon {
    height: 50px;
    width: 50px;
    color: #ffffff;
    font-size: 24px;
    background: #4CAF50;
    border-radius: 30%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.div1 {
    grid-area: 1 / 1 / 2 / 4;
    font-size: 20px;
    font-weight: bold;
    color: #333;
}

.div2{
    grid-area: 2 / 1 / 3 / 4;
    font-size: 16px;
    color: #555;
}

.div3 {
    grid-area: 3 / 1 / 4 / 4;
    font-size: 16px;
    color: #555;
}

.div4 {
    grid-area: 1 / 4 / 3 / 5;
    display: flex;
    align-items: normal;
    justify-content: flex-end;
}

.device-info-box {
    margin-top: 10px;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, auto);
    row-gap: 20px;
    column-gap: 30px;
}

.div5, .div6, .div7, .div8 {
    font-size: 20px;
    color: #4CAF50;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #ffffff;
    border: #ffffff 1px solid;
    border-radius: 10px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    padding: 10px;
    text-align: center;
    transition: border 0.5s linear, background 0.5s linear, color 0.5s linear;
}

.div5:hover, .div6:hover, .div7:hover, .div8:hover {
    border: #ffffff 1px solid;
    background: #4CAF50;
    color: #ffffff;
}

.div5 {
    grid-area: 1 / 1 / 2 / 2;
}
.div6 { grid-area: 1 / 2 / 2 / 3; }
.div7 { grid-area: 2 / 1 / 3 / 2; }
.div8 { grid-area: 2 / 2 / 3 / 3; }
</style>
