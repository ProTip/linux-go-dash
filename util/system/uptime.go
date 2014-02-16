package system

import "io/ioutil"
import "bytes"
import "time"

type UpTime struct {
	UpTime   time.Duration
	IdleTime time.Duration
}

func GetUpTime() *UpTime {
	upTimeBytes, err := ioutil.ReadFile("/proc/uptime")
	_ = err
	upTimeBytes = upTimeBytes[:len(upTimeBytes)-1]
	upTimeSplit := bytes.Split(upTimeBytes, []byte(" "))

	upTime, err := time.ParseDuration(string(upTimeSplit[0]) + "s")
	_ = err
	idleTime, err := time.ParseDuration(string(upTimeSplit[1]) + "s")
	_ = err

	return &UpTime{
		upTime,
		idleTime,
	}
}
