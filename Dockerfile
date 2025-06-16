FROM alpine:latest AS builder

RUN mkdir -p /dbfile
RUN apk add --no-cache tzdata

FROM scratch

COPY earthquake-crawler /app

COPY --from=builder /dbfile /dbfile
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT [ "/app" ]