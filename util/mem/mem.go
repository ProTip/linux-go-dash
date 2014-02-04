package mem

import "io/ioutil"
import "bytes"
import "strconv"

const (
	HugePageSize = "Hugepagesize"
	MemTotal     = "MemTotal"
	MemFree      = "MemFree"
	Buffers      = "Buffers"
	Cached       = "Cached"
	SwapCached   = "SwapCached"
	Active       = "Active"
	Inactive     = "Inactive"
)

func GetkB(name string) (mem int) {
	return getMemInfo()[name]
}

func GetB(name string) (mem int) {
	return getMemInfo()[name] * 1024
}

func GetMB(name string) (mem int) {
	return getMemInfo()[name] / 1024
}

func GetUsedMB() (mem int) {
	return GetUsed() / 1024
}

func GetUsed() (mem int) {
	return getMemInfo()[MemTotal] -
		getMemInfo()[MemFree] -
		getMemInfo()[Buffers] -
		getMemInfo()[Cached]
}

func GetFree() (mem int) {
	return getMemInfo()[MemFree] + getMemInfo()[Cached] + getMemInfo()[Buffers]
}

func GetFreeMB() (mem int) {
	return GetFree() / 1024
}

func getMemInfo() (memInfo map[string]int) {
	memInfo = make(map[string]int)
	procMemInfo, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		panic(err.Error())
	}
	lines := bytes.Split(procMemInfo, []byte("\n"))
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		memParts := bytes.Split(line, []byte(":"))
		memName := memParts[0]
		memSizeBytes := bytes.Split(bytes.TrimSpace(memParts[1]), []byte(" "))[0]
		memSizeInt, err := strconv.Atoi(string(memSizeBytes))
		if err != nil {
			panic(err.Error())
		}
		memInfo[string(memName)] = memSizeInt
	}
	return memInfo
}
