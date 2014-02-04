package cpu

import "io/ioutil"
import "bytes"

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
