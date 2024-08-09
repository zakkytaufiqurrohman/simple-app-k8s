FROM golang:1.22-alpine AS temp
WORKDIR /app
COPY . .
RUN go build -o main .
FROM alpine:latest
WORKDIR /root/
COPY --from=temp /app/main .
EXPOSE 4000
CMD ["./main"]