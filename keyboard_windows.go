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
		case keys.Tab:
			connect.Write([]byte{'\x09'})
		case keys.CtrlC:
			connect.Write([]byte{'\x03'})
		case keys.Enter:
			connect.Write([]byte{'\r'})
		case keys.Backspace:
			connect.Write([]byte{'\x08'})
		case keys.Delete:
			connect.Write([]byte{'\x08'})
		case keys.Left:
			connect.Write([]byte{'\x1b', '[', 'D'})
		case keys.Right:
			connect.Write([]byte{'\x1b', '[', 'C'})
		case keys.Up:
			connect.Write([]byte{'\x1b', '[', 'A'})
		case keys.Down:
			connect.Write([]byte{'\x1b', '[', 'B'})
		case keys.Space:
			connect.Write([]byte{' '})
		case keys.CtrlA:
			connect.Write([]byte{'\x01'})
		case keys.CtrlB:
			connect.Write([]byte{'\x02'})
		case keys.CtrlD:
			connect.Write([]byte{'\x04'})
		case keys.CtrlE:
			connect.Write([]byte{'\x05'})
		case keys.CtrlF:
			connect.Write([]byte{'\x06'})
		case keys.CtrlK:
			connect.Write([]byte{'\x0b'})
		case keys.CtrlU:
			connect.Write([]byte{'\x15'})
		case keys.RuneKey:
			connect.Write([]byte(string(key.Runes)))
		}

		return false, nil
	})
}
