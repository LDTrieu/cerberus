GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGENERATE=$(GOCMD) generate
GOFMT=gofmt
GOIMPORT=goimports

REMOVECMD=rm
ECHOCMD=echo

generate:
	$(GOGENERATE) ./...

test: generate
	$(ECHOCMD) "Running unit testing..."
	$(GOTEST) -short ./...

lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
	./bin/golangci-lint run --timeout 5m
	$(REMOVECMD) -rf bin

code-format:
	$(GOFMT) -w config/ package/
	$(GOIMPORT) -w config/ package/

clean:
	$(GOCLEAN) -testcache ./...
	$(REMOVECMD) -rf bin *.out

.PHONY: generate lint tests unit-tests clean

tag:
	git tag -a v$(VERSION) -m "Release version $(VERSION)"
	git push origin v$(VERSION)