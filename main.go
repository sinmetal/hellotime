package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const backendURL = "10.6.0.10:8080"

func main() {
	res, err := http.Get(backendURL)
	if err != nil {
		fmt.Printf("%+v:%+v", time.Now(), err)
		return
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(backendURL)
	if err != nil {
		fmt.Printf("%+v:%+v", time.Now(), err)
		return
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	fmt.Fprintf(w, "Hello Backend %s", string(b))
}
