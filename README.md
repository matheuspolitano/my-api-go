
© 2023 Matheus Politano. All rights reserved.
Repository: https://github.com/matheuspolitano/my-api-go


# My API in Go

This API is a simple Golang RESTful service to manage and interact with \`personalidades\` (personalities). It uses \`gorm\` as an ORM for PostgreSQL, and \`gorilla/mux\` for routing.

## Directory and File Overview

- \`controllers/controllers.go\`: Houses the main CRUD (Create, Read, Update, Delete) operations for the \`personalidades\` entity.
  
- \`database/db.go\`: Contains the database connection logic using \`gorm\`.

- \`middleware/middleware.go\`: Contains middleware functions, specifically for setting response headers.
  
- \`migration/docker-database-initial.sql\`: A SQL script to setup and seed the \`personalidades\` table in the database.

- \`models/personalidades.go\`: Contains the Go struct that represents the \`personalidades\` entity and its mapping.

- \`main.go\`: The main entry point of the application where the server starts and routes are initialized.

- \`docker-compose.yml\`: Provides configurations to spin up PostgreSQL and pgAdmin services using Docker.

## Detailed Breakdown

## Docker-Compose Setup

This project includes a \`docker-compose.yml\` file that facilitates setting up a PostgreSQL database and a pgAdmin instance for database management.

### Services

1. **postgres**: 
    - Uses the official \`postgres\` image.
    - Mapped to port \`5432\`.
    - Initializes with a script (\`docker-database-initial.sql\`) that sets up and seeds the \`personalidades\` table.

2. **pgadmin-compose**:
    - Uses the \`dpage/pgadmin4\` image for a web-based PostgreSQL admin interface.
    - Mapped to port \`54321\`.
    - Requires the \`postgres\` service to be running.

### Getting Started with Docker

1. Ensure you have Docker and Docker-Compose installed.
2. Navigate to the project directory.
3. Run the command \`docker-compose up -d\` to start the services.
4. Access the PostgreSQL database via pgAdmin by navigating to \`http://localhost:54321/\`.

## Running the Application

1. Ensure you have Go installed.
2. Clone the repository.
3. Navigate to the project directory.
4. Run the server with: \`go run main.go\`.