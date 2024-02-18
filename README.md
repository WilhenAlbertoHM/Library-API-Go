# Library-API-Go
## Overview
A simple library API created with Go, where we can revise all available books in the library, add a book, check in and check out books, and get books by ID. This project uses Gin, a high performance and simple web framework that is used to create APIs with Go.  

## How to Use
1. Clone the repository with: 
```bash
git clone "https://github.com/WilhenAlbertoHM/Library-API-Go.git"
```

2. Run the following command to start the server:
```bash
go run main.go
```

3. On another terminal, use the following commands to interact with the API:
To check all available books:
```bash
# Get all books
curl -X GET http://localhost:8080/books
```

To add a book:
```bash
# Add a book
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
```

To check out a book:
```bash
# Check out a book
curl -X PATCH http://localhost:8080/checkout?id=1
```

To check in a book:
```bash
# Check in a book
curl -X PATCH http://localhost:8080/books/checkin?id=1
```

To get a book by ID:
```bash
# Get a book by ID
curl -X GET http://localhost:8080/books/3
```

## Acknowledgments
This project was done with the intent to learn more about creating APIs with Go. This project was brought by Tech With Tim, and all the code was shared by him.
