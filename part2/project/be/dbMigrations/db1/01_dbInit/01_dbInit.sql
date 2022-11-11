CREATE SCHEMA IF NOT EXISTS todo;

CREATE TABLE IF NOT EXISTS todo.todos (
    id            varchar(40) CONSTRAINT firstkey PRIMARY KEY,
    content       text NOT NULL,
    is_done       boolean NOT NULL
);
