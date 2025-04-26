# MyApp

**MyApp** is a RESTful API built with **Go**, designed with the principles of **Domain-Driven Design (DDD)** and **Clean Architecture**.\
The project emphasizes clean code practices, separation of concerns, and the application of well-known design principles such as **SOLID**.

---

## 📚 Project Structure

```
myapp/
├── cmd/                   # Main program entry point
│   └── main.go
├── internal/              # Private application and domain logic
│   ├── domain/            # Domain layer (Entities, Value Objects, Repositories, Services)
│   │   └── user/
│   │       ├── entity/
│   │       ├── valueobject/
│   │       ├── repository/
│   │       └── service/
│   ├── application/       # Application layer (Use Cases / Services)
│   │   └── user/
│   │       └── user_service.go
│   ├── infrastructure/    # Infrastructure layer (Persistence, APIs, etc.)
│   │   ├── persistence/
│   │   │   └── user_repository_pg.go
│   │   └── api/
│   │       └── user_handler.go
├── pkg/                   # Utilities and shared libraries
├── go.mod                 # Go module definition
└── go.sum                 # Go module checksums
```

---

## 🛠️ Main Concepts

- **Domain-Driven Design (DDD):** Focus on the core business domain with a rich and expressive model.
- **Clean Architecture:** Decouple business rules from frameworks, databases, and UI.
- **SOLID Principles:** Ensure high-quality, maintainable, and scalable code.

---

## 🚀 How to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/myapp.git
   cd myapp
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run ./cmd/main.go
   ```

---

## 📂 About Each Main Folder

- ``: Entry point of the application (main executable).
- ``: Contains business rules (entities, value objects, domain services, and repository interfaces).
- ``: Application services that orchestrate use cases and business logic.
- ``: External integrations like database persistence, API handlers, etc.
- ``: Reusable libraries and helper functions across the project.

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).

