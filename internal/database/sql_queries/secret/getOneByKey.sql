-- Get one secret by the given key.
SELECT `created_at`,
    `expires_at`,
    `access_code`,
    `name`,
    `key`,
    `value`,
    `is_expire_after_first_unlock`
FROM `secret_sharer_data`
WHERE `key` = $1