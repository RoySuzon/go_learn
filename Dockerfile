# ==========================================
# STAGE 1: Build Stage (Compiler & Tools)
# ==========================================
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency files first for optimal Docker layer caching
COPY go.mod ./
RUN go mod download

# Copy full source code
COPY . .

# Build lightweight, static CGO-disabled binary for production
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/rest-api ./projects/01_rest_api/main.go

# ==========================================
# STAGE 2: Production Minimal Stage (~15MB)
# ==========================================
FROM alpine:3.19

# Install ca-certificates for secure HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/rest-api .

# Expose API port
EXPOSE 8080

# Command to execute binary
ENTRYPOINT ["./rest-api"]
