FROM golang:1.16-alpine

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags="-s -w" -gcflags="-trimpath=$PWD" -asmflags="-trimpath=$PWD"

FROM alpine

COPY --from=0 /go/src/app/itms-api /usr/local/bin/itms-api

ENTRYPOINT ["itms-api"]
