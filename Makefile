# Go parameters
GOCMD = go
GOBUILD = CGO_ENABLED=0 $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GORELEASERCMD = goreleaser

# Output binary name
BINARY_NAME = templater

# Targets
build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./...
	cp ./sample/* bin/

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -rf dist

snapshot:
	$(GORELEASERCMD) build --snapshot --clean

release-snapshot:
	$(GORELEASERCMD) release --snapshot --clean

release:
	$(GORELEASERCMD) release --clean
