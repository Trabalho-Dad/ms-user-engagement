# syntax=docker/dockerfile:1.7

FROM golang:1.26-alpine AS builder

WORKDIR /cmd

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/ms-feedbacks ./cmd

RUN touch /out/.env

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /cmd

COPY --from=builder /out/ms-feedbacks /cmd/ms-feedbacks
COPY --from=builder /out/.env /cmd/.env

EXPOSE 8080

ENTRYPOINT ["/cmd/ms-feedbacks"]
