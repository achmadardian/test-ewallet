# Take home interview test

A simple e-wallet REST API

## Tech Stack

- **Go**
- **Gin**
- **PostgreSQL**
- **GORM**
- **JWT**
- **RabbitMQ**
- **Docker Compose**
- **Go Migrate**
- **Docker**
- **Docker Compose**

---

## API

| Endpoint              | Description                      |
| --------------------- | -------------------------------- |
| `POST /registers`     | Register a new user              |
| `POST /login`         | Authenticate and get JWT         |
| `POST /topup`         | Top up user balance              |
| `POST /payments`      | Make a payment                   |
| `POST /transfers`     | Transfer balance to another user |
| `GET /transactions`   | List transaction history         |
| `PUT /update-profile` | Update user profile              |
| `GET /account`        | View current user account        |

---

## Setup & Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/achmadardian/test-ewallet.git
```

### 2. Copy environment

```bash
cp .env.example .env
```

### 3. Run migrations

```bash
migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/ewallet?sslmode=disable" up
```

### 4. Run app

```bash
docker compose up -d
```

## App can be access at

```bash
http://localhost:80/
```
