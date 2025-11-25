# FROM golang:1.25.3 as builder


# RUN apk update
# RUN apk add git openssh tzdata build-base python3 net-tools

# WORKDIR /app

# COPY .env.example .env
# COPY . .

# RUN go install github.com/buu700/gin@latest
# RUN go mod tidy

# RUN make build

# FROM alpine:latest

# RUN apk update && apk upgrade && \
#     apk --update --no-cache add tzdata && \
#     apk --no-cache add curl && \
#     mkdir /app

# WORKDIR /app

# EXPOSE 8001

# COPY --from=builder /app /app

# ENTRYPOINT ["/app/user-service"]


FROM golang:1.25.3 AS builder

RUN apt-get update && apt-get install -y \
    git openssh-client tzdata build-essential python3 net-tools make

WORKDIR /app

COPY .env.example .env
COPY . .

RUN go install github.com/buu700/gin@latest
RUN go mod tidy

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y tzdata curl && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8001

ENTRYPOINT ["/app/user-service"]
