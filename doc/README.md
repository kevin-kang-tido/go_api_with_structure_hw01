
---

# Go REST API with SQLite

This is a simple RESTful API built with Go (Golang), Gin, GORM, and SQLite. It follows a clean architecture with a structured MVC (Model-View-Controller) approach. The API provides CRUD (Create, Read, Update, Delete) operations for products.

## Features

- **CRUD operations** for products.
- **SQLite database** for lightweight storage.
- **GORM ORM** for easy database interaction.
- **Gin framework** for HTTP routing and middleware.
- **Modular structure** with separate layers for controllers, services, and repositories.
- **Docker support** for containerization.

## Prerequisites

- **Go** (>= 1.19)
- **SQLite** (optional, since it's included as a library)
- **Docker** (optional, for containerization)

## Project Structure

```
/your-go-rest-api
├── base
│   └── baseErorr.go
├── cmd
│   └── api
│       ├── main.go
│       └── product.db
├── config
│   └── config.go
├── controllers
│   └── product_controller.go
├── doc
│   └── README.md
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── middleware
│   └── auth_middleware.go
├── models
│   └── product.go
├── repositories
│   └── product_repository.go
├── routes
│   └── router.go
├── services
│   └── product_service.go
└── utils
    └── jwt.go

```

## Installation and Setup

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/your-go-rest-api.git
cd your-go-rest-api
```

### 2. Install Go dependencies

Run the following command to install all necessary Go modules:

```bash
go mod tidy
```

### 3. Run the Application

To run the API locally:

```bash
go run cmd/api/main.go
```

The API will start running at `http://localhost:8080`.

### 4. Run with Docker

To run the project inside a Docker container, you can use the included `Dockerfile` and `docker-compose.yaml` files.

1. Build and run the container:

```bash
docker-compose up --build
```

2. The API will be available at `http://localhost:8080`.

### 5. SQLite Database

The project uses SQLite for the database. A file named `app.db` will be created in the root of the project once the API is started. GORM will automatically handle the migration of the `Product` model.

### 6. API Endpoints

#### Product CRUD Endpoints

- **GET** `/products/` - Get all products
- **GET** `/products/:id` - Get a single product by ID
- **POST** `/products/` - Create a new product
- **PUT** `/products/:id` - Update an existing product
- **DELETE** `/products/:id` - Delete a product by ID

#### Example JSON Payload for Creating/Updating a Product:

```json
{
    "name": "Sample Product",
    "description": "This is a sample product",
    "price": 99.99
}
```

### Example API Requests

#### 1. Create a Product

```bash
curl -X POST http://localhost:8080/products/ \
-H "Content-Type: application/json" \
-d '{"name": "Sample Product", "description": "This is a sample product", "price": 99.99}'
```

#### 2. Get All Products

```bash
curl http://localhost:8080/products/
```

#### 3. Get a Product by ID

```bash
curl http://localhost:8080/products/1
```

#### 4. Update a Product

```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Product", "description": "Updated description", "price": 79.99}'
```

#### 5. Delete a Product

```bash
curl -X DELETE http://localhost:8080/products/1
```

## Technologies Used

- **Go**: Primary programming language.
- **Gin**: Web framework for routing and middleware.
- **GORM**: ORM for interacting with SQLite.
- **SQLite**: Embedded database for lightweight data storage.
- **Docker**: Containerization platform.

## Docker Setup

If you're using Docker, ensure that you have Docker and Docker Compose installed.

To start the application in a Docker container:

```bash
docker-compose up --build
```

To stop and remove the container:

```bash
docker-compose down
```

## Future Improvements

- Add JWT authentication for securing API endpoints.
- Implement pagination and filtering for the product list.
- Add more unit and integration tests.
  
## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---