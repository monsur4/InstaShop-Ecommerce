# 📚 **E-commerce API Backend**

Welcome to the **E-commerce API**, a backend application built with **Golang** and the **Gin Framework**. This API handles **user authentication**, **product management**, **order processing**, and includes **role-based access control** with **JWT authentication**.

---

## 🚀 **Features**
- ✅ **User Authentication:** JWT-based login and registration.
- ✅ **Product Management:** Admin-only CRUD operations.
- ✅ **Order Management:** Place, update, and view orders.
- ✅ **Role-Based Access Control:** User and Admin roles.
- ✅ **API Documentation:** Swagger integration.
- ✅ **Environment Configuration:** `.env` file for sensitive data.

---

## 🛠️ **Technologies Used**
- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** MySQL
- **Authentication:** JWT
- **Documentation:** Swagger
- **Environment Management:** godotenv

---

## 🚀 **Getting Started**

### ✅ **1. Prerequisites**
- **Go** (1.19+)
- **MySQL** (8.0+)
- **Git**

### ✅ **2. Clone the Repository**
```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```

### ✅ **3. Set Up Environment Variables**
Create a `.env` file in the root directory and add the following environment variables:
    
    DB_NAME=ecommerce_db
    DB_USER=root
    DB_PASSWORD=your_password
    SECRET_KEY=your_secret_key

### ✅ **4. Install Dependencies**
Ensure all necessary dependencies are installed and available in your project environment.
```bash
go mod tidy
```

### ✅ **5. Database Setup**
Log into MySQL and create the application database. Grant appropriate permissions to the database user.
```sql
CREATE DATABASE ecommerce_db;
GRANT ALL PRIVILEGES ON ecommerce_db.* TO 'root'@'localhost';
FLUSH PRIVILEGES;
```

### ✅ **6. Run Database Migrations**
Start the application to apply database migrations and set up initial database schemas.
```bash
go run main.go
```

### ✅ **7. Start the Application**
Run the application locally.
```bash
go run main.go
```
Once started, it will be available at: `http://localhost:8080`

---

## 📖 **API Documentation**
- **Swagger UI:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)  
- **Swagger JSON:** [http://localhost:8080/swagger-docs/doc.json](http://localhost:8080/swagger-docs/doc.json)

---

## 🧪 **Testing**
- Run unit and integration tests across all modules.
  ```bash
  go test ./...
  ```
- Generate a test coverage report.
  ```bash
  go test ./... -cover
  ``` 
- View the detailed coverage report in your browser.
  ```bash
  go test ./... -coverprofile=coverage.out
  ```

---

## 🛡️ **Authentication**
- Protected endpoints require a **Bearer Token** for access.  
- Include the token in the `Authorization` header as follows:
```makefile
Authorization: Bearer <token>
```

---

## 👤 **User Roles**
- **User:** Place and view orders.
- **Admin:** Manage products and orders.

---

## 🤝 **Contributing**
1. Fork the repository.  
2. Create a new feature branch.  
3. Make changes and commit them.  
4. Push your branch to the remote repository.  
5. Open a Pull Request for review.

---

## 📄 **License**
This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more details.

---

## 💬 **Support**
For feedback, questions, or support:  
- **Email:** okuniyimonsuru@yahoo.com 
- **Issues:** [Open an Issue](https://github.com/monsur4/ecommerce-api/issues)

---

**Happy Coding! 🚀✨**
