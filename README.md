# 📘 Book Management REST API (Golang)

A simple REST API written in **Go** to manage book data using **in-memory storage**, supporting full **CRUD** operations, **pagination**, **sorting**, and **filtering** by criteria.  
It uses only Go's standard `net/http` library and includes asynchronous logging with **goroutines** and **channels**.

---

## ⚙️ Requirements

- Go 1.18 or higher  
- OS: Windows / Linux / macOS

---

## 📦 How to Run the App

### Option 1: Clone the Repository

```bash
git clone https://github.com/rizkyyjun/book-management-api.git
go run main.go
```

### Option 2: Using ZIP File

1. Download the ZIP file (`book-management-api.zip`)
2. Unzip/extract the file into any folder
3. Open the folder in Visual Studio Code or terminal
4. Run the app:

```bash
go run main.go
```

> ℹ️ The app will start at `http://localhost:8080`  
> 📝 The log file is saved as `log.txt` in the project folder and cleared on every app start.

---

## 📫 API Endpoints

### 📗 Create a Book

**POST** `/books`

#### Request Body

```json
{
  "title": "Book Title",
  "author": "Author Name",
  "isbn": "1234567890",
  "release_date": "2024-01-01"
}
```

#### Response

```json
{
  "message": "Book successfully created",
  "book": {
    "title": "Book Title",
    "author": "Author Name",
    "isbn": "1234567890",
    "release_date": "2024-01-01"
  }
}
```

---

### 📘 Get All Books

**GET** `/books?page=1&limit=10`

- `page`: page number (default: 1)
- `limit`: items per page (default: 10)

#### Response

```json
[
  {
    "title": "A Book",
    "author": "John Doe",
    "isbn": "1234567890",
    "release_date": "2024-01-01"
  }
]
```

---

### 🔍 Get Book by ISBN

**GET** `/books/{isbn}`

#### Response

```json
{
  "title": "A Book",
  "author": "John Doe",
  "isbn": "1234567890",
  "release_date": "2024-01-01"
}
```

#### Error

```
book with ISBN 1234567890 is not found
```

---

### ✏️ Update Book

**PUT** `/books/{isbn}`

#### Request Body

```json
{
  "title": "Updated Title",
  "author": "New Author",
  "isbn": "1234567890",
  "release_date": "2024-02-01"
}
```

#### Response

```json
{
  "message": "Book successfully updated",
  "book": {
    "title": "Updated Title",
    "author": "New Author",
    "isbn": "1234567890",
    "release_date": "2024-02-01"
  }
}
```

---

### ❌ Delete Book

**DELETE** `/books/{isbn}`

#### Response

```json
{
  "message": "Book with ISBN 1234567890 successfully deleted"
}
```

---

### 🔎 Get Books by Criteria

**POST** `/books/get-by-criteria`

#### Request Body

```json
{
  "title": "",
  "author": "",
  "isbn": "",
  "release_date": "",
  "sort_by": "release_date",
  "order": "desc",
  "page": 1,
  "limit": 5
}
```

> All fields are optional. If not provided, the default values are:  
> - `sort_by`: `"title"` (can be one of `"title"`, `"author"`, `"isbn"`, or `"release_date"`)  
> - `order`: `"asc"` (ascending order)  
> - `page`: `1`  
> - `limit`: `10`  
>  
> The endpoint filters, sorts, and paginates books based on the given criteria.


#### Response

```json
[
  {
    "title": "Go Basics",
    "author": "Alice",
    "isbn": "1111111111",
    "release_date": "2023-01-01"
  }
]
```

---

## 🧠 How It Works

- In-memory store with thread-safe `sync.RWMutex`
- Asynchronous logging with `chan string` and a background goroutine
- Criteria filtering, sorting, and ordering handled in `GetByCriteria`

---

## 📂 Folder Structure

```
book/
├── book.go
├── filter.go
├── handler.go
├── logger.go
└── store.go
go.mod
main.go
log.txt
README.md
```

---

## 📝 License

This project is open-source and available under the MIT License.

---

## 🙋 Author

Made with 💻 by Rizky Juniastiar  
GitHub: [@rizkyyjun](https://github.com/rizkyyjun)
