SHELL = bash
OSARCHES := "darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 linux/ppc64 linux/ppc64le linux/s390x"
OUTPUT := "build/bifrost-$(VERSION)-{{.OS}}-{{.Arch}}/bifrost"

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