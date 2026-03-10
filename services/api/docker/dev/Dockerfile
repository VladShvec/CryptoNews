FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git bash

RUN go install github.com/air-verse/air@latest

COPY services/api/docker/dev/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]