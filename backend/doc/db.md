
# 数据库说明

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

## 2. 连接数据库

运行项目时，自动使用迁移模式创建数据库表。请确保电脑中存在连接MySQL数据库的驱动。

使用如下命令安装MySQL驱动：

```bash
cd backend
go mod tidy
```