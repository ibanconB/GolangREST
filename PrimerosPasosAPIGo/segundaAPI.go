package main

import (
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func main() {
	currencies := "EURUSD, GBPUSD"
	api_key := "vLh6BpRvNWr8baJaGBYl"
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(body))

}
