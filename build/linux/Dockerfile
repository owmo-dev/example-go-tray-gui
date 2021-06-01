FROM ubuntu

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y sudo wget

RUN useradd -m docker && echo "docker:docker" | chpasswd && adduser docker sudo

RUN wget -c https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apt-get install -y wget gcc libgtk-3-dev libappindicator3-dev make

WORKDIR /example-tray-gui
COPY ../go.mod .
COPY ../go.sum .
RUN go mod download

COPY ../ .