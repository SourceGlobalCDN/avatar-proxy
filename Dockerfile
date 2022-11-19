FROM golang:alpine AS Builder
WORKDIR /app/avatar-proxy/

RUN apk add build-base

COPY . .
RUN go mod download

RUN go build -tags=sonic -o avatar-proxy .

FROM alpine AS Runner
WORKDIR /app/avatar-proxy/

COPY --from=Builder /app/avatar-proxy/avatar-proxy avatar-proxy

# Add Tini
RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "--"]

ENV LISTEN :9000

RUN chmod +x avatar-proxy
CMD ./avatar-proxy