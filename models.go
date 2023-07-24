package main

import (
	"time"

	"github.com/Alphonnse/rssagg/internal/database"
	"github.com/google/uuid"
)

// This whole file for that we can specify the needed names in json (like snake case)
type User struct {
	ID 			uuid.UUID 	`json:"id"`
	CreatedAt	time.Time 	`json:"created_ad"`
	UpdatedAt	time.Time 	`json:"updated_ad"`
	Name		string		`json:"name"`
	APIKey		string		`json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User {
		ID: 		dbUser.ID,
		CreatedAt: 	dbUser.CreatedAt,
		UpdatedAt: 	dbUser.UpdatedAt,
		Name: 		dbUser.Name,
		APIKey:		dbUser.ApiKey,
	}
}
