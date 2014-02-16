package process

import "io/ioutil"
import "strconv"
import "os"

type Process struct {
	Pid     int
	cmdline string
}

func (p *Process) GetCmdLine() (string, error) {
	if p.cmdline == "" {
		cmdline, err := ioutil.ReadFile("/proc/" + strconv.Itoa(p.Pid) + "/cmdline")
		if err == nil {
			p.cmdline = string(cmdline)
			return p.cmdline, nil
		} else {
			return "", err
		}
	} else {
		return p.cmdline, nil
	}
}

func (p *Process) GetCwd() (string, error) {
	return os.Readlink("/proc/" + strconv.Itoa(p.Pid) + "/cwd")
}
