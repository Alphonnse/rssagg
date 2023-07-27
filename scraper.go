package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Alphonnse/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrancy int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrancy, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C { // C there is the chanal. This mean -- run this loop every tick and than it will append to time duration from timeBetweenRequest
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrancy),
		)
		if err != nil {
			log.Println("error fetching feeds:",err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
		
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{} // this because on (1) we need a nullstring. And here also we are check are is the empty string
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate) // There we are prsing the date because its just a string
		if err != nil {
			if item.PubDate == "" {
				log.Println("its empty string")
			}
			log.Printf("couldn't parse date %v with err %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), 
			database.CreatePostParams{
				ID:				uuid.New(),
				CreatedAt:		time.Now().UTC(),
				UpdatedAt:		time.Now().UTC(),
				Title:			item.Title,
				Description:	description,
				PublishedAt:	pubAt,
				Url:			item.Link,
				FeedID:			feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post:", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
