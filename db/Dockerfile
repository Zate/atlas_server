FROM golang:latest as builder
WORKDIR /
RUN git clone https://github.com/dustin/seriesly.git && \
    cd seriesly && \
    export GO111MODULE=on && \ 
	go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o seriesly . && \
    mkdir /db

FROM scratch as atlas-d
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /db /db
COPY --from=builder /seriesly .
EXPOSE 3133
CMD ["/seriesly", "-logAccess"]
