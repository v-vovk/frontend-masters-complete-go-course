# Complete Go Course

This is a code repository for a workout tracking API project built in the [Complete Go](https://frontendmasters.com/courses/complete-go) course on Frontend Masters.
[![Frontend Masters](https://static.frontendmasters.com/assets/brand/logos/full.png)](https://frontendmasters.com/courses/complete-go)

Based on the [original project](https://github.com/Melkeydev/fem-project-live) by [Melkeydev](https://github.com/Melkeydev).

## Setup

The API project is built from scratch. Before watching the course, you should install:

- [Go](https://go.dev/doc/install) (version 1.24.2 or higher)
- [Postgres](https://www.postgresql.org/download/) and any DB tool like psql or Sequel Ace to run SQL queries.
- [Docker and Docker Compose](https://www.docker.com/)

## Setup Tips

- In the [Postgres Database Container lesson](https://frontendmasters.com/courses/complete-go/postgres-database-docker-container/), the Docker container exposes Postgres on the default port of `5432`. If you already have Postgres or something else running on that port and you get a connection error, you can use an alternate port by updating the `docker-compose.yml` to be something like `"5433:5432"`.
- In the [SQL Migrations with Goose lesson](https://rc.frontendmasters.com/courses/complete-go/sql-migrations-with-goose/), if you get a "command not found" error when running `goose -version`, it's because the `$HOME/go/bin` directory is not added to your `PATH`. You can fix this temporarily by running `export PATH=$HOME/go/bin:$PATH`, but this will not persist if you close your terminal. A permanent fix would require adding `export PATH=$HOME/go/bin:$PATH` to your `.zshrc` or `.bashrc`.

## Project Structure

```
/
├── main.go              # Entry point for the application
├── docker-compose.yml   # Docker configuration for Postgres
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── internal/            # Internal application packages
│   ├── api/             # API handlers
│   │   ├── token_handler.go
│   │   ├── user_handler.go
│   │   └── workout_handler.go
│   ├── app/             # Application setup
│   │   └── app.go
│   ├── middleware/      # HTTP middleware
│   │   └── middleware.go
│   ├── routes/          # API routes
│   │   └── routes.go
│   ├── store/           # Data store interfaces and implementations
│   │   ├── database.go
│   │   ├── tokens.go
│   │   ├── user_store.go
│   │   ├── workout_store.go
│   │   └── workout_store_test.go
│   ├── tokens/          # Token handling
│   │   └── tokens.go
│   └── utils/           # Utility functions
│       └── utils.go
├── migrations/          # Database migrations
│   ├── 00001_users.sql
│   ├── 00002_workouts.sql
│   ├── 00003_workout_entries.sql
│   ├── 00004_tokens.sql
│   ├── 00005_user_id_alter.sql
│   └── fs.go
└── curl/                # Example API requests
    └── examples.txt
```

## Course Content Progression

The course follows a logical progression through building different parts of the API:

- **Parsing Command-Line Flags** - Setting up the application entry point
- **Getting Workouts By ID** - Implementing the first API endpoint
- **Deleting Workouts** - Adding CRUD operations for workouts
- **Logging & JSON Error Responses** - Improving error handling and logging
- **Token Authentication & OAuth 2.0** - Adding authentication to the API
- **Testing the Authentication Routes** - Implementing tests for the auth system

You can refer to the [original repository](https://github.com/Melkeydev/fem-project-live) for the complete final code.

## Tests

After the `workout_store_test.go` migration is added, the test will fail due to a foreign key violation. This is because the tests create a workout without a `user_id`. Creating a test user for the tests will fix this issue. The `main` branch of the original repository has the working tests.

## Common Issues and Solutions

### Chi Router Version Mismatch

If you encounter routing issues with URL parameters not being correctly extracted (e.g., "ERROR: ReadIDParam: id parameter is missing"), check that you're using consistent Chi router imports throughout your code:

```go
// Use this consistently
import "github.com/go-chi/chi/v5"

// Not this older version
// import "github.com/go-chi/chi"
```

### Goose Migration Command Not Found

If you get "command not found: goose" error when running migrations, add the Go bin directory to your PATH:

```bash
export PATH=$HOME/go/bin:$PATH
```

For a permanent solution, add this line to your .zshrc or .bashrc file.

---

#GoLang #FrontendMasters #ContinuousLearning #GopherLife
