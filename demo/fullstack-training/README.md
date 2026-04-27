# 全栈工程师转型培训项目

这是一个为期8周的全栈工程师转型培训项目，旨在帮助前端开发人员掌握Golang后端开发技能，帮助后端开发人员掌握Vue3+Vite+TypeScript前端开发技能，实现真正的前后端互通能力。

## 项目结构

```
fullstack-training/
├── backend/                # Golang后端服务
│   ├── cmd/
│   │   └── main.go         # 主入口
│   ├── config/             # 配置文件
│   ├── database/           # 数据库连接
│   ├── model/              # 数据模型
│   ├── handler/            # HTTP处理器
│   ├── service/            # 业务逻辑
│   ├── middleware/         # 中间件
│   ├── utils/              # 工具函数
│   └── go.mod              # Go模块文件
├── frontend/               # Vue3前端应用
│   └── my-admin/           # 后台管理系统
│       ├── src/
│       │   ├── assets/     # 静态资源
│       │   ├── components/ # 组件
│       │   ├── layout/     # 布局组件
│       │   ├── router/     # 路由配置
│       │   ├── stores/     # 状态管理
│       │   ├── views/      # 页面组件
│       │   ├── utils/      # 工具函数
│       │   ├── api/        # API服务
│       │   └── main.ts     # 应用入口
│       ├── vite.config.ts  # Vite配置
│       ├── package.json    # 依赖配置
│       └── tsconfig.json   # TypeScript配置
└── docs/                   # 培训文档
    ├── API参考.md
    ├── 项目结构说明.md
    ├── 部署指南.md
    └── 学习指南.md
```

## 技术栈

- **前端**: Vue3 + Vite + TypeScript + Element Plus + Pinia + Axios + Vue Router
- **后端**: Golang + Gin + GORM + MySQL + JWT + bcrypt

## 启动项目

### 后端服务

1. 确保已安装Go 1.19+
2. 进入后端目录：`cd backend`
3. 初始化数据库：`mysql -u root -p -e "CREATE DATABASE fullstack_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"`
4. 启动服务：`go run cmd/main.go`

### 前端应用

1. 确保已安装Node.js 18+(推荐nvm管理node版本)
2. 进入前端目录：`cd frontend/my-admin`
3. 安装依赖：`npm install`
4. 启动开发服务器：`npm run dev`

## 功能模块

1. **登录认证**: 基于JWT的登录认证系统
2. **用户管理**: 用户的增删改查，支持角色权限
3. **菜单管理**: 动态菜单管理，支持权限控制
4. **权限控制**: 基于角色的接口和页面权限控制

## 使用说明

### 前端开发人员

- 学习Golang基础语法
- 理解Gin框架的路由和中间件
- 掌握GORM数据库操作
- 学习JWT令牌生成与验证
- 了解bcrypt密码加密

### 后端开发人员

- 学习Vue3基础语法
- 掌握TypeScript类型系统
- 理解Vite构建工具
- 学习Element Plus组件库
- 掌握Pinia状态管理
- 理解Vue Router路由配置
- 学习Axios网络请求

## 培训目标

- 前端人员能独立开发完整的Golang后端服务
- 后端人员能独立开发完整的Vue3前端应用
- 双方都能理解并实现完整的前后端协作流程
- 掌握企业级后台管理系统的开发规范

> 本项目用于全栈工程师转型培训
> 2026年4月
