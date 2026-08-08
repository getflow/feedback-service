FROM golang:1.24-alpine AS build
RUN apk --no-cache add upx git
WORKDIR /app
COPY go.mod go.sum main.go /app/
COPY internal /app/internal/
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v ./... && go build -ldflags="-s -w" -o ./out/app /app/ && upx ./out/app

FROM alpine:3.11 AS runtime

ENV FB_PORT=3000

EXPOSE 3000

COPY --from=build /app/out/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]
