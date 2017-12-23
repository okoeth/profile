FROM golang:1.9
USER 1001
WORKDIR /go/src/profile
COPY . /go/src/profile
RUN go test 
RUN go install 
EXPOSE 8017
ENTRYPOINT [ "/go/bin/profile"]
