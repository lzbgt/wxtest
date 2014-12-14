FROM google/golang

WORKDIR /gopath/src/myapp
ADD . /gopath/src/myapp/
RUN go get -d -v
RUN go install -v 
ENV PATH $PATH:/gopath/bin/
CMD myapp
#ENTRYPOINT ["/gopath/bin/myapp"]
