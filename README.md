# Movies API

A simple RESTful API built with Go for managing a movie collection. This project demonstrates clean architecture with separated concerns using handlers, routes, and models.

## 📋 Features

- **CRUD Operations**: Create, Read, Update, and Delete movies
- **RESTful Design**: Standard HTTP methods (GET, POST, PUT, DELETE)
- **In-Memory Storage**: Movies stored in memory for quick prototyping
- **Clean Architecture**: Organized code structure with separate packages

## 🏗️ Project Structure

```
Movies/
├── main.go                 # Application entry point
├── routes/
│   └── routes.go          # Route definitions
├── handlers/
│   └── movie_handlers.go  # Request handlers
├── models/
│   └── movie.go           # Data models
├── go.mod                 # Go module file
└── go.sum                 # Dependency checksums
```

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- Git (optional)

### Installation

1. Clone or navigate to the project directory:
```bash
cd /Users/omartamer/Desktop/Movies
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8000`

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/movies` | Get all movies |
| `GET` | `/movies/{id}` | Get a specific movie by ID |
| `POST` | `/movies` | Create a new movie |
| `PUT` | `/movies/{id}` | Update an existing movie |
| `DELETE` | `/movies/{id}` | Delete a specific movie |
| `DELETE` | `/movies` | Delete all movies |

## 💡 Usage Examples

### Get All Movies
```bash
curl http://localhost:8000/movies
```

### Get a Specific Movie
```bash
curl http://localhost:8000/movies/1
```

### Create a New Movie
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Inception",
    "isbn": "978-3-16-148410-0",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

### Update a Movie
```bash
curl -X PUT http://localhost:8000/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Inception Updated",
    "isbn": "978-3-16-148410-0",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

### Delete a Movie
```bash
curl -X DELETE http://localhost:8000/movies/1
```

### Delete All Movies
```bash
curl -X DELETE http://localhost:8000/movies
```

## 📦 Data Model

### Movie
```json
{
  "id": "string",
  "title": "string",
  "isbn": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

## 🛠️ Technologies Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router and URL matcher

## 📝 Sample Data

The application comes with two pre-loaded movies:

1. **Catch me if you can** - Directed by Steven Spielberg
2. **Suicide Squad** - Directed by David Ayre

## 🔧 Development

### Adding New Routes

1. Add handler function in `handlers/movie_handlers.go`
2. Register route in `routes/routes.go`
3. Restart the server

### Testing

You can use tools like:
- **cURL** - Command line
- **Postman** - GUI application

## 👤 Author

Omar Tamer

---
