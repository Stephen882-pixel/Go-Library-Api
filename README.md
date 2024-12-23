# Library Management API

A simple RESTful API built with Go and Gin Framework for managing a library's book inventory. This API allows you to perform basic library operations such as adding books, checking out books, and returning books.

## Features

- Get list of all books
- Create new books
- Get book details by ID
- Checkout books (decrease quantity)
- Return books (increase quantity)
- In-memory storage (no database required)
- Error handling for invalid requests

## API Endpoints

### Get All Books
```
GET /books
```
Returns a list of all books in the library.

**Response Example:**
```json
[
    {
        "id": "1",
        "title": "In Search of Lost Time",
        "author": "Stephen Omondi",
        "quantity": 2
    },
    {
        "id": "2",
        "title": "The Great Story",
        "author": "Stephen Ondeyo",
        "quantity": 5
    }
]
```

### Create New Book
```
POST /books
```
Add a new book to the library.

**Request Body:**
```json
{
    "id": "4",
    "title": "Kigogo",
    "author": "Pauline Kea",
    "quantity": 7
}
```

### Get Book by ID
```
GET /books/:id
```
Retrieve details of a specific book by its ID.

**Response Example:**
```json
{
    "id": "2",
    "title": "The Great Story",
    "author": "Stephen Ondeyo",
    "quantity": 5
}
```

### Checkout Book
```
PATCH /checkout?id=<book_id>
```
Decrease the quantity of a book by 1 when it's checked out.

**Response Example:**
```json
{
    "id": "2",
    "title": "The Great Story",
    "author": "Stephen Ondeyo",
    "quantity": 4
}
```

### Return Book
```
PATCH /return?id=<book_id>
```
Increase the quantity of a book by 1 when it's returned.

## Error Handling

The API includes error handling for various scenarios:

- Invalid book ID
- Missing query parameters
- Book not found
- Other potential errors

**Error Response Example:**
```json
{
    "message": "Missing id query parameter."
}
```

## Technical Details

- Built with Go programming language
- Uses Gin Web Framework
- In-memory storage (no database)
- Running on `localhost:8081`

## Setup and Installation

1. Ensure you have Go installed on your system
2. Install the Gin framework:
   ```bash
   go get -u github.com/gin-gonic/gin
   ```
3. Run the application:
   ```bash
   go run main.go
   ```
4. The server will start on `localhost:8081`

## Data Structure

Books are stored using the following structure:
```go
type Book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
}
```

## Future Improvements

- Database integration
- Authentication and authorization
- Additional book metadata
- Search functionality
- Pagination for book listings
- Rate limiting
- Request validation

## Contributing

Feel free to submit issues and enhancement requests!

## License

[Add your chosen license here]
