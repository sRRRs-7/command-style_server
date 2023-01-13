# Build stage
FROM --platform=linux/amd64 postgres:15-alpine

COPY ./script /docker-entrypoint-initdb.d
COPY ./conf /etc/postgresql/postgresql.conf

EXPOSE 5432