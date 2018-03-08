# 🌈 bifrost
[![Go Report Card](https://goreportcard.com/badge/github.com/ishuah/bifrost)](https://goreportcard.com/report/github.com/ishuah/bifrost)

Bifrost is a tiny terminal emulator for serial port communication.

Note: Only Linux and OSX are currenly supported. Windows will be supported in subsequent releases.

## Installation
- Download the latest version from the releases page (https://github.com/ishuah/bifrost/releases)

On linux:
- Unzip and copy binary to `/usr/bin/`
```
unzip bifrost-<version>-linux-amd64.zip
cd bifrost-<version>-linux-amd64
sudo cp bifrost /usr/bin/
sudo chown root:root /usr/bin/bifrost
sudo chmod 755 /usr/bin/bifrost
```
On OSX:
- Unzip and copy binary to `/usr/local/bin/`
```
sudo mkdir -p /usr/local/bin
unzip bifrost-<version>-darwin-amd64.zip
cd bifrost-<version>-darwin-amd64
sudo cp bifrost /usr/local/bin/
```

- Run `bifrost -help` to confirm bifrost was installed correctly.