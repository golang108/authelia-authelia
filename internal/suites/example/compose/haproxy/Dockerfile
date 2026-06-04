FROM haproxy:3.4.0-alpine@sha256:5614ec450485ce1f9f8c25d231cf7fbab9326302a395f2355e05cbbc2dd7468b
USER root

RUN \
apk add --no-cache \
  lua-json4

USER haproxy
