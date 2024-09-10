# Sample Fiber API Application

This is a simple web application built with [Fiber](https://gofiber.io), a web framework for Go. The application provides a basic RESTful API to manage users.

## Prerequisites

- Go (1.20 or later)
- [Fiber](https://gofiber.io) framework


# API Endpoints
The following API endpoints are available:

### GET /users - Retrieve all users

```
curl -X GET http://localhost:3000/users
```

### POST /users - Create a new user
```
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{"id": "3", "name": "Ali", "age": 28}'
```

### PUT /users/:id - Update an existing user

```
curl -X PUT http://localhost:3000/users/3 \
-H "Content-Type: application/json" \
-d '{"name": "Ali", "age": 29}'
```
### DELETE /users/:id - Delete a user

```
curl -X DELETE http://localhost:3000/users/3
```
