CREATE TYPE event_type AS ENUM ('create', 'update', 'delete');

CREATE TABLE answer_map (
    answer_key TEXT PRIMARY KEY,
    answer_value TEXT
);

CREATE TABLE answer_event (
    event_id UUID PRIMARY KEY,
    event_type event_type NOT NULL,
    event_timestamp TIMESTAMPTZ NOT NULL,
    answer_key TEXT NOT NULL REFERENCES answer_map,
    answer_value TEXT
);