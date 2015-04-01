FROM golang:1.4.2

#logster install
RUN apt-get update
RUN apt-get install -y logcheck
RUN git clone https://github.com/etsy/logster.git
RUN cd logster ; python setup.py install 

#logster-docker-runner

COPY . /go/src/logster-docker-runner
WORKDIR /go/src/logster-docker-runner

ENV GOPATH /go/src/logster-docker-runner/Godeps/_workspace:$GOPATH
RUN go install -v 

CMD ["logster-docker-runner"]
#ENTRYPOINT ["logster-docker-runner"]
