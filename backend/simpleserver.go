package server

import "net/http"

var handlers = map[string]map[HTTPMethod]MethodHandler{}

// HTTPMethod represents an HTTP method
type HTTPMethod string

// All available HTTP methods
const (
	HTTPGET      HTTPMethod = "GET"
	HTTPPOST     HTTPMethod = "POST"
	HTTPPUT      HTTPMethod = "PUT"
	HTTPPATCH    HTTPMethod = "PATCH"
	HTTPDELETE   HTTPMethod = "DELETE"
	HTTPCOPY     HTTPMethod = "COPY"
	HTTPHEAD     HTTPMethod = "HEAD"
	HTTPOPTIONS  HTTPMethod = "OPTIONS"
	HTTPLINK     HTTPMethod = "LINK"
	HTTPUNLINK   HTTPMethod = "UNLINK"
	HTTPPURGE    HTTPMethod = "PURGE"
	HTTPLOCK     HTTPMethod = "LOCK"
	HTTPUNLOCK   HTTPMethod = "UNLOCK"
	HTTPPROPFIND HTTPMethod = "PROPFIND"
	HTTPVIEW     HTTPMethod = "VIEW"
)

// MethodHandler is the definition for a function that handles a HTTP method request
type MethodHandler func(*http.Request) []byte

// RegisterHandler registers the given `methodHandler` with the given `pattern` and HTTP `method`
func RegisterHandler(pattern string, method HTTPMethod, methodHandler MethodHandler) {
	patternHandlers, ok := handlers[pattern]
	if !ok {
		patternHandlers = map[HTTPMethod]MethodHandler{}
		handlers[pattern] = patternHandlers

		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			methodHandler := handlers[pattern][HTTPMethod(r.Method)]
			if methodHandler != nil {
				w.Write(methodHandler(r))
			} else {
				http.NotFound(w, r)
			}
		})
	}
	handlers[pattern][method] = methodHandler
}
