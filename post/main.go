package main

import (
	"log"
	"net/http"
	"net/url"
)

func main(){
	values := url.Values{
		"test":{"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values) //x-www-form-urlencoded
	if err != nil{
		panic(err)
	}
	log.Println("status:", resp.Status)
}