FROM golang:1.8

WORKDIR /go/src/github.com/nightlegend/consul-dashboard

COPY . .

RUN export http_proxy=http://10.222.124.58:13128 && \
    export https_proxy=http://10.222.124.58:13128 && \
    go get -d -v ./...

# RUN go install

CMD ["consul-dashboard"]
