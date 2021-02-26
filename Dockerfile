###############
### BUILDER ###
###############

FROM golang:1.14-alpine3.11 as builder

RUN apk add git build-base

RUN mkdir -p /usr/src/microservice-ads
WORKDIR /usr/src/microservice-ads

RUN apk update 

ADD cmd ./cmd
ADD bin ./bin
ADD pkg ./pkg
ADD db ./db
ADD docker ./docker
ADD app.env .
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD Makefile .
# ADD LICENSE .
RUN make build

###############
### RELEASE ###
###############

FROM alpine:3.11

LABEL Author="Stanley Salvatierra <stanlee321@gmail.com>"

COPY --from=builder /usr/src/microservice-ads/ /microservice-ads

RUN cd /microservice-ads \
    && apk add make git \
    && make install \
    # && ls -l \
    # && pwd \
    && mv /microservice-ads/app.env  / \
    && rm -rf /microservice-ads \
    && apk del make git 

# Configuration
COPY ./docker/docker-entrypoint.sh /usr/local/bin/

RUN chmod 777 /usr/local/bin/docker-entrypoint.sh \
    && ln -s /usr/local/bin/docker-entrypoint.sh /

RUN ln -s /usr/local/bin/docker-entrypoint.sh /entrypoint.sh # backwards compat


ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["microservice-ads"]