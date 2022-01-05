FROM golang:1.17

WORKDIR /test-with-revel

# Grab the source code and add it to the workspace.
ADD . /test-with-revel

# RUN apk add --no-cache bash

# Install revel and the revel CLI.
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
# RUN tar xzvf test-with-revel.tar.gz
# RUN revel package test-with-revel

# Use the revel CLI to start up our application.
ENTRYPOINT revel run -a test-with-revel
# ENTRYPOINT  ["/bin/bash", "run.sh"]

# Open up the port where the app is running.
EXPOSE 8080

# Final stage
# FROM alpine:3.15
# EXPOSE 8080
# WORKDIR /
# ADD . /
# #COPY /go/app /
# RUN tar xzvf test-with-revel.tar.gz
# ENTRYPOINT /run.sh