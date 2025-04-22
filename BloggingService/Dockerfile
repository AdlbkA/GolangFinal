FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o blogging ./
RUN ls -l blogging
RUN chmod +x blogging

EXPOSE 8080

CMD ["./blogging"]