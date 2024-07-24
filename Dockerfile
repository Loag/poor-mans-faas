FROM golang:alpine as builder

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /usr/local/src

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o dist cmd/app/main.go

EXPOSE 3000

CMD ["./dist"]