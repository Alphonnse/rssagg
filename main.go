package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env") // there we just importing the .env config file

	portString := os.Getenv("PORT") // importing the port from .env
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter() // defining the router
	
	
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:		[]string{"https://*", "http://*"},
		AllowedMethods: 	[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: 	[]string{"*"},
		ExposedHeaders: 	[]string{"Link"},
		AllowCredentials: 	false,
		MaxAge: 			300,
	}))

	// all this construction is for that if we make braking changes in the future we can have two different handlers { 
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // full path for this request will be /v1/healthz
	v1Router.Get("/err", handlerErr) // path for this is the /v1/err

	router.Mount("/v1", v1Router)
	// }

	srv := &http.Server { // defining the server that will use that first router
		Handler: 	router,
		Addr:		":" + portString,
	}

	log.Printf("Server starting in port %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
