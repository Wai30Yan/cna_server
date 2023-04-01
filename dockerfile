FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# COPY cmd/ .
COPY . .
RUN go build -o ./cna-server ./cmd/main.go ./cmd/middleware.go ./cmd/routes.go
CMD [ "/app/cna-server" ]