INSERT INTO users (username, email, interests, borough, created_on, last_active) VALUES
('michael', 'michael@email.com', 'museums', 'queens', '2017-01-01 10:23:54', '2018-12-01 01:02:12'),
('mick', 'jake@email.com', 'bmx bikes', 'manhattan', '2017-01-01 10:23:54', '2018-11-04 03:22:24'),
('marco', 'marco@email.com', 'movies', 'queens', '2017-01-01 10:23:54', '2018-01-01 09:13:54'),
('amanda', 'amanda@email.com', 'eating', 'brooklyn', '2017-01-01 10:23:54', '2018-01-02 10:13:44');

INSERT INTO messages (body, timestamp, sender_id, recipient_id) VALUES
('hey amanda', '2017-12-11 09:22:51', 4, 1),
('hey michael', '2018-02-21 08:13:13', 2, 2),
('hey michael', '2018-03-01 02:26:04', 2, 3);

INSERT INTO followings (follower_id, followee_id) VALUES
(2, 1),
(4, 2),
(4, 3);
