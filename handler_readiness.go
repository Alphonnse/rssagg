package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // this handler should respond if the server is alive
	// respondWithJSON(w, 200, struct{}{})
	respondWithJSON(w, 200, struct{}{})
}
