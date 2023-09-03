FROM ubuntu
USER root
COPY go1.20.4.linux-amd64.tar.gz /root/go1.20.4.linux-amd64.tar.gz
COPY ./module3 /root/module3
COPY ./go.mod /root/go.mod
COPY ./go.sum /root/go.sum
RUN tar -C /usr/local -xzf /root/go1.20.4.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin | tee -a /etc/profile && \
    export PATH=$PATH:/usr/local/go/bin >> ~/.profile && \
    cd /root/module3  && go build
WORKDIR /root/module3/
EXPOSE 8080
ENTRYPOINT ["./module3"]
