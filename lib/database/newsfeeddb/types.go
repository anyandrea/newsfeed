package newsfeeddb

import "time"

type Feed struct {
	Id            int            `json:"id" xml:"id,attr"`
	Title         string         `json:"title" xml:"title"`
	Link          string         `json:"url" xml:"url"`
	Updated       time.Time      `json:"updated" xml:"updated"`
	Fetched       time.Time      `json:"fetched" xml:"fetched"`
	Items         []Item         `json:"items" xml:"items"`
	Subscriptions []Subscription `json:"subscriptions" xml:"subscriptions"`
}

type Item struct {
	Id        int       `json:"id" xml:"id,attr"`
	FeedId    int       `json:"feed_id" xml:"feed_id,attr"`
	Title     string    `json:"title" xml:"title"`
	Link      string    `json:"url" xml:"url"`
	Updated   time.Time `json:"updated" xml:"updated"`
	Published time.Time `json:"published" xml:"published"`
}

type Subscription struct {
	FeedId      int `json:"feed_id" xml:"feed_id,attr"`
	UserId      int `json:"user_id" xml:"user_id,attr"`
	ShowEntries int `json:"show_entries" xml:"show_entries"`
}

type User struct {
	Id            int            `json:"id" xml:"id,attr"`
	Password      string         `json:"password" xml:"password"`
	Name          string         `json:"name" xml:"name"`
	Email         string         `json:"email" xml:"email"`
	Role          string         `json:"role" xml:"role"`
	Active        bool           `json:"active" xml:"active"`
	Subscriptions []Subscription `json:"subscriptions" xml:"subscriptions"`
}
