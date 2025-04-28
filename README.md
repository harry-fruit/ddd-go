# DDD-GO

**DDD-GO** is a RESTful API built with **Go**, designed following the principles of **Domain-Driven Design (DDD)** and **Clean Architecture**.  
The project emphasizes clean code practices, separation of concerns, and the application of well-known design principles such as **SOLID**.

---

## 📚 Project Structure

```
DDD-GO/
├── cmd/                     # Main program entry point
│   └── main.go
├── config/                  # Configuration files
│   └── config.go
├── internal/                # Private application and domain logic
│   ├── application/         # Application layer (DTOs and Use Cases)
│   │   ├── dto/
│   │   │   └── product/
│   │   │       ├── create_product_dto.go
│   │   │       └── update_product_dto.go
│   │   └── use_case/
│   │       └── product/
│   │           ├── create_product_use_case.go
│   │           ├── delete_product_use_case.go
│   │           ├── get_product_by_id_use_case.go
│   │           ├── get_products_use_case.go
│   │           └── update_product_use_case.go
│   ├── domain/              # Domain layer (Entities, Repositories, Services)
│   │   ├── entity/
│   │   │   └── product_entity.go
│   │   ├── repository/
│   │   │   └── product_repository.go
│   │   └── service/
│   │       └── product_service.go
│   ├── infrastructure/      # Infrastructure layer (HTTP Handlers and Server)
│   │   └── api/
│   │       └── http/
│   │           ├── entrypoint/
│   │           ├── handler/
│   │           │   └── product_handler.go
│   │           └── router/
│   │               ├── product_route.go
│   │               └── server.go
│   └── repository/          # Shared repository interfaces
│       └── shared/
│           └── application/
│               └── use_case.go
├── pkg/                     # Utilities and shared libraries
│   └── pagination.go
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── README.md                 # Project documentation
```

---

## 🛠️ Main Concepts

- **Domain-Driven Design (DDD):** Focused on modeling the core domain of the business with expressive entities, services, and repositories.
- **Clean Architecture:** Separation between domain, application, infrastructure, and frameworks.
- **SOLID Principles:** Ensuring high-quality, maintainable, and extensible codebase.

---

## 🚀 How to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/harry-fruit/DDD-GO.git
   cd DDD-GO
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

- **cmd/**: Entry point of the application.
- **config/**: Application configuration settings.
- **internal/domain/**: Domain entities, repository interfaces, and domain services.
- **internal/application/**: DTOs and use cases that orchestrate domain logic.
- **internal/infrastructure/**: Frameworks and delivery mechanisms (HTTP handlers, routers, servers).
- **internal/repository/shared/application/**: Shared application-level interfaces.
- **pkg/**: Utility packages that can be reused across the application.

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).