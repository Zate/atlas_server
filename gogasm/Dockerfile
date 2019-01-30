FROM golang:latest as builder
WORKDIR /go/src/github.com/zate/

RUN git clone https://github.com/Zate/gogasm.git && \
        cd gogasm/cmd && \
	export GO111MODULE=on && \ 
	go mod tidy && \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gogasm .

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/zate/gogasm/cmd/gogasm .
#EXPOSE 2086
CMD ["/gogasm", "-web", "8081"]
