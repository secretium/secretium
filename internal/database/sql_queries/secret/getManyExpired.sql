-- Get all expired records.
SELECT `id`,
    `created_at`,
    `expires_at`,
    `name`,
    `key`
FROM `secret_sharer_data`
WHERE `expires_at` <= datetime('now', 'localtime')
ORDER BY `created_at` DESC