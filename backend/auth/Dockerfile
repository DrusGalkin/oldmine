FROM golang:1.25-alpine

WORKDIR /auth

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download

EXPOSE 8123

CMD ["air"]