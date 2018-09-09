-- user
INSERT INTO user (password, name, email, role, active)
VALUES('$2a$12$jlO46pJTt9xmwszAHaVi4OvqgyFVxko/lNCYwE2sLJtQ4mo97YQ9S', 'SimRacer', 'simracer@localhost', 'user', 1); -- iracing

-- feed
INSERT INTO feed (title, link, feed_link)
VALUES('Sim Racing News – iRacing.com', 'https://www.iracing.com/', 'https://www.iracing.com/category/news/sim-racing-news/feed/');
INSERT INTO feed (title, link, feed_link)
VALUES('iRacing', 'https://www.reddit.com/r/iracing', 'https://www.reddit.com/r/iracing.rss');
INSERT INTO feed (title, link, feed_link)
VALUES('Sim Racing', 'https://www.reddit.com/r/simracing', 'https://www.reddit.com/r/simracing.rss');

-- subscription
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.iracing.com/category/news/sim-racing-news/feed/'),
	(select pk_user_id from user where name = 'SimRacer'),
	10
);
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.reddit.com/r/iracing.rss'),
	(select pk_user_id from user where name = 'SimRacer'),
	10
);
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.reddit.com/r/simracing.rss'),
	(select pk_user_id from user where name = 'SimRacer'),
	10
);
