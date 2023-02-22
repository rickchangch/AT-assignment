FROM postgres:15-alpine

ENV POSTGRES_DB=${POSTGRES_DB:-test}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-root}
ENV POSTGRES_USER=${POSTGRES_USER:-test}

COPY sql/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
