package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		// Fallback: try loading from cmd/api/.env
		if errFallback := godotenv.Load("cmd/api/.env"); errFallback != nil {
			log.Println("Warning: No .env file found. Using default/system environment variables.")
		}
	}

	_, err = sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println("Error----------", err)
		return
	}

	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"
	if _, err := os.Stat(cert); os.IsNotExist(err) {
		if _, errFallback := os.Stat("cmd/api/" + cert); errFallback == nil {
			cert = "cmd/api/" + cert
			key = "cmd/api/" + key
		}
	}

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
	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	router := router.Router()
	secureMux := mw.SecurityHeaders(router)
	// secureMux := mw.SecurityHeaders(router.Router())

	// END ----------- 036 ------------

	// Create custom server
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
