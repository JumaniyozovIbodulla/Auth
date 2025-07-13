# 🛡️ Authorization Service

This project is a basic yet scalable authorization service designed for internal system users such as admins, accountants, marketers, and other roles. It supports essential authentication and role-based access control (RBAC) features and is structured with simplicity and clarity in mind.

---

## 📦 Project Setup

To get started with the project, follow these steps:

1. **Clone the repository**  
   ```bash
   git clone https://github.com/yourusername/auth-service.git
   cd auth-service


2. **Install dependencies**  
   ```bash
   go mod tidy


3. **Configure environment**  
   Add a .env file (or use config.yaml, depending on your setup) with necessary configurations such as database connection, JWT secret, etc.


4. **Run the application**
    ```bash
    go run main.go



## 🧱 Project Architecture
    ```bash
  /cmd
        main.go      → App entrypoint
  /api               → HTTP handlers (routes, middleware)
  /service           → Business logic
  /storage           → Database layer
  /model             → Request/response structs and DB models
  /pkg               → Utility functions