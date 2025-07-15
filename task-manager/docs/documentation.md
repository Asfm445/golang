# Task Manager API Documentation

## Overview

The Task Manager API allows users to manage tasks with basic CRUD operations. Users can create, read, update, and delete tasks.

## Base URL

http://localhost:8080

## Endpoints

### 1. Get All Tasks

- **Endpoint:** GET /tasks
- **Description:** Retrieves a list of all tasks.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    [
    {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2025-07-15T01:34:06.88226-07:00",
    "status": "Pending"
    },
    {
    "id": "2",
    "title": "Task 2",
    "description": "Second task",
    "due_date": "2025-07-16T01:34:06.88226-07:00",
    "status": "In Progress"
    },
    {
    "id": "3",
    "title": "Task 3",
    "description": "Third task",
    "due_date": "2025-07-17T01:34:06.8827763-07:00",
    "status": "Completed"
    }
    ]

### 2. Get Task by ID

- **Endpoint:** GET /tasks/:id
- **Description:** Retrieves a specific task by its ID.
- **Parameters:**
  - id (string): The ID of the task to retrieve.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2025-07-15T01:34:06.88226-07:00",
    "status": "Pending"
    }
  - **Status Code:** 404 Not Found
  - **Body:**
    {
    "message": "task not found"
    }

### 3. Create a New Task

- **Endpoint:** POST /tasks
- **Description:** Creates a new task.
- **Request Body:**
  - **Content-Type:** application/json
  - **Body:**
    {
    "id": "4",
    "title": "Task 4",
    "description": "First task",
    "due_date": "2025-07-15T01:34:06.88226-07:00",
    "status": "Pending"
    }
- **Response:**
  - **Status Code:** 201 Created
  - **Body:**
    {
    "id": "3",
    "name": "New Task"
    }

### 4. Update an Existing Task

- **Endpoint:** PUT /tasks
- **Description:** Updates an existing task.
- **Request Body:**
  - **Content-Type:** application/json
  - **Body:**
    {
    "id": "1",
    "name": "Updated Sample Task"
    }
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    {
    "id": "1",
    "title": "updated Task 1",
    "description": "First task",
    "due_date": "2025-07-15T01:34:06.88226-07:00",
    "status": "Pending"
    }
  - **Status Code:** 404 Not Found
  - **Body:**
    {
    "message": "task not found"
    }

### 5. Delete a Task

- **Endpoint:** DELETE /tasks/:id
- **Description:** Deletes a specific task by its ID.
- **Parameters:**
  - id (string): The ID of the task to delete.
- **Response:**
  - **Status Code:** 204 No Content
  - **Body:**
    {
    "message": "task deleted"
    }
  - **Status Code:** 404 Not Found
  - **Body:**
    {
    "message": "task not found"
    }

## Error Responses

Common error responses include:

- **400 Bad Request**: Indicates that the request was malformed or invalid.
- **404 Not Found**: Indicates that the specified resource could not be found.

## Example Usage

### Get All Tasks

bash
curl -X GET http://localhost:8080/tasks

### Create a New Task

bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"id": "4","title": "Task 4","description": "Fourth task","due_date": "2025-07-15T01:34:06.88226-07:00","status": "Pending"}'

### Update an Existing Task

bash
curl -X PUT http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"id": "4","title": "Task 4","description": "Fourth task","due_date": "2025-07-15T01:34:06.88226-07:00","status": "Pending"}'

### Delete a Task

bash
curl -X DELETE http://localhost:8080/tasks/1

## Conclusion

This API provides a simple interface for managing tasks. Ensure to handle errors gracefully and validate input data when using the API.
