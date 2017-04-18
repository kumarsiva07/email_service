# Email Service

[![Build Status](https://travis-ci.org/lileio/email_service.svg?branch=master)](https://travis-ci.org/lileio/email_service)

A simple gRPC based email service that sends emails via SMTP. Connections are kept open whilst sending email then closed for after 30 seconds of inactivity.

```
service EmailService {
  rpc Send (Email) returns (EmailResponse) {}
}
```

## Docker

Email Service is available via Docker hub

```
docker pull lileio/email_service
```

## Config

You can set SMTP details with the following env vars.

```
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=2525
SMTP_USERNAME=apikey
SMTP_PASSWORD=somepassword
```

Set `DEBUG=true` to see timing and connection information
