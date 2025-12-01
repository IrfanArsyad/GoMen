# GulMen

A production-ready REST API starter kit built with Go, inspired by Laravel/Lumen architecture.

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

## Project Structure

```
gulmen/
├── app/
│   ├── controllers/     # HTTP request handlers
│   ├── middlewares/     # HTTP middlewares
│   ├── models/          # Database models
│   ├── requests/        # Request validation structs
│   ├── responses/       # Response helpers
│   └── services/        # Business logic
├── config/              # Configuration files
├── database/
│   ├── migrations/      # Database migrations
│   └── seeders/         # Database seeders
├── helpers/             # Helper functions
├── routes/              # Route definitions
├── main.go              # Application entry point
├── .env.example         # Environment template
└── README.md
```

## Requirements

- Go 1.18+
- MySQL/PostgreSQL/SQLite

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd gulmen
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Configure your database in `.env`:
```env
DB_DRIVER=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=gulmen
DB_USERNAME=root
DB_PASSWORD=your_password
```

4. Install dependencies:
```bash
go mod tidy
```

5. Run migrations:
```bash
go run main.go -migrate
```

6. (Optional) Seed database:
```bash
go run main.go -seed
```

7. Run the server:
```bash
go run main.go
```

## API Endpoints

### Health Check
```
GET /health
```

### Authentication
```
POST   /api/v1/auth/register       # Register new user
POST   /api/v1/auth/login          # Login user
GET    /api/v1/auth/profile        # Get current user profile (auth required)
PUT    /api/v1/auth/profile        # Update profile (auth required)
POST   /api/v1/auth/change-password # Change password (auth required)
POST   /api/v1/auth/refresh        # Refresh token (auth required)
```

### Users (Auth Required)
```
GET    /api/v1/users               # Get all users (paginated)
GET    /api/v1/users/:id           # Get user by ID
POST   /api/v1/users               # Create new user
PUT    /api/v1/users/:id           # Update user
DELETE /api/v1/users/:id           # Delete user
```

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

### 1. Create Model
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

### 2. Create Request
```go
// app/requests/post_request.go
package requests

type CreatePostRequest struct {
    Title   string `json:"title" validate:"required,min=3,max=255"`
    Content string `json:"content" validate:"required"`
}
```

### 3. Create Service
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

### 4. Create Controller
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

### 5. Register Routes
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

### 6. Add Migration
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
| APP_NAME | Application name | GulMen |
| APP_ENV | Environment (development/production) | development |
| APP_PORT | Server port | 8080 |
| APP_DEBUG | Debug mode | true |
| DB_DRIVER | Database driver (mysql/postgres/sqlite) | mysql |
| DB_HOST | Database host | 127.0.0.1 |
| DB_PORT | Database port | 3306 |
| DB_DATABASE | Database name | gulmen |
| DB_USERNAME | Database username | root |
| DB_PASSWORD | Database password | |
| JWT_SECRET | JWT signing secret | your-secret-key |

## License

MIT License
