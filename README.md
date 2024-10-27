# Project test_go

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

### Ideal File Tree Structure with SQLC
.
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go          # Configuration management
│   ├── database
│   │   ├── db.go             # Database connection
│   │   ├── queries           # SQLC queries by domain
│   │   │   ├── auth.sql
│   │   │   ├── users.sql
│   │   │   ├── products.sql
│   │   │   └── orders.sql
│   │   └── models            # SQLC generated models
│   └── server
│       ├── handlers
│       │   ├── auth
│       │   │   ├── login.go
│       │   │   ├── logout.go
│       │   │   ├── register.go
│       │   │   └── handler.go    # Auth handler struct and constructor
│       │   ├── users
│       │   │   ├── profile.go
│       │   │   ├── settings.go
│       │   │   ├── admin.go
│       │   │   └── handler.go
│       │   ├── products
│       │   │   ├── catalog.go
│       │   │   ├── inventory.go
│       │   │   ├── pricing.go
│       │   │   └── handler.go
│       │   └── orders
│       │       ├── cart.go
│       │       ├── checkout.go
│       │       ├── history.go
│       │       └── handler.go
│       ├── middleware
│       │   ├── auth.go          # Authentication middleware
│       │   ├── logger.go        # Logging middleware
│       │   ├── cors.go          # CORS middleware
│       │   └── rate_limiter.go  # Rate limiting middleware
│       ├── routes
│       │   ├── auth.go
│       │   ├── users.go
│       │   ├── products.go
│       │   ├── orders.go
│       │   └── routes.go        # Main route setup
│       ├── types
│       │   ├── auth
│       │   │   ├── requests.go
│       │   │   └── responses.go
│       │   ├── users
│       │   │   ├── requests.go
│       │   │   └── responses.go
│       │   └── products
│       │       ├── requests.go
│       │       └── responses.go
│       ├── errors
│       │   ├── errors.go        # Custom error definitions
│       │   └── handlers.go      # Error handlers
│       ├── validation
│       │   └── validator.go     # Request validation helpers
│       └── server.go            # Server setup and configuration
├── migrations                   # Database migrations
│   ├── 000001_init_users.up.sql
│   ├── 000001_init_users.down.sql
│   ├── 000002_init_products.up.sql
│   └── 000002_init_products.down.sql
├── pkg
│   └── utils
│       ├── jwt.go              # JWT utilities
│       ├── password.go         # Password hashing utilities
│       └── validator.go        # Common validation utilities
└── sqlc.yaml                   # SQLC configuration
