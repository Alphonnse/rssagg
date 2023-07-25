package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Alphonnse/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFallow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID	uuid.UUID `json:"feed_id"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),	
		UserID:		user.ID,
		FeedID:		params.FeedID,
	}) // we can do so because the sqlc create for us the create user function from quaries

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFallows(w http.ResponseWriter, r *http.Request, user database.User) {
		feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	
		if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFallow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:		feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
