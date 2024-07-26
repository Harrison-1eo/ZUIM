# 运行方式


## 1. 创建数据库

初次使用时，需要创建数据库，数据库名为`imdsb`，字符集为`utf8`，排序规则为`utf8_general_ci`。

登录到MySQL数据库，执行以下SQL语句：

```bash
mysql -u root -p

// 输入密码
```

创建数据库的SQL语句如下：

```sql
CREATE DATABASE imdsb DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE USER 'imdsbMaster'@'localhost' IDENTIFIED BY 'imdsbPassword';

GRANT ALL PRIVILEGES ON IMDSB.* TO 'imdsbMaster'@'localhost';

FLUSH PRIVILEGES;
```

运行项目时，自动使用迁移模式创建数据库表。请确保电脑中存在连接MySQL数据库的驱动。

使用如下命令安装MySQL驱动：

```bash
cd backend
go mod tidy
```


## 2. 后端

确保已经安装了`Go 1.22`以上版本

```bash
cd backend
go mod tidy
go run ./cmd/main.go
```

后端运行端口为`8000`



## 3. 前端

确保已经安装了`Node.js 14.0`以上版本

```bash
npm install -g @vue/cli
cd frontend
npm install
npm run serve
```

前端运行端口为`8001`



