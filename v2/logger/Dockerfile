# This Dockerfile builds a custom Fluentd image that layers in the Fluentd
# plugins we rely upon on top of a standard Fluentd image.

FROM fluent/fluentd:v1.14.2-debian-1.0

USER root

RUN buildDeps="make gcc g++ libc-dev libffi-dev" \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends \
    $buildDeps \
    net-tools \
  && gem install \
    bson:4.12.1 \
    fluent-plugin-rewrite-tag-filter:2.4.0 \
    fluent-plugin-mongo:1.5.0 \
    fluent-plugin-kubernetes_metadata_filter:2.9.1 \
    fluent-plugin-multi-format-parser:1.0.0 \
  && apt-get purge -y --auto-remove \
    $buildDeps \
  && rm -rf /var/lib/apt/lists/* \
  && gem sources --clear-all \
  && rm -rf /tmp/* /var/tmp/* /usr/lib/ruby/gems/*/cache/*.gem
