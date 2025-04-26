## App should be based on following structure

myapp/
├── cmd/                 # Main program
│   └── main.go
├── internal/
│   ├── domain/
│   │   └── user/
│   │       ├── entity/
│   │       ├── valueobject/
│   │       ├── repository/
│   │       └── service/
│   ├── application/
│   │   └── user/
│   │       └── user_service.go   # Application service
│   ├── infrastructure/
│   │   ├── persistence/
│   │   │   └── user_repository_pg.go
│   │   └── api/
│   │       └── user_handler.go
├── pkg/                  # Utilities (shared code)
├── go.mod
└── go.sum
