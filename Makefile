SHELL = bash
OSARCHES := "darwin/amd64 darwin/386 linux/386 linux/amd64 linux/arm linux/arm64 linux/mips linux/mips64 linux/mips64le linux/mipsle linux/ppc64 linux/ppc64le linux/s390x"
OUTPUT := "build/bifrost-v0.1.20-{{.OS}}-{{.Arch}}/bifrost"

build_all:
	mkdir -p build
	if [ -d "build/" ]; then \
    	rm -rf build/*; \
	fi
	gox -osarch=$(OSARCHES) -cgo -output=$(OUTPUT)