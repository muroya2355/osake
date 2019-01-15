FROM fluent/fluentd:v0.12-debian

RUN gem install fluent-plugin-elasticsearch --no-rdoc --no-ri --version 1.9.2
COPY ./docker/fluentd/fluent.conf /fluentd/etc/fluent.conf