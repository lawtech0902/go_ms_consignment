#FROM golang:1.12.5 as builder
#WORKDIR /go/src/go_projects/learngo/shippy_demo/user_cli
#COPY . .
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensure
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
#
#
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#RUN mkdir /app
#WORKDIR /app
#COPY --from=builder /go/src/go_projects/learngo/shippy_demo/user_cli/user_cli .
#ENTRYPOINT ["./user_cli"]
#CMD ["./user_cli"]

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD user_cli /app/user_cli

CMD ["./user_cli"]