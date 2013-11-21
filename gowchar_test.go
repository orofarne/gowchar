package gowchar

import (
	"testing"
)

func TestGowcharSimple(t *testing.T) {
	str1 := "Привет, 世界"
	wstr, size := StringToWcharT(str1)
	if size != 11 {
		t.Errorf("size(%v) != 11", size)
	}
	str2, err := WcharTToString(wstr)
	if err != nil {
		t.Error("WcharTToString error: %v", err)
	}
	if str1 != str2 {
		t.Errorf("\"%s\" != \"%s\"", str1, str2)
	}
}

func TestGowcharSimpleN(t *testing.T) {
	str1 := "Привет, 世界"
	wstr, size := StringToWcharT(str1)
	if size != 11 {
		t.Errorf("size(%v) != 11", size)
	}
	str2, err := WcharTNToString(wstr, 10)
	if err != nil {
		t.Error("WcharTToString error: %v", err)
	}
	if str1 != str2 {
		t.Errorf("\"%s\" != \"%s\"", str1, str2)
	}
}
