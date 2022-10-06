FROM golang:1.16.5-alpine3.12 AS backendBuilder
RUN apk add --no-cache build-base git

ENV WORKDIR=/dogu
RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

COPY go.mod go.sum ${WORKDIR}/
RUN go mod download

COPY main.go ${WORKDIR}/
COPY config ${WORKDIR}/config
COPY Makefile ${WORKDIR}/
COPY build ${WORKDIR}/build
COPY .git ${WORKDIR}/.git

RUN go mod vendor
RUN make compile-generic

FROM registry.cloudogu.com/official/base:3.12.4-1

LABEL maintainer="hello@cloudogu.com" \
    NAME="tetsing/oidc-client" \
    VERSION="0.1.0-1" \
    SERVICE_TAGS=webapp

COPY --from=backendBuilder /dogu/target/oidc-client /app/oidc-client
COPY resources /

EXPOSE 8080

RUN set -x \
    && addgroup -S oidc && adduser -S oidc -G oidc \
    && chmod -R 775 /app /etc/ssl/certs \
    && chown -R oidc:oidc /app /etc/ssl/certs
USER oidc

HEALTHCHECK CMD doguctl healthy cas-oidc-client || exit 1

WORKDIR /app
CMD /app/startup.sh