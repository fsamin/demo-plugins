FROM golang:1.8-onbuild

COPY . /go/src/github.com/golang-rennes/demo-plugins/go-1-8

RUN go build -buildmode=plugin /go/src/github.com/golang-rennes/demo-plugins/go-1-8/world/world.go
RUN mkdir /opt/plugins/
RUN mv world.so /opt/plugins/world.so

RUN go build -buildmode=plugin /go/src/github.com/golang-rennes/demo-plugins/go-1-8/reverse/reverse.go
RUN mv reverse.so /opt/plugins/reverse.so

RUN go get -d -v
RUN go build  -o go-1-8 github.com/golang-rennes/demo-plugins/go-1-8 
RUN cp go-1-8 /opt/
RUN ls -lR /opt/

CMD /opt/go-1-8 foo bar