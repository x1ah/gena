FROM golang:1.14.4-alpine3.12 as builder

ARG APK_REPO=mirrors.aliyun.com
RUN sed -i "s|//dl-cdn.alpinelinux.org|//${APK_REPO}|g" /etc/apk/repositories; \
    apk add --no-cache make curl git build-base

WORKDIR /app
COPY . .
ARG GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN CGO_ENABLED=0 make build

FROM scratch

WORKDIR /opt/code
COPY --from=builder /app/server /opt/code/

EXPOSE 800

HEALTHCHECK --interval=30s --timeout=1s CMD ["/opt/code/server", "healthcheck"]

ENTRYPOINT [ "/opt/code/server"]
