# go-playground

```
docker-compose up -d --build
docker-compose exec golang /bin/bash
go run main.go

curl http://localhost:8080
curl http://localhost:8080/books
curl http://localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
curl http://localhost:8080/books/2
curl http://localhost:8080/books/999
```