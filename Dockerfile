FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

ENV PORT=${PORT}

RUN go build -o main ./main.go

RUN chmod +x main

EXPOSE ${PORT} 

CMD ["./main"]