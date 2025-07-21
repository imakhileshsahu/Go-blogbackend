# 📖 Go Blog Backend

A simple and scalable blog backend built with **Go** and **Fiber**, connected to a MySQL database using GORM.  
This project provides REST APIs for user authentication, blog creation, and management.

---

## 🚀 Live Deployment

✅ **Backend URL:**  
[https://go-blog-backend-12bo.onrender.com](https://go-blog-backend-12bo.onrender.com)

---

## 🛠️ Tech Stack

- [Go](https://golang.org/) — Backend language
- [Fiber](https://gofiber.io/) — Fast HTTP web framework
- [GORM](https://gorm.io/) — ORM for database interactions
- [MySQL](https://www.mysql.com/) — Relational database
- [Render](https://render.com/) — Hosting platform

---

## 📂 Project Structure

.
├── controller/ # Handlers for API endpoints
├── database/ # DB connection logic
├── middleware/ # Authentication & other middleware
├── models/ # GORM models
├── routes/ # Route definitions
├── main.go # Entry point
└── .env # Environment variables


---

## 📋 Features

✨ User Registration & Login  
✨ JWT-based Authentication  
✨ Create, Read, Update, Delete (CRUD) blogs  
✨ MySQL database support  
✨ CORS enabled for frontend communication  
✨ Deployed & live on Render

---

## 🧪 Running Locally

### Prerequisites
- Go installed (`>=1.20`)
- MySQL running
- `git` installed

### Clone & Run

git clone https://github.com/imakhileshsahu/Go-blogbackend.git
cd Go-blogbackend

### Setup Environment

Create a .env file in the project root:

PORT=3000
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret

### Install Dependencies & Run

go mod tidy
go run main.go

Now app will start at http://localhost:3000.



## 🌐 API Endpoints

| Method | Endpoint         | Description         |
| ------ | ---------------- | ------------------- |
| POST   | `/api/register`  | Register a new user |
| POST   | `/api/login`     | Login & get JWT     |
| GET    | `/api/blogs`     | Get all blogs       |
| POST   | `/api/blogs`     | Create a new blog   |
| PUT    | `/api/blogs/:id` | Update a blog       |
| DELETE | `/api/blogs/:id` | Delete a blog       |

## 📄 License
This project is licensed under the MIT License — feel free to use & modify!

## 👨‍💻 Author
Made with ❤️ by Akhilesh Sahu
