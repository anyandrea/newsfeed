package newsfeeddb

import (
	"database/sql"

	"github.com/anyandrea/newsfeed/lib/database"
)

type NewsFeedDB interface {
	Housekeeping(int) error
}

type newsfeedDB struct {
	*sql.DB
	DatabaseType string
}

func NewNewsFeedDB(adapter database.Adapter) NewsFeedDB {
	return &newsfeedDB{adapter.GetDatabase(), adapter.GetType()}
}

func (db *newsfeedDB) Housekeeping(entries int) (err error) {
	// TODO: housekeeping logic: select count(*) from feed order by timestamp desc
	// if count > $entries, then: delete from feed until count <= $entries
	return nil
}
