FROM postgres:13

ENV POSTGRES_USER bot
ENV POSTGRES_PASSWORD pass
ENV POSTGRES_DB botdb

ADD initial.sql /docker-entrypoint-initdb.d/
