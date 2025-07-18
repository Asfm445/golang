# Task Manager API

A simple RESTful task manager built with **Go**, **Gin**, and **MongoDB**.  
It includes JWT-based authentication and role-based authorization (`admin` and `user` roles).

## 📦 Features

- First registered user becomes **admin**
- All other users register as **user**
- Admins can promote other users to **admin**
- JWT-based login & authentication
- Role-based access control for endpoints
- Users can view tasks; Admins can manage tasks

---

## 📁 Routes

### 🔓 Public Routes (No Auth Required)

#### `POST /register`

Register a new user.  
**Note**: The first registered user will automatically have the `admin` role.

- **Body JSON**:
  ```json
  {
    "id": "123",
    "email": "user@example.com",
    "password": "password123"
  }
  ```

#### `POST /login`

Login and receive a JWT token.

- **Body JSON**:

  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

- **Response**:
  ```json
  {
    "message": "User logged in successfully",
    "token": "<jwt_token>"
  }
  ```

---

## 🔐 Protected Routes (Auth Required)

Include JWT in the `Authorization` header as:

```
Bearer <token>
```

### 🧑 User Role

#### `GET /tasks`

Get all tasks.  
Accessible to `user` and `admin`.

#### `GET /tasks/:id`

Get a specific task by ID.

---

### 🔒 Admin Role

#### `POST /tasks`

Create a new task.

#### `PUT /tasks/:id`

Update an existing task.

#### `DELETE /tasks/:id`

Delete a task.

#### `PATCH /promote`

Promote a user to **admin** using their email.  
Only an existing **admin** can promote others.

- **Body JSON**:
  ```json
  {
    "email": "user@example.com"
  }
  ```

Example scenario:

1. `admin@example.com` registers first → becomes admin automatically.
2. `user@example.com` registers → becomes a regular user.
3. `admin@example.com` logs in and sends a PATCH `/promote` request with `user@example.com`'s email.
4. Now `user@example.com` is also an admin.

---

## 🔑 Authentication & Authorization

### JWT Payload

The token includes:

- `user_id`
- `email`
- `role`
- `exp`, `iat`

### Middleware

All protected routes use:

```go
middleware.AuthMiddleware("role")
```

This checks:

- If the token is valid
- If the user's role matches the required one (`admin`, `user`, etc.)

---

## 🧪 Admin Promotion Flow Summary

1. First user is assigned `admin` role during registration.
2. Admin can promote others using their email via `PATCH /promote`.
3. Promoted users gain access to admin routes.

---

## 📌 Technologies Used

- Go (Golang)
- Gin Web Framework
- MongoDB
- JWT (Authentication)
- bcrypt (Password hashing)

---

## 🗂 Project Structure

```
.
├── controllers/
│   └── task-controller.go
├── middleware/
│   └── auth_middleware.go
├── models/
│   └── user.go, task.go
├── data/
│   └── task_service.go, user_service.go
├── router/
│   └── router.go
├── main.go
├── docs/
│   └── documentation.md
├── go.mod
└── go.sum
```

---

## 🛡️ Security Notes

- Passwords are hashed using `bcrypt`
- JWT tokens include expiration (`exp`) and issue time (`iat`)
- Sensitive routes are protected by role-based middleware

---

## 📬 Contact

If you have any questions or want to contribute, feel free to open an issue or reach out.
