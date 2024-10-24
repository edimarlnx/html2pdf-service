# Build the manager binary
FROM golang:1.22.7-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace

COPY . .

RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -o html2pdf \
    -ldflags="-X 'main.Version=$(head -n1 .version)'" main.go

FROM alpine:3.18

RUN apk add --no-cache chromium

WORKDIR /workspace

COPY --from=builder /workspace/html2pdf /usr/local/bin/

COPY --from=builder /workspace/index.html /workspace/index.html

CMD ["html2pdf"]
