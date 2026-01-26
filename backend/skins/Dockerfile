FROM golang:1.25-alpine

WORKDIR /skins

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download

EXPOSE 8122

CMD ["air"]