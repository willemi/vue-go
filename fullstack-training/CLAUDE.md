# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Full-stack training project for cross-skilling engineers: frontend devs learn Golang, backend devs learn Vue3+TS. It's an admin panel with JWT auth, user management, and menu management.

## Tech Stack

- **Backend**: Go 1.26 + Gin + GORM + MySQL + JWT (golang-jwt/jwt/v5) + bcrypt
- **Frontend**: Vue 3.5 + Vite 8 + TypeScript 6 + Element Plus + Pinia + Axios + Vue Router 5

## Commands

### Backend (from `backend/`)
```bash
go run cmd/main.go          # Start dev server on :8080
go build -o fullstack-backend cmd/main.go  # Production build
go mod tidy                 # Install/update dependencies
```

### Frontend (from `frontend/my-admin/`)
```bash
npm install                 # Install dependencies
npm run dev                 # Start dev server on :5173
npm run build               # Type-check + production build
```

### Database

**Windows (PowerShell):**
```powershell
mysql -u root -p -e "CREATE DATABASE fullstack_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

**macOS / Linux / Git Bash / WSL:**
```bash
mysql -u root -p -e "CREATE DATABASE fullstack_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

GORM auto-migrates tables on startup. Default admin user created via `cmd/init_admin.go`.

## Architecture

### Backend (module: `fullstack-backend`)

Layered architecture: `cmd` -> `server` -> `handler` -> `service` -> `model` + `database`

- **cmd/main.go**: Entry point, calls `config.Init()` then `server.New()`
- **server/server.go**: Wires routes and middleware. Routes split into public (`/api/user/login`) and protected (`/api/*` with AuthMiddleware). Admin-only routes use additional `AdminMiddleware`.
- **handler/**: HTTP handlers (user_handler, menu_handler) — parse request, call service, return JSON
- **service/**: Business logic (user_service, menu_service) — called by handlers
- **model/**: GORM models (User, Menu) and response helpers (ErrorResponse, SuccessResponse)
- **middleware/**: `auth.go` (JWT validation + admin check), `cors.go` (CORS)
- **utils/**: `jwt.go` (token generation/parsing), `password.go` (bcrypt hashing)
- **config/**: Global config — holds `DB *gorm.DB` and `Port`
- **database/**: `InitDB()` connects MySQL, auto-migrates models

### Frontend (app: `my-admin`)

- **src/api/user.ts**: Axios instance with request interceptor (injects Bearer token) and response interceptor (401 -> logout). All API calls defined here.
- **src/stores/user.ts**: Pinia store — token + userInfo, persisted to localStorage
- **src/router/index.ts**: Route definitions with `meta.requiresAuth` and `meta.role` for guard logic. `beforeEach` guard checks token and role.
- **src/views/**: LoginView, DashboardView, UserManagementView, MenuManagementView, NotFoundView
- **src/layout/Layout.vue**: Shell layout with sidebar + top nav + content area

### API Convention

All routes under `/api/`. Auth expects `Authorization: Bearer <token>` header. Responses use standard format from `model/response.go`.

### Dev Proxy

Vite proxies `/api` -> `http://localhost:8080` (strips `/api` prefix). Frontend API calls use `baseURL: 'http://localhost:8080/api'` directly.

## Ignored Files

Do not read, edit, or reference the following file:
- `docs/环境错误集.md`

## Key Details

- Database DSN is hardcoded in `database/database.go` (root:password@localhost:3306/fullstack_db) — must match local MySQL setup
- JWT secret is hardcoded in `utils/jwt.go` (`fullstack-secret-key-2024`), tokens expire in 24h
- User roles: `admin` or `user`. Menu roles: `admin`, `user`, or `all`
- Both User and Menu models use GORM soft delete (`DeletedAt gorm.DeletedAt`)
- No test suites currently exist
