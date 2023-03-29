FROM golang:1.16-alpine

WORKDIR /app/cmd
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/ .
COPY . .
RUN go build -o cna-server ./cmd/*.go
CMD [ "cna-server" ]