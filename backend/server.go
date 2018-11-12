package server

import (
	"fmt"
	"net/http"
	"os"
)

func index(request *http.Request) []byte {
	return []byte("Test " + request.Method)
}

// StartServer starts the server at localhost:PORT, where PORT is an env variable, or defaulted to 8000
func StartServer(staticDirectory string) {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	fs := http.FileServer(http.Dir(staticDirectory))
	http.Handle("/", http.StripPrefix("/", fs))

	RegisterHandler("/api/", HTTPGET, index)

	fmt.Println("Serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
