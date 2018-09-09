-- subscription
DELETE FROM subscription
WHERE fk_user_id = (select pk_user_id from user where name = 'SimRacer');

-- feed
DELETE FROM feed
WHERE feed_link = 'https://www.iracing.com/category/news/sim-racing-news/feed/';
DELETE FROM feed
WHERE feed_link = 'https://www.reddit.com/r/iracing.rss';
DELETE FROM feed
WHERE feed_link = 'https://www.reddit.com/r/simracing.rss';

-- user
DELETE FROM user
WHERE name = 'SimRacer';
