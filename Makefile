SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

bin/wendover: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/wendover cmd/wendover/main.go

.PHONY: install
install: bin/wendover
	go run build/wendover/install.go $(CURDIR)
	cp bin/wendover ${HOME}/.local/bin/

.PHONY: dev-install
dev-install: bin/wendover
	-rm -r ./testdata
	mkdir -p ./testdata/config
	mkdir -p ./testdata/cache
	WENDOVER_CONFIG=./testdata/config WENDOVER_CACHE=./testdata/cache go run build/wendover/install.go $(CURDIR)

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/


# For future reference when I set this up with a server:
#.PHONY: run
#run: bin/ag7if
#	air -d -c .air.conf