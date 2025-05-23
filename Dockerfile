# 1. Use an official Go image
FROM golang:1.21-alpine

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy your Go files into the container
COPY . .

# 4. Build your Go app
RUN go build -o main .

# 5. Expose the port your app uses (change if it's not 8080)
EXPOSE 8080

# 6. Start the app
CMD ["./main"]