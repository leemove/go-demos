package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type maka struct {
	Month     string
	Num       int
	SafeTitle string `json:"safe_title"`
}

func main() {
	var url = "https://xkcd.com/571/info.0.json"
	imgPath := request(url)
	fmt.Println(imgPath)
}

func request(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "12314234"
	}
	var res maka
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Printf("%#v", err)
	}
	fmt.Printf("%+v\n", res)
	return "123"
}
