FROM FROM golang:1.20-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o wick ./cmd

FROM busybox as runner
COPY --from=builder /app/wick /app
ENTRYPOINT ["/app"]
