// Intructions

go mod init github.com/fabiofmo/go-api-postgres-docker

// Import REST library
go get github.com/gin-gonic/gin

cd cmd
go run .

// Response test:
curl http://localhost:8080


// Subir BD no Docker
docker compose up -d go_db


// Create IMAGE on Docker
docker build -t go-api-app . 

// Subir DB e App
docker compose up -d

// docker rmi go-api-app:latest -f