package newsfeeddb

import (
	"time"

	"github.com/mmcdole/gofeed"
)

func (db *newsfeedDB) FetchAllFeeds() error {
	users, err := db.GetUsers()
	if err != nil {
		return err
	}

	// figure out all feed Ids that active users are subscribed to.
	feeds := make(map[int]bool, 0)
	for _, user := range users {
		if user.Active {
			for _, subscription := range user.Subscriptions {
				feeds[subscription.FeedId] = true
			}
		}
	}

	// fetch/update all those feeds
	for feedId, _ := range feeds {
		if err := db.FetchFeed(feedId); err != nil {
			return err
		}
	}

	return nil
}

func (db *newsfeedDB) FetchFeed(feedId int) error {
	feed, err := db.GetFeedById(feedId)
	if err != nil {
		return err
	}

	feedParser := gofeed.NewParser()
	parsedFeed, err := feedParser.ParseURL(feed.FeedLink)
	// if err != nil {
	// 	return err
	// }
	if err == nil {
		feed.Fetched = time.Now()

		feed.Title = parsedFeed.Title
		feed.Link = parsedFeed.Link
		feed.Updated = parsedFeed.UpdatedParsed

		for _, item := range parsedFeed.Items {
			feed.Items = append(feed.Items, Item{
				FeedId:    feed.Id,
				Title:     item.Title,
				Link:      item.Link,
				Updated:   item.UpdatedParsed,
				Published: item.PublishedParsed,
			})
		}
	}

	return feed.Store(db)
}
