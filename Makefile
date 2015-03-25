default: test

.PHONY: default test

test:
	go test -test.timeout 10s
	go test ./numtheory -test.timeout 10s
	go test ./poly -test.timeout 10s
	go test ./float -test.timeout 10s
	go test ./decimal -test.timeout 10s
