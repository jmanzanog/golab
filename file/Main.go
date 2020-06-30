package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	//fileinfo, err := ioutil.ReadDir(".")

	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, value := range fileinfo {
	//
	//	fmt.Print(value)
	//}

	content, err := ioutil.ReadFile("./file/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents:\n%s", content)


}
