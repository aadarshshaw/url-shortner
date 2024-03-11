FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY index.html ./

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]
