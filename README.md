# Go-Fiber-MongoDB-HR-System-


# Employee Management Application

This is a simple employee management application built with **Go** and the **Fiber** web framework. The application provides a RESTful API for managing employee records, allowing you to perform CRUD operations (Create, Read, Update, Delete).

## Technologies Used

- **Go**: Programming language for building the application.
- **Fiber**: Fast and minimalist web framework for Go.
- **MongoDB**: NoSQL database for storing employee data.
- **Gorilla Mux**: HTTP request router for managing routes (if used).
- **dotenv**: For loading environment variables.

## Features

- **Create Employee**: Add a new employee record.
- **Get All Employees**: Retrieve a list of all employees.
- **Get Employee**: Retrieve details of a specific employee by ID.
- **Update Employee**: Modify an existing employee's details.
- **Delete Employee**: Remove an employee record from the database.

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd employee-management
   ```

2. **Install dependencies:**
   Make sure you have Go installed on your machine, then run:
   ```bash
   go mod tidy
   ```

3. **Set up MongoDB:**
   Ensure you have a MongoDB instance running. You can use a local instance or a cloud service like MongoDB Atlas.

4. **Configure environment variables:**
   Create a `.env` file in the root directory and add your MongoDB connection string:
   ```env
   MONGODB_URI=mongodb://localhost:27017/yourdbname
   ```

## Usage

1. **Run the application:**
   ```bash
   go run main.go
   ```

2. **API Endpoints:**
   - **POST /api/employees**: Create a new employee.
     - **Request Body:**
       ```json
       {
         "name": "John Doe",
         "position": "Developer",
         "department": "Engineering",
         "salary": 60000
       }
       ```
   - **GET /api/employees**: Retrieve all employees.
   - **GET /api/employees/:id**: Retrieve an employee by ID.
   - **PUT /api/employees/:id**: Update an employee's details.
     - **Request Body:**
       ```json
       {
         "name": "John Doe",
         "position": "Senior Developer",
         "department": "Engineering",
         "salary": 80000
       }
       ```
   - **DELETE /api/employees/:id**: Delete an employee by ID.

## Example

To create a new employee, you can use a tool like **Postman** or **cURL**. Here's an example using cURL:

```bash
curl -X POST http://localhost:3000/api/employees \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "position": "Developer", "department": "Engineering", "salary": 60000}'
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to submit issues, fork the repository, and create pull requests if you would like to contribute to this project.
