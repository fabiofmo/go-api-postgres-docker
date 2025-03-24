Intrictions

go mod init github.com/fabiofmo/go-api-postgres-docker

go get github.com/gin-gonic/gin


cd cmd
go run .

response test:
curl http://localhost:8080

Run on Docker
docker build -t go-app-api . 