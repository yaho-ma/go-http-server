# 📊 Calculation Service (Go)

This project shows a **simple and extensible architecture** for building a calculation API using Go, PostgreSQL, Docker, and a layered service structure.

---

## 🏗️ Project Structure

```
├── cmd
├── internal
│   └── calculationService
│       ├── orm
│       ├── repository
│       └── service
├── db
└── handlers
```

### Architecture Overview

* **cmd** – Entry point of the application
* **internal/calculationService**

  * **orm** – Database models
  * **repository** – Database queries and persistence logic
  * **service** – Business logic layer
* **db** – Database configuration
* **handlers** – HTTP handlers (API endpoints)


---

## 🗄️ Database

This project uses **PostgreSQL** with **Docker** for containerization.

### Run PostgreSQL Container

```bash
docker run --name postgres-container \
-e POSTGRES_PASSWORD=yourpassword \
-d -p 5432:5432 postgres
```

---

## 🚀 Running the Project Locally

### 1️⃣ Clone Repository

```bash
git clone <your-repo-link>
cd <repo-name>
```

---

### 2️⃣ Install Required Libraries

```bash
go get github.com/google/uuid
go get github.com/Knetic/govaluate
go get github.com/labstack/echo/v4
go get gorm.io/driver/postgres
go get gorm.io/gorm
```

---

### 3️⃣ Start PostgreSQL in Docker

```bash
docker run --name postgres-container \
-e POSTGRES_PASSWORD=yourpassword \
-d -p 5432:5432 postgres
```

---

### 4️⃣ Run the Application

```bash
go run main.go
```

Or click **Run ▶️** if you are using GoLand IDE.

---

## 📬 Testing API

Use the Postman collection to test endpoints and calculate expressions.

👉 **Postman Collection Link:**
`https://github.com/yaho-ma/go-http-server/blob/main/Go%20Calculator%20API.postman_collection.json`



