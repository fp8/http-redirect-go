package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const DEFAUTL_SERVER_REDIRECT = "https://www.google.com/"
const DEFAULT_SERVER_REDIRECT_CODE = 301
const DEFAULT_HEALTH_ENDPOINT = "/healthz"
const DEFAULT_SERVER_PORT = ":8080"

/*
Return the system env variable, and if it's empty, return provided
default
*/
func GetenvOrDefault(e string, d string) string {
	var env = os.Getenv(e)
	if len(env) == 0 {
		return d
	} else {
		return env
	}
}

/*
Simple healthCheck that return "ok" with 200 status code
*/
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

/*
Redirect to URL provided by "SERVER_REDIRECT" env variable
*/
func redirect(w http.ResponseWriter, r *http.Request) {
	server := GetenvOrDefault("SERVER_REDIRECT", DEFAUTL_SERVER_REDIRECT)
	intCode := DEFAULT_SERVER_REDIRECT_CODE
	code := GetenvOrDefault("SERVER_REDIRECT_CODE", "")

	if len(code) > 0 {
		cvCode, err := strconv.Atoi(code)
		if err == nil {
			intCode = cvCode
		} else {
			log.Fatalf("Invalide SERVER_REDIRECT_CODE of %v providing.  Defaulting to %v", code, DEFAULT_SERVER_REDIRECT_CODE)
			intCode = DEFAULT_SERVER_REDIRECT_CODE
		}
	}

	http.Redirect(w, r, server, intCode)
}

/*
Start the server at the host/port defined in "SERVER_PORT",
set the health check url per "HEALTH_ENDPOINT"
and redirect to url defined by "SERVER_REDIRECT"
*/
func main() {
	var serverPort = GetenvOrDefault("SERVER_PORT", DEFAULT_SERVER_PORT)
	var healthEndpoint = GetenvOrDefault("HEALTH_ENDPOINT", DEFAULT_HEALTH_ENDPOINT)

	mux := http.NewServeMux()
	mux.HandleFunc(healthEndpoint, healthCheck)
	mux.HandleFunc("/", redirect)

	log.Printf("Http-Redirect listening at %v", serverPort)
	if err := http.ListenAndServe(serverPort, mux); err != nil {
		panic(err)
	}

}
