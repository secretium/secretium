-- Update one secret's access code by the given key.
UPDATE `secret_sharer_data`
SET `access_code` = $1
WHERE `key` = $2