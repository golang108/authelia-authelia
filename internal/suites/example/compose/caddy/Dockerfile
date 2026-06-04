FROM caddy:2.11.4-builder@sha256:f2b98918658f949a3c533f2c73bd0806e3f2576ccf8eb182c8b1690c977007ea AS builder

RUN xcaddy build fix-empty-copy-headers

FROM caddy:2.11.4@sha256:eb37f58646a901dc7727cf448cae36daaefaba79de33b5058dab79aa4c04aefb

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
