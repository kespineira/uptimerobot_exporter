FROM alpine AS builder
RUN apk add --no-cache ca-certificates

FROM scratch AS final
USER 65535:65535
COPY  --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./uptimerobot_exporter /
EXPOSE 9429
ENTRYPOINT ["/uptimerobot_exporter"]