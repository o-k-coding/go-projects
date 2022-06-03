//go:build unit

// To use the build tag, remove the word "COMMENTED" this is there because vscode flips out with the build tag addeded
package data

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
