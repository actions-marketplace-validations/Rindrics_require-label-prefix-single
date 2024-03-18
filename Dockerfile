FROM golang:1.22 AS builder

WORKDIR /dist

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch

LABEL org.opencontainers.image.source="https://github.com/Rindrics/require-label-prefix"

COPY --from=builder /dist/app /app

CMD ["/app"]
