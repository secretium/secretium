-- Add a new secret.
INSERT INTO `secret_sharer_data` (
        `created_at`,
        `expires_at`,
        `access_code`,
        `name`,
        `key`,
        `value`,
        `is_expire_after_first_unlock`
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)