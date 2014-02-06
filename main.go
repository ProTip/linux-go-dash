// linux-go-dash project main.go
package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"github.com/protip/linux-go-dash/util/cpu"
	"github.com/protip/linux-go-dash/util/mem"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/sh/hostname", hostnameHandler)
	http.HandleFunc("/sh/uptime", uptimeHandler)
	http.HandleFunc("/sh/mem", memHandler)
	http.HandleFunc("/sh/loadavg", loadavgHandler)
	http.HandleFunc("/sh/numberofcores", numberOfCoresHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err.Error())
	}
	w.Write(indexHtml)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	var buffer bytes.Buffer
	buffer.WriteString(`"`)
	buffer.WriteString(hostname)
	buffer.WriteString(`\n"`)
	w.Write(buffer.Bytes())
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {

}

func memHandler(w http.ResponseWriter, r *http.Request) {
	total := mem.GetMB(mem.MemTotal)
	used := mem.GetUsedMB()
	free := mem.GetFreeMB()

	var response = map[string]string{
		"Total": strconv.Itoa(total),
		"Used":  strconv.Itoa(used),
		"Free":  strconv.Itoa(free),
	}

	enc := json.NewEncoder(w)
	enc.Encode(response)
}

func loadavgHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	loadavg := cpu.LoadAvg()
	numCores := cpu.Count()
	var response = map[string]string{
		"OneMinute":      strconv.FormatFloat(float64(loadavg.OneMinute), 'f', 0, 32),
		"FiveMinute":     strconv.FormatFloat(float64(loadavg.FiveMinute), 'f', 0, 32),
		"FifteenMinute":  strconv.FormatFloat(float64(loadavg.FifteenMinute), 'f', 0, 32),
		"OneMinute%":     strconv.FormatFloat(float64(loadavg.OneMinute)*100/float64(numCores), 'f', 0, 32),
		"FiveMinute%":    strconv.FormatFloat(float64(loadavg.FiveMinute)*100/float64(numCores), 'f', 0, 32),
		"FifteenMinute%": strconv.FormatFloat(float64(loadavg.FifteenMinute)*100/float64(numCores), 'f', 0, 32),
	}
	enc.Encode(response)
}

func numberOfCoresHandler(w http.ResponseWriter, r *http.Request) {
	count := cpu.Count()
	countString := strconv.Itoa(count)
	enc := json.NewEncoder(w)
	enc.Encode(countString)
}
