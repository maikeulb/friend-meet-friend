CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    password_hash bytea NULL,
    password_salt bytea NULL,
    bio TEXT NOT NULL,
    interests TEXT NOT NULL,
    created_on timestamp NOT NULL,
    last_active timestamp NOT NULL
);

CREATE TABLE messages (
    id serial PRIMARY KEY,
    body TEXT NOT NULL,
    timestamp timestamp NOT NULL,
    sender_id integer NOT NULL,
    recipient_id integer NOT NULL
);

CREATE TABLE followings (
    follower_id integer NOT NULL,
    followed_id integer NOT NULL,
    PRIMARY KEY (follower_id, followed_id)
);

ALTER TABLE messages
ADD CONSTRAINT fk_messages_senders
FOREIGN KEY (sender_id)
REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE messages
ADD CONSTRAINT fk_messages_recipients
FOREIGN KEY (recipient_id)
REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE followings
ADD CONSTRAINT fk_followers_users
FOREIGN KEY (follower_id)
REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE followings
ADD CONSTRAINT fk_followed_users
FOREIGN KEY (followed_id)
REFERENCES users(id)
ON DELETE CASCADE;

