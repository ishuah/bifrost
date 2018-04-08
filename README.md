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
On OSX/MacOS:
- Unzip and copy binary to `/usr/local/bin/`
```
sudo mkdir -p /usr/local/bin
unzip bifrost-<version>-darwin-amd64.zip
cd bifrost-<version>-darwin-amd64
sudo cp bifrost /usr/local/bin/
```

- Run `bifrost -help` to confirm bifrost was installed correctly.

## Usage
Bifrost takes `-port-path` and `-baud` as parameters. By default `-port-path` is set to `/dev/tty.usbserial`
and `-baud` is set to 115200.

Example usage:
    ```
    bifrost -port-path="/dev/ttyUSB0" -baud=115200
    ```

On Linux the serial port adapter path is /dev/ttyUSB0, /dev/ttyUSB1 and so on. Some USB serial port adapters may appear as /dev/ttyACM0.

On OSX/MacOS the serial port adapter path is /dev/tty.usbserial.