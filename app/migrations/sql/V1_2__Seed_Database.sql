INSERT INTO users (username, email, bio) VALUES
('michael', 'michael@email.com', 'i am cool'),
('mick', 'jake@email.com', 'i am mick'),
('marco', 'marco@email.com', 'i am marco'),
('amanda', 'amanda@email.com', 'i am amanda');

INSERT INTO messages (body, sender_id, recipient_id) VALUES
('hey amanda', 4, 1),
('hey michael', 2, 2),
('hey michael', 2, 3);

INSERT INTO followings (follower_id, followed_id) VALUES
(2, 1),
(4, 2),
(4, 3);
