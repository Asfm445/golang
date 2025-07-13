# 📚 Library Management System – Documentation

## 📝 Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Project Structure](#project-structure)
4. [How It Works](#how-it-works)
5. [Installation & Running](#installation--running)
6. [Usage Guide](#usage-guide)
7. [Code Design](#code-design)
8. [Future Improvements](#future-improvements)

---

## 📌 Overview

The **Library Management System** is a command-line application written in Go, designed to simulate basic library operations such as:

- Registering new members
- Logging in
- Adding and removing books
- Borrowing and returning books
- Viewing available and borrowed books

The application uses in-memory data structures to manage books and users without persistent storage.

---

## 🌟 Features

- 📚 Add / Remove Books
- 👤 Register / Login Members
- 📖 Borrow / Return Books
- 📄 List Available Books
- 📄 List Member's Borrowed Books
- 🚪 Logout / Exit

---

## 📁 Project Structure

```plaintext
library_management/
│
├── main.go                    # Entry point
├── controllers/
│   └── control.go             # User interface and menu controller
│
├── models/
│   ├── book.go                # Book struct
│   └── member.go              # Member struct and methods
│
├── services/
│   └── library.go             # Core logic of library operations
│
└── DOCUMENTATION.md           # 📄 You are here
```

---

## ⚙️ How It Works

- On start, the system prompts the user to **login or register**.
- Once logged in, the user sees a **menu** of operations.
- Each operation (add book, borrow, return, etc.) is handled by the **services layer**, with user input managed in **controllers**.
- Books and Members are stored in memory using maps.

---

## 🚀 Installation & Running

### Requirements:

- Go 1.18+

### Run the project:

```bash
go run main.go
```

> You will be prompted via the CLI for actions like registering, adding books, etc.

---

## 🧑‍💻 Usage Guide

### 🔐 Register or Login

You’ll be prompted to:

```plaintext
1. Login
2. Register
```

### 📚 Add a Book

From menu:

```plaintext
1. Enter new book
→ Enter book title and author.
```

### ❌ Remove a Book

```plaintext
2. Remove book
→ Only works if book is not currently borrowed.
```

### 📖 Borrow a Book

```plaintext
3. Borrow book
→ Requires book ID and login.
```

### 🔁 Return a Book

```plaintext
4. Return the book
→ User must be the one who borrowed it.
```

### 📄 View Lists

```plaintext
5. List available books
6. List borrowed books (only for logged-in user)
```

### 🔒 Logout or Exit

```plaintext
7. Logout
8. Exit
```

---

## 🧱 Code Design

### ✅ `models/`

- **Book**: ID, Title, Author, Status (available/borrowed)
- **Member**: ID, Name, BorrowedBooks map

### ✅ `services/`

- Core `Library` type with:
  - `store_book`: book storage map
  - `members`: member storage map
  - Functions for all logic (borrow, return, list, etc.)

### ✅ `controllers/`

- Handles input/output (menus, prompts)
- Controls app flow and calls service functions

---

## 🔮 Future Improvements

| Feature           | Benefit                                        |
| ----------------- | ---------------------------------------------- |
| JSON persistence  | Save/load books and members between sessions   |
| Admin role        | Only admin can add/remove books                |
| Borrow limit      | Prevent users from borrowing more than N books |
| Book search       | Search by title or author                      |
| Better validation | Avoid duplicates, ensure ID integrity          |
| Testing           | Add unit tests using `testing` package         |

---

## 👨‍💻 Author

**Awel Abubekar**  
Adama Science and Technology University (ASTU)  
Backend Developer | Go, Python (Django), React
