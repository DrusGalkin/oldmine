FROM golang:1.25-alpine

WORKDIR /gateway

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download

EXPOSE 4000

CMD ["air"]