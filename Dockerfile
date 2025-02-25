# Stage 1: Build the TypeScript files
FROM node:16 AS build-ts
WORKDIR /app

# Copy the TypeScript files and tsconfig.json
COPY static/ ./static/
COPY tsconfig.json ./

# Install TypeScript and transpile the files
RUN npm install -g typescript
RUN tsc

# Stage 2: Build the Go application
FROM golang:1.23 AS build-go
WORKDIR /app

# Copy the Go files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o executus main.go

# Stage 3: Create the final image
FROM alpine:latest
WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the built Go application
COPY --from=build-go /app/executus .

# Copy the static files
COPY --from=build-ts /app/static ./static

# Expose the port
EXPOSE 8081

# Run the Go application
CMD ["./executus"]