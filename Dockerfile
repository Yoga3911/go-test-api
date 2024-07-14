FROM golang:1.22.5-alpine

WORKDIR /temp

COPY . .

RUN go build -tags netgo -o main.app main.go

WORKDIR /app

RUN mv /temp/main.app .

RUN rm -rf /temp

CMD ["./main.app"]