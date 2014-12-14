FROM google/golang

WORKDIR /gopath/src/myapp
ADD . /gopath/src/myapp/
RUN cd wxmp; go build; go install; cd -
RUN cd myutils; go build; go install; cd -
RUN go get -d -v
RUN go install -v 
CMD [myapp]
ENTRYPOINT ["/gopath/bin/myapp"]
