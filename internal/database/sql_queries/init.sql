-- Init secret sharer DB.
CREATE TABLE IF NOT EXISTS `secret_sharer_data` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expires_at` datetime NOT NULL,
    `access_code` varchar(64) NOT NULL UNIQUE,
    `name` varchar(32) NOT NULL,
    `key` varchar(16) NOT NULL UNIQUE,
    `value` text NOT NULL,
    `is_expire_after_first_unlock` boolean NOT NULL DEFAULT false
)