FROM golang:1.20 AS build

RUN useradd -u 10001 benthos

WORKDIR /build/
COPY . /build/

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/geo-benthos -mod=vendor 

FROM busybox AS package

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /build/geo-benthos .
COPY ./config/example_1.yaml /benthos.yaml

USER benthos

EXPOSE 4195

ENTRYPOINT ["/geo-benthos"]

CMD ["-c", "/benthos.yaml"]
