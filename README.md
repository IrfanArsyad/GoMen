# GoMen

```
   ____       __  __
  / ___| ___ |  \/  | ___ _ __
 | |  _ / _ \| |\/| |/ _ \ '_ \
 | |_| | (_) | |  | |  __/ | | |
  \____|\___/|_|  |_|\___|_| |_|
```

A production-ready REST API starter kit built with Go, inspired by Laravel/Lumen architecture.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Code Generator (CLI)](#code-generator-cli)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Usage Examples](#usage-examples)
- [Response Format](#response-format)
- [Adding New Features](#adding-new-features)
- [Environment Variables](#environment-variables)
- [License](#license)

## Features

- **Gin Framework** - High-performance HTTP web framework
- **GORM** - ORM with support for MySQL, PostgreSQL, SQLite
- **JWT Authentication** - Secure token-based authentication
- **Request Validation** - Using go-playground/validator
- **Structured Response** - Consistent JSON API responses
- **Middlewares** - Auth, CORS, Logger, Rate Limiter, Recovery
- **Database Migrations** - Auto-migrate with GORM
- **Database Seeding** - Seed initial data
- **Environment Config** - Using .env file
- **MVC-like Architecture** - Clean code structure
- **Code Generator** - CLI tool for scaffolding

## Requirements

- Go 1.18+
- MySQL / PostgreSQL / SQLite

## Installation

### Quick Install (Recommended)

Install GoMen CLI globally with one command:

```bash
git clone https://github.com/IrfanArsyad/GoMen.git
cd GoMen
./install.sh
```

This will:
1. Check Go installation (requires Go 1.18+)
2. Build the GoMen CLI
3. Install `gomen` to `/usr/local/bin` (requires sudo)

After installation, you can use `gomen` command from anywhere:

```bash
gomen help
gomen make:controller User
gomen make:resource Product
```

### Manual Installation

#### 1. Clone the repository

```bash
git clone https://github.com/IrfanArsyad/GoMen.git
cd GoMen
```

#### 2. Copy environment file

```bash
cp .env.example .env
```

#### 3. Configure your database in `.env`

```env
DB_DRIVER=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=gomen
DB_USERNAME=root
DB_PASSWORD=your_password
```

#### 4. Install dependencies and build CLI

```bash
go mod tidy
make build
```

#### 5. Run migrations

```bash
gomen migrate
```

#### 6. (Optional) Seed database

```bash
gomen seed
```

#### 7. Run the server

```bash
gomen serve
```

## Project Structure

```
gomen/
├── app/
│   ├── controllers/     # HTTP request handlers
│   ├── middlewares/     # HTTP middlewares
│   ├── models/          # Database models
│   ├── requests/        # Request validation structs
│   ├── responses/       # Response helpers
│   └── services/        # Business logic
├── bin/                 # Compiled CLI binary
├── cmd/
│   └── gomen/            # CLI tool source code
├── config/              # Configuration files
├── database/
│   ├── migrations/      # Database migrations
│   └── seeders/         # Database seeders
├── helpers/             # Helper functions
├── internal/
│   └── generator/       # Code generator templates
├── routes/              # Route definitions
├── main.go              # Application entry point
├── Makefile             # Make commands
├── .env.example         # Environment template
└── README.md
```

## Code Generator (CLI)

GoMen menyediakan code generator CLI (`gomen`) untuk mempermudah pembuatan controller, model, migration, dan lainnya.

### Setup

Build CLI terlebih dahulu:

```bash
make build
```

Ini akan membuat binary `gomen` di folder `bin/`.

### Cara Penggunaan

Ada dua cara menggunakan generator:

#### 1. Menggunakan Make (Recommended)

```bash
make <generator> name=<Name>
```

#### 2. Menggunakan CLI langsung

```bash
./bin/gomen <command> <Name>
```

---

### Generator Commands

#### Create Controller

Membuat file controller baru di `app/controllers/`.

```bash
# Menggunakan Make
make controller name=Product

# Menggunakan CLI
./bin/gomen make:controller Product
```

**Output:** `app/controllers/product_controller.go`

---

#### Create Model

Membuat file model baru di `app/models/`.

```bash
# Menggunakan Make
make model name=Product

# Menggunakan CLI
./bin/gomen make:model Product
```

**Output:** `app/models/product.go`

---

#### Create Migration

Membuat file migration baru di `database/migrations/`.

```bash
# Menggunakan Make
make migration name=create_products_table

# Menggunakan CLI
./bin/gomen make:migration create_products_table
```

**Output:** `database/migrations/YYYYMMDDHHMMSS_create_products_table.go`

---

#### Create Service

Membuat file service baru di `app/services/`.

```bash
# Menggunakan Make
make service name=Product

# Menggunakan CLI
./bin/gomen make:service Product
```

**Output:** `app/services/product_service.go`

---

#### Create Request

Membuat file request validation baru di `app/requests/`.

```bash
# Menggunakan Make
make request name=Product

# Menggunakan CLI
./bin/gomen make:request Product
```

**Output:** `app/requests/product_request.go`

---

#### Create Middleware

Membuat file middleware baru di `app/middlewares/`.

```bash
# Menggunakan Make
make middleware name=RateLimit

# Menggunakan CLI
./bin/gomen make:middleware RateLimit
```

**Output:** `app/middlewares/rate_limit_middleware.go`

---

#### Create Seeder

Membuat file seeder baru di `database/seeders/`.

```bash
# Menggunakan Make
make seeder name=Product

# Menggunakan CLI
./bin/gomen make:seeder Product
```

**Output:** `database/seeders/product_seeder.go`

---

#### Create Resource (Full CRUD)

Membuat model, controller, service, dan request sekaligus.

```bash
# Menggunakan Make
make resource name=Product

# Menggunakan CLI
./bin/gomen make:resource Product
```

**Output:**
- `app/models/product.go`
- `app/controllers/product_controller.go`
- `app/services/product_service.go`
- `app/requests/product_request.go`

---

### Quick Reference

| Command | Deskripsi | Output |
|---------|-----------|--------|
| `make controller name=<Name>` | Buat controller | `app/controllers/<name>_controller.go` |
| `make model name=<Name>` | Buat model | `app/models/<name>.go` |
| `make migration name=<name>` | Buat migration | `database/migrations/<timestamp>_<name>.go` |
| `make service name=<Name>` | Buat service | `app/services/<name>_service.go` |
| `make request name=<Name>` | Buat request | `app/requests/<name>_request.go` |
| `make middleware name=<Name>` | Buat middleware | `app/middlewares/<name>_middleware.go` |
| `make seeder name=<Name>` | Buat seeder | `database/seeders/<name>_seeder.go` |
| `make resource name=<Name>` | Buat full resource | Model + Controller + Service + Request |

---

### Other Make Commands

| Command | Deskripsi |
|---------|-----------|
| `make build` | Build CLI tool ke `bin/gomen` |
| `make run` | Jalankan aplikasi |
| `make dev` | Jalankan dengan hot reload (perlu install [air](https://github.com/air-verse/air)) |
| `make migrate` | Jalankan database migrations |
| `make seed` | Jalankan database seeders |
| `make clean` | Bersihkan build artifacts |
| `make help` | Tampilkan semua commands |
| `make list` | Lihat semua generator commands |
| `make version` | Lihat versi CLI |

---

### CLI Commands

```bash
gomen help                    # Tampilkan bantuan
gomen list                    # Lihat semua commands
gomen version                 # Lihat versi CLI
```

## Running the Application

### Development Mode

```bash
# Standard run
gomen serve

# Hot reload (requires air)
go install github.com/air-verse/air@latest
make dev
```

### Production Build

```bash
go build -o app main.go
./app
```

## API Endpoints

### Health Check

```
GET /health
```

### Authentication

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/api/v1/auth/register` | Register new user | No |
| POST | `/api/v1/auth/login` | Login user | No |
| GET | `/api/v1/auth/profile` | Get current user profile | Yes |
| PUT | `/api/v1/auth/profile` | Update profile | Yes |
| POST | `/api/v1/auth/change-password` | Change password | Yes |
| POST | `/api/v1/auth/refresh` | Refresh token | Yes |

### Users (Auth Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/users` | Get all users (paginated) |
| GET | `/api/v1/users/:id` | Get user by ID |
| POST | `/api/v1/users` | Create new user |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

## Usage Examples

### Register

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "password_confirm": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Get Profile (with token)

```bash
curl -X GET http://localhost:8080/api/v1/auth/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get Users (paginated)

```bash
curl -X GET "http://localhost:8080/api/v1/users?page=1&per_page=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {}
}
```

### Paginated Response

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [],
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "errors": {}
}
```

## Adding New Features

### Quick Start dengan Generator

```bash
# 1. Buat resource lengkap
gomen make:resource Post

# 2. Edit model sesuai kebutuhan
# app/models/post.go

# 3. Buat migration
gomen make:migration create_posts_table

# 4. Register routes di routes/api.go

# 5. Jalankan migration
gomen migrate
```

### Manual Steps

#### 1. Create Model

```go
// app/models/post.go
package models

type Post struct {
    BaseModel
    Title   string `json:"title" gorm:"size:255;not null"`
    Content string `json:"content" gorm:"type:text"`
    UserID  uint   `json:"user_id"`
    User    User   `json:"user" gorm:"foreignKey:UserID"`
}
```

#### 2. Create Request

```go
// app/requests/post_request.go
package requests

type CreatePostRequest struct {
    Title   string `json:"title" validate:"required,min=3,max=255"`
    Content string `json:"content" validate:"required"`
}
```

#### 3. Create Service

```go
// app/services/post_service.go
package services

type PostService struct{}

func NewPostService() *PostService {
    return &PostService{}
}

func (s *PostService) Create(req *requests.CreatePostRequest) (*models.Post, error) {
    // Implementation
}
```

#### 4. Create Controller

```go
// app/controllers/post_controller.go
package controllers

type PostController struct {
    postService *services.PostService
}

func NewPostController() *PostController {
    return &PostController{
        postService: services.NewPostService(),
    }
}
```

#### 5. Register Routes

```go
// routes/api.go
func setupPostRoutes(rg *gin.RouterGroup) {
    postController := controllers.NewPostController()

    posts := rg.Group("/posts")
    posts.Use(middlewares.AuthMiddleware())
    {
        posts.GET("", postController.Index)
        posts.POST("", postController.Store)
        // ...
    }
}
```

#### 6. Add Migration

```go
// database/migrations/migrate.go
err := db.AutoMigrate(
    &models.User{},
    &models.Post{}, // Add new model
)
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_NAME` | Application name | GoMen |
| `APP_ENV` | Environment (development/production) | development |
| `APP_PORT` | Server port | 8080 |
| `APP_DEBUG` | Debug mode | true |
| `DB_DRIVER` | Database driver (mysql/postgres/sqlite) | mysql |
| `DB_HOST` | Database host | 127.0.0.1 |
| `DB_PORT` | Database port | 3306 |
| `DB_DATABASE` | Database name | gomen |
| `DB_USERNAME` | Database username | root |
| `DB_PASSWORD` | Database password | |
| `JWT_SECRET` | JWT signing secret | your-secret-key |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Author

**Irfan Arsyad** - [GitHub](https://github.com/IrfanArsyad)
