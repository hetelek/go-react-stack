package server

import (
	"fmt"
	"net/http"
	"os"
)

func gethostname(request *http.Request) []byte {
	hostname, err := os.Hostname()
	if err != nil {
		return []byte("error getting hostname")
	}
	return []byte(hostname)
}

// StartServer starts the server at localhost:PORT, where PORT is an env variable, or defaulted to 8000
func StartServer(staticDirectory string) {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	fs := http.FileServer(http.Dir(staticDirectory))
	http.Handle("/", http.StripPrefix("/", fs))

	RegisterHandler("/api/gethostname", HTTPGET, gethostname)

	fmt.Println("Serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
