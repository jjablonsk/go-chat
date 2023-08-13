FROM golang:1.20

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

ARG USER
ARG UID=1000
RUN useradd -ms /bin/bash ${USER} -u ${UID}

USER ${USER}
ENV HOME=/home/${USER}