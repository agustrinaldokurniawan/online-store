## How to run the project
Run on windows use wsl

### Init Go Mod Tidy
```
cd ./backend
./script/init-go-mod-tidy.sh
```

### Create docker-compose.yml file in backend folder
```
version: '4.0'
services:
  api-gateway:
    build: ./api-gateway
    ports:
      - "8080:8080"
    depends_on:
      - user-service
      - product-service
      - order-service
  
  user-service:
    build: 
      context: ./services/user-service
      dockerfile: Dockerfile
    ports:
      - "50051:50051"

  product-service:
    build: 
      context: ./services/product-service
      dockerfile: Dockerfile
    ports:
      - "50052:50052"
      
  order-service:
    build: 
      context: ./services/order-service
      dockerfile: Dockerfile
    ports:
      - "50053:50053"
  
  postgres:
    image: postgres
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: ecommerce
    ports:
      - "5432:5432"

```

### Create Dockerfile in api-gateway folder
```
# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Stage 2: Run
FROM ubuntu:22.04

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    ca-certificates \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]

# Expose port (if necessary)
EXPOSE 8080
```

### Create Dockerfile in user-service folder
```
# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o user-service ./cmd

# Stage 2: Run
FROM ubuntu:22.04

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    ca-certificates \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/user-service /app/user-service

ENTRYPOINT ["/app/user-service"]

# Expose port (if necessary)
EXPOSE 50051
```

### Create Dockerfile in product-service folder
```
# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o product-service ./cmd

# Stage 2: Run
FROM ubuntu:22.04

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    ca-certificates \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/product-service /app/product-service

ENTRYPOINT ["/app/product-service"]

# Expose port (if necessary)
EXPOSE 50052
```

### Create Dockerfile in order-service folder
```
# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o order-service ./cmd

# Stage 2: Run
FROM ubuntu:22.04

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    ca-certificates \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/order-service /app/order-service

ENTRYPOINT ["/app/order-service"]

# Expose port (if necessary)
EXPOSE 50053
```

### Run docker-compose
```
cd ./backend
./scripts/start.sh
```

### Curl to test
```
curl -x GET http://localhost:8080/v1/user/1
```

### Stop docker-compose
```
cd ./backend
./scripts/stop.sh
```

### Generate protos
```
cd ./backend
./scripts/generate-protos.sh
```

