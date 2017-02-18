FROM alpine:latest

RUN apk add --no-cache ca-certificates
ADD build/email_service /
CMD ["/email_service"]
