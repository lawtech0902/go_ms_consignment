FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment_cli /app/consignment_cli

CMD ["./consignment_cli"]