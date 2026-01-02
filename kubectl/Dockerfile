FROM ubuntu:latest as kc

WORKDIR /
RUN apt update && apt install -y ca-certificates \
    curl

ENV VERSION="1.35.0"
RUN curl -sfSOL "https://dl.k8s.io/release/v$VERSION/bin/linux/amd64/kubectl"

FROM ubuntu:latest

RUN apt update && apt install -y ca-certificates \
    jq vim curl

# Copy over kubectl
COPY --from=kc /kubectl /usr/bin/
RUN chmod +x /usr/bin/kubectl