-- Get all active records.
SELECT `id`,
    `created_at`,
    `expires_at`,
    `name`,
    `key`,
    `is_expire_after_first_unlock`
FROM `secret_sharer_data`
WHERE `expires_at` > datetime('now', 'localtime')
ORDER BY `created_at` DESC