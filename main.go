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

	memInfo := make([]string, 0, 0)
	memInfo = append(memInfo, "Mem:")
	memInfo = append(memInfo, strconv.Itoa(total))
	memInfo = append(memInfo, strconv.Itoa(used))
	memInfo = append(memInfo, strconv.Itoa(free))
	enc := json.NewEncoder(w)
	enc.Encode(memInfo)
}

func numberOfCoresHandler(w http.ResponseWriter, r *http.Request) {
	count := cpu.Count()
	countString := strconv.Itoa(count)
	enc := json.NewEncoder(w)
	enc.Encode(countString)
}
