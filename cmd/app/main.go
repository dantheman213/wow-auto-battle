package main

import (
	"github.com/JamesHovious/w32"
	"github.com/dantheman213/wow-auto-battle/pkg/hotkeys"
)

func main() {
	_ = hotkeys.Register()
	_ = waitForInputLoop()
}

func waitForInputLoop() (err error) {
	var msg w32.MSG
	quit := false

	for !quit {
		_ = w32.GetMessage(&msg, 0,0,0)

		switch msg.Message {
		case w32.WM_HOTKEY:
			if msg.WParam == 1 {
				// TODO
				w32.MessageBox(0, "Test","Test", 0)
			}
		}
	}

	return nil
}
