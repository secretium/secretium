-- Delete one secret by the given key.
DELETE FROM `secret_sharer_data`
WHERE `key` = $1