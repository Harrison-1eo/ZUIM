import { userCipherWebsocketBackend, userCipherWebsocketFrontend } from '@/utils/encrypt.js';
import { ElMessage } from "element-plus";

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
                    // console.log('WS 收到的数据 ', e.data)
                    let data = JSON.parse(e.data)
                        // let data = e.data
                    if (data.code === 0) {
                        try {
                            data.data.content = userCipherWebsocketBackend.decrypt(data.en_data.data, data.en_data.position)
                                // console.log('解密后的数据', data.data.content)
                        } catch (e) {
                            console.log('解密失败：', e)
                            console.log('解密失败：', data.en_data.data, data.en_data.position)
                            ElMessage.error('解密失败，请重新登录获取密钥');
                        }
                    }
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
        /*
        data的格式为：
        {
            room_id: 0,
            type: 'text',
            content: 'hello'
         }
         将data.data进行加密
         */
        // console.log('WS 发送的数据 ', JSON.stringify(data))
        const dd = userCipherWebsocketFrontend.encrypt(data.content);
        data.content = "";
        data.en_data = {
            data: dd.cipherText,
            position: dd.position,
            mac: '',
            length: 0
        };
        /*
        发送的数据格式为：
        {
            room_id: 0,
            type: 'text',
            data: {
                data: '加密后的数据',
                position: '加密位置',
                mac: '',
                length: 0
            }
         }
         */
        // console.log('WS 发送的数据 ', data)
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
            data: ''
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