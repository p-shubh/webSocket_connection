package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func zenquotes() Quotes {

// 	url := "https://zenquotes.io/api/random"
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)

// 	if err != nil {
// 		fmt.Println(err)
// 		// return
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		// return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		// return
// 	}
// 	// fmt.Println(string(body))

// 	// Unmarshal(data []byte, v interface{}) error
// 	var a Quotes
// 	json.Unmarshal(body, &a)
// 	return a
// }
