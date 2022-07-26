package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sms(Quotes string) {

	apiurl := "https://api.twilio.com/2010-04-01/Accounts/AC2eb104c9dedee690bbc937c6d56f7063/Messages.json"
	method := "POST"

	// payload := strings.NewReader("To=%2B918340535268&MessagingServiceSid=MG8d7957cb0ce6bbfd61bffd55dea8315c&Body=checking%20the%20message%20is%20sent")

	//prepare payload post req body
	data := url.Values{}
	data.Set("To", "+918340535268")
	data.Set("MessagingServiceSid", "MG8d7957cb0ce6bbfd61bffd55dea8315c")
	data.Set("Body", Quotes)

	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUMyZWIxMDRjOWRlZGVlNjkwYmJjOTM3YzZkNTZmNzA2MzpkZmIxNDhkNGMwZWEyNWEyZDExYTE5NzM2NDI1NGRjMA==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func whatsapp(Quotes string) {

	apiurl := "https://api.twilio.com/2010-04-01/Accounts/AC2eb104c9dedee690bbc937c6d56f7063/Messages.json"
	method := "POST"

	//   payload := strings.NewReader("To=whatsapp%3A%2B918340535268&From=whatsapp%3A%2B14155238886&Body=ap%20jante%20hai%20apki%20kisse%20baat%20ho%20rahi%20hai%20salkumari%20ji")

	data := url.Values{}
	data.Set("To", "+917008963113")
	data.Set("MessagingServiceSid", "MG8d7957cb0ce6bbfd61bffd55dea8315c")
	data.Set("Body", Quotes)

	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUMyZWIxMDRjOWRlZGVlNjkwYmJjOTM3YzZkNTZmNzA2MzpkZmIxNDhkNGMwZWEyNWEyZDExYTE5NzM2NDI1NGRjMA==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
