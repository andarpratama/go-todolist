# ToDo List Project

A simple ToDo List application built with Go and MySQL. This project allows users to create, read, and manage their tasks efficiently.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/go-todolist.git
   ```
   
2. **Change into the project directory**:
   ```bash
   cd go-todolist
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Setup the MySQL Database**:
   - Create a database named `todolist`.
   - Create a `todos` table with the following structure:
     ```sql
     CREATE TABLE todos (
         id INT AUTO_INCREMENT PRIMARY KEY,
         title VARCHAR(255) NOT NULL,
         completed BOOLEAN DEFAULT FALSE
     );
     ```

## Usage

Run the application:
```bash
go run main.go
```

The server will start running on `http://localhost:8080`.


### API Endpoints

- **Create Todo**: `POST /todos/create`
  - Request body:
    ```json
    {
      "title": "Your Todo Title"
    }
    ```
  - Response: 201 Created with a success message.

- **Get All Todos**: `GET /todos`
  - Response: JSON array of all todos.

- **Get Todo by ID**: `GET /todo?id=<todo_id>`
  - Response: JSON object of the requested todo.

## Features

- Create new todos.
- Retrieve a list of all todos.
- Get details of a specific todo by ID.

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some YourFeature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
