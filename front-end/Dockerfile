# builder
FROM golang:1.20-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY go.mod /app/
RUN go mod download
COPY ./cmd /app/cmd
RUN CGO_ENABLED=0 go build -o front-end /app/cmd/web/main.go

# runner
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app /app
CMD ["/app/front-end"]