// Package input resolves the active input source (args, stdin, clipboard).
package input

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/GigiTiti-Kai/ja2en/internal/clipboard"
)

// Source describes which input channels are enabled by the caller.
type Source struct {
	Args    []string
	UseClip bool
}

// Resolve picks the active input. Precedence:
//  1. --clip explicit flag → clipboard
//  2. positional args      → args joined by space
//  3. piped stdin          → all of stdin
//
// An empty/whitespace-only result yields an error so callers can stop early.
func Resolve(s Source) (string, error) {
	if s.UseClip {
		text, err := clipboard.Read()
		if err != nil {
			return "", fmt.Errorf("read clipboard: %w", err)
		}
		text = strings.TrimSpace(text)
		if text == "" {
			return "", fmt.Errorf("clipboard is empty")
		}
		return text, nil
	}

	if len(s.Args) > 0 {
		text := strings.TrimSpace(strings.Join(s.Args, " "))
		if text != "" {
			return text, nil
		}
	}

	if isStdinPiped() {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("read stdin: %w", err)
		}
		text := strings.TrimSpace(string(data))
		if text != "" {
			return text, nil
		}
	}

	return "", fmt.Errorf("no input. pass text as argument, pipe to stdin, or use --clip")
}

func isStdinPiped() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) == 0
}
