package process

import "fmt"
import "io/ioutil"
import "strconv"

const ()

type Stat struct {
	Pid               int
	Command           string
	State             byte
	PPID              int
	PGRP              int
	Session           int
	TTY               int
	TPGID             int
	Flags             uint
	MinFlt            uint
	CMinFlt           uint
	MajFlt            uint
	CMajFlt           uint
	UTime             uint
	STime             uint
	CuTime            int
	CsTime            int
	Priority          int
	Nice              int
	Threads           int
	ItRealValue       uint
	StartTime         uint
	VSize             uint
	RSS               int
	RSSLim            uint
	StartCode         uint
	EndCode           uint
	StartStack        uint
	KStkESP           uint
	KStkEIP           uint
	Signal            uint
	Blocked           uint
	SigIgnore         uint
	SigCatch          uint
	WChan             uint
	NSwap             uint
	CNSwap            uint
	ExitSignal        int
	Processor         int
	RTPriority        uint
	Policy            uint
	DelayedBlkIOTicks uint
	GuestTime         uint
	CGuestTime        int
}

func (p *Process) GetStat() *Stat {
	stat := &Stat{}
	statBytes, err := ioutil.ReadFile("/proc/" + strconv.Itoa(p.Pid) + "/stat")
	_ = err
	statBytes = statBytes[:len(statBytes)-1]
	num, err := fmt.Sscanf(string(statBytes),
		`%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d`,
		&stat.Pid, &stat.Command, &stat.State, &stat.PPID, &stat.PGRP, &stat.Session, &stat.TTY, &stat.TPGID, &stat.Flags, &stat.MinFlt,
		&stat.CMinFlt, &stat.MajFlt, &stat.CMajFlt, &stat.UTime, &stat.STime, &stat.CuTime, &stat.CsTime, &stat.Priority, &stat.Nice,
		&stat.Threads, &stat.ItRealValue, &stat.StartTime, &stat.VSize, &stat.RSS, &stat.RSSLim, &stat.StartCode, &stat.EndCode, &stat.StartStack,
		&stat.KStkESP, &stat.KStkEIP, &stat.Signal, &stat.Blocked, &stat.SigIgnore, &stat.SigCatch, &stat.WChan, &stat.NSwap, &stat.CNSwap,
		&stat.ExitSignal, &stat.Processor, &stat.RTPriority, &stat.Policy, &stat.DelayedBlkIOTicks, &stat.GuestTime, &stat.CGuestTime)
	_ = num
	_ = err
	return stat
}
