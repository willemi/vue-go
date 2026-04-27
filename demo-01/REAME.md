# Vue3(Vite) + Axios + Golang(Gin) + CORS

## 整体架构（超清晰）

- 前端：Vue3 + Vite（纯客户端渲染，无 SSR）
- 后端：Golang + Gin 提供 API
- 通信：axios /fetch 调用 HTTP 接口
- 跨域：开发环境 Go 开启 CORS，生产用 Nginx 代理

## 环境准备

- Node.js 18+
- Go 1.21+
- pnpm

### 目录结构：

```plaintext
plaintext
vue3-go/
├── client/ # Vue3 前端
└── server/ # Go 后端
```

## server

- server目录下执行

### 安装依赖

```bash
# go mod tidy 会自动检查缺少的依赖
# 自动下载 gin 和 cors
# 自动生成 go.sum
go mod tidy
```

### 启动后端

```bash
go run main.go
```

#### 接口地址：http://localhost:8080/api/hello

## 执行

### 开发

```bash
# server\
go run main.go

# client\
pnpm dev
```

### 生产

```bash
# server\
go build -o app main.go

# client\
pnpm build
```

### Nginx 配置（无跨域）

```nginx
erver {
  listen 80;
  server_name your.com;

  # 前端
  location / {
    root /dist;
    try_files $uri $uri/ /index.html;
  }

  # 后端 API
  location /api {
    proxy_pass http://127.0.0.1:8080;
  }
}
```
