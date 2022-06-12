FROM golang:1.18-alpine as development

ENV LANG C.UTF-8
ENV APP_ROOT /var/app

RUN apk add gcc g++

WORKDIR ${APP_ROOT}
COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@v1.29.0

RUN go mod download

CMD ["air", "-c", ".air.toml"]
