package newsfeeddb

func (feed *Feed) Store(db *newsfeedDB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		insert into feed (pk_feed_id, title, link, feed_link, updated, fetched)
		values (?, ?, ?, ?, ?, ?) on duplicate key update
		title = ?, link = ?, feed_link = ?, updated = ?, fetched = ?`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(
		feed.Id, feed.Title, feed.Link, feed.FeedLink, feed.Updated, feed.Fetched,
		feed.Title, feed.Link, feed.FeedLink, feed.Updated, feed.Fetched); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// save feed items
	for _, item := range feed.Items {
		if err := item.Store(db); err != nil {
			return err
		}
	}
	return nil
}

func (item *Item) Store(db *newsfeedDB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// see if we can find the item already inside the database
	if item.Id < 1 {
		stmt, err := tx.Prepare(`
			select i.pk_feed_item_id
			from item i
			where i.fk_feed_id = ?
			and i.title = ?
			order by i.pk_feed_item_id, i.updated desc, i.published desc`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		rows, err := stmt.Query(item.FeedId, item.Title)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer rows.Close()

		// get item id if found
		for rows.Next() {
			if err := rows.Scan(&item.Id); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	stmt, err := tx.Prepare(`
		insert into item (pk_feed_item_id, fk_feed_id, title, link, updated, published)
		values (?, ?, ?, ?, ?, ?) on duplicate key update
		title = ?, link = ?, updated = ?, published = ?`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(
		item.Id, item.FeedId, item.Title, item.Link, item.Updated, item.Published,
		item.Title, item.Link, item.Updated, item.Published); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
