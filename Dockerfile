FROM golang:1.18-alpine as dev

RUN apk update \
    && apk add gcc musl-dev

WORKDIR /opt/app

RUN go install golang.org/x/tools/cmd/goimports@latest \
      && go install mvdan.cc/gofumpt@latest

################
# Base Builder #
################

FROM dev as builder-base

ADD ./integration /opt/app/integration
ADD ./internal /opt/app/internal
ADD ./main.go /opt/app/main.go
ADD ./go.mod /opt/app/go.mod
ADD ./go.sum /opt/app/go.sum

RUN go build -o bin/app .

###################
# Command Builder #
###################

FROM dev as builder-plugin-command

ADD ./integration /opt/app/integration
ADD ./plugin /opt/app/plugin

RUN go build -buildmode=plugin -o bin/command.so plugin/command.go

#####################
# WebServer Builder #
#####################

FROM dev as builder-plugin-web-server

ADD ./integration /opt/app/integration
ADD ./plugin /opt/app/plugin

RUN go build -buildmode=plugin -o bin/web-server.so plugin/web-server.go

######################
# Final Base Builder #
######################

FROM alpine:latest as final-base

WORKDIR /opt/app

COPY --from=builder-base /opt/app/bin /opt/app/bin

################################
# Final Command Plugin Builder #
################################

FROM kz/plugin-playground-base:latest as final-plugin-command

WORKDIR /opt/app

COPY --from=builder-plugin-command /opt/app/bin/*.so /opt/app/bin

CMD bin/app plugin --module=command

##################################
# Final WebServer Plugin Builder #
##################################

FROM kz/plugin-playground-base:latest as final-plugin-web-server

WORKDIR /opt/app

COPY --from=builder-plugin-web-server /opt/app/bin/*.so /opt/app/bin

CMD bin/app plugin --module=web-server
