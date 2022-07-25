package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func zenquotes2() string {

	url := "https://zenquotes.io/api/random"
	method := "GET"

	payload := strings.NewReader(`[ {"q":"When asked, how do you write? I invariably answer, one word at a time.","a":"Stephen King","h":"<blockquote>&ldquo;When asked, how do you write? I invariably answer, one word at a time.&rdquo; &mdash; <footer>Stephen King</footer></blockquote>"} ]`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		// return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(string(body))

	a := []Quotes{}
	json.Unmarshal(body, &a)
	return a[0].Q
}
