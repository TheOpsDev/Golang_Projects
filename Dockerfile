FROM mcr.microsoft.com/vscode/devcontainers/go

RUN apt-get update && \
    export DEBIAN_FRONTEND=noninteractive && \
    apt-get -y install git