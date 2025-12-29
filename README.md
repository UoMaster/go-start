# Go Gin Clean Starter

一个开箱即用的 Go + Gin 项目脚手架，采用 Clean Architecture 架构，集成了用户认证、JWT、依赖注入等常用功能，帮助你快速启动后端开发。

## ✨ 特性

- 🚀 **Gin 框架** - 高性能 HTTP Web 框架
- 🏗️ **Clean Architecture** - 清晰的分层架构，易于维护和测试
- 💉 **依赖注入** - 使用 [samber/do](https://github.com/samber/do) 实现 IoC 容器
- 🔐 **JWT 认证** - 完整的 Access Token + Refresh Token 机制
- 🗄️ **GORM + PostgreSQL** - 强大的 ORM 支持
- 🐳 **Docker 支持** - 一键启动开发环境
- 🔄 **热重载** - 使用 Air 实现开发时代码热更新
- 🌐 **Nginx 反向代理** - 生产级部署配置

## 📁 项目结构

```
.
├── cmd/                    # 应用入口
│   └── main.go
├── config/                 # 配置管理
│   └── database.go
├── database/               # 数据库相关
│   ├── entities/           # 数据库实体
│   └── migration.go        # 数据库迁移
├── docker/                 # Docker 配置
│   ├── air/                # Air 热重载配置
│   ├── nginx/              # Nginx 配置
│   └── Dockerfile
├── middlewares/            # 中间件
│   ├── authentication.go   # JWT 认证中间件
│   └── cors.go             # CORS 跨域中间件
├── modules/                # 业务模块
│   ├── auth/               # 认证模块
│   │   ├── controller/
│   │   ├── dto/
│   │   ├── repository/
│   │   ├── service/
│   │   └── routes.go
│   └── user/               # 用户模块
│       ├── controller/
│       ├── dto/
│       ├── repository/
│       ├── service/
│       └── routes.go
├── pkg/                    # 公共包
│   ├── constants/          # 常量定义
│   ├── helpers/            # 辅助函数
│   └── utils/              # 工具函数
├── providers/              # 依赖注入提供者
├── docker-compose.yml
├── Makefile
└── go.mod
```

## 🚀 快速开始

### 环境要求

- Go 1.23+
- Docker & Docker Compose
- Make (可选)

### 1. 克隆项目

```bash
git clone <repository-url>
cd go-start
```

### 2. 配置环境变量

创建 `.env` 文件：

```env
# App
APP_NAME=go-gin-clean-starter
APP_ENV=localhost
GOLANG_PORT=8888

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=go_gin_clean_starter

# JWT
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRED=72              # Access Token 过期时间（小时）
JWT_REFRESH_EXPIRED=168     # Refresh Token 过期时间（小时）

# Nginx
NGINX_PORT=81
```

### 3. 启动服务

#### 使用 Docker（推荐）

```bash
# 首次启动（构建镜像）
make init-docker

# 后续启动
make up

# 查看日志
make logs

# 停止服务
make down
```

#### 本地开发

```bash
# 安装依赖
make dep

# 运行项目
make run
```

### 4. 数据库初始化（首次启动需要）

```bash
# 创建数据库
make create-db

# 初始化 UUID 扩展
make init-uuid
```

## 📡 API 接口

### 认证模块 `/api/auth`

| 方法 | 路径 | 描述 | 认证 |
|------|------|------|------|
| POST | `/register` | 用户注册 | ❌ |
| POST | `/login` | 用户登录 | ❌ |
| POST | `/refresh` | 刷新 Token | ❌ |

### 用户模块 `/api/user`

| 方法 | 路径 | 描述 | 认证 |
|------|------|------|------|
| GET | `/me` | 获取当前用户信息 | ✅ |
| PUT | `/me` | 更新当前用户信息 | ✅ |
| DELETE | `/me` | 删除当前用户 | ✅ |

### 请求示例

#### 注册

```bash
curl -X POST http://localhost:8888/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "telp_number": "08123456789"
  }'
```

#### 登录

```bash
curl -X POST http://localhost:8888/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

#### 获取用户信息（需要认证）

```bash
curl -X GET http://localhost:8888/api/user/me \
  -H "Authorization: Bearer <your-access-token>"
```

## 🛠️ 开发命令

```bash
# 安装/更新依赖
make dep

# 运行项目
make run

# 构建项目
make build

# 构建并运行
make run-build

# 运行测试
make test

# Docker 相关
make init-docker    # 首次启动（构建镜像）
make up             # 启动容器
make down           # 停止容器
make logs           # 查看日志

# 进入容器
make container-go       # 进入 Go 应用容器
make container-postgres # 进入 PostgreSQL 容器
```

## 🏗️ 添加新模块

1. 在 `modules/` 下创建新模块目录
2. 创建 `controller/`、`dto/`、`repository/`、`service/` 子目录
3. 在 `database/entities/` 中定义实体
4. 在 `providers/` 中注册依赖
5. 创建 `routes.go` 并在 `cmd/main.go` 中注册路由

## 📝 License

MIT License
