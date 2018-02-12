FROM golang:1.8

WORKDIR /go/src/github.com/nightlegend/consul-dashboard

COPY . .

RUN go get -d -v ./...

# RUN go install

CMD ["consul-dashboard"]
