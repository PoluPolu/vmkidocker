FROM jenkins/jenkins:lts

USER root

# Install Go
RUN curl -OL https://golang.org/dl/go1.16.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz \
    && rm go1.16.5.linux-amd64.tar.gz

# Set Go environment variables
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

USER jenkins

