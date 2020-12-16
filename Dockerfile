# build stage
FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN go build -o app ./cmd/app

# build frontend
FROM node:alpine AS builder-f2e

WORKDIR /src

COPY web/package.json ./
RUN yarn install

COPY web/public ./public
COPY web/src ./src
RUN yarn build

# final stage
FROM alpine:3

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

WORKDIR /app

COPY --from=builder /src/app ./
COPY --from=builder-f2e /src/build ./web/build

ENTRYPOINT ./app
