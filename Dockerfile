# Dockerfile

FROM golang:1.16-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
