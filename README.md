# 📝 Golang Notes API

A secure and efficient RESTful API for managing personal notes, built with **Golang**. This project implements complete **user authentication**, including **registration** and **login**, along with full CRUD functionality for user-specific notes.

---

## 🔧 Features

- ✅ **User Registration** – Create a secure account with hashed passwords.
- ✅ **User Login** – Authenticate using JWT tokens.
- ✅ **Token-Based Authentication** – Secure all endpoints using middleware.
- ✅ **Create Notes** – Add personal notes linked to your account.
- ✅ **Read Notes** – Retrieve notes specific to each authenticated user.
- ✅ **Update Notes** – Edit your existing notes.
- ✅ **Delete Notes** – Remove notes securely.
- ✅ **Built Entirely in Go** – Using Fiber, GORM, and MySQL.

---

## 📦 Tech Stack

- **Language**: Go (Golang)
- **Framework**: [Fiber](https://gofiber.io/) – Fast HTTP web framework
- **Database**: MySQL with [GORM](https://gorm.io/) ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Environment Config**: `.env` file for secrets and DB settings

---

## 🚀 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/Agrim236/Golang
cd Golang
