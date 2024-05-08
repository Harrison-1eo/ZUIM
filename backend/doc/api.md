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
    "CreatedAt": "2024-05-07T15:23:30.8072479+08:00",
    "UpdatedAt": "2024-05-07T15:23:30.8072479+08:00",
    "DeletedAt": null,
    "username": "user12345",
    "password": "password12345"
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "用户名已存在"
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
    "ID": 1,
    "token": "abc.eyJleHAiOjE.55Pvhb4Nsjv"
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "用户名或密码错误"
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
    "room_name": "room1",
    "room_info": "test room"
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
    "room_name": "room1",
    "room_info": "test room"
  }
}
```

失败返回：
```json
{
  "code": 1,
  "msg": "创建失败"
}
```

