# Golang Flutter E-Commerce

This project is a robust REST API built with Golang, featuring a dynamic front-end integration using flutter. It supports user authentication, product and order management, and integrates additional features like payment processing, cloud storage, and database flexibility.

---

## Features

- **User Authentication**: Cookie-based JWT authentication for secure session management.
- **User Registration**: New users can register with their email and password.
- **Password Reset**: Password reset functionality via email integration.
- **Role-based Access Control**: Admin users can manage products.
- **Order Management**: Authenticated users can place orders for available products.
- **Payment Integration**: Razorpay integration for seamless payment processing.
- **Database Options**: 
  - SQLite for local testing.
  - MySQL for production environments, with easy switching between the two.
- **Cloud Storage**: Integration with Cloudinary for media management.
- **Flutter Frontend**: A mobile app front-end under development in Flutter.

---

## Tech Stack

### Backend
- **Language**: [Go](https://golang.org/)
- **Framework**: [Fiber](https://gofiber.io/) - High-performance web framework inspired by Express.js.
- **ORM**: [GORM](https://gorm.io/) - Elegant ORM for Go.
- **Authentication**: Cookie-based JWT for session management.

### Frontend
- **Flutter**: Cross-platform mobile app (under development).

### Database
- **SQLite**: Lightweight database for local testing.
- **MySQL**: Relational database for production environments.

### Payment Integration
- **Razorpay**: Seamless payment gateway for online transactions.

### Cloud Storage
- **Cloudinary**: Cloud-based media storage and delivery.

### Tools and Libraries
- **bcrypt**: Password hashing and validation.
- **dotenv**: Environment variable management.
- **golang-jwt**: Secure token-based authentication.
- **docker**: Dockerizing the Rest API.

## Setup and Installation

### Docker
1. **To build docker image**
```
    docker build -t golang-fiber-api .
```
2. **To run a docker image container**
```
    docker run --env-file .env -p 3000:3000 golang-fiber-api
```

### Prerequisites
- Go 1.19+ installed.
- SQLite (optional for local testing) or access to a MySQL database.
- Razorpay (test mode) and Cloudinary accounts for configuration.

### Steps
1. **Clone the Repository**:
   ```bash
   https://github.com/edgarAndrew/Golang-Ecommerce.git
   cd Golang-Ecommerce
   ```
2. **Setup env**:
   ```bash
   CLOUDINARY_NAME=<your cloudinary name>
   CLOUDINARY_API_KEY=<your cloudinary api key>
   CLOUDINARY_API_SECRET=<your cloudinary secret>
   JWT_SECRET=<your jwt secret key>
   JWT_EXPIRATION_HOURS=<hours>
   PORT=<your port>
   ENABLE_CORS=<true/false>
   USE_SQLITE_DB=<true/false> (set to true if you want to use a sqlite local db)
   MYSQL_URI=<your mysql uri> (mandatory if USE_SQLITE_DB is false)
   SHOW_SQL=<true/false>
   ```

3. **Start server**:
   ```bash
   go run main.go
   ```

4. **Explore endpoints**:
##### Import the postman collection json in postman