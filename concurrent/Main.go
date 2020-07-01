package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type DNS struct {
	URL     string
	Channel chan map[string]string
}

var nextInt = intSeq()

func main() {

	channel := make(chan map[string]string)

	facebokDNS := new(DNS)
	facebokDNS.Channel = channel
	facebokDNS.URL = "https://www.facebook.com/"

	googleDNS := new(DNS)
	googleDNS.Channel = channel
	googleDNS.URL = "https://www.google.com/"

	netflixDNS := new(DNS)
	netflixDNS.Channel = channel
	netflixDNS.URL = "https://www.netflix.com/"

	golangDNS := new(DNS)
	golangDNS.Channel = channel
	golangDNS.URL = "https://golang.org/"

	digitaloceanDNS := new(DNS)
	digitaloceanDNS.Channel = channel
	digitaloceanDNS.URL = "https://www.digitalocean.com/"

	cloudflareDNS := new(DNS)
	cloudflareDNS.Channel = channel
	cloudflareDNS.URL = "https://www.cloudflare.com/"

	listDNS := []DNS{
		*facebokDNS,
		*googleDNS,
		//*netflixDNS,
		*golangDNS,
		*digitaloceanDNS,
		*cloudflareDNS,
	}

	for _, dns := range listDNS {
		go HealthCheckServer(dns.URL, dns.Channel)
	}

	primero := <-channel
	segundo := <-channel
	tercero := <-channel
	cuarto := <-channel
	quinto := <-channel

	Result := []map[string]string{
		primero,
		segundo,
		tercero,
		cuarto,
		quinto,
	}

	//netflixDNS.Delay = <-channel
	//golangDNS.Delay = <-channel

	for _, value := range Result {
		fmt.Printf("%s -> %s, posicion %s\n", value["Server"], value["Delay"], value["Position"])

	}

	//fmt.Printf("%s -> %s, posicion %s\n", primero["Server"], primero["Delay"],primero["Position"])
	//fmt.Printf("%s -> %s, posicion %s\n", segundo["Server"], segundo["Delay"],segundo["Position"])

	//fmt.Printf("%s -> %v\n", netflixDNS.Name, netflixDNS.Delay)
	//fmt.Printf("%s -> %v\n", golangDNS.Name, golangDNS.Delay)

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
	for i := 0; i < 100; i++ {
		if err != nil {
			return false
		}
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, requestURI.String(), nil)

		if err != nil {
			fmt.Println(err)
			return false
		}
		res, err := client.Do(req)
		fmt.Printf("Ping # %d al server %s %s\n", i+1, requestURI.String(), res.Status)
		defer res.Body.Close()
		if res.StatusCode == 200 {
			healthy = true
		} else {
			return false
		}

	}
	return healthy
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
