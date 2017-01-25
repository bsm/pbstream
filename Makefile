PKG=$(shell glide nv)

default: vet test

vet:
	go vet $(PKG)

test:
	go test -v $(PKG)

bench:
	go test -test.run=NONE -test.bench=. -test.benchmem $(PKG)

proto:
	protoc testdata/test.proto --gogo_out=. # go get -u github.com/gogo/protobuf/protoc-gen-gogo
