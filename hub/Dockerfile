FROM golang:1.21.6-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go env -w GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o /hub hub.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /hub .
ENV PORT 3990
EXPOSE 3990
CMD [ "/hub" ]
