// Intructions

go mod init github.com/fabiofmo/go-api-postgres-docker

// Import REST library
go get github.com/gin-gonic/gin

cd cmd
go run . -d

// Response test:
curl http://localhost:8000/ping


// Subir BD ou App no Docker (nao necessário, já está no compose para subirem juntos)
docker compose up -d go_db
docker compose up -d go-api-app 

// Criar IMAGE do App
docker build -t go-api-app . 

// Subir DB e App juntos
docker compose up -d

// em teste
docker run go-api-app
docker run --name=go-api-app -p 8080:8080 go-api-app

// docker rmi go-api-app:latest -f