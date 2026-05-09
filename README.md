🔐 Auth Service (Go + Gin + PostgreSQL)

A production-style authentication service built using Go, Gin, and PostgreSQL, focusing on clean architecture, security practices, and backend engineering fundamentals.

This project is part of my journey to become a production-ready backend engineer.

🚀 Features (Current Progress)
User registration API
Secure password hashing using bcrypt
User login with password verification
PostgreSQL integration using pgx
Clean layered architecture (Handler → Service → Repository)
Proper separation of concerns
Secure authentication flow design
🧱 Tech Stack
Language: Go
Framework: Gin
Database: PostgreSQL
Driver: pgx
Security: bcrypt password hashing
🏗️ Architecture
Client → Handler → Service → Repository → Database
Layers:
🔹 Handler (HTTP Layer)
Handles request/response
Input validation
Returns HTTP status codes
🔹 Service (Business Logic)
Authentication logic
Password hashing & verification
Domain rules
🔹 Repository (Data Layer)
PostgreSQL queries
Data persistence logic only
🔐 Authentication Flow
Register User
POST /user
Receive email & password
Hash password using bcrypt
Store user in PostgreSQL
Login User
POST /login
Fetch user by email
Compare password using bcrypt
Return success or invalid credentials
🧠 Key Learnings
Why backend systems use layered architecture
Why password hashing is mandatory (bcrypt)
Why raw database errors should not be exposed to API layer
Why authentication errors must be generic (security: user enumeration prevention)
How to structure scalable Go backend projects
🔒 Security Practices
Passwords are never stored in plain text
bcrypt used for secure hashing
Generic login error responses to prevent user enumeration
Separation of internal and external error handling
📂 Project Structure
auth-service/
│
├── cmd/api
├── internal/
│   ├── config
│   ├── db
│   ├── handler
│   ├── service
│   ├── repository
│   └── routes
├── go.mod
📌 Current Status

✔ Working authentication system (register + login)
✔ Clean architecture implemented
🔄 Next: JWT authentication
🔄 Next: middleware (protected routes)
🔄 Next: deployment (Docker + VPS)

🎯 Roadmap
User registration
Login system
Password hashing
Clean architecture
JWT authentication
Middleware (auth guard)
Refresh tokens
Docker deployment
CI/CD pipeline
📈 Goal

To evolve this project into a production-grade authentication service that demonstrates real backend engineering skills used in companies.

🤝 Feedback

Open to suggestions from backend engineers on:

architecture improvements
security best practices
performance optimizations
