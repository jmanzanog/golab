package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type DNS []struct {
	URL string `json:"url"`
}

var nextInt = intSeq()
var channel = make(chan map[string]string)
var methodGET = "GET"

func main() {

	content, err := ioutil.ReadFile("./concurrent/dns.js")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents:\n%s", content)

	var listDNS DNS
	if err := json.Unmarshal(content, &listDNS); err != nil {
		panic(err)
	}
	fmt.Println(listDNS)

	for _, dns := range listDNS {
		go HealthCheckServer(dns.URL, channel)
	}
	WaitAndPrintGoRoutines(len(listDNS), channel)

}

func WaitAndPrintGoRoutines(goRoutineStackLen int, channel chan map[string]string) {
	var listValues []map[string]string
	for i := 0; i < goRoutineStackLen; i++ {
		listValues = append(listValues, <-channel)
	}
	for _, value := range listValues {
		fmt.Printf("%s -> %s, posicion %s\n", value["Server"], value["Delay"], value["Position"])

	}
}

func HealthCheckServer(serverDns string, channel chan map[string]string) (healthy bool) {
	init := time.Now()
	defer func() {
		channel <- map[string]string{
			"Delay":    time.Now().Sub(init).String(),
			"Position": strconv.Itoa(nextInt()),
			"Server":   serverDns}
	}()
	requestURI, err := url.ParseRequestURI(serverDns)
	if err != nil {
		return false
	}
	client := &http.Client{}
	req, err := http.NewRequest(methodGET, requestURI.String(), nil)
	if err != nil {
		return errorFunc(err)
	}
	for i := 0; i < 100; i++ {
		if res, err := client.Do(req); err != nil {
			return errorFunc(err)
		} else {
			fmt.Printf("Ping # %d to server %s %s\n", i+1, requestURI.String(), res.Status)
			//defer res.Body.Close()
			if res.StatusCode == 200 {
				healthy = true
			} else {
				return false
			}
		}

	}
	return healthy
}

func errorFunc(err error) bool {
	log.Fatal(err)
	return false
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
