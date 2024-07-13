FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# install essential library librdkafka, enabling comms through kafka in Go
# https://github.com/confluentinc/librdkafka
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]