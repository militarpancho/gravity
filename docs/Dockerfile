# This Dockerfile makes the container used to build Gravity documentation
#
# A multi step build is used to keep the go toolchain outside the final container

# milv-builder, contains the whole go toolchain
FROM quay.io/gravitational/debian-venti:go1.14.4-buster AS milv-builder
RUN GO111MODULE=on go get -u -v github.com/magicmatatjahu/milv@v0.0.6

# docbox, contains everything for building gravity documentation
FROM quay.io/gravitational/mkdocs-base:1.0.3-1

ARG UID
ARG GID
ARG WORKDIR
ARG PORT

RUN groupadd builder --gid=$GID -o && useradd builder --uid=$UID --gid=$GID --create-home --shell=/bin/bash
COPY --from=milv-builder /gopath/bin/milv /usr/bin/milv

VOLUME [$WORKDIR]
EXPOSE $PORT
