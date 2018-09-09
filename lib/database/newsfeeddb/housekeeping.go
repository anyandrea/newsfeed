package newsfeeddb

func (db *newsfeedDB) Housekeeping(entries int) (err error) {
	// TODO: housekeeping logic: select count(*) from feed order by timestamp desc
	// if count > $entries, then: delete from feed until count <= $entries
	return nil
}
