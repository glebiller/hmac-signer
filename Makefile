build:
	goreleaser build --rm-dist --single-target --snapshot --output hmac-signer
.PHONY: build

package: build
	docker build -t hmac-signer:local .
.PHONY: package

clean:
	rm -rf dist/
.PHONY: build
