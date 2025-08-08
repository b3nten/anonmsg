-- name: CreateInbox :one
INSERT INTO inboxes (private_key)
VALUES ($1)
RETURNING id, created_at, updated_at, private_key, public_key, active;

-- name: GetInbox :one
SELECT
  id,
  created_at,
  updated_at,
  private_key,
  public_key,
  active,
  (SELECT COUNT(*) FROM messages WHERE inbox_id = inboxes.id) AS message_count
FROM inboxes
WHERE private_key = $1;

-- name: DeleteInbox :exec
DELETE FROM inboxes WHERE private_key = $1;

-- name: SetInboxActive :one
UPDATE inboxes SET active = $2 WHERE private_key = $1
returning active;

-- name: AddPublicKeyToInbox :one
UPDATE inboxes SET public_key = $2 WHERE private_key = $1
returning public_key;

-- name: RemovePublicKeyFromInbox :exec
UPDATE inboxes SET public_key = NULL WHERE private_key = $1;

-- name: GetMessagesByPrivateKey :many
WITH inbox_check AS (
  SELECT id
  FROM inboxes
  WHERE private_key = $1
)
SELECT
  messages.msg_content,
  messages.created_at,
  messages.id
FROM inbox_check
LEFT JOIN messages ON inbox_check.id = messages.inbox_id;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE messages.id = $1
  AND inbox_id = (
    SELECT inbox.id
    FROM inboxes inbox
    WHERE inbox.private_key = $2
  );

-- name: CreatePrivateMessage :one
INSERT INTO messages (inbox_id, msg_content, created_at)
SELECT i.id, $2, NOW()
FROM inboxes i
WHERE i.private_key = $1
AND i.active = true
RETURNING id, created_at;

-- name: CreatePublicMessage :one
INSERT INTO messages (inbox_id, msg_content, created_at, sent_with_public_key)
SELECT i.id, $2, NOW(), true
FROM inboxes i
WHERE i.public_key = $1
AND i.active = true
AND i.public_key IS NOT NULL
RETURNING id, created_at;
