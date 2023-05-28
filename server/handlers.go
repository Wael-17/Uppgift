package server

import "net/http"

// helloHandler returns a simple HTTP handler function which writes a response.
func (s server) helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Det är 24+ grader ute och jag sitter och gör denna uppgift"))
	})
}
