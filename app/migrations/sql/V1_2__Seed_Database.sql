INSERT INTO users (nickname, email, password_hash, interests, borough) VALUES
('michael', 'michael@email.com','$2a$10$5kDwYNJsKpQf00fGbAyEs.JeSGfn8GL7SgbgbGppAT8tD8K40iyjq', 'museums', 'queens'),
('mick','jake@email.com','$2a$10$5kDwYNJsKpQf00fGbAyEs.JeSGfn8GL7SgbgbGppAT8tD8K40iyjq', 'bmx bikes', 'manhattan'),
('marco','marco@email.com','$2a$10$5kDwYNJsKpQf00fGbAyEs.JeSGfn8GL7SgbgbGppAT8tD8K40iyjq', 'movies', 'queens'),
('amanda','amanda@email.com','$2a$10$5kDwYNJsKpQf00fGbAyEs.JeSGfn8GL7SgbgbGppAT8tD8K40iyjq', 'eating', 'brooklyn');

INSERT INTO messages (body, timestamp, sender_id, recipient_id) VALUES
('hey amanda', '2017-12-11 09:22:51', 4, 1),
('hey michael', '2018-02-21 08:13:13', 2, 2),
('hey michael', '2018-03-01 02:26:04', 2, 3);

INSERT INTO followings (follower_id, followee_id) VALUES
(2, 1),
(4, 2),
(4, 3);
