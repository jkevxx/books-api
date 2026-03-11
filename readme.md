## Books Api Project

Add this dependency

```bash
go get github.com/mattn/go-sqlite3
```

Then execute

```bash
go mod tidy

go run main.go
```

send a request

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"Go Programming", "author": "John Doe"}' http://localhost:8080/books

curl -X GET http://localhost:8080/books

curl -X GET http://localhost:8080/books/1

curl -X PUT -H "Content-Type: application/json" -d '{"title":"Go Programming", "author": "Jane Doe"}' http://localhost:8080/books/2

curl -X DELETE http://localhost:8080/books/2
```
