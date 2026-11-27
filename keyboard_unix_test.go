//go:build !windows && !darwin

package main

import "github.com/micmonay/keybd_event"

const vkForwardDelete = keybd_event.VK_DELETE
