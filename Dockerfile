FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install curl \
    apt-get install iputils-ping

CMD ["tail", "-f", "/dev/null"]