export default class WebsocketClass {
    /**
     * @description: 初始化参数
     * @param {*} url ws资源路径
     * @param token token
     * @param {*} callback 服务端信息回调
     */
    constructor(url, token, callback, ifENCRYPT = false) {
        this.url = url + '?token=Bearer ' + token
        this.token = token
        this.callback = callback
        this.ws = null // websocket 对象
        this.status = 0 // 连接状态: 0-关闭 1-连接 2-手动关闭
        this.ping = 10000 // 心跳时长
        this.pingInterval = null // 心跳定时器
        this.reconnect = 5000 // 重连间隔
        this.ifENCRYPT = ifENCRYPT
        this.textDecoder = new TextDecoder('utf-8')
    }

    /**
     * @description: 连接
     */
    connect() {
        this.ws = new WebSocket(this.url)
        // 监听socket连接
        this.ws.onopen = () => {
            this.status = 1
            this.heartHandler()
        }
        // 监听socket消息
        this.ws.onmessage = (e) => {
            if (this.ifENCRYPT) {
                console.log('WS 收到的数据 ', e.data)
                let data = JSON.parse(e.data)
                let decode = JSON.parse(this.textDecoder.decode(Uint8Array.from(atob(data.data), c => c.charCodeAt(0))));
                data.data = decode
                this.callback(data)
            } else {
                this.callback(JSON.parse(e.data))
            }
        }
        // 监听socket错误信息
        this.ws.onerror = (e) => {
            console.log(e)
        }
        // 监听socket关闭
        this.ws.onclose = (e) => {
            this.onClose(e)
        }
        console.log('WS 已连接')
    }

    /**
     * @description: 发送消息
     */
    send(data) {
        console.log('WS 发送的数据 ', JSON.stringify(data))
        return this.ws.send(JSON.stringify(data))
    }

    /**
     * @description: 关闭 weibsocket 主动关闭不会触发重连
     */
    close() {
        this.status = 2
        this.ws.close()
        clearInterval(this.pingInterval)
        console.log('WS 已关闭')
    }

    /**
     * @description: socket关闭事件
     */
    onClose(e) {
        if (this.status === 2) {
            return
        }
        console.error(e)
        console.log(this)
        this.status = this.status === 2 ? this.status : 0
        setTimeout(() => {
            if (this.status === 0) {
                this.connect()
            }
        }, this.reconnect)
    }

    /**
     * @description: 心跳机制
     */
    heartHandler() {
        const data = {
            room_id: 0,
            type: 'ping',
            data: 'ping'
        }
        this.pingInterval = setInterval(() => {
            if (this.status === 1) {
                this.ws.send(JSON.stringify(data))
            } else {
                clearInterval(this.pingInterval)
            }
        }, this.ping)
    }
}
