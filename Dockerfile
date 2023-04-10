# Builder
FROM golang:1.20-alpine AS builder
ARG APP=goulash-menu-api
WORKDIR /home/${APP}

RUN apk add --update make
COPY Makefile Makefile
COPY . .
RUN make build

# Server
FROM alpine:latest as server
ARG APP=goulash-menu-api
WORKDIR /root/

COPY --from=builder /home/${APP}/bin/${APP} .
COPY --from=builder /home/${APP}/config.yml .
COPY --from=builder /home/${APP}/migrations/ ./migrations

RUN chown root:root ${APP}


EXPOSE 8080

CMD ["./goulash-menu-api"]