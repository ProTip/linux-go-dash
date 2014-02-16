package process

import "testing"
import _ "fmt"

func BenchmarkGetstat(b *testing.B) {
	p := &Process{Pid: 1}
	for i := 0; i < b.N; i++ {
		stat := p.GetStat()
		_ = stat
		//fmt.Println(stat)
	}
}
