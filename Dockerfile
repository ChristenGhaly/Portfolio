# 1. Use an official Go image
FROM golang:1.21-alpine

# 2. Set the working directory inside the container
WORKDIR /app

# Separate layer for dependencies to enable Docker cache
COPY go.mod go.sum ./
RUN go mod download

# 3. Copy your Go files into the container
COPY . .

# 4. Build your Go app
RUN go build -o main .

# 5. Expose the port your app uses (change if it's not 8080)
EXPOSE $PORT

# 6. Start the app
CMD ["./main"]