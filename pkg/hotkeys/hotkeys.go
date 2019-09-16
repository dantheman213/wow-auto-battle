package hotkeys

import (
    "fmt"
    "github.com/JamesHovious/w32"
)

func Register() error {
    err := w32.RegisterHotKey(0, 1, w32.MOD_NOREPEAT, w32.VK_DELETE)
    if err != nil {
        fmt.Println("Error registering hotkey:", err.Error())
        return err
    }

    return nil
}
