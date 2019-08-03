package logs

import "testing"

func TestTTKVEncoder(t *testing.T) {
	encoder := NewTTLogKVEncoder()
	encoder.Write([]byte("123 "))
	encoder.AppendKVs("name", "zyj")
	encoder.AppendKVs(456, 7891)
	encoder.AppendKVs(2)
	encoder.EndRecord()
	msg := encoder.String()
	if msg != "123 name=zyj 456=7891 \n" {
		t.Fatal("error")
	}
}
