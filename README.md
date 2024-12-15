
# GoProductManager-Product-Management-System-with-Async-Processing

This is a backend system for managing products, built with Golang, including asynchronous image processing, caching, and structured logging. The system uses PostgreSQL for relational data storage, Redis for caching, and RabbitMQ for message queuing.

## Setup Instructions

Follow these steps to set up and run the project on your local system.

### 1. Clone the Repository
```bash
git clone <repository-url>
cd product-management
```

### 2. Install Dependencies
Ensure you have **Go 1.19+** installed. Install project dependencies:
```bash
go mod tidy
```

### 3. Set Up PostgreSQL

#### Install PostgreSQL:

- On **Ubuntu**:
```bash
sudo apt update
sudo apt install postgresql
```

- On **Mac**:
```bash
brew install postgresql
brew services start postgresql
```

- On **Windows**: Download PostgreSQL from [PostgreSQL Official Website](https://www.postgresql.org/download/).

#### Start PostgreSQL:
```bash
sudo service postgresql start
```

#### Create the database and user:
```sql
CREATE DATABASE product_management;
CREATE USER dev_user WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE product_management TO dev_user;
```

#### Run the migrations (schema setup):
```bash
psql -U dev_user -d product_management -f migrations/init.sql
```

### 4. Set Up Redis

#### Install Redis:

- On **Ubuntu**:
```bash
sudo apt update
sudo apt install redis-server
```

- On **Mac**:
```bash
brew install redis
brew services start redis
```

- On **Windows**: Download Redis from [Redis Windows](https://github.com/microsoftarchive/redis/releases).

#### Start Redis:
```bash
redis-server
```

#### Verify Redis is running:
```bash
redis-cli ping
```
If it responds with `PONG`, Redis is running successfully.

### 5. Set Up RabbitMQ

#### Install RabbitMQ:

- On **Ubuntu**:
```bash
sudo apt update
sudo apt install rabbitmq-server
```

- On **Mac**:
```bash
brew install rabbitmq
brew services start rabbitmq
```

- On **Windows**: Download RabbitMQ from [RabbitMQ Installation](https://www.rabbitmq.com/install-windows.html).

#### Start RabbitMQ:
```bash
sudo service rabbitmq-server start
```

#### Verify RabbitMQ is running:
- Access the management UI at [http://localhost:15672](http://localhost:15672).
- Default credentials:
  - Username: `guest`
  - Password: `guest`

### 6. Configure the Project

Update the configuration file at `config/config.yaml`:
```yaml
DBUser: "dev_user"
DBPassword: "password"
DBName: "product_management"
DBHost: "localhost"
DBPort: "5432"
RabbitMQ: "amqp://guest:guest@localhost:5672/"
RedisHost: "localhost:6379"
```

### 7. Run the Project

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`.

---

## API Endpoints

### 1. Create a Product
- **URL**: `POST /api/products`
- **Request Body**:
```json
{
  "user_id": 1,
  "product_name": "Smartphone",
  "product_description": "A high-end smartphone with 128GB storage.",
  "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
  "product_price": 699.99
}
```
- **Response**:
```json
{
  "data": {
    "id": 1,
    "user_id": 1,
    "product_name": "Smartphone",
    "product_description": "A high-end smartphone with 128GB storage.",
    "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
    "product_price": 699.99,
    "created_at": "2024-12-14T00:00:00Z"
  }
}
```

### 2. Get Product by ID
- **URL**: `GET /api/products/:id`
- **Response**:
```json
{
  "data": {
    "id": 1,
    "user_id": 1,
    "product_name": "Smartphone",
    "product_description": "A high-end smartphone with 128GB storage.",
    "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
    "compressed_product_images": [],
    "product_price": 699.99
  }
}
```

---

## Testing the API

Use **Postman**, **cURL**, or any API client to test the endpoints.

### Example `cURL` Request to Create a Product:
```bash
curl -X POST http://localhost:8080/api/products \
-H "Content-Type: application/json" \
-d '{
  "user_id": 1,
  "product_name": "Smartphone",
  "product_description": "A high-end smartphone.",
  "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
  "product_price": 699.99
}'
```

---

## Future Enhancements

- Implement image compression in the RabbitMQ consumer.
- Add more filters for the `GET /api/products` endpoint.
- Include authentication and authorization using JWT.
- Add benchmark tests for API performance.

---

## Project Structure

```plaintext
product-management/
├── config/                     # Configuration files
│   └── config.go               # App configuration loader
├── controllers/                # Controllers (API handlers)
│   └── product_controller.go   # Product controller logic
├── database/                   # Database and ORM setup
│   └── database.go             # Database connection and migrations
├── models/                     # Models (DB schemas)
│   └── product.go              # Product model definition
├── services/                   # Business logic
│   └── product_service.go      # Product service (business rules)
├── middlewares/                # Middlewares (e.g., logging, error handling)
│   └── logging.go              # Logging middleware
├── rabbitmq/                   # RabbitMQ producer and consumer
│   ├── producer.go             # RabbitMQ producer logic
│   └── consumer.go             # RabbitMQ consumer logic
├── utils/                      # Utility functions
│   ├── logger.go               # Structured logger setup
│   └── response.go             # Unified API response
├── routes/                     # Route definitions
│   └── routes.go               # API route grouping
├── main.go                     # Application entry point
├── go.mod                      # Go modules file
├── go.sum                      # Dependency lockfile
└── README.md                   # Project documentation
```

