package system

import "testing"
import "fmt"

func TestGetUptime(t *testing.T) {
	upTime := GetUpTime()
	fmt.Println(upTime)
}
