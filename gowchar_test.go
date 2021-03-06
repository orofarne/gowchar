package gowchar

import (
	"fmt"
	"testing"
)

func TestGowcharSimple(t *testing.T) {
	str1 := "Привет, 世界. 𪛖"
	wstr, size := StringToWcharT(str1)
	switch SIZEOF_WCHAR_T {
	case 2:
		if size != 15 {
			t.Errorf("size(%v) != 15", size)
		}
	case 4:
		if size != 14 {
			t.Errorf("size(%v) != 14", size)
		}
	default:
		panic(fmt.Sprintf("sizeof(wchar_t) = %v", SIZEOF_WCHAR_T))
	}
	str2, err := WcharTToString(wstr)
	if err != nil {
		t.Errorf("WcharTToString error: %v", err)
	}
	if str1 != str2 {
		t.Errorf("\"%s\" != \"%s\"", str1, str2)
	}
}

func TestGowcharSimpleN(t *testing.T) {
	str1 := "Привет, 世界. 𪛖"
	wstr, size := StringToWcharT(str1)
	switch SIZEOF_WCHAR_T {
	case 2:
		if size != 15 {
			t.Errorf("size(%v) != 15", size)
		}
	case 4:
		if size != 14 {
			t.Errorf("size(%v) != 14", size)
		}
	default:
		panic(fmt.Sprintf("sizeof(wchar_t) = %v", SIZEOF_WCHAR_T))
	}

	str2, err := WcharTNToString(wstr, size-1)
	if err != nil {
		t.Errorf("WcharTToString error: %v", err)
	}
	if str1 != str2 {
		t.Errorf("\"%s\" != \"%s\"", str1, str2)
	}
}
