FROM golang:1.19-alpine3.17 AS builder
WORKDIR /project
COPY . .
RUN go build -o advandev

FROM alpine:3.17
COPY --from=builder /project/advandev /advandev
ENTRYPOINT [ "advandev" ]
EXPOSE 80
