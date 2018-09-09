-- user
INSERT INTO user (password, name, email, role, active)
VALUES('dudeli', 'Admin', 'admin@localhost', 'admin', 1);

-- feed
INSERT INTO feed (title, link, feed_link)
VALUES('Sim Racing News – iRacing.com', 'https://www.iracing.com/', 'https://www.iracing.com/category/news/sim-racing-news/feed/');
INSERT INTO feed (title, link, feed_link)
VALUES('iRacing', 'https://www.reddit.com/r/iracing', 'https://www.reddit.com/r/iracing.rss');

-- subscription
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.iracing.com/category/news/sim-racing-news/feed/'),
	(select pk_user_id from user where name = 'Admin'),
	12
);
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.reddit.com/r/iracing.rss'),
	(select pk_user_id from user where name = 'Admin'),
	12
);
