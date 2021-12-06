FROM golang:1.17.1-alpine as builder

LABEL maintainer="Shuang Pan <shuang.pan@sjsu.edu>"
ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

COPY backend /go/src/MochiNote/backend

WORKDIR /go/src/MochiNote/backend/source/generated/cmd/coreapi-server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s'


FROM scratch

COPY --from=builder /go/src/MochiNote/backend/source/generated/cmd/coreapi-server/coreapi-server .

ENTRYPOINT ["/coreapi-server"]