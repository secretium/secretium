package attachments

import "embed"

//go:embed static_files/*
var staticFiles embed.FS

// Attachments contains all SQL queries for attachments.
type Attachments struct {
	StaticFiles embed.FS
}

// New returns a new instance of the embedded files.
func New() *Attachments {
	return &Attachments{
		StaticFiles: staticFiles,
	}
}
