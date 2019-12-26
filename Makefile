BINARY="apkparser"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

default:
	CGO_ENABLED=0 GOARCH=amd64 go build -o ${BINARY} .

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} .

install:
	ssh root@dl.5u2.net "rm -f /usr/local/bin/apkparser" && scp -r apkparser templates root@dl.5u2.net:/usr/local/bin/

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default linux install clean
