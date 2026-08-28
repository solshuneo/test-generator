FROM golang:1.26

WORKDIR /app

# cached dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "."]
