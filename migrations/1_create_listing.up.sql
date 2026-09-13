CREATE TABLE listings (
    id          UUID        DEFAULT gen_random_uuid() PRIMARY KEY,
    title       TEXT        NOT NULL,
    description TEXT        NOT NULL,
    price       BIGINT      NOT NULL,
    city        TEXT        NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL
);