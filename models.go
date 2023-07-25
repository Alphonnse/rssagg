package main

import (
	"time"

	"github.com/Alphonnse/rssagg/internal/database"
	"github.com/google/uuid"
)

// This whole file for that we can specify the 
// needed names in json (like snake case)

// The pretty output when adding the user
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

// The pretty output when adding the feed
type Feed struct {
	ID 			uuid.UUID 	`json:"id"`
	CreatedAt	time.Time 	`json:"created_ad"`
	UpdatedAt	time.Time 	`json:"updated_ad"`
	Name		string		`json:"name"`
	Url			string		`json:"url"`
	UserID		uuid.UUID	`json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID: 		dbFeed.ID,
		CreatedAt: 	dbFeed.CreatedAt,
		UpdatedAt: 	dbFeed.UpdatedAt,
		Name: 		dbFeed.Name,
		Url: 		dbFeed.Url,
		UserID: 	dbFeed.UserID,
	}
}

// The pretty output when getting the feeds
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}
