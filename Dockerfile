FROM golang:1.14

ENV RUN_IN_CLUSTER="true"
ENV NAMESPACE="default"

RUN mkdir -p /go/src/github.com/kgysu/oc-apm
WORKDIR /go/src/github.com/kgysu/oc-apm

COPY . /go/src/github.com/kgysu/oc-apm
RUN go build

RUN chmod -R 777 /go/src/github.com/kgysu/oc-apm

EXPOSE 8080

USER nobody

CMD ["./oc-apm", "start"]