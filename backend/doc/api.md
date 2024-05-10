# 前后端接口说明

## 1. 用户相关

### 1.1 注册

地址：`/register`

方法：`POST`

参数：
```json
{
  "username": "user12345",
  "password": "password12345"
}
```

成功返回：
```json
{
  "code": 0,
  "msg": "注册成功",
  "data": {
    "ID": 4,
    "username": "user12345"
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "用户名已存在",
  "data": null
}
```

### 1.2 登录

地址：`/login`

方法：`POST`

参数：
```json
{
    "username": "user123",
    "password": "password123"
}
```

成功返回：
```json
{
  "code": 0,
  "msg": "登录成功",
  "data": {
    "token": "eyJhb.eyJleHO.KNzZ",
    "user": {
      "ID": 5,
      "username": "123"
    }
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "用户名或密码错误",
  "data": null
}
```

## 2. 房间相关

认证：需要在请求头中添加`Authorization`字段，值为登录成功后返回的`token`

### 2.1 创建房间

地址：`/api/room/create`

方法：`POST`

参数：
```json
{
    "name": "room1",
    "description": "test room"
}
```

成功返回：
```json
{
  "code": 0,
  "msg": "创建成功",
  "data": {
    "ID": 1,
    "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
    "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
    "DeletedAt": null,
    "name": "room1",
    "description": "test room"
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "创建失败",
  "data": null
}
```

### 2.2 获取房间列表

地址：`/api/room/list`

方法：`GET`

成功返回：
```json
{
  "code": 0,
  "msg": "获取成功",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "DeletedAt": null,
      "name": "room1",
      "description": "test room"
    },
    {
      "ID": 2,
      "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "DeletedAt": null,
      "name": "room2",
      "description": "test room"
    }
  ]
}
```

### 2.3 查看房间中的用户

地址：`/api/room/members?room_id=1`

方法：`GET`

参数：GET方法，参数填写在URL中，room_id 表示房间的ID

成功返回：
```json
{
  "code": 0,
  "msg": "获取成功",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "DeletedAt": null,
      "username": "user1"
    },
    {
      "ID": 2,
      "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
      "DeletedAt": null,
      "username": "user2"
    }
  ]
}
```


### 2.4 添加其他人到房间

地址：`/api/room/add_user`

方法：`POST`

参数：
```json
{
  "room_id": 5,
  "user_id": 3,
  "role": "member"
}
```

成功返回：
```json
{
  "code": 0,
  "msg": "添加用户到聊天室成功",
  "data": null
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "添加用户到聊天室失败，用户已在聊天室中",
  "data": null
}
```

## 3. 消息相关

### 3.1 发送消息（不用）

地址：`/api/message/send`

方法：`POST`

参数：
```json
{
  "room_id": 1,
  "type": "text",
  "content": "hello"
}
```

### 3.2 获取消息

地址：`/api/message/list`

方法：`GET`

参数：
```json
{
  "room_id": 1,
  "last_message_id": 0,
  "limit": 20
}
```

last_message_id 为 0 时，表示获取最新的消息，limit 表示获取的消息数量

last_message_id 不为 0 时，表示获取 last_message_id **之前（不含）**的消息

成功返回：
```json
{
  "code": 0,
  "msg": "历史消息获取成功",
  "data": [
    {
      "ID": 1,
      "send_time": "2024-05-07 15:23:30",
      "room_id": 1,
      "sender_id": 1,
      "sender_name": "user1",
      "type": "text",
      "content": "hello"
    },
    {
        "ID": 2,
        "send_time": "2024-05-07 15:23:30",
        "room_id": 1,
        "sender_id": 2,
        "sender_name": "user2",
        "type": "text",
        "content": "world"
    }
  ]
}
```


### 3.3 WebSocket 消息推送（未测试）

地址：`/api/message/ws`

方法：`GET`

使用 WebSocket 连接到服务器，连接成功后，服务器会返回一个消息：
```json
{
  "code": 0,
  "msg": "连接成功",
  "data": null
}
```

连接成功后，向服务器发送消息，即向房间发送消息：
```json
{
  "room_id": 1,
  "type": "text",
  "content": "hello"
}
```

服务器返回新消息，即推送给房间内的所有用户，用户收到新消息后，可以在页面上显示：
```json
{
  "code": 0,
  "msg": "新消息",
  "data": {
    "ID": 1,
    "send_time": "2024-05-07 15:23:30",
    "room_id": 1,
    "sender_id": 1,
    "sender_name": "user1",
    "type": "text",
    "content": "hello"
  }
}
```
