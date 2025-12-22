# syntax=docker/dockerfile:1
ARG GO_VERSION=1.25.5

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -trimpath -ldflags "-s -w" -o /out/upwatch ./cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/upwatch /app/upwatch

EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["/app/upwatch"]
