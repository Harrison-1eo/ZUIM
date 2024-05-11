export default class WebsocketClass {
  /**
   * @description: 初始化参数
   * @param {*} url ws资源路径
   * @param token token
   * @param {*} callback 服务端信息回调
   */
  constructor(url, token, callback) {
    this.url = url + '?token=Bearer ' + token
    this.token = token
    this.callback = callback
    this.ws = null // websocket 对象
    this.status = 0 // 连接状态: 0-关闭 1-连接 2-手动关闭
    this.ping = 10000 // 心跳时长
    this.pingInterval = null // 心跳定时器
    this.reconnect = 5000 // 重连间隔
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
      this.callback(JSON.parse(e.data))
    }
    // 监听socket错误信息
    this.ws.onerror = (e) => {
      console.log(e)
    }
    // 监听socket关闭
    this.ws.onclose = (e) => {
      this.onClose(e)
    }
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
  }

  /**
   * @description: socket关闭事件
   */
  onClose(e) {
    console.error(e)
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
