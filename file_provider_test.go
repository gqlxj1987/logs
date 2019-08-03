package logs

import (
	"os"
	"testing"
)

func TestFileProvider(t *testing.T) {
	err := os.Mkdir(TestDir, 0755)
	if err != nil && !os.IsExist(err) {
		t.Fatalf("Mkdir %s error: %s\n", TestDir, err)
	}
	provider := NewFileProvider(TestDir+"/file.log", HourDur, 10<<20)
	err = provider.Init()
	if err != nil {
		t.Fatalf("Init FileProvider error: %s\n", err)
	}
	provider.SetLevel(LevelInfo)
	if provider.level != LevelInfo {
		t.Error("Test FileProvider SetLevel error")
	}

	for i := 0; i < 1000*1000*2; i++ {
		provider.WriteMsg("Hello Log Test!\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}

func TestFileProvider2(t *testing.T) {
	provider := NewFileProvider(TestDir+"/file2.log", NoDur, 10<<20)
	if err := provider.Init(); err != nil {
		t.Fatalf("Init FileProvider2 error: %s\n", err)
	}
	for i := 0; i < 1000*1000; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}

func TestFileProvider3(t *testing.T) {
	provider := NewFileProvider(TestDir+"/file3.log", DayDur, 0)
	if err := provider.Init(); err != nil {
		t.Fatalf("Init FileProvider3 error: %s\n", err)
	}
	for i := 0; i < 100*100; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}

func TestFileProvider4(t *testing.T) {
	provider := NewFileProvider(TestDir+"/file4.log", NoDur, 0)
	if err := provider.Init(); err != nil {
		t.Fatalf("Init FileProvider4 error: %s\n", err)
	}
	for i := 0; i < 100*100; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}

func TestFileProvider5(t *testing.T) {
	provider := NewFileProvider(TestDir+"/file5.log", HourDur, 0)
	if err := provider.Init(); err != nil {
		t.Fatalf("Init FileProvider5 error: %s\n", err)
	}
	for i := 0; i < 100*100; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	//provider.preDur = time.Now().Hour() - 1
	for i := 0; i < 10*10; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}

func TestFileProvider6(t *testing.T) {
	provider := NewFileProvider(TestDir+"/file6.log", DayDur, 0)
	if err := provider.Init(); err != nil {
		t.Fatalf("Init FileProvider6 error: %s\n", err)
	}
	for i := 0; i < 100*100; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	//provider.preDur = time.Now().Day() - 1
	for i := 0; i < 10*10; i++ {
		provider.WriteMsg("Hello FileProvider Test\n", LevelInfo)
	}
	provider.Flush()
	provider.Destroy()
}
