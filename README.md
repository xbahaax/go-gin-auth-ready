# Go Gin Auth Ready

A lightweight, ready-to-use authentication backend built with Go, Gin, JWT, and GORM. Features user registration, login, and JWT-based authentication middleware with SQLite database support (easily configurable for other databases).

## Features

- 🔐 JWT-based authentication
- 📝 User registration and login
- 🛡️ Authentication middleware for protected routes
- 🗃️ SQLite database (easily switchable to PostgreSQL, MySQL, etc.)
- 🚀 Ready-to-run backend server
- 🍪 Cookie support for token storage

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/xbahaax/go-gin-auth-ready.git
cd go-gin-auth-ready
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

#### Register a new user
```http
POST /register
Content-Type: application/json

{
  "username": "your_username",
  "password": "your_password"
}
```

**Response:**
```json
{
  "message": "Registered"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "username": "your_username",
  "password": "your_password"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Access protected route
```http
GET /protected
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "message": "You are authenticated!"
}
```

## Configuration

### Database Configuration

The application uses SQLite by default. You can change the database by setting environment variables:

```bash
# For SQLite (default)
export DB_DIALECT=sqlite
export DB_DSN=auth.db

# For PostgreSQL
export DB_DIALECT=postgres
export DB_DSN="host=localhost user=username password=password dbname=mydb port=5432 sslmode=disable"

# For MySQL
export DB_DIALECT=mysql
export DB_DSN="username:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
```

### JWT Secret Key

Change the JWT secret key in `internal/auth.go`:
```go
var jwtKey = []byte("your_secret_key_here")
```

## Project Structure

```
go-fram/
├── main.go                 # Main application entry point
├── config/
│   └── config.go          # Database configuration
├── internal/
│   ├── models.go          # User model
│   └── auth.go            # Authentication handlers and middleware
├── go.mod                 # Go module file
└── README.md              # This file
```

## Usage Examples

### Using curl

1. Register a user:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

2. Login:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

3. Access protected route:
```bash
curl -X GET http://localhost:8080/protected \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Security Features

- JWT tokens with 24-hour expiration
- Bearer token authentication
- Password storage (Note: Consider adding password hashing for production)
- CORS-ready structure

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Future Enhancements

- [ ] Password hashing with bcrypt
- [ ] Refresh token support
- [ ] Rate limiting
- [ ] Email verification
- [ ] Password reset functionality
- [ ] User roles and permissions
- [ ] Database migrations
