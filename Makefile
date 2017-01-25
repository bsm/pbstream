PKG=$(shell glide nv)

default: vet test

vet:
	go vet $(PKG)

test:
	go test -v $(PKG)

bench:
	go test -test.run=NONE -test.bench=. -test.benchmem $(PKG)
