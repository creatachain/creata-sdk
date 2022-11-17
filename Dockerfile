# Simple usage with a mounted data directory:
# > docker build -t creataapp .
#
# Server:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.creataapp:/root/.creataapp creataapp creatad init test-chain
# TODO: need to set validator in genesis so start runs
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.creataapp:/root/.creataapp creataapp creatad start
#
# Client: (Note the creataapp binary always looks at ~/.creataapp we can bind to different local storage)
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.creataappcli:/root/.creataapp creataapp creatad keys add foo
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.creataappcli:/root/.creataapp creataapp creatad keys list
# TODO: demo connecting rest-server (or is this in server now?)
FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git licp-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/creatachain/creata-sdk

# Add source files
COPY . .

# install creataapp, remove packages
RUN make build-linux


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/creatachain/creata-sdk/build/creatad /usr/bin/creatad

EXPOSE 26656 26657 1317 9090

# Run creatad by default, omit entrypoint to ease using container with simcli
CMD ["creatad"]
