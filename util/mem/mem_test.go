package mem

import "testing"
import _ "fmt"

func TestGetkB(t *testing.T) {
	hugePageSize := GetkB(HugePageSize)
	if hugePageSize == 2048 {
		return
	} else {
		t.Fatalf("Not working!")
	}
}

func TestGetB(t *testing.T) {
	hugePageSize := GetB(HugePageSize)
	if hugePageSize != 2097152 {
		t.Fatalf("Bytes not correct!")
	}
}

func TestGetMB(t *testing.T) {
	hugePageSize := GetMB(HugePageSize)
	if hugePageSize != 2 {
		t.Fatalf("MegaBytes not correct!")
	}
}

func TestGetMemInfo(t *testing.T) {
	memInfo := getMemInfo()
	if memInfo[HugePageSize] != 2048 {
		t.Fatalf("No match!")
	}
}
