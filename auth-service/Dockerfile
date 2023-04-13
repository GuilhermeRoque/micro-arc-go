
# builder
FROM golang:1.20-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download
COPY ./src /app/src
RUN CGO_ENABLED=0 go build -o auth-service /app/src

# runner
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/auth-service /app
CMD ["/app/auth-service"]

