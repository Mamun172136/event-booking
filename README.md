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

## Authentication

- **POST /signup**: Register a new user.
- **POST /login**: Authenticate a user and get a token.

## Events

- **GET /events**: Retrieve a list of all events.
- **GET /events/{id}**: Retrieve details of a specific event.
- **POST /events/{id}**: Update details of a specific event.
- **PUT /events/{id}**: Modify an existing event.
- **DELETE /events/{id}**: Delete a specific event.

## Event Registration

- **POST /events/{id}/register**: Register a user for an event.
