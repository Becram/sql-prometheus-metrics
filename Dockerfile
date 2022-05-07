# syntax=docker/dockerfile:experimental
# ---
FROM golang:1.18 AS build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /work
COPY . /work

# Build admission-webhook
RUN --mount=type=cache,target=/root/.cache/go-build,sharing=private \
  go build -o bin/sql-prometheus-metrics .

# ---
FROM scratch AS run

COPY --from=build /work/bin/sql-prometheus-metrics /usr/local/bin/

CMD ["/usr/local/bin/sql-prometheus-metrics"]
