ARG SERVICE_NAME

FROM golang:1.22 AS builder
ARG SERVICE_NAME

COPY go.mod go.sum /go/src/github.com/amnestia/${SERVICE_NAME}/service/

WORKDIR /go/src/github.com/amnestia/${SERVICE_NAME}/service/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/main ./cmd/${SERVICE_NAME}/main.go


FROM scratch
ARG SERVICE_NAME

ENV SERVICE_ENV=dev

COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/src/github.com/amnestia/${SERVICE_NAME}/service/cmd/${SERVICE_NAME}/config/server /etc/${SERVICE_NAME}/config/server

EXPOSE 80
ENTRYPOINT ["/go/bin/main"]
