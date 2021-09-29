FROM registry.cloudogu.com/official/base:3.12.4-1

LABEL maintainer="hello@cloudogu.com" \
    NAME="tetsing/cas-oidc-client" \
    VERSION="0.1.0-1" \
    SERVICE_TAGS=webapp

COPY target/oidc-client /app/oidc-client
COPY resources /

EXPOSE 8080

RUN set -x \
    && addgroup -S oidc && adduser -S oidc -G oidc \
    && chmod -R 777 /app /etc/ssl/certs \
    && chown -R oidc:oidc /app /etc/ssl/certs
USER oidc

HEALTHCHECK CMD doguctl healthy cas-oidc-client || exit 1
WORKDIR /app
CMD /app/startup.sh