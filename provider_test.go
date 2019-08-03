package logs

import (
	"testing"
)

func TestLogProvider(t *testing.T) {
	var _ LogProvider = NewConsoleProvider()
	var _ LogProvider = NewFileProvider(TestDir+"/test.log", NoDur, 0)
}
