FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY . .
RUN go build -o /out/ai-platform-operator ./cmd/operator

FROM alpine:3.20
RUN adduser -D -u 10001 operator
USER operator
COPY --from=builder /out/ai-platform-operator /usr/local/bin/ai-platform-operator
ENTRYPOINT ["/usr/local/bin/ai-platform-operator"]
