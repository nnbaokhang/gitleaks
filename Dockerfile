FROM golang:1.15.5 AS build
WORKDIR /go/src/github.com/nnbaokhang/gitleaks
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -o bin/gitleaks
FROM alpine:3.11
RUN apk add --no-cache bash
COPY --from=build /go/src/github.com/nnbaokhang/gitleaks/bin/* /usr/bin/
COPY --from=build /go/src/github.com/nnbaokhang/gitleaks/examples/* /opt/
ENTRYPOINT ["gitleaks"]

# How to use me :

# docker build -t gitleaks .
# docker run --rm --name=gitleaks gitleaks --repo=https://github.com/zricethezav/gitleaks

# This will check for secrets in https://github.com/zricethezav/gitleaks
