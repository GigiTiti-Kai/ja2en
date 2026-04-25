// Package clipboard wraps atotto/clipboard for read/write operations.
// On WSL, the underlying library shells out to powershell.exe.
package clipboard

import "github.com/atotto/clipboard"

// Read returns the current clipboard text. On WSL, atotto/clipboard
// internally invokes powershell.exe Get-Clipboard which is fine for
// our use case.
func Read() (string, error) {
	return clipboard.ReadAll()
}

// Write replaces the clipboard with the given text.
func Write(text string) error {
	return clipboard.WriteAll(text)
}
