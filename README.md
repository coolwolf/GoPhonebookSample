# Go Phonebook Application

A simple web-based phonebook application built with Go to demonstrate fundamental Go web development concepts including HTTP handling, database operations, authentication, and templating.

## 🎯 Learning Objectives

This project is designed to help developers learn:
- Go web server basics with `net/http`
- SQLite database integration
- HTML templating
- User authentication with bcrypt
- MVC architecture pattern in Go
- Middleware implementation
- Environment configuration
- Logging with structured logging

## 🏗️ Project Structure

```
phonebook/
├── main.go              # Application entry point and server setup
├── go.mod              # Go modules file
├── go.sum              # Go modules checksums
├── phonebook.db        # SQLite database file
├── db/
│   ├── db.go           # Database connection logic
│   └── schema.go       # Database schema creation
├── models/
│   ├── contact.go      # Contact model and database operations
│   └── user.go         # User model and authentication logic
├── handlers/
│   ├── auth.go         # Authentication handlers
│   ├── contact.go      # Contact CRUD handlers
│   ├── user.go         # User management handlers
│   ├── general.go      # General utility handlers
│   └── templates.go    # Template rendering utilities
├── templates/          # HTML templates
│   ├── layout.html     # Base layout template
│   ├── contacts.html   # Contact listing page
│   ├── login.html      # Login form
│   └── ...            # Other template files
└── static/
    └── css/
        └── style.css   # Application styles
```

## 🚀 Features

### Core Functionality
- **Contact Management**: Create, read, update, and delete contacts
- **User Management**: User registration and profile management
- **Authentication**: Login/logout with password hashing
- **Search**: Search contacts by name or phone number
- **Responsive UI**: Clean HTML interface with CSS styling

### Technical Features
- **Soft Delete**: Records are marked as inactive instead of being permanently deleted
- **Audit Trail**: Track who created/updated records and when
- **Session Management**: Cookie-based authentication
- **Logging**: Structured logging with different levels
- **Environment Configuration**: Development/production environment support

## 📋 Prerequisites

- Go 1.24.4 or later
- SQLite (handled by the modernc.org/sqlite driver)

## 🛠️ Installation & Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/coolwolf/GoPhonebookSample.git
   cd phonebook
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Create environment file**
   Create a `.env` file in the root directory:
   ```bash
   APP_ENV=development
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```
   or
   ```bash
   go run .
   ```

6. **Access the application**
   Open your browser and navigate to `http://localhost:8080`

## 🎓 Key Learning Concepts

### 1. HTTP Server Setup
```go
mux := http.NewServeMux()
mux.HandleFunc("/", handlers.ListContactsHandler)
http.ListenAndServe(":8080", loggingMiddleware(mux))
```

### 2. Database Operations
- Using SQLite with prepared statements
- Connection pooling with `sql.DB`
- CRUD operations with proper error handling

### 3. Template Rendering
```go
handlers.Tmpl.ExecuteTemplate(w, "main-layout", data)
```

### 4. Middleware Implementation
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Logging logic
        next.ServeHTTP(w, r)
    })
}
```

### 5. Password Security
- Using bcrypt for password hashing
- Secure password storage and verification

### 6. Session Management
- Cookie-based authentication
- User session handling

## 📚 Dependencies

- **modernc.org/sqlite**: Pure Go SQLite driver
- **golang.org/x/crypto**: For bcrypt password hashing
- **github.com/sirupsen/logrus**: Structured logging
- **github.com/joho/godotenv**: Environment variable loading

## 🔧 Configuration

The application uses environment variables for configuration:

- `APP_ENV`: Set to "development" for debug logging, "production" for warn level

## 📊 Database Schema

### Users Table
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    in_use INTEGER DEFAULT 1,
    inserted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    inserted_by INTEGER,
    updated_at DATETIME,
    updated_by INTEGER
);
```

### Contacts Table
```sql
CREATE TABLE contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    in_use INTEGER DEFAULT 1,
    inserted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    inserted_by INTEGER,
    updated_at DATETIME,
    updated_by INTEGER
);
```

## 🛣️ Routes

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/` | List all contacts (home page) |
| GET | `/contacts` | List all contacts |
| GET | `/contacts/new` | Show new contact form |
| POST | `/contacts/create` | Create new contact |
| GET | `/contacts/edit` | Show edit contact form |
| POST | `/contacts/update` | Update existing contact |
| POST | `/contacts/delete` | Delete contact (soft delete) |
| GET | `/users` | List all users |
| GET | `/users/new` | Show new user form |
| POST | `/users/create` | Create new user |
| GET | `/users/edit` | Show edit user form |
| POST | `/users/update` | Update existing user |
| POST | `/users/delete` | Delete user (soft delete) |
| GET | `/login` | Show login form |
| POST | `/dologin` | Process login |
| GET | `/logout` | Logout user |

## 🏃‍♂️ Running the Application

1. Start the server: `go run main.go`
2. The application will create the SQLite database and tables automatically
3. Access the web interface at `http://localhost:8080`
4. Create a user account to start managing contacts

## 🧪 Next Steps for Learning

To extend your Go knowledge, consider adding:

1. **Testing**: Write unit tests and integration tests
2. **API Endpoints**: Add JSON API endpoints alongside HTML interface
3. **Validation**: Implement input validation and error handling
4. **Pagination**: Add pagination for large contact lists
5. **Docker**: Containerize the application
6. **Configuration**: Advanced configuration management
7. **Graceful Shutdown**: Implement proper server shutdown handling

## 🤝 Contributing

This is a learning project! Feel free to:
- Fork the repository
- Add new features
- Improve the code structure
- Add tests
- Update documentation

## 📝 License

This project is created for educational purposes. Feel free to use it for learning and teaching Go web development.

---

**Happy Learning! 🚀**

This phonebook application demonstrates core Go web development patterns in a simple, understandable way. Each component is designed to showcase different aspects of Go programming while building a functional web application.
