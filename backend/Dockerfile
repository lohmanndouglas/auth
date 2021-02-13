FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy && go mod vendor && go build -o main .
CMD ["./main"]
