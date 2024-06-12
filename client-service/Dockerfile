FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /client-service

EXPOSE 8082

CMD [ "/client-service" ]
