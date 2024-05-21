# 关于加密过程的说明

## 1. 模块说明

## 2. 从前端到后端的加密过程（√）

前端调用`@/utils/axios-config.js`中的`axios_config`实例进行数据通信时，会添加Token头，并对数据进行加密处理。

从前端发向后端的数据格式为：
```json
{
    "data": "JzU4la9VWuczY5FBLk7NTI2kWkZ0/5aQ5WM0M0lHZSk4NXfE+wUYpDBjxCR8TtpSk6UKDnSq2MWYbg==",
    "position": 0,
    "mac": "",
    "length": 0
}
```
其中，`data`字段为加密并进行Base64编码后的数据，`position`字段为数据的位置，`mac`字段为数据的校验码，`length`字段为数据的长度。

后端接收到数据后，会对数据进行解密处理，解密后的数据直接重新赋值给`c.Request.Body`，供后续处理。在上述例子中，解密后的数据为：
```json
{
    "username":"admin",
    "old_password":"2",
    "new_password":"2"
}
```

## 3. 从后端到前端的加密过程

后端返回数据时，会对数据进行加密处理，加密后的数据格式为：
```json
{
    "code": 1,
    "msg": "修改密码失败，原密码错误",
    "data": {},
    "en_data": {
        "data": "JzU4la9VWuczY5FBLk7NTI2kWkZ0/5aQ5WM0M0lHZSk4NXfE+wUYpDBjxCR8TtpSk6UKDnSq2MWYbg==",
        "position": 0,
        "mac": "",
        "length": 0
    }
}
```

前端接收到数据后，会对数据进行解密处理，解密后的数据直接重新赋值给`response.data`，供后续处理。在上述例子中，解密后的数据为：
```json
{
    "code": 1,
    "msg": "修改密码失败，原密码错误",
    "data": {
        "username":"admin",
        "old_password":"2",
        "new_password":"2"
    },
    "en_data": {
        "data": "JzU4la9VWuczY5FBLk7NTI2kWkZ0/5aQ5WM0M0lHZSk4NXfE+wUYpDBjxCR8TtpSk6UKDnSq2MWYbg==",
        "position": 0,
        "mac": "",
        "length": 0
    }
}
```