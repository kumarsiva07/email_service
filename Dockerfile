FROM alpine:latest

RUN apk add --no-cache ca-certificates
ADD build/email_service /bin
CMD ["email_service", "server"]
