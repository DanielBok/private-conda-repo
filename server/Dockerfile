FROM golang:1.13-alpine as builder

RUN apk update && apk upgrade

WORKDIR /app

COPY . .
RUN go build -o pcr .

FROM alpine:3

RUN apk update && apk upgrade && apk add docker
RUN mkdir -p /var/condapkg

WORKDIR /app
COPY --from=builder /app/pcr            pcr
COPY --from=builder /app/config.yaml    /var/private-conda-repo/config.yaml

ENTRYPOINT ["./pcr"]
