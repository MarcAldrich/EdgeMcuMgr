# Author: Marc Aldrich
#
# Date Last Modified: 2020 Oct 05
# Date Created: 2020 Sep 20
# Summary: Builds the docker container for
# Options: Arch left as option

FROM golang:alpine as BUILDER

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm \
    GOARM=7
# Setup environment for Bluetooth
RUN set -xe \
    && echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
#     && echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
#     && echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
    && apk add --no-cache --purge -uU curl \
        bluez bluez-libs bluez-firmware python3 openrc dbus-openrc libc6-compat \
    && rm -rf /var/cache/apk/* /tmp/*

# Build go application
RUN mkdir /build
ADD ./ /build/
WORKDIR /build
RUN go build -o main ./newtmgr/
ENTRYPOINT ./entrypoint.sh
