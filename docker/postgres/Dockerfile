FROM postgres:12.4
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.UTF-8

ENV POSTGRES_USER osake_user
ENV POSTGRES_PASSWORD osake_user
ENV POSTGRES_DB osake

COPY ./docker/postgres/init/*.sql /docker-entrypoint-initdb.d/