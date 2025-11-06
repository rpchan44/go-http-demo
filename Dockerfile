
FROM golang:1.23-alpine AS builder

# Install build tools
RUN apk add --no-cache build-base musl-dev

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

# Build static binary (no glibc)
ENV CGO_ENABLED=1
RUN go build -ldflags="-s -w -linkmode external -extldflags '-static'" -o app .

# Export-only stage for binary extraction
FROM scratch AS export-stage
COPY --from=builder /src/app /app

# Dummy command so docker create works
ENTRYPOINT ["/app"]

