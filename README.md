## 运行方式

### 1. 后端

确保已经安装了`Go 1.22`以上版本

```bash
cd backend
go mod tidy
go run ./cmd/main.go
```

后端运行端口为`8000`

### 2. 前端

确保已经安装了`Node.js 14.0`以上版本

```bash
npm install -g @vue/cli
cd frontend
npm install
npm run serve
```

前端运行端口为`8001`



