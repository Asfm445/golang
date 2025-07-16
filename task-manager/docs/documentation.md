# 📘 Task Manager API Documentation

## Overview

The **Task Manager API** allows clients to perform **CRUD** operations on tasks stored in a **MongoDB** database. Tasks contain fields like `id`, `title`, `description`, `due_date`, and `status`.

## Base URL

```
http://localhost:8081
```

## Endpoints

### 1. 🔍 Get All Tasks

- **Endpoint:** `GET /tasks`
- **Description:** Retrieve all tasks.
- **Response:**
  - **Status Code:** `200 OK`
  - **Body:**
    ```json
    [
      {
        "id": "1",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2025-07-15T01:34:06.88226-07:00",
        "status": "Pending"
      },
      ...
    ]
    ```

### 2. 🔍 Get Task by ID

- **Endpoint:** `GET /tasks/:id`
- **Description:** Retrieve a task by its ID.
- **Parameters:**
  - `id` (string): Task ID to fetch.
- **Responses:**
  - **200 OK**
    ```json
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2025-07-15T01:34:06.88226-07:00",
      "status": "Pending"
    }
    ```
  - **404 Not Found**
    ```json
    {
      "message": "task not found"
    }
    ```

### 3. 🆕 Create a Task

- **Endpoint:** `POST /tasks`
- **Description:** Create a new task.
- **Request:**
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "id": "4",
      "title": "Task 4",
      "description": "Fourth task",
      "due_date": "2025-07-15T01:34:06.88226-07:00",
      "status": "Pending"
    }
    ```
- **Response:**
  - **201 Created**
    ```json
    {
      "id": "4",
      "title": "Task 4",
      "description": "Fourth task",
      "due_date": "2025-07-15T01:34:06.88226-07:00",
      "status": "Pending"
    }
    ```
  - **500 Internal Server Error**
    ```json
    {
      "message": "could not create task"
    }
    ```

### 4. 🔁 Update an Existing Task

- **Endpoint:** `PUT /tasks/:id`
- **Description:** Update an existing task.
- **Request:**
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "id": "4",
      "title": "Updated Task 4",
      "description": "Updated fourth task",
      "due_date": "2025-07-16T01:34:06.88226-07:00",
      "status": "In Progress"
    }
    ```
- **Response:**
  - **200 OK**
    ```json
    {
      "id": "4",
      "title": "Updated Task 4",
      "description": "Updated fourth task",
      "due_date": "2025-07-16T01:34:06.88226-07:00",
      "status": "In Progress"
    }
    ```
  - **404 Not Found**
    ```json
    {
      "message": "task not found"
    }
    ```

### 5. 🗑️ Delete a Task

- **Endpoint:** `DELETE /tasks/:id`
- **Description:** Delete a task by its ID.
- **Parameters:**
  - `id` (string): ID of the task to delete.
- **Response:**
  - **200 OK**
    ```json
    {
      "message": "task deleted"
    }
    ```
  - **404 Not Found**
    ```json
    {
      "message": "task not found"
    }
    ```

## 🧩 Error Responses

- **400 Bad Request**: Invalid or malformed request.
- **404 Not Found**: Task not found.
- **500 Internal Server Error**: Database error or internal issue.

## 💡 Example Usage

### ✅ Get All Tasks

```bash
curl -X GET http://localhost:8081/tasks
```

### ✅ Get Task by ID

```bash
curl -X GET http://localhost:8081/tasks/4
```

### ✅ Create a New Task

```bash
curl -X POST http://localhost:8081/tasks   -H "Content-Type: application/json"   -d '{"id": "4", "title": "Task 4", "description": "Fourth task", "due_date": "2025-07-15T01:34:06.88226-07:00", "status": "Pending"}'
```

### ✅ Update a Task

```bash
curl -X PUT http://localhost:8081/tasks/4   -H "Content-Type: application/json"   -d '{"id": "4", "title": "Updated Task 4", "description": "Updated task", "due_date": "2025-07-16T01:34:06.88226-07:00", "status": "Completed"}'
```

### ✅ Delete a Task

```bash
curl -X DELETE http://localhost:8081/tasks/4
```

## ✅ Conclusion

This API provides a RESTful interface to manage tasks. All data is persisted in MongoDB. Make sure to send well-formed JSON and handle all responses accordingly.
