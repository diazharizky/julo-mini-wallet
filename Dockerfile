#################
# BUILDER IMAGE #
#################

FROM golang:1.20-alpine AS builder

RUN apk update && \
  apk add --no-cache make

WORKDIR /usr/local/build

COPY . .

RUN go mod download

RUN make compile

###############
# FINAL IMAGE #
###############

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/build/bin/app /usr/bin/app

ENTRYPOINT ["/usr/bin/app"]
