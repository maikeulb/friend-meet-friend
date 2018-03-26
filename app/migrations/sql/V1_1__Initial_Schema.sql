CREATE TABLE states (
    id serial PRIMARY KEY,
    name varchar(50) NULL
);

CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL,
    email TEXT NOT NULL,
    password_hash varchar(50) NOT NULL,
    password_salt varchar(50) NOT NULL,
    bio integer NOT NULL,
    created_on NOT NULL,
    last_active NOT NULL,
    last_message_read_time integer NULL
    follower_id integer NULL
    followed_id_ integer NULL
    message_sent_id integer NULL
    message_recieved_id integer NULL
);

CREATE TABLE messages (
    id serial PRIMARY KEY,
    body integer NOT NULL,
    timestamp integer NULL
    sender_id integer NOT NULL,
    recipient_id integer NULL
);

CREATE TABLE followings (
    follower_id serial PRIMARY KEY,
    followed_id integer NULL
);

