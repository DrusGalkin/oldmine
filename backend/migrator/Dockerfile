FROM golang:1.25-alpine

WORKDIR /migrator

COPY . .

RUN go mod download

EXPOSE 8122

CMD ["go", "run", "cmd/main.go"]