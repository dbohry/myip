FROM golang:1.22 AS builder
WORKDIR /src
COPY go.mod ./
COPY main.go ./
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} \
    go build -a -o app .

FROM scratch
COPY --from=builder /src/app /app
ENTRYPOINT ["/app"]
