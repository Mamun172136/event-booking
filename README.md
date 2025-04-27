# Go Project

This is a Go project that can be compiled and executed locally using the `go build` command. It does not require Docker or installing Go globally on your system.

## Prerequisites

Before you begin, make sure that you have Go installed on your machine.

1. **Install Go**:  
   Download and install Go from the official website: [Go Downloads](https://golang.org/dl/)

2. **Verify Go installation**:  
   After installing Go, verify the installation by running the following command in your terminal and make sure that you have **Go 1.23.4** (or a compatible version) installed on your machine.:

   ```bash
   go version
   git clone <https://github.com/Mamun172136/event-booking.git>
   cd <project-directory>
   go get -u
   go run .


# API Endpoints
- **http://localhost:8080
## Authentication

- **POST /signup**: Register a new user.
- payload {"email":"test@example.com",
- "password":"test"}
- **POST /login**: Authenticate a user and get a token.

## Events

- **GET /events**: Retrieve a list of all events.
- **GET /events/:id**: Retrieve details of a specific event.
- **POST /events**: Update details of a specific event.
- authorization: ""
- payload {
  "name": " test event",
  "descripton": "A test event",
  "location": "Test location ",
  "dateTime": "2025-04-21T15:30:00z"
  }
- **PUT /events/:id**: Modify an existing event.
-  authorization: ""
 -payload {
  "name": "updatd test event",
  "descripton": "A test event",
  "location": "Test location updated",
  "dateTime": "2025-04-21T15:30:00z"
  }
- **DELETE /events/:id**: Delete a specific event.
-  authorization: ""

## Event Registration

- **POST /events/:id/register**: Register a user for an event.
