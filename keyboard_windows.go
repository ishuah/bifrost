//go:build windows

package main

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func KeyboardListener(connect *Connect) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Esc:
			connect.Write([]byte{'\x1b'})
		case keys.CtrlBackslash:
			fmt.Println("\nbye!")
			return true, nil
		case keys.RuneKey:
			connect.Write([]byte(string(key.Runes)))
		}

		return false, nil
	})
}