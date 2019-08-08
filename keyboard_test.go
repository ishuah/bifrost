package main

import (
	"runtime"
	"testing"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestKey struct {
	Key         int
	CtrlEnabled bool
}
type TestKeyMap map[KeyType]TestKey

func TestPollKeyEvents(t *testing.T) {
	testKeys := TestKeyMap{
		Enter: TestKey{keybd_event.VK_ENTER, false},
		Esc:   TestKey{keybd_event.VK_ESC, false},
		Tab:   TestKey{keybd_event.VK_TAB, false},
		CtrlA: TestKey{keybd_event.VK_A, true},
		CtrlB: TestKey{keybd_event.VK_B, true},
		//CtrlC:      TestKey{keybd_event.VK_C, true},
		//CtrlBackslash: TestKey{keybd_event.VK_BACKSLASH, true},
		Space:      TestKey{keybd_event.VK_SPACE, false},
		UpArrow:    TestKey{keybd_event.VK_UP, false},
		DownArrow:  TestKey{keybd_event.VK_DOWN, false},
		LeftArrow:  TestKey{keybd_event.VK_LEFT, false},
		RightArrow: TestKey{keybd_event.VK_RIGHT, false},
	}
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		t.Skipf("Could not run KeyBonding: %v", err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)
	err = kb.Launching()
	require.NoError(t, err)

	key := pollKeyEvents()
	assert.Equal(t, []byte("ab"), key.Value)

	for k, v := range testKeys {
		kb.SetKeys(v.Key)
		kb.HasCTRL(v.CtrlEnabled)

		err = kb.Launching()
		require.NoError(t, err)

		key = pollKeyEvents()
		assert.Equal(t, k, key.Type)
	}

}
