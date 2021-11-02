FROM ruby:2.6.0 AS ruby
WORKDIR /website
COPY website ./
COPY looks.yml tags.yml ./data/
RUN make build

FROM golang AS gomod
WORKDIR /go/src/
COPY service/go.mod service/go.sum ./
RUN go mod download

FROM golang AS go
WORKDIR /go/src/
COPY service ./
COPY looks.yml tags.yml ./data/
COPY --from=gomod /go/pkg/mod /go/pkg/mod
COPY --from=ruby /website/build ./static
RUN CGO_ENABLED=0 go install

FROM scratch
COPY --from=go /go/bin/service /service
COPY --from=go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/service"]
