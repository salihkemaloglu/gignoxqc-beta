FROM golang:1.9

RUN go get gopkg.in/mgo.v2/bson 
RUN go get github.com/spf13/pflag
RUN go get github.com/rs/cors
RUN go get github.com/dgrijalva/jwt-go
RUN go get goji.io
# set environment path
ENV PATH /go/bin:$PATH

# cd into the api code directory
WORKDIR /go/src/github.com/salihkemaloglu/gignoxqc-beta-001

# create ssh directory
RUN mkdir ~/.ssh
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# allow private repo pull
RUN git config --global url."https://e4d5159cc774d99744024453431f00ddbb8d7b1d:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# copy the local package files to the container's workspace
ADD . /go/src/github.com/salihkemaloglu/gignoxqc-beta-001

# install the program
RUN go install github.com/salihkemaloglu/gignoxqc-beta-001

# expose default port
EXPOSE 80 443
# start application
CMD ["go","run","main.go"] 