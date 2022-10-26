package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

//We need to create a request handler. This is a func. that going to be used for every request that's made to our server.
//

func main() {
	//This tell HTTP that when a request is made we want to use our handler
	// "/" - Path name
	// Routing
	// http.HandleFunc("/", formHandler)
	http.HandleFunc("/", func(w http.formHandler, r *http.ResponseWriter))

	// Routing
	http.HandleFunc("/ascii-art", resultHandler)

	// Then we need to tell HTTP to listen and serve on port 8080
	fmt.Println("Server is starting at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("\nCannot start server")
	}
}

// w - response writer
func formHandler(w http.ResponseWriter, r *http.Request) {
	//
	if r.Method == "GET" && r.URL.Path != "/" {
		showError(w, "400 BAD REQUEST", http.StatusBadRequest)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		showError(w, "404 BANNER NOT FOUND", http.StatusNotFound)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		banner := r.FormValue("banner")
		text := r.FormValue("request")

		// Required for Unit Testing to Parse param/////////////
		// params format banner=shadow&request=A
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			params := strings.Split(string(reqBody), "&")
			for _, param := range params {
				paramSplit := strings.Split(param, "=")
				if len(paramSplit) != 2 {
					break
				}
				key := paramSplit[0]
				value := paramSplit[1]
				if key == "banner" && banner == "" {
					banner = value
				}
				if key == "request" && text == "" {
					text = value
				}
			}
		}
		////////////////////////

		b, err := readFile(banner + ".txt")
		if err != nil {
			showError(w, "404 BANNER NOT FOUND", http.StatusNotFound)
			return
		}

		myMap := parseBanner(b)

		printMessage(w, text, myMap)
		return

	}
	if r.Method == "GET" {
		showError(w, "400 BAD REQUEST", http.StatusBadRequest)
		return
	}

	showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
}

func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles("templates/error.html")
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
	}
}
