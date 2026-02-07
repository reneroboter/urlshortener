FROM golang:1.25.4

WORKDIR /app

#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /urlshortener ./cmd/main.go

EXPOSE 8888

CMD ["/urlshortener"]