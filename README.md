<<<<<<< HEAD
# ampl-assignment
=======
# Task Management API

This is a simple Task Management API built using **Golang**, **Gin**, and **SQLite**. It allows users to create, view, update, and delete tasks. The API includes both public and protected endpoints, with authentication middleware for protected routes.

---

## Table of Contents
1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Setup and Installation](#setup-and-installation)
4. [Running the Application](#running-the-application)
5. [API Endpoints](#api-endpoints)
6. [Testing](#testing)
7. [Sample Requests](#sample-requests)
8. [Technologies Used](#technologies-used)

---

## Features
- **Public Endpoints:**
  - Retrieve a list of all tasks.
- **Protected Endpoints:**
  - Create, update, and delete tasks (requires authentication).
- **Authentication:**
  - Simple API key-based authentication for protected routes.
- **Database:**
  - SQLite database for storing tasks.
- **Testing:**
  - Unit tests and integration tests for core functionality.

---

## Prerequisites
Before running the project, ensure you have the following installed:
1. **Go** (version 1.20 or higher): [Install Go](https://golang.org/doc/install)
2. **SQLite**: [Install SQLite](https://sqlite.org/download.html)
3. **Git**: [Install Git](https://git-scm.com/downloads)

---

## Setup and Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/task-management-api.git
   cd task-management-api

    Install dependencies:
    bash
    Copy

    go mod download

    Set up the database:

        The application uses SQLite, and the database will be automatically created in the project directory (tasks.db) when you run the application.

Running the Application

    Start the server:
    bash
    Copy

    go run main.go

    Verify the server is running:

        The API will be available at http://localhost:8080.

API Endpoints
Public Endpoints

    GET /public/tasks: Retrieve a list of all tasks.
    bash
    Copy

    curl -X GET http://localhost:8080/public/tasks

Protected Endpoints

    GET /tasks/{id}: Retrieve a task by its ID.
    bash
    Copy

    curl -X GET http://localhost:8080/tasks/1 \
    -H "X-API-Key: your-secret-api-key"

    POST /tasks: Create a new task.
    bash
    Copy

    curl -X POST http://localhost:8080/tasks \
    -H "X-API-Key: your-secret-api-key" \
    -H "Content-Type: application/json" \
    -d '{"title": "Task 1", "description": "Description 1", "status": "pending"}'

    PUT /tasks/{id}: Update an existing task.
    bash
    Copy

    curl -X PUT http://localhost:8080/tasks/1 \
    -H "X-API-Key: your-secret-api-key" \
    -H "Content-Type: application/json" \
    -d '{"title": "Updated Task", "description": "Updated Description", "status": "in-progress"}'

    DELETE /tasks/{id}: Delete a task.
    bash
    Copy

    curl -X DELETE http://localhost:8080/tasks/1 \
    -H "X-API-Key: your-secret-api-key"

Testing
Unit Tests

Run unit tests for the handlers:
bash
Copy

go test -v ./handlers/...

Integration Tests

Run integration tests for the API:
bash
Copy

go test -v ./tests/integration/...

Sample Requests
1. Get All Tasks (Public)
bash
Copy

curl -X GET http://localhost:8080/public/tasks

2. Create a Task (Protected)
bash
Copy

curl -X POST http://localhost:8080/tasks \
-H "X-API-Key: your-secret-api-key" \
-H "Content-Type: application/json" \
-d '{"title": "Task 1", "description": "Description 1", "status": "pending"}'

3. Update a Task (Protected)
bash
Copy

curl -X PUT http://localhost:8080/tasks/1 \
-H "X-API-Key: your-secret-api-key" \
-H "Content-Type: application/json" \
-d '{"title": "Updated Task", "description": "Updated Description", "status": "in-progress"}'

4. Delete a Task (Protected)
bash
Copy

curl -X DELETE http://localhost:8080/tasks/1 \
-H "X-API-Key: your-secret-api-key"

Technologies Used

    Golang: Backend programming language.

    Gin: Web framework for building the API.

    SQLite: Relational database for storing tasks.

    GORM: ORM for database operations.

    Testing: Go's testing package and httptest for unit and integration tests.
>>>>>>> 119e254 (project done)
