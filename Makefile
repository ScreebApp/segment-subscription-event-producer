
BINARY=segment-subscription-event-producer

ASSETS_LOCATION=views
ASSETS_FILE=assets.go

all: build

#
# Deps
#
deps-tools:
	go get -x github.com/cespare/reflex
	go get -x github.com/rakyll/gotest
	go get -x github.com/psampaz/go-mod-outdated
	go get -x github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -x github.com/sonatype-nexus-community/nancy

deps: deps-tools
	go mod download -x

cleanup-deps:
	go mod tidy

audit: cleanup-deps
	go list -json -m all | nancy sleuth

outdated: cleanup-deps
	go list -u -m -json all | go-mod-outdated -update -direct

#
# Build
#
build:
	go build -o ${BINARY} ./*.go

watch:
	reflex -t 50ms -s -- sh -c 'echo \\nBUILDING && go run ./*.go && echo Exited \(0\)'

#
# Quality
#
lint:
	golangci-lint run --fix --timeout 600s --path-prefix=./ ./...

test:
	gotest -v ./tests/...

watch-test: deps
	reflex -t 50ms -s -- sh -c 'make test'

#
# Cleaning
#
clean:
	rm -f ${BINARY}

re: clean all
