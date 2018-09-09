-- user
INSERT INTO user (password, name, email, role, active)
VALUES('$2a$14$BHkC8UDmVJ3YbUOjwZaEa.4T.kG54L2bRc1561R0067CG5MHok04S', 'Admin', 'admin@localhost', 'admin', 1); -- dudeli

-- feed
INSERT INTO feed (title, link, feed_link)
VALUES('Hacker News', 'https://news.ycombinator.com/', 'https://news.ycombinator.com/rss');
INSERT INTO feed (title, link, feed_link)
VALUES('heise online News', 'https://www.heise.de/newsticker/', 'https://www.heise.de/newsticker/heise-atom.xml');

-- subscription
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://news.ycombinator.com/rss'),
	(select pk_user_id from user where name = 'Admin'),
	20
);
INSERT INTO subscription (fk_feed_id, fk_user_id, show_entries)
VALUES(
	(select pk_feed_id from feed where feed_link = 'https://www.heise.de/newsticker/heise-atom.xml'),
	(select pk_user_id from user where name = 'Admin'),
	20
);
