default: test

.PHONY: default test

test:
	go test -test.timeout 10s
	go test ./poly -test.timeout 10s
	go test ./experimental/numtheory -test.timeout 10s
	go test ./experimental/float -test.timeout 10s
	go test ./experimental/decimal -test.timeout 10s
