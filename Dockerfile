FROM --platform=$BUILDPLATFORM alpine:latest AS builder

RUN apk add --no-cache tzdata

FROM scratch

COPY earthquake-crawler /app

# 没有打包data文件夹，建议使用-v挂载，或者取消下一行注释
# COPY data data

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT [ "/app" ]