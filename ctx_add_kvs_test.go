package logs

import (
	"context"
	"testing"
)

func TestAddKVs(t *testing.T) {
	ctx := context.Background()
	ctx = CtxAddKVs(ctx, "hello", "world", 123, 4.56)
	ctx = CtxAddKVs(ctx, "a", "a")
	ctx = CtxAddKVs(ctx, "b", "b")
	ctx = CtxAddKVs(ctx, "c", "c", "c") // ignored

	// Info 2018-04-24 13:52:33,495 v1(6) ctx_add_kvs_test.go:13 10.2.202.0 - - default - b=b a=a hello=world 123=4.560 bytedance
	CtxInfo(ctx, "byte%s", "dance")
	Flush()
}
