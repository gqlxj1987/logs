package logs

import (
	"testing"
)

func TestConsoleProvider(t *testing.T) {
	provider := NewConsoleProvider()
	provider.SetLevel(LevelInfo)
	provider.Init()
	if provider.level != LevelInfo {
		t.Error("Test ConsoleProvider SetLevel Error")
	}
	provider.WriteMsg("Test ConsoleProvider Log Message", LevelInfo)
	provider.Flush()
	provider.Destroy()
}
