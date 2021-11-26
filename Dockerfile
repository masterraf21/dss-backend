FROM golang:1.17-alpine as builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod ./
COPY go.sum ./

RUN go mod download
#RUN go get -u
RUN go mod tidy

COPY . .

ENV GO111MODULE=on
# ENV PROJECTNAME account

RUN go mod tidy
RUN go build -o main

FROM alpine:latest as final

RUN apk update \
    && apk add --no-cache ca-certificates \
    && apk --no-cache add tzdata \
    && rm -rf /var/cache/apk/* \
    && mkdir -p /app/keys

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8800

ENTRYPOINT [ "./main" ]


