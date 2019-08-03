package logs

import (
	"testing"

	"github.com/cgCodeLife/logs/context"
)

func TestDefaultLogger(t *testing.T) {
	SetLevel(LevelTrace)
	AddProvider(NewConsoleProvider())
	Trace("Trace")
	Debug("Debug")
	Info("Info")
	Notice("Notice")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")
	Flush()
	t.Log("Test Logger Pass.")
}

func TestCtxLogger(t *testing.T) {
	ctx := context.Background()
	CtxTrace(ctx, "Trace")
	CtxDebug(ctx, "Debug")
	CtxInfo(ctx, "Info")
	CtxNotice(ctx, "Notice")
	CtxWarn(ctx, "Warn")
	CtxError(ctx, "Error")
	CtxFatal(ctx, "Fatal")
	ctx = context.Background()
	ctx = context.WithValue(ctx, "K_LOCALIP", "127.0.0.1")
	ctx = context.WithValue(ctx, "K_SNAME", "service_name")
	ctx = context.WithValue(ctx, "K_LOGID", "logid")
	CtxTraceKvs(ctx, "name", "anni", "msg", "am i here?")
	CtxWarnKvs(ctx, "name", "anni", "msg", "am i here?")
	CtxInfoKvs(ctx, "name", "anni", "msg", "am i here?")
	CtxNoticeKvs(ctx, "name", "anni", "msg", "am i here?")
	CtxDebugKvs(ctx, "name", "anni", "msg", "am i here?")
	CtxFatalKvs(ctx, "name", "anni", "msg", "am i here?")
	Flush()
	Stop()
	t.Log("Test Ctx logger Pass.")
}

func TestCtxDynamicLogger(t *testing.T) {
	SetLevel(LevelInfo)
	Trace("Trace")
	Debug("Debug")
	Info("Info")
	Info("UserDynamicLevelDebug")
	ctx := context.WithValue(context.Background(), DynamicLogLevelKey, LevelDebug)
	CtxTrace(ctx, "CtxTrace")
	CtxDebug(ctx, "CtxDebug")
	CtxInfo(ctx, "CtxInfo")
	Info("Enable DynamicLog")
	EnableDynamicLogLevel()
	CtxTrace(ctx, "CtxTrace")
	CtxDebug(ctx, "CtxDebug")
	CtxInfo(ctx, "CtxInfo")
	CtxNotice(ctx, "CtxNotice")
	CtxWarn(ctx, "CtxWarn")
	CtxError(ctx, "CtxError")
	CtxFatal(ctx, "CtxFatal")
	Info("UserLevelError")
	ctx = context.WithValue(ctx, DynamicLogLevelKey, LevelError)
	CtxTrace(ctx, "CtxTrace")
	CtxDebug(ctx, "CtxDebug")
	CtxInfo(ctx, "CtxInfo")
	CtxNotice(ctx, "CtxNotice")
	CtxWarn(ctx, "CtxWarn")
	CtxError(ctx, "CtxError")
	CtxFatal(ctx, "CtxFatal")
	Flush()
	Stop()
	t.Log("Test CtxDynamicLogger Pass.")
}
