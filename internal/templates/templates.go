package templates

import (
	"github.com/a-h/templ"
	"github.com/secretium/secretium/internal/database"
)

// TemplateOptions is the options for the template.
type TemplateOptions struct {
	PageTitle            string
	LogoVariant          string
	Header, Main, Footer *ElementStyle
	Component            templ.Component
}

// ElementStyle is the style of an element.
type ElementStyle struct {
	IsHidden bool
	CSSClass string
}

// DashboardComponentOptions is the options for the dashboard component.
type DashboardComponentOptions struct {
	State, Username, ShareURL string
	Secret                    *database.Secret
	Data                      map[string]string
}
