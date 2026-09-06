FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/ai-platform-operator ./cmd/operator

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /out/ai-platform-operator /ai-platform-operator
USER nonroot:nonroot
EXPOSE 8081
ENTRYPOINT ["/ai-platform-operator"]
