# 🌈 bifrost
[![Go Report Card](https://goreportcard.com/badge/github.com/ishuah/bifrost)](https://goreportcard.com/report/github.com/ishuah/bifrost)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fishuah%2Fbifrost.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fishuah%2Fbifrost?ref=badge_shield)

Bifrost is a tiny terminal emulator for serial port communication. Supports USB type-C out of the box (2016+ Macbook friendly).

Currently supports Linux, MacOS, and Windows.

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
On MacOS:
- Unzip and copy binary to `/usr/local/bin/`
```
sudo mkdir -p /usr/local/bin
unzip bifrost-<version>-darwin-amd64.zip
cd bifrost-<version>-darwin-amd64
sudo cp bifrost /usr/local/bin/
```

On Windows:
- Unzip your bindary to a directory of your choice. Example - `C:\Users\ish\utils`
- On Powershell, create an alias for bifrost
```
// If you don't have a Powershell profile, run the following command
New-Item -Path $profile -Type File -Force

// Open your Powershell profile
notepad $PROFILE

// Create your new alias
 New-Alias -Name bifrost -Value C:\Users\ish\utils\bifrost.exe

// Save and close the file. Reload your Powershell profile by running the following command
. $PROFILE

// You should be able to run bifrost from your Powershell terminal
bifrost --help
```

- Run `bifrost -help` to confirm bifrost was installed correctly.

## Usage
Bifrost takes `-port-path` and `-baud` as parameters. By default `-port-path` is set to `/dev/tty.usbserial`
and `-baud` is set to 115200.

Example usage:

```
bifrost -port-path="/dev/ttyUSB0" -baud=128000
```

On Linux the serial port adapter path is /dev/ttyUSB0, /dev/ttyUSB1 and so on. Some USB serial port adapters may appear as /dev/ttyACM0.

On OSX/MacOS the serial port adapter path is /dev/tty.usbserial.

The default baud rate 115200 works for most serial connection but you may want to confirm the optimal baud rate for the device you're connecting to.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fishuah%2Fbifrost.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fishuah%2Fbifrost?ref=badge_large)