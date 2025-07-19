FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o orders_microservice
ENTRYPOINT [ "./orders_microservice" ]