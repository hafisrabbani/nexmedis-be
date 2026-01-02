# =========================
# Build stage
# =========================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# install git (for go mod)
RUN apk add --no-cache git

# copy go mod first (cache friendly)
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/server

# =========================
# Runtime stage
# =========================
FROM alpine:3.19

WORKDIR /app

# ca-certificates for HTTPS / Redis / DB
RUN apk add --no-cache ca-certificates

# copy binary
COPY --from=builder /app/app /app/app

# env
ENV APP_PORT=3000

EXPOSE 3000

# run
CMD ["/app/app"]
