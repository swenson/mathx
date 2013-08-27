default: run

.PHONY: default run test

run: test
	GOPATH=`pwd` go install ntag

test:
	GOPATH=`pwd` go test ntag -test.timeout 0.1s
