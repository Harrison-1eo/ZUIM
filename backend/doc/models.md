# 对象说明

## User 用户

```go
package models

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
```

## Room 房间

```go
package models

type Room struct {
    gorm.Model  // 会自动添加ID,ID为主键
    Roomname string `json:"roomname"`
	
}
```

