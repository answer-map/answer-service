CREATE TYPE event_type AS ENUM ('create', 'update', 'delete');

CREATE TABLE answer_user (
    user_id UUID PRIMARY KEY
);

CREATE TABLE answer_map (
    map_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES answer_user,
    answer_key TEXT NOT NULL,
    answer_value TEXT,
    CONSTRAINT unique_map_for_user UNIQUE (user_id, answer_key)
);

CREATE TABLE answer_event (
    event_id UUID PRIMARY KEY,
    event_type event_type NOT NULL,
    event_timestamp TIMESTAMPTZ NOT NULL,
    map_id UUID NOT NULL REFERENCES answer_map,
    answer_value TEXT
);