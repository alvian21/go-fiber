# Go Fiber GORM PostgreSQL Project

A robust RESTful API boilerplate built with Go, Fiber, GORM, and PostgreSQL. Features automatic database migrations using `golang-migrate` and an integrated seeding system.

## 🚀 Features

- **Framework**: [Fiber v2](https://gofiber.io/) - Fast and low memory footprint.
- **ORM**: [GORM](https://gorm.io/) - Developer-friendly ORM for Go.
- **Database**: PostgreSQL.
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate) - Versioned SQL migrations.
- **Seeding**: Custom integrated seeder for initial data (e.g., admin users).
- **Testing**: Unit tests for models/entities using `testify`.
- **JWT**: Pre-configured JWT middleware for authentication.
- **File Upload**: Utilities for handling single and multiple file uploads.

## 🛠️ Prerequisites

- [Go](https://go.dev/dl/) (v1.24+ recommended)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (optional, for manual migration management)

## 📦 Installation & Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/alvian21/go-fiber.git
   cd go-fiber
   ```

2. **Setup environment variables**:
   Copy `.env.example` to `.env` and update the database credentials:

   ```bash
   cp .env.example .env
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

## 🚀 Running the Application

Simply run:

```bash
go run main.go
```

The application will automatically:

1. Connect to PostgreSQL.
2. Run any pending **Migrations** from `database/migration/sql`.
3. Run **Seeders** to populate initial data if not already present.
4. Start the server on `http://localhost:8080`.

## 🧪 Testing

Run unit tests for entities:

```bash
go test -v ./model/entity/...
```

## 🛣️ API Endpoints

| Method | Endpoint   | Description               | Auth Required |
| ------ | ---------- | ------------------------- | ------------- |
| `GET`  | `/`        | Health check / Status     | No            |
| `POST` | `/login`   | User login                | No            |
| `GET`  | `/user`    | Get all users             | Yes (JWT)     |
| `POST` | `/user`    | Create user               | No            |
| `POST` | `/book`    | Create book (with upload) | No            |
| `POST` | `/gallery` | Upload multiple photos    | No            |

## 📂 Project Structure

- `database/`: DB connection, migrations, and seeders.
- `handler/`: Request handlers (Controllers).
- `model/entity/`: GORM models and unit tests.
- `route/`: API route definitions.
- `utils/`: Helpers (JWT, Password hashing, File uploads).
- `middleware/`: Authentication and other middlewares.

## 📝 Creating New Migrations

If you have `migrate` CLI installed:

```bash
migrate create -ext sql -dir database/migration/sql -seq your_migration_name
```

Alternatively, manually create `.up.sql` and `.down.sql` files in `database/migration/sql/` following the existing numbering.
