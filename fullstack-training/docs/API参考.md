# API参考文档

## 基础信息

- **基础URL**: `http://localhost:8080/api`
- **请求格式**: JSON
- **响应格式**: JSON
- **认证方式**: JWT Bearer Token

## 通用响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

- `code`: 状态码，200表示成功，其他表示错误
- `message`: 信息描述
- `data`: 返回数据，成功时存在，失败时为空

## 认证相关API

### 登录

- **URL**: `/user/login`
- **方法**: POST
- **请求参数**:

```json
{
  "username": "string",
  "password": "string"
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "string",
    "user": {
      "id": 1,
      "username": "string",
      "role": "string"
    }
  }
}
```

- **错误响应**:

```json
{
  "code": 401,
  "message": "Invalid username or password"
}
```

## 用户管理API

### 获取用户列表

- **URL**: `/user/list`
- **方法**: GET
- **请求参数**:

```json
{
  "username": "string", // 可选，用户名模糊查询
  "page": 1, // 当前页码，默认1
  "page_size": 10 // 每页数量，默认10，最大100
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 100,
    "users": [
      {
        "id": 1,
        "username": "string",
        "role": "string",
        "created_at": "2026-04-24T10:00:00Z"
      }
    ]
  }
}
```

### 创建用户

- **URL**: `/user/add`
- **方法**: POST
- **请求参数**:

```json
{
  "username": "string",
  "password": "string",
  "role": "string" // 可选，默认为'user'
}
```

- **响应**:

```json
{
  "code": 201,
  "message": "success",
  "data": {
    "id": 1,
    "username": "string",
    "role": "string",
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 更新用户

- **URL**: `/user/edit`
- **方法**: PUT
- **请求参数**:

```json
{
  "id": 1,
  "username": "string",
  "password": "string", // 可选，如果提供则更新密码
  "role": "string" // 可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "string",
    "role": "string",
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 删除用户

- **URL**: `/user/delete/:id`
- **方法**: DELETE
- **路径参数**:
  - `id`: 用户ID

- **响应**:
  - 成功: 204 No Content
  - 失败: 500 Internal Server Error

## 菜单管理API

### 获取菜单列表

- **URL**: `/menu/list`
- **方法**: GET
- **请求参数**: 无

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "string",
      "path": "string",
      "icon": "string",
      "parent_id": 0,
      "sort": 0,
      "hidden": false,
      "role": "string",
      "created_at": "2026-04-24T10:00:00Z",
      "updated_at": "2026-04-24T10:00:00Z"
    }
  ]
}
```

### 创建菜单

- **URL**: `/menu/add`
- **方法**: POST
- **请求参数**:

```json
{
  "title": "string",
  "path": "string",
  "icon": "string", // 可选
  "parent_id": 0, // 可选，默认0
  "sort": 0, // 可选，默认0
  "hidden": false, // 可选，默认false
  "role": "string" // 可选，默认'user'
}
```

- **响应**:

```json
{
  "code": 201,
  "message": "success",
  "data": {
    "id": 1,
    "title": "string",
    "path": "string",
    "icon": "string",
    "parent_id": 0,
    "sort": 0,
    "hidden": false,
    "role": "string",
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 更新菜单

- **URL**: `/menu/edit`
- **方法**: PUT
- **请求参数**:

```json
{
  "id": 1,
  "title": "string",
  "path": "string",
  "icon": "string", // 可选
  "parent_id": 0, // 可选
  "sort": 0, // 可选
  "hidden": false, // 可选
  "role": "string" // 可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "title": "string",
    "path": "string",
    "icon": "string",
    "parent_id": 0,
    "sort": 0,
    "hidden": false,
    "role": "string",
    "created_at": "2026-04-24T10:00:00Z",
    "updated_at": "2026-04-24T10:00:00Z"
  }
}
```

### 删除菜单

- **URL**: `/menu/delete/:id`
- **方法**: DELETE
- **路径参数**:
  - `id`: 菜单ID

- **响应**:
  - 成功: 204 No Content
  - 失败: 500 Internal Server Error

## 权限说明

- **登录接口** (`/user/login`): 公开接口，无需认证
- **用户管理接口** (`/user/*`): 需要登录认证，管理员和普通用户均可访问
- **菜单管理接口** (`/menu/*`): 需要登录认证，仅管理员可访问
- **其他所有接口**: 需要登录认证

## 错误码说明

| 代码 | 说明            |
| ---- | --------------- |
| 400  | 请求参数错误    |
| 401  | 未授权/令牌无效 |
| 403  | 权限不足        |
| 500  | 服务器内部错误  |

> 本API文档生成，基于项目代码
> 2026年4月
