package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://quotes.rest/qod.json?category=inspire")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	strContents := string(contents)
	var jsonContents interface{}
	json.Unmarshal([]byte(strContents), &jsonContents)
	fmt.Println(strContents, jsonContents)
}
