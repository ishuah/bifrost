package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

const version = "v0.2.2"

var header = fmt.Sprintf("\nBifrost %s\n", version)
var helpText = fmt.Sprintf(`%s
Bifrost is a tiny terminal emulator for serial port communication.

    Usage:
	  bifrost [flags]
	  
    Flags:
      -port-path	Name/path of the serial port
      -baud		The baud rate to use on the connection
      -save-config	Save a connection configuration
      -load-config	Load a connection configuration
      -help		This help message
	`, header)

var configDir = os.Getenv("HOME") + "/.bifrost/"

const configFile = "config.ini"

func welcomeMessage(portPath string, baud int) string {
	return fmt.Sprintf(`%s
Options:
    Port:	%s
    Baud rate:	%d

Press Ctrl+\ to exit
		`, header, portPath, baud)
}

func writeConfig(configName string, portPath string, baud int) error {
	configPath := configDir + configFile

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(configPath, []byte(""), os.ModePerm)
		if err != nil {
			return err
		}
	}

	config, err := ini.Load(configPath)
	if err != nil {
		return err
	}

	_, err = config.NewSection(configName)
	if err != nil {
		return err
	}

	config.Section(configName).Key("port").SetValue(portPath)
	config.Section(configName).Key("baud").SetValue(strconv.Itoa(baud))

	config.SaveTo(configPath)
	return nil
}

func readConfig(configName string) (portPath string, baud int, err error) {
	configPath := configDir + configFile

	config, err := ini.Load(configPath)
	if err != nil {
		return
	}

	section, err := config.GetSection(configName)
	if err != nil {
		return
	}

	portPath = section.Key("port").String()
	baud, _ = section.Key("baud").Int()

	return
}

func main() {
	var portPath string
	var baud int
	var saveConfig bool
	var loadConfig string
	var help bool

	flag.StringVar(&portPath, "port-path", "/dev/tty.usbserial", "Name/path of the serial port")
	flag.IntVar(&baud, "baud", 115200, "The baud rate to use on the connection")
	flag.BoolVar(&saveConfig, "save-config", false, "Save a connection configuration")
	flag.StringVar(&loadConfig, "load-config", "", "Load a connection configuration")
	flag.BoolVar(&help, "help", false, "A brief help message")
	flag.Parse()

	if saveConfig {
		var configName string
		fmt.Println("What name do you want to save this config under?")
		fmt.Scanln(&configName)
		err := writeConfig(configName, portPath, baud)
		if err != nil {
			log.Printf("Failed to save config. FatalError: %v\n", err)
			return
		}
		fmt.Printf("Config saved! You can view and edit your configurations at %s%s.\n", configDir, configFile)
		return
	}

	if loadConfig != "" {
		fmt.Printf("Loading config %s...\n", loadConfig)
		cfgPortPath, cfgBaud, err := readConfig(loadConfig)
		if err != nil {
			log.Printf("Failed to load config. FatalError: %v\n", err)
			return
		}
		portPath = cfgPortPath
		baud = cfgBaud
		fmt.Println("Config loaded successfully.")
	}

	if help {
		fmt.Println(helpText)
		return
	}

	connect, err := NewConnection(portPath, baud)
	if err != nil {
		log.Printf("FatalError: %v\n", err)
		return
	}
	// Welcome message
	fmt.Print(welcomeMessage(portPath, baud))

	go connect.Start()

	KeyboardListener(connect)
}
