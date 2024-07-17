FROM golang:1.22-alpine3.20 as protoc-builder

ENV PROTOC_VERSION=27.2
ARG TARGETARCH

RUN apk --no-cache add git curl unzip \
    && if [ $TARGETARCH = "amd64" ]; then export DL_ARCH=x86_64 ; fi \
    && if [ $TARGETARCH = "arm64" ]; then export DL_ARCH=aarch_64 ; fi \
    && curl -f -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${DL_ARCH}.zip \
    && unzip /tmp/protoc.zip && rm /tmp/protoc.zip \
    && mv bin/protoc /usr/bin && rm -rf bin include \
    && mkdir -p /usr/include \
    # cloning well-known-types
    && git clone --depth=1 https://github.com/protocolbuffers/protobuf.git /protobuf-repo \
    && mv /protobuf-repo/src/ /usr/include/protobuf/ \
    # cloning googleapis-types
    && git clone --depth=1 https://github.com/googleapis/googleapis.git /usr/include/googleapis \
    # cleanup
    && rm -rf /protobuf-repo \
    && find /usr/include/protobuf -not -name "*.proto" -type f -delete \
    && find /usr/include/googleapis -not -name "*.proto" -type f -delete

FROM golang:1.22-alpine3.20 AS builder

RUN go install -v -ldflags "-s -w" google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install -v -ldflags "-s -w" google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . /src

WORKDIR /src/protoc-gen-gripmock

RUN go install -v -ldflags "-s -w"

FROM golang:1.22-alpine3.20

LABEL org.opencontainers.image.source=https://github.com/bavix/gripmock
LABEL org.opencontainers.image.description="gRPC Mock Server"
LABEL org.opencontainers.image.licenses=Apache-2.0

ARG version

COPY --from=protoc-builder /usr/bin/protoc /usr/bin/protoc
COPY --from=protoc-builder /usr/include/protobuf /protobuf
COPY --from=protoc-builder /usr/include/googleapis /googleapis

COPY --from=builder $GOPATH/bin/protoc-gen-go $GOPATH/bin/protoc-gen-go
COPY --from=builder $GOPATH/bin/protoc-gen-go-grpc $GOPATH/bin/protoc-gen-go-grpc
COPY --from=builder $GOPATH/bin/protoc-gen-gripmock $GOPATH/bin/protoc-gen-gripmock

COPY --from=builder /src /go/src/github.com/bavix/gripmock

WORKDIR /go/src/github.com/bavix/gripmock

RUN go install -v -ldflags "-X 'github.com/bavix/gripmock/cmd.version=${version:-dev}' -s -w"

EXPOSE 4770 4771

HEALTHCHECK CMD gripmock check --silent

ENTRYPOINT ["gripmock"]
