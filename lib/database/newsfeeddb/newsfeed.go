package newsfeeddb

import (
	"database/sql"

	"github.com/anyandrea/newsfeed/lib/database"
)

type NewsFeedDB interface {
	GetUsers() ([]User, error)
	GetSubscriptionsByUserId(int) ([]Subscription, error)
	GetSubscriptionsByFeedId(int) ([]Subscription, error)
	GetFeedById(int) (Feed, error)
	GetFeedsByUserId(int) ([]Feed, error)
	GetItemsByFeedId(int) ([]Item, error)
	Housekeeping(int) error
}

type newsfeedDB struct {
	*sql.DB
	DatabaseType string
}

func NewNewsFeedDB(adapter database.Adapter) NewsFeedDB {
	return &newsfeedDB{adapter.GetDatabase(), adapter.GetType()}
}

func (db *newsfeedDB) GetUsers() ([]User, error) {
	rows, err := db.Query(`
		select
			u.pk_user_id,
			u.password,
			u.name,
			u.email,
			u.role,
			u.active
		from user u
		order by u.email asc, u.name asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		// get user
		var user User
		if err := rows.Scan(&user.Id, &user.Password, &user.Name, &user.Email, &user.Role, &user.Active); err != nil {
			return nil, err
		}

		// get users subscriptions
		var err error
		user.Subscriptions, err = db.GetSubscriptionsByUserId(user.Id)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (db *newsfeedDB) GetSubscriptionsByUserId(userId int) ([]Subscription, error) {
	stmt, err := db.Prepare(`
		select
			s.fk_feed_id,
			s.fk_user_id,
			s.show_entries
		from subscription s
		where s.fk_user_id = ?
		order by s.fk_user_id asc, s.fk_feed_id asc`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		// get subscription
		var subscription Subscription
		if err := rows.Scan(&subscription.FeedId, &subscription.UserId, &subscription.ShowEntries); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func (db *newsfeedDB) GetSubscriptionsByFeedId(feedId int) ([]Subscription, error) {
	stmt, err := db.Prepare(`
		select
			s.fk_feed_id,
			s.fk_user_id,
			s.show_entries
		from subscription s
		where s.fk_feed_id = ?
		order by s.fk_user_id asc, s.fk_feed_id asc`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(feedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		// get subscription
		var subscription Subscription
		if err := rows.Scan(&subscription.FeedId, &subscription.UserId, &subscription.ShowEntries); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func (db *newsfeedDB) GetFeedById(feedId int) (Feed, error) {
	stmt, err := db.Prepare(`
		select
			f.pk_feed_id,
			f.title,
			f.link,
			f.updated,
			f.fetched
		from feed f
		where f.pk_feed_id = ?
		order by f.title asc, f.link asc`)
	if err != nil {
		return Feed{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(feedId)
	if err != nil {
		return Feed{}, err
	}
	defer rows.Close()

	var feed Feed
	for rows.Next() {
		// get feed
		if err := rows.Scan(&feed.Id, &feed.Title, &feed.Link, &feed.Updated, &feed.Fetched); err != nil {
			return Feed{}, err
		}

		// get subscriptions
		var err error
		feed.Subscriptions, err = db.GetSubscriptionsByFeedId(feed.Id)
		if err != nil {
			return Feed{}, err
		}

		// get items
		feed.Items, err = db.GetItemsByFeedId(feed.Id)
		if err != nil {
			return Feed{}, err
		}
	}
	return feed, nil
}

func (db *newsfeedDB) GetFeedsByUserId(userId int) ([]Feed, error) {
	// get subscriptions
	subscriptions, err := db.GetSubscriptionsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var feeds []Feed
	// get feeds
	for _, subscription := range subscriptions {
		feed, err := db.GetFeedById(subscription.FeedId)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}
	return feeds, nil
}

func (db *newsfeedDB) GetItemsByFeedId(feedId int) ([]Item, error) {
	stmt, err := db.Prepare(`
		select
			i.pk_feed_item_id,
			i.fk_feed_id,
			i.title,
			i.link,
			i.updated,
			i.published
		from item i
		where i.fk_feed_id = ?
		order by i.fk_user_id asc, i.fk_feed_id asc`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(feedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		// get item
		var item Item
		if err := rows.Scan(&item.Id, &item.FeedId, &item.Title, &item.Link, &item.Updated, &item.Published); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (db *newsfeedDB) Housekeeping(entries int) (err error) {
	// TODO: housekeeping logic: select count(*) from feed order by timestamp desc
	// if count > $entries, then: delete from feed until count <= $entries
	return nil
}
