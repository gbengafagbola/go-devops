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

func main() {

	// we want to parse the below json respone, into a go Words struct to better handle the task

	// {
	// 	page: "words",
	// 	input: "word1",
	// 	words: [
	// 		"word1"
	// 	]
	// }

	type Words struct {
		Page  string   `json:"page"`  // `json:"page"` this are metadata that can be used by packages that converts JSON string to structs and vice-versa.
		Input string   `json:"input"` // so the package/JSON-library that would convert the JSON string into a struct would look for the metadata corresponds to the attribute,
		Words []string `json:"words"` // it would then input it into the struct field.
	}

	// for example
	//		JSON					  		JSON Library											Struct
	// { page: "page" }        --> 	{<- identifies `json:"page"` tag ->}	-->		type Words struct {Page: "page"}

	// the below code is a continuation of the 01 http-get request 
	args := os.Args;

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url> \n")
		os.Exit(1)	
	}  

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("invalid url format: %s \n", err);
		os.Exit(1)
	}

	response, err := http.Get(args[1]);
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
 
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s \n", response.StatusCode, body)
		os.Exit(1)
	}

	// the code executing to this part shows no error has occured so far so we declare a variable words of type Words (the strruct defined at the top)
	var words Words
	// json library --> json.Unmarshall would parse the json string into a go struct..
	// func json.Unmarshal(data []byte, v any) error
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	// hence &words and takes in two arguments the data and pointer


	// unmarshall returns an error. so since err has been defined earlier no need for :=, hence we can override it since it's already defined 
	// for line 51 where err has been defined previously [body, err := io.ReadAll(response.Body)] since body has not be defined yet we have to use the :=
	err = json.Unmarshal(body, &words)
	  
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\n Page: %s \n Words: %s \nv \n", words.Page, strings.Join(words.Words, ", "))
	//strings.Join(words.Words, ", ") this would concantinate all the strings in an array slice
}
