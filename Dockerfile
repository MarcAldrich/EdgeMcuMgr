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

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -ldflags="-extldflags=-static" -o main


# RUNTIME container needs only
FROM scratch as RUNNER
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
