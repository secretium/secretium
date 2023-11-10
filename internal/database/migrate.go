package database

// Migrate migrates the given DB schema from the embedded SQL file.
func (d *Database) Migrate(schema string) error {
	// Read the DB schema from the embedded SQL file.
	s, err := d.SQLQueries.ReadFile(schema)
	if err != nil {
		return err
	}

	// Migrate the DB schema.
	d.Connection.MustExec(string(s))

	return nil
}
