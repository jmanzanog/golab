package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var wg sync.WaitGroup

type DNS struct {
	Name    string
	Url     string
	Channel chan time.Duration
	Delay   time.Duration
}

func main() {

	channel := make(chan time.Duration)

	facebokDNS := new(DNS)
	facebokDNS.Channel = channel
	facebokDNS.Name = "Facebook"
	facebokDNS.Url = "https://www.facebook.com/"

	googleDNS := new(DNS)
	googleDNS.Channel = channel
	googleDNS.Name = "Google"
	googleDNS.Url = "https://www.google.com/"

	netflixDNS := new(DNS)
	netflixDNS.Channel = channel
	netflixDNS.Name = "Netflix"
	netflixDNS.Url = "https://www.netflix.com/"

	listDNS := []DNS{*facebokDNS, *googleDNS, *netflixDNS}
	print(len(listDNS))

	go HealthCheckServer(googleDNS.Url, googleDNS.Channel)
	go HealthCheckServer(facebokDNS.Url, facebokDNS.Channel)
	go HealthCheckServer(netflixDNS.Url, netflixDNS.Channel)

	googleDNS.Delay = <-channel
	facebokDNS.Delay = <-channel
	netflixDNS.Delay = <-channel

	fmt.Printf("%s -> %v\n", googleDNS.Name, googleDNS.Delay)
	fmt.Printf("%s -> %v\n", facebokDNS.Name, facebokDNS.Delay)
	fmt.Printf("%s -> %v\n", netflixDNS.Name, netflixDNS.Delay)

}
func HealthCheckServer(serverDns string, channel chan time.Duration) (healthy bool) {

	init := time.Now()
	defer func() { channel <- time.Now().Sub(init) }()
	requestURI, err := url.ParseRequestURI(serverDns)
	for i := 0; i < 10; i++ {
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
