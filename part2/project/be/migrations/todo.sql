CREATE TABLE todos (
    id            varchar(40) CONSTRAINT firstkey PRIMARY KEY,
    content       varchar(40) NOT NULL,
    is_done       boolean NOT NULL
);

