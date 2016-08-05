FROM gliderlabs/alpine:3.2
EXPOSE 8080
ENTRYPOINT ["/bin/roloca"]

COPY *.go /go/src/github.com/nalexeric/roloca/
COPY loc.db /

RUN apk-install -t build-deps go git mercurial gcc g++ libc-dev \
	&& cd /go/src/github.com/nalexeric/roloca \
	&& export GOPATH=/go \
	&& go get \
	&& go build -ldflags "-X main.Version $(cat VERSION)" -o /bin/roloca \
	&& rm -rf /go \
	&& apk del --purge build-deps
