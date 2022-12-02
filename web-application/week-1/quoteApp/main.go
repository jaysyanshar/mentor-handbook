package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Quotes struct {
	Quote     string `json:"q"`
	Speaker   string `json:"a"`
	QuoteHTML string `json:"h"`
}

func main() {

	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	req, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
	if err != nil {
		panic(err)
	}

	// Request dengan method `Do()` pada instance http.Client yang sudah dibuat
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Cetak status code request
	fmt.Println("Status: ", resp.Status)

	// Kita bisa membaca response body menggunakan package ioutil.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert body type menjadi string dan cetak
	fmt.Println(string(responseData))

	var quote []Quotes
	var err2 = json.Unmarshal(responseData, &quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(quote[0].Quote))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><head><title>My quote</title></head><body>"+quote[0].QuoteHTML+"</body></html>")
	})

	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", nil)

}
