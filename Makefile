SHELL = bash
OSARCHES := "darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 linux/ppc64 linux/ppc64le linux/s390x"
OUTPUT := "build/bifrost-$(VERSION)-{{.OS}}-{{.Arch}}/bifrost"


test:
	@socat pty,link=/tmp/bifrostmaster,echo=0,crnl pty,link=/tmp/bifrostslave,echo=0,crnl & echo "$$!" > "socat.pid"
	go test
	@if [ -a socat.pid ]; then \
		kill -TERM $$(cat socat.pid) || true; \
	fi
	@rm socat.pid

test_coverage:
	@socat pty,link=/tmp/bifrostmaster,echo=0,crnl pty,link=/tmp/bifrostslave,echo=0,crnl & echo "$$!" > "socat.pid"
	go test -coverprofile=cover.out && go tool cover -html=cover.out
	@if [ -a socat.pid ]; then \
		kill -TERM $$(cat socat.pid) || true; \
	fi
	@rm socat.pid

build_all:
	if [ -z $(VERSION) ]; then \
	  echo "You need to specify a VERSION"; \
	  exit 1; \
	fi

	mkdir -p build
	if [ -d "build/" ]; then \
    	rm -rf build/*; \
	fi
	gox -osarch=$(OSARCHES) -cgo -output=$(OUTPUT)
	echo "compressing build files"
	cd build && for d in */; do filepath=$${d%/*}; echo $$filepath; zip "$${filepath##*/}.zip" "$${filepath##*/}/bifrost"; done