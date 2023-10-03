ARG GOVERSION=1.21
ARG GOARCH
FROM golang:${GOVERSION} as builder
ARG GOARCH
ENV GOARCH=${GOARCH}
WORKDIR /app
COPY . /app/

RUN make build-local

FROM gcr.io/distroless/static:latest-${GOARCH}
COPY --from=builder /app/systemtrading /

USER nobody

ENTRYPOINT ["/systemtrading"]
