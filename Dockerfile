FROM golang:1.17-alpine3.14 as build

WORKDIR /usr/local/opengauss_exporter

ADD . .

RUN ./build.sh

FROM alpine:3.14 as prod

WORKDIR /usr/local/opengauss_exporter

ENV CONFIG_FILE "config.yaml.example"

COPY --from=build /usr/local/opengauss_exporter/opengauss_exporter .
COPY --from=build /usr/local/opengauss_exporter/config.yaml.example ./config.yaml.example

ENTRYPOINT ["/usr/local/opengauss_exporter/opengauss_exporter", "-c", "$CONFIG_FILE"]