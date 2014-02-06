package cpu

import "io/ioutil"
import "bytes"
import "strconv"

type LoadAverage struct {
	OneMinute     float32
	FiveMinute    float32
	FifteenMinute float32
}

func Count() (count int) {
	presentFile, err := ioutil.ReadFile("/sys/devices/system/cpu/present")
	if err != nil {
		panic(err.Error())
	}
	presentFile = presentFile[:len(presentFile)-1]
	presentSplit := bytes.Split(presentFile, []byte("-"))
	count = len(presentSplit)
	return count
}

func LoadAvg() *LoadAverage {
	loadavgFile, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err.Error())
	}

	loadavgFile = loadavgFile[:len(loadavgFile)-1]
	loadavgSplit := bytes.Split(loadavgFile, []byte(" "))

	oneMinute, err := strconv.ParseFloat(string(loadavgSplit[0]), 32)
	_ = err
	fiveMinute, err := strconv.ParseFloat(string(loadavgSplit[1]), 32)
	_ = err
	fifteenMinute, err := strconv.ParseFloat(string(loadavgSplit[2]), 32)
	_ = err
	loadAvg := &LoadAverage{
		float32(oneMinute),
		float32(fiveMinute),
		float32(fifteenMinute),
	}
	return loadAvg
}
