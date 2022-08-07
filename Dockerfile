FROM golang:1.18-alpine3.14

ENV CGO_ENABLED=0

RUN go install github.com/cespare/reflex@latest && \
  go install github.com/mitranim/gow@latest
