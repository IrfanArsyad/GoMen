# GoMen

REST API starter kit dengan Go, terinspirasi dari Laravel/Lumen.

## Installation

```bash
git clone https://github.com/IrfanArsyad/GoMen.git
cd GoMen
./install.sh
```

Script akan otomatis:
- Check Go installation
- Build CLI
- Install `gomen` ke `/usr/local/bin`

## Quick Start

```bash
cp .env.example .env        # Copy & edit config database
./bin/gomen migrate         # Create tables
./bin/gomen seed            # (Optional) Seed admin user
./bin/gomen serve           # Start server (port 8080)
```

## Database

### Migration
```bash
./bin/gomen migrate         # Jalankan semua migrations
```

### Seeding
```bash
./bin/gomen seed            # Jalankan seeder (default: admin user)
```

Default admin user setelah seed:
- Email: `admin@example.com`
- Password: `password123`

## Postman Collection

Import file `Postman.json` ke Postman untuk testing API.

## API Endpoints

### Auth (Public)
```
POST /api/v1/auth/register  - Register user
POST /api/v1/auth/login     - Login
```

### Auth (Token Required)
```
GET  /api/v1/auth/profile   - Get profile
PUT  /api/v1/auth/profile   - Update profile
POST /api/v1/auth/refresh   - Refresh token
```

### Users & Products (Token Required)
```
GET    /api/v1/users        - List all
GET    /api/v1/users/:id    - Get by ID
POST   /api/v1/users        - Create
PUT    /api/v1/users/:id    - Update
DELETE /api/v1/users/:id    - Delete
```

## Contoh Request

**Register:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com","password":"password123","password_confirm":"password123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

**Request dengan Token:**
```bash
curl http://localhost:8080/api/v1/auth/profile \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Code Generator

```bash
./bin/gomen make:resource Product     # Model + Controller + Service + Request
./bin/gomen make:model Product        # Model saja
./bin/gomen make:controller Product   # Controller saja
./bin/gomen make:service Product      # Service saja
./bin/gomen make:request Product      # Request validation saja
./bin/gomen make:migration create_x   # Migration file
./bin/gomen make:seeder Product       # Seeder file
./bin/gomen help                      # Lihat semua commands
```

## Environment (.env)

```env
APP_PORT=8080
DB_DRIVER=postgres          # postgres/mysql/sqlite
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=gomen
DB_USERNAME=root
DB_PASSWORD=secret
JWT_SECRET=your-secret-key
```

## License

MIT - [Irfan Arsyad](https://github.com/IrfanArsyad)
