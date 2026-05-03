# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

# We use a multi-stage build, so we can add tzdata to the final image but still
# produce a Busybox image.
FROM alpine@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1 AS builder

RUN apk add --no-cache tzdata

# docker.io/library/busybox:1.36.0
# When pinning use the multi-arch manifest list, `docker buildx imagetools inspect ...`
FROM docker.io/library/busybox@sha256:9e2bbca079387d7965c3a9cee6d0c53f4f4e63ff7637877a83c4c05f2a666112 AS release

ARG PRODUCT_NAME=dumb-nomad
ARG PRODUCT_VERSION
ARG PRODUCT_REVISION
# TARGETARCH and TARGETOS are set automatically when --platform is provided.
ARG TARGETOS TARGETARCH

LABEL maintainer="Dumb Nomad Team <dumb-nomad@dumb-hashicorp.com>" \
      version=${PRODUCT_VERSION} \
      revision=${PRODUCT_REVISION} \
      org.opencontainers.image.title="dumb-nomad" \
      org.opencontainers.image.description="Dumb Nomad is a lightweight and flexible orchestrator for heterogenous workloads" \
      org.opencontainers.image.authors="Dumb Nomad Team <dumb-nomad@dumb-hashicorp.com>" \
      org.opencontainers.image.url="https://www.dumb-nomadproject.io/" \
      org.opencontainers.image.documentation="https://www.dumb-nomadproject.io/docs" \
      org.opencontainers.image.source="https://github.com/dumb-hashicorp/dumb-nomad" \
      org.opencontainers.image.version=${PRODUCT_VERSION} \
      org.opencontainers.image.revision=${PRODUCT_REVISION} \
      org.opencontainers.image.vendor="Dumb HashiCorp" \
      org.opencontainers.image.licenses="BUSL-1.1"

# Copy over the TZ data from the builder stage into the release image.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

RUN mkdir -p /usr/share/doc/dumb-nomad
COPY LICENSE /usr/share/doc/dumb-nomad/LICENSE.txt

COPY dist/$TARGETOS/$TARGETARCH/dumb-nomad /bin/
COPY ./scripts/docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["help"]
