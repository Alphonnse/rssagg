package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Alphonnse/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Its an db driver. _ there means include this even if i dnot use is directly
)

type apiConfig struct {
	DB *database.Queries // its defined in db.go
}


func main() {
	godotenv.Load(".env") // there we just importing the .env config file

	portString := os.Getenv("PORT") // importing the port from .env
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL") // importing the DB_URL from .env
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL) // we are connecting to DB
	if err != nil {
		log.Fatal("Cant connect to database:", err)
	}

	// its for that we can pass our api config to the handlers so that they can have access to our database
	apiCfg := apiConfig {
		DB: database.New(conn), // converting conn of type sql.db to database.queries

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

	// the v1 router is for that if we make braking changes in the future we can have two different handlers
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/healthz", handlerReadiness) // full path for this request will be /v1/healthz
	v1Router.Get("/err", handlerErr) // path for this is the /v1/err
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	srv := &http.Server { // defining the server that will use that first router
		Handler: 	router,
		Addr:		":" + portString,
	}

	log.Printf("Server starting in port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
