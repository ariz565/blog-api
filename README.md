# 📝 Go Blog REST API

A simple, modular REST API built in **Go** using the **Gin** web framework, with **JWT-based authentication**. This project is a basic blog platform with user login, and protected endpoints for managing posts.

---

## 📁 Project Structure


---

## 🚀 Features

- 🔐 User login with **JWT auth**
- ✏️ Full CRUD for blog posts (`/posts`)
- 🧱 Modular file structure
- 🧪 Basic route testing (`GET /posts`)
- ⚙️ Environment variables support (`.env`)

---

## 🧰 Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **JWT**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- **Env loader**: [joho/godotenv](https://github.com/joho/godotenv)
- **Testing**: Built-in `testing` + `httptest`

---

## ⚙️ Setup Instructions

### 1. 📥 Clone the repo

```bash
git clone https://github.com/yourusername/blog-api.git
cd blog-api


2. 📦 Initialize dependencies
Go automatically downloads dependencies defined in go.mod

go mod tidy

4. ▶️ Run the server
go run .
OR (explicitly):

bash

go run main.go app.go auth.go
Server will run at: http://localhost:8080

🔐 Authentication
Login: POST /login

Body: { "username": "admin", "password": "password" }

Returns: { "token": "..." }

Use JWT token in Authorization header for protected routes:

makefile
Authorization: <token>
📚 API Endpoints
Method	Route	Description	Auth Required
POST	/login	Get JWT token	❌
GET	/posts	List all posts	✅
GET	/posts/:id	Get a post by ID	✅
POST	/posts	Create a new post	✅
PUT	/posts/:id	Update an existing post	✅
DELETE	/posts/:id	Delete a post	✅
🧪 Run Tests
Basic test for GET /posts:

go test -v
