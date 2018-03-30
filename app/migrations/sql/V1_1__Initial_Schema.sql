CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    password_hash bytea NULL,
    password_salt bytea NULL,
    interests TEXT NULL,
    borough varchar(50) NULL,
    created_on timestamp NULL,
    last_active timestamp NULL
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
    followee_id integer NOT NULL,
    PRIMARY KEY (follower_id, followee_id)
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
ADD CONSTRAINT fk_followee_users
FOREIGN KEY (followee_id)
REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE users
ALTER COLUMN created_on
SET DEFAULT now();

ALTER TABLE users
ALTER COLUMN last_active
SET DEFAULT now();
