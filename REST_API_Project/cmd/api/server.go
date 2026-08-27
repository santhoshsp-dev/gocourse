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
	"restapi/pkg/utils"

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
		utils.ErrorHandler(err, "")
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

	router := router.MainRouter()
	secureMux := mw.SecurityHeaders(router)

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
