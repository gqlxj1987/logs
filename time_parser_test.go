package logs

import (
	"strings"
	"testing"
	"time"
)

func BenchmarkDate1(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	t := time.Now()

	for i := 0; i < b.N; i++ {
		now := t.Format("2006-01-02 15:04:05.000")
		now = strings.Replace(now, ".", ",", 1)
	}
}

func BenchmarkDate2(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	t := time.Now()
	for i := 0; i < b.N; i++ {
		timeDate(t)
	}
}

func TestDate(t *testing.T) {
	MillisecondPerDay := 24 * 60 * 60 * 1000
	tm := time.Now()
	for i := 0 - 100*MillisecondPerDay; i < 1000*MillisecondPerDay; i += 1234567 {
		realTm := tm.Add(time.Duration(i) * time.Millisecond)
		val := timeDate(realTm)
		tmS := realTm.Format("2006-01-02 15:04:05.000")
		tmS = strings.Replace(tmS, ".", ",", 1)
		if tmS != string(val[:]) {
			t.Errorf("TimeDate error: expect is %s, actual is %s\n", tmS, string(val[:]))
		}
	}
}
