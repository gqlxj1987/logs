package logs

import (
	"fmt"
	"sync"
	"testing"
)

func TestPushNotice(t *testing.T) {
	ntc := NewNoticeKVs()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			for j := 0; j < 10; j++ {
				x := fmt.Sprintf("%v_%v", id, j)
				ntc.PushNotice(x, x)
			}
			wg.Done()
		}(i)
	}

	kvMap := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			x := fmt.Sprintf("%v_%v", i, j)
			kvMap[x] = struct{}{}
		}
	}

	wg.Wait()
	kvs := ntc.KVs()
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		v := kvs[i+1]
		if k != v {
			t.Fatal("err")
		}

		str := k.(string)
		delete(kvMap, str)
	}

	if len(kvMap) != 0 {
		fmt.Println(len(kvMap))
		t.Fatal("err")
	}
}
