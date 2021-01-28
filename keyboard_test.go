package main

import (
	"fmt"
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
		Enter:      TestKey{keybd_event.VK_ENTER, false},
		Esc:        TestKey{keybd_event.VK_ESC, false},
		Tab:        TestKey{keybd_event.VK_TAB, false},
		CtrlA:      TestKey{keybd_event.VK_A, true},
		CtrlB:      TestKey{keybd_event.VK_B, true},
		Space:      TestKey{keybd_event.VK_SPACE, false},
		Delete:     TestKey{keybd_event.VK_DELETE, false},
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

	kb.SetKeys(keybd_event.VK_B, keybd_event.VK_I, keybd_event.VK_F,
		keybd_event.VK_R, keybd_event.VK_O,
		keybd_event.VK_S, keybd_event.VK_T)
	err = kb.Launching()
	require.NoError(t, err)

	var testOutput []byte

	for i := 0; i < 7; i++ {
		key := pollKeyEvents()
		testOutput = append(testOutput, key.Value...)
	}

	assert.Equal(t, []byte("bifrost"), testOutput)

	for k, v := range testKeys {
		kb.SetKeys(v.Key)
		kb.HasCTRL(v.CtrlEnabled)

		err = kb.Launching()
		require.NoError(t, err)

		key := pollKeyEvents()
		assert.Equal(t, k, key.Type)
	}

	// Test CtrlC signal interrupt
	fmt.Println("\n\nPress CtrlC to proceed...")
	key := pollKeyEvents()
	assert.Equal(t, CtrlC, key.Type)

	// Test CtrlBackslash signal interrupt
	fmt.Println("Press Ctrl-\\ to proceed...")
	key = pollKeyEvents()
	assert.Equal(t, CtrlBackslash, key.Type)

	// Test Backspace signal interrupt
	fmt.Println("Press Backspace to proceed...")
	key = pollKeyEvents()
	assert.Equal(t, Backspace, key.Type)

}
