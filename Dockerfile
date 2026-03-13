FROM docker.1ms.run/library/golang:1.25.4-alpine AS builder

WORKDIR "/app"

COPY . .

RUN chmod +x ./build/build.sh && \
    ./build/build.sh

FROM docker.1ms.run/library/busybox:1.36
COPY --from=builder /app/local-api-gateway /app/local-api-gateway
ENTRYPOINT ["/app/local-api-gateway"]