package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	casesMap := make(map[string]string)

	str, err := readFile("example00.txt")
	if err == nil {
		casesMap["{123}/n<Hello> (World)!"] = str
	}

	casesMap["123"] = `                    
 _   ____    _____  
/ | |___ \  |___ /  
| |   __) |   |_ \  
| |  / __/   ___) | 
|_| |_____| |____/  
                    
                    
`

	for k, v := range casesMap {

		data := url.Values{}
		data.Set("request", k)
		data.Set("banner", "standard")
		writer := strings.NewReader(data.Encode())
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/ascii-art", writer)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(resultHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		fmt.Println(rr.Code)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := v
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
