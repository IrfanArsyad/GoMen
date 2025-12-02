# GoMen

**Lightning-Fast Go Framework with Laravel Elegance**

GoMen adalah micro-framework Go yang menggabungkan kecepatan raw performance Go dengan keanggunan dan kemudahan development ala Laravel/Lumen.

## Kenapa GoMen?

🚀 **Blazing Fast** - Dibangun di atas Go dan Gin, GoMen memberikan performa hingga 40x lebih cepat dibanding framework PHP tradisional.

🎯 **Familiar & Intuitive** - Jika kamu sudah familiar dengan Laravel/Lumen, kamu akan merasa seperti di rumah. Struktur folder, naming convention, dan workflow yang sudah kamu kenal.

⚡ **CLI Generator** - Buat model, controller, migration, dan resource hanya dengan satu perintah. Tidak perlu boilerplate berulang-ulang.

🔐 **Built-in Authentication** - JWT authentication siap pakai out-of-the-box. Register, login, dan middleware auth sudah tersedia.

🛠️ **Developer Experience First** - Hot reload, structured logging, dan error handling yang jelas membuat debugging jadi menyenangkan.

📦 **Lightweight & Minimal** - Tidak ada bloat. Hanya fitur yang kamu butuhkan, tanpa overhead yang tidak perlu.

🏗️ **Clean Architecture** - Separation of concerns yang jelas: Controllers, Services, Models, Requests, dan Middlewares terorganisir rapi.

## Perfect For

- REST API & Microservices
- Backend untuk aplikasi mobile
- High-traffic applications yang butuh performa tinggi
- Developer Laravel/Lumen yang ingin migrasi ke Go
- Startup yang butuh scalability tanpa kompleksitas

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

---

**"Write Go code with Laravel comfort. Ship faster, scale better."**
