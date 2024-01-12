package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	Page string `json:"page"`
	Input string `json:clear
	"input"`
	Words []string `json:"words"`
}

func main(){
	args := os.Args;

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url> \n")
		os.Exit(1)
	}
	// ParseRequestURI checks if the passed argument is a valid url and error returns nil
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("invalid url format: %s \n", err);
		os.Exit(1)
	}

	response, err := http.Get(args[1]);

	if err != nil {
		log.Fatal(err)
	}

	// response ships with the http-status-code, the Body wrapped in the stream variable
	// The response body is streamed on demand as the Body field is read 
	defer response.Body.Close()

	// io.ReadAll reads all the data from the response body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s \n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words
	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed \nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "));


	//TWO
	mySlowReader := SlowReader {
		Contents: "hello",
	}

	out, err := io.ReadAll(mySlowReader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)
}  