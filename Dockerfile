FROM golang:latest

WORKDIR /app

# RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .

# Expose the Fiber app port
EXPOSE 3000

# Default command
CMD ["go", "run", "cmd/main.go"]
