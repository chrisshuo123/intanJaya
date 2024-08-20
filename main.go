package main

import (
	"log"
	"net/http"
)

func ServeStatic(route string, directory string) {
	// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at directory
	fs := http.FileServer(http.Dir(directory))
	// StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the request URL's Path
	http.Handle(route, http.StripPrefix(route, fs))
}

func main() {
	// serve static files for set1 at /static1/ path
	ServeStatic("/static1/", "./static/set1")

	// serve static files for set2 at /static2/ path
	ServeStatic("/static2/", "./static/set2")

	// Default handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Switch between static sets based on URL query
		switch r.URL.Query().Get("set") {
		case "2":
			http.ServeFile(w, r, "templates/index2.html")
		default:
			http.ServeFile(w, r, "templates/index1.html")
		}
	})

	// Start the server
	log.Println("Listening on : 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
