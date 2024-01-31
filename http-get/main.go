package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main(){
	// capturing argument passed in the cli can be performed with the aid of os package 
	//where Args holds the command line arguments starting with the program name i.e [main argument1]
	args := os.Args;

	//to ensure we are always making use of the first argument passed, 
	//that is the second item in the array after the name of the program. 
	//an if check would suffice ensuring that we have at least 2 items else the code exits with a status of 1
	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url> \n")
	// >`echo $?` would return the exit code of the last command ran in your terminal
		os.Exit(1)	
	}  

	// ParseRequestURI checks if the passed argument is a valid url and returns the url & error 
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("invalid url format: %s \n", err);
		os.Exit(1)
	}

	//once we are sure the second value of Args is a valid uri we then perform our get request which returns a url and error
	response, err := http.Get(args[1]);
	// if error since it a system error not a user error then log the error
	if err != nil {
		log.Fatal(err)
	}

	// if no error, response would have the response data (http-status-code, body)
	// response ships with the http-status-code and the Body which can be bigger than the avalable memory, 
	// so the body might not contain all the data from the http output, so it wrapped in a different varible called the stream and once our program exist,
	// we need to close that stream and to do that we invoke the following command which the defer would execute the function after all other functions have ran successfully
	// The response body is streamed on demand as the Body field is streamed on demand (continously) hence need to close the stream of the response body
	defer response.Body.Close()
 
	//since we now know our body would fit into our available memory, then we can read the data all at once
	// io.ReadAll reads all the data from the response body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s \n", response.StatusCode, body)
		os.Exit(1)
	}

	fmt.Printf("HTTP Status code: %d\n Body:%s\n", response.StatusCode, body)
}


// running our server locally; 
//> go run main.go http://localhost:8080/words
//