default: run

.PHONY: default run test

run: test
	GOPATH=`pwd` go install mathx

test:
	GOPATH=`pwd` go test mathx -test.timeout 1s
	GOPATH=`pwd` go test mathx/float -test.timeout 1s
