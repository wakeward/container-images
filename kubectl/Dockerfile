FROM ubuntu:latest AS kc

WORKDIR /
RUN apt update && apt install -y ca-certificates \
    curl

ENV VERSION="1.35.0"
RUN curl -sfSOL "https://dl.k8s.io/release/v$VERSION/bin/linux/amd64/kubectl"

FROM ubuntu:latest

# add useful tools
RUN apt update && apt install -y ca-certificates \
    jq vim curl wget netcat-traditional

# Copy over kubectl
COPY --from=kc /kubectl /usr/bin/
RUN chmod +x /usr/bin/kubectl