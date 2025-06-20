FROM --platform=$BUILDPLATFORM alpine:latest AS builder

RUN apk add --no-cache tzdata

FROM scratch

COPY earthquake-crawler /app

# COPY data data

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT [ "/app" ]