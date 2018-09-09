-- subscription
DELETE FROM subscription
WHERE fk_user_id = (select pk_user_id from user where name = 'Admin');

-- feed
DELETE FROM feed
WHERE feed_link = 'https://news.ycombinator.com/rss';
DELETE FROM feed
WHERE feed_link = 'https://www.heise.de/newsticker/heise-atom.xml';

-- user
DELETE FROM user
WHERE name = 'Admin';
