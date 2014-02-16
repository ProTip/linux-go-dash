package cpu

import "testing"
import "fmt"

func TestLoadAvg(t *testing.T) {
	fmt.Println(LoadAvg())
}
