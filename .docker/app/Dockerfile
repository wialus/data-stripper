FROM golang:1 AS base

WORKDIR /app
COPY . /app

FROM base AS build

ENV DOCKER_TAG=0.0.0-dev

RUN CGO_ENABLED=0 go build -o data-stripper -ldflags "-X main.Version=${DOCKER_TAG}"

FROM scratch

WORKDIR /app

COPY --from=build /app/data-stripper /app/data-stripper

USER 33:33

ENTRYPOINT [ "/app/data-stripper" ]
