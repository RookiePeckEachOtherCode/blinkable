FROM golang:1.21.6-alpine
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY server ./server

RUN go mod tidy

WORKDIR /app/server/service/article

CMD sh build.sh && sh output/bootstrap.sh

EXPOSE 10003
