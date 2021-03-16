package lagan

import "testing"

func TestLoad(t *testing.T) {
	_ = Load(0)
}

func TestPrint(t *testing.T) {
	_ = Load(0)
	SetFilterLevel(LevelDebug)
	EnableColor(true)
	Print("test", LevelOff, "TestPrintOut1:%d", 100)
	Print("test", LevelDebug, "TestPrintOut1:%d", 100)
	Print("test", LevelInfo, "TestPrintOut1:%d", 100)
	Print("test", LevelWarn, "TestPrintOut1:%d", 100)
	Print("test", LevelError, "TestPrintOut1:%d", 100)
}

func TestPrintHex(t *testing.T) {
	_ = Load(0)
	EnableColor(true)
	s := make([]uint8, 100)
	for i := 0; i < 100; i++ {
		s[i] = uint8(i)
	}
	PrintHex("test", LevelError, s)
}

func TestCase1(t *testing.T) {
	_ = Load(0)
	Debug("test", "case1:%d,a=%d", 101, 102)
	Print("test", LevelWarn, "case1:%d,b=%d", 101, 102)
}

func TestCase2(t *testing.T) {
	_ = Load(0)
}
