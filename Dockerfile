FROM golang:1.15.5 AS build
WORKDIR /go/src/github.com/dsadadsa/gitleaks
ARG ldflags
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -o bin/gitleaks

FROM alpine:3.11
RUN apk add --no-cache bash git openssh
COPY --from=build /go/src/github.com/dsadsad/gitleaks/bin/* /usr/bin/
ENTRYPOINT ["gitleaks"]

# How to use me :

# docker build -t gitleaks .
# docker run --rm --name=gitleaks gitleaks --repo=https://github.com/zricethezav/gitleaks

# This will check for secrets in https://github.com/zricethezav/gitleaks
