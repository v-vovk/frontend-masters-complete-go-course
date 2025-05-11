# Complete Go Course

This is a code repository for a workout tracking API project built in the [Complete Go](https://frontendmasters.com/courses/complete-go) course on Frontend Masters.
[![Frontend Masters](https://static.frontendmasters.com/assets/brand/logos/full.png)](https://frontendmasters.com/courses/complete-go)

## Setup

The API project is built from scratch. Before watching the course, you should install:

- [Go](https://go.dev/doc/install) (version 1.24.2 or higher)
- [Postgres](https://www.postgresql.org/download/) and any DB tool like psql or Sequel Ace to run SQL queries.
- [Docker and Docker Compose](https://www.docker.com/)

## Setup Tips

- In the [Postgres Database Container lesson][database], the Docker container exposes Postgres on the default port of `5432`. If you already have Postgres or something else running on that port and you get a connection error, you can use an alternate port but updating the `docker-compose.yml` to be something like `"5433:5432"`.
- In the [SQL Migrations with Goose lesson][goose], if you get a "command not found" error when running `goose -version`, it's because the `$HOME/go/bin` directory is not added to your `PATH`. You can fix this temporarily by running `export PATH=$HOME/go/bin:$PATH`, but this will not persist if you close your terminal. A permanent fix would require adding `export PATH=$HOME/go/bin:$PATH` to your `.zshrc` or `.bashrc`.

## Code Checkpoints

The commits on the main branch are noted in lessons throughout the course and can be used to check your code or get the latest version of the application.

- Starting code for the **Parsing Command-Line Flags** lesson
- Starting code for the **Getting Workouts By ID** lesson
- Starting code for the **Deleting Workouts** lesson
- Starting code for the **Logging & JSON Error Responses** lesson
- Starting code for the **Token Authentication & OAuth 2.0** lesson
- Starting code for the **Testing the Authentication Routes** lesson
- Final code for the course

[database]: https://frontendmasters.com/courses/complete-go/postgres-database-docker-container/
[goose]: https://rc.frontendmasters.com/courses/complete-go/sql-migrations-with-goose/

## Tests

After the `workout_store_test.go` migration is added, the test will fail due to a foreign key violation. This is becasue the tests create a workout without a `user_id`. Creating a test user for the tests will fix this issue. The `main` branch has the working tests.
