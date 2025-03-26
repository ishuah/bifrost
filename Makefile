SHELL = bash
OSARCHES := "darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 linux/ppc64 linux/ppc64le linux/s390x"
DARWIN_ARCHES = amd64 arm64
LINUX_ARCHES = 386 amd64 arm arm64 ppc64 ppc64le s390x
WINDOWS_ARCHES = 386 amd64 arm arm64
OUTPUT := "build/bifrost-$(VERSION)-{{.OS}}-{{.Arch}}/bifrost"


test:
	@socat pty,link=/tmp/bifrostport1,echo=0,crnl pty,link=/tmp/bifrostport2,echo=0,crnl & echo "$$!" > "socat.pid"
	go test -v
	@if [ -a socat.pid ]; then \
		kill -TERM $$(cat socat.pid) || true; \
		rm socat.pid || true; \
	fi


test_coverage:
	@socat pty,link=/tmp/bifrostport1,echo=0,crnl pty,link=/tmp/bifrostport2,echo=0,crnl & echo "$$!" > "socat.pid"
	go test -v -coverprofile=cover.out && go tool cover -html=cover.out
	@if [ -a socat.pid ]; then \
		kill -TERM $$(cat socat.pid) || true; \
	fi
	@rm socat.pid


# build darwin/amd64 separately when building on Linux
# use cross compiler xgo
# xgo --targets=darwin/amd64 github.com/ishuah/bifrost
build_all:
	if [ -z $(VERSION) ]; then \
	  echo "You need to specify a VERSION"; \
	  exit 1; \
	fi

	mkdir -p build
	if [ -d "build/" ]; then \
    	rm -rf build/*; \
	fi
	# Build Darwin first
	for arch in $(DARWIN_ARCHES); do \
		echo $$arch; \
		GOOS=darwin GOARCH=$$arch go build -o "build/bifrost-$(VERSION)-darwin-$$arch/bifrost"; \
	done

	# Build Linux
	for arch in $(LINUX_ARCHES); do \
		echo $$arch; \
		GOOS=linux GOARCH=$$arch go build -o "build/bifrost-$(VERSION)-linux-$$arch/bifrost"; \
	done

	echo "compressing build files for Darwin and Linux"
	cd build && for d in */; do filepath=$${d%/*}; echo $$filepath; zip "$${filepath##*/}.zip" "$${filepath##*/}/bifrost"; done

	mkdir -p build/windows

	# Build Windows
	for arch in $(WINDOWS_ARCHES); do \
		echo $$arch; \
		GOOS=windows GOARCH=$$arch go build -o "build/windows/bifrost-$(VERSION)-windows-$$arch/bifrost.exe"; \
	done
	echo "compressing build files for Windows"
	cd build/windows && for d in */; do filepath=$${d%/*}; echo $$filepath; zip "$${filepath##*/}.zip" "$${filepath##*/}/bifrost.exe"; done
