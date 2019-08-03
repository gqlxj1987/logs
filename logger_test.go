package logs

import (
	"bytes"
	"testing"

	"github.com/cgCodeLife/logs/context"
)

const TestDir = "testdata"

func TestNewLogger(t *testing.T) {
	logger := NewLogger(1024)
	logger.SetLevel(LevelTrace)

	fileProvider := NewFileProvider(TestDir+"/test.log", HourDur, 8<<20)
	fileProvider.SetLevel(LevelInfo)
	if err := logger.AddProvider(fileProvider); err != nil {
		t.Fatalf("Add fileProvider error: %s\n", err)
	}

	consoleProvider := NewConsoleProvider()
	consoleProvider.SetLevel(LevelDebug)
	logger.AddProvider(consoleProvider)
	logger.SetCallDepth(3)

	logger.StartLogger()
	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.DisableCallDepth()
	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Flush()
	logger.Stop()

	logger = NewLogger(1024)
	logger.AddProvider(fileProvider)
	logger.StartLogger()
	for i := 0; i < 100*00; i++ {
		logger.Info("Info")
	}
	logger.Flush()
	logger.Stop()
	for i := 0; i < 10; i++ {
		logger.Flush()
		logger = NewLogger(1024)
		logger.SetLevel(LevelDebug)
		InitLogger(logger)
		logger.StartLogger()
	}
	logger.Flush()
	logger.Stop()
}

func BenchmarkLogger(b *testing.B) {
	discardProvider := NewDiscardProvider()
	discardProvider.SetLevel(LevelTrace)
	AddProvider(discardProvider)
	format := "Benchmark logger like %s, you know !"
	bs := make([]byte, 200)
	for i := 0; i < 200; i++ {
		bs[i] = 'a'
	}
	value := string(bs)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "K_LOCALIP", "127.0.0.1")
	ctx = context.WithValue(ctx, "K_SNAME", "service_name")
	ctx = context.WithValue(ctx, "K_LOGID", "logid")
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CtxInfo(ctx, format, value)
		// CtxInfoKvs(ctx, "name", "anni", "msg", value)
	}
}

func BenchmarkPrefix(b *testing.B) {
	b.StartTimer()
	b.ReportAllocs()
	logger := NewLogger(1024)
	logger.SetLevel(LevelTrace)
	buf := bytes.NewBuffer(make([]byte, 1024))
	for i := 0; i < b.N; i++ {
		buf.Reset()
		logger.prefixV1(nil, LevelError, "test:0", buf)
	}
}
