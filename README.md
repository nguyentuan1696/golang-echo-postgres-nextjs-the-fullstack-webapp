# Golang Echo API Template 

A production-ready template for building scalable REST APIs using Golang Echo framework with dependency injection principles.

## Architecture Overview
- Three-Tier Architecture:
  - API Layer - Echo handlers for HTTP request/response handling
  - Business Layer - Services for business logic and data processing
  - Data Layer - Repositories for database operations
- Dependency Injection with Wire for better testability and maintainability
- Interface-based Design for loose coupling
- Feature-based Package Organization
- Middleware-based Request Processing
  - Authentication
  - Request Validation
  - Error Handling
  - Logging
- Centralized Configuration Management
- Database Integration:
  - PostgreSQL for persistent storage
  - Redis for caching and search functionality

## Core Features
- Graceful Shutdown Handling
- API Documentation with Swagger
- Structured Logging
- Centralized Error Handling
- Request Validation
- Database Migrations

## Technical Stack

### Backend (Golang)
- Echo Framework - High performance web framework
- Wire - Compile-time Dependency Injection
- SQLX - Database Access Library
- PostgreSQL - Primary Database
- Log - Structured Logging
- Validator - Request Validation
- Swagger - API Documentation
- Docker - Containerization

### Frontend (Next.js)
- Next.js 13 (App Router)
- TypeScript
- Tailwind CSS
- shadcn/ui Components
- React Query
- Axios

## Project Structure
```tree
.
├── cmd/
│   └── api/                    # Application entry point
│       └── main.go            # Main application setup
├── internal/
│   ├── model/                 # Data models & DTOs
│   │   └── product.go        
│   ├── repository/            # Database operations
│   │   └── product/
│   │       ├── interface.go   # Repository interface
│   │       └── postgres.go    # PostgreSQL implementation
│   ├── service/               # Business logic
│   │   └── product/
│   │       ├── interface.go   # Service interface
│   │       └── service.go     # Service implementation
│   ├── handler/               # HTTP handlers
│   │   └── product/
│   │       └── handler.go     # Product endpoints
│   └── middleware/            # Custom middlewares
├── pkg/
│   ├── config/               # Configuration management
│   │   └── config.go
│   ├── database/             # Database connections
│   │   └── postgres.go
│   ├── logger/               # Logging setup
│   │   └── logger.go
│   ├── validator/            # Request validation
│   │   └── validator.go
│   └── server/              # HTTP server setup
│       └── server.go
├── config/                   # Configuration files
│   ├── app.yaml             # Application config
│   └── database.yaml        # Database config
└── migrations/              # Database migrations
    └── postgres/
```