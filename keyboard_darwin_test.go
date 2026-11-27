package main

import "github.com/micmonay/keybd_event"

// On macOS, keybd_event.VK_DELETE is kVK_Delete (the key labelled "delete"),
// which behaves as Backspace. Forward delete is kVK_ForwardDelete.
const vkForwardDelete = keybd_event.VK_ForwardDelete
