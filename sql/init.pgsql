-- Clean up previous instance

DROP TABLE IF EXISTS demo_record;

CREATE TABLE demo_record (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL
);
