FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .

FROM scratch
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8080

CMD ["./server"]