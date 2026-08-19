package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Execs Route"))
	}
}

// Start ---------- 031 -----------
// END ----------- 031 ------------
func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// Start ---------- 036 -----------
	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux)))))) // no need
	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)
	// END ----------- 036 ------------

	// Create custom server
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

// Start ---------- 036 -----------
// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// END ----------- 036 ------------
