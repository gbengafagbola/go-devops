### HTTP-GET-JSON REQUEST

#### this is a continuation of the  01. http-get request to capture a JSON object. let take for instance the JSON response below, we would parse this in go by first creating a type in our function to parse this JSON response from our test server.
 ``
{
    page: "words",
    input: "word1",
    words: [
        "word1"
    ]
}
``    

Using External API's,
1. make a connection request to the server (http request)
2. Retrieve the response data
3. Parse the data