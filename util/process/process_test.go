package process

import "testing"
import "fmt"

func TestCwd(t *testing.T) {
	p := &Process{Pid: 1}
	cwd, err := p.GetCwd()
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(cwd)
	}
}

func TestCmdLine(t *testing.T) {
	p := &Process{Pid: 1}
	cmdline, err := p.GetCmdLine()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cmdline)
}
