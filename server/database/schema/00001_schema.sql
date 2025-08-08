-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $func$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$func$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TABLE inboxes (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  private_key TEXT UNIQUE NOT NULL,
  public_key TEXT UNIQUE,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TRIGGER update_inboxes_updated_at
    BEFORE UPDATE ON inboxes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE messages (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  inbox_id BIGSERIAL NOT NULL REFERENCES inboxes(id) ON DELETE CASCADE,
  msg_content TEXT NOT NULL,
  msg_format TEXT NOT NULL DEFAULT 'text/plain',
  sent_with_public_key BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TRIGGER update_messages_updated_at
    BEFORE UPDATE ON messages
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- +goose Down

DROP TRIGGER IF EXISTS update_messages_updated_at ON messages;
DROP TRIGGER IF EXISTS update_inboxes_updated_at ON inboxes;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS inboxes;
DROP FUNCTION IF EXISTS update_updated_at_column();
