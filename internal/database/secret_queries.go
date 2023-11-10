package database

import (
	"time"
)

// Secret represents a secret record.
type Secret struct {
	ID                       int       `db:"id"`
	CreatedAt                time.Time `db:"created_at"`
	ExpiresAt                time.Time `db:"expires_at"`
	AccessCode               string    `db:"access_code"`
	Name                     string    `db:"name"`
	Key                      string    `db:"key"`
	Value                    string    `db:"value"`
	IsExpireAfterFirstUnlock bool      `db:"is_expire_after_first_unlock"`
}

// QueryAddSecret adds a new secret to the database.
func (d *Database) QueryAddSecret(s *Secret) error {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/add.sql")
	if err != nil {
		return err
	}

	// Add the record to the database.
	_, err = d.Connection.Exec(
		string(query),
		s.CreatedAt, s.ExpiresAt,
		s.AccessCode, s.Name, s.Key, s.Value,
		s.IsExpireAfterFirstUnlock,
	)
	if err != nil {
		return err
	}

	return nil
}

// QueryGetSecretByKey returns the secret by its key from the database.
func (d *Database) QueryGetSecretByKey(key string) (secret Secret, err error) {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/getOneByKey.sql")
	if err != nil {
		return secret, err
	}

	// Get the record by its key from the database.
	if err := d.Connection.Get(&secret, string(query), key); err != nil {
		return secret, err
	}

	return secret, nil
}

// QueryUpdateExpiresAtFieldByKey updates the 'expires_at' field of the secret by its key in the database.
func (d *Database) QueryUpdateExpiresAtFieldByKey(key string, expiredAt time.Time) error {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/updateExpiresAtFieldOneByKey.sql")
	if err != nil {
		return err
	}

	// Refresh the record by its key from the database.
	_, err = d.Connection.Exec(string(query), expiredAt, key)
	if err != nil {
		return err
	}

	return nil
}

// QueryUpdateAccessCodeFieldByKey updates the 'access_code' field of the secret by its key in the database.
func (d *Database) QueryUpdateAccessCodeFieldByKey(key, accessCode string) error {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/updateAccessCodeFieldOneByKey.sql")
	if err != nil {
		return err
	}

	// Refresh the record by its key from the database.
	_, err = d.Connection.Exec(string(query), accessCode, key)
	if err != nil {
		return err
	}

	return nil
}

// QueryDeleteSecretByKey deletes a secret by its key from the database.
func (d *Database) QueryDeleteSecretByKey(key string) error {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/deleteOneByKey.sql")
	if err != nil {
		return err
	}

	// Delete the record by its key from the database.
	_, err = d.Connection.Exec(string(query), key)
	if err != nil {
		return err
	}

	return nil
}

// QueryGetActiveSecrets returns the active secrets from the database.
func (d *Database) QueryGetActiveSecrets() (secrets []*Secret, err error) {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/getManyActive.sql")
	if err != nil {
		return nil, err
	}

	// Get the record by its key from the database.
	if err := d.Connection.Select(&secrets, string(query)); err != nil {
		return nil, err
	}

	return secrets, nil
}

// QueryGetExpiredSecrets returns the expired secrets from the database.
func (d *Database) QueryGetExpiredSecrets() (secrets []*Secret, err error) {
	// Create a query from the embedded SQL file.
	query, err := d.SQLQueries.ReadFile("sql_queries/secret/getManyExpired.sql")
	if err != nil {
		return nil, err
	}

	// Get the record by its key from the database.
	if err := d.Connection.Select(&secrets, string(query)); err != nil {
		return nil, err
	}

	return secrets, nil
}
