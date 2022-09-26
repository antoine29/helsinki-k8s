CREATE SCHEMA IF NOT EXISTS pingpong;

CREATE TABLE IF NOT EXISTS pingpong.counts (
	id			SERIAL PRIMARY key,
    count       integer NOT NULL,
    count_date  timestamp,
    hash        varchar
);
