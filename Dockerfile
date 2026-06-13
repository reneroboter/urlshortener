# ---- build stage ----
FROM golang:1.25.4 AS builder

WORKDIR /app

# Improve build speed
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-buildvcs=false"

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build using cache mounts
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o urlshortener ./cmd/main.go


# ---- runtime stage ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/urlshortener .

EXPOSE 8888

CMD ["./urlshortener"]