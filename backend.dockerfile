FROM golang:1.12-alpine

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOPROXY=https://proxy.golang.org

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/Backend-Golang-RESTful-API
COPY . .


RUN go mod init DACN-GithubTrending
RUN GOOS=linux go build -o github-trending


ENTRYPOINT ["./github-trending"]

EXPOSE 3000
