FROM golang:1.12.5 as builder
WORKDIR /go/src/go_projects/learngo/shippy_demo/user_service
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go token_service.go repository.go handler.go datastore.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/go_projects/learngo/shippy_demo/user_service/user_service .
CMD ["./user_service"]