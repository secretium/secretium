-- Update one secret's expiration by the given key.
UPDATE `secret_sharer_data`
SET `expires_at` = $1
WHERE `key` = $2