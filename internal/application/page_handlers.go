package application

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/helpers"
	"github.com/secretium/secretium/internal/templates"
	"github.com/secretium/secretium/internal/templates/pages"
)

// PageIndexHandler renders the index page (GET).
func (a *Application) PageIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check, if the user is authenticated.
	if a.Session.Manager.GetBool(r.Context(), "authenticated") {
		// Redirect to the dashboard page.
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	// Create template options.
	templateOptions := &templates.TemplateOptions{
		PageTitle: "Login to your account",
		Header:    &templates.ElementStyle{},
		Main: &templates.ElementStyle{
			CSSClass: "index",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "index",
		},
		Component: pages.Index(),
	}

	// Render the index page.
	_ = templates.Layout(templateOptions).Render(r.Context(), w)
}

// PageSecretHandler renders the secret page (GET).
func (a *Application) PageSecretHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create template options.
	templateOptions := &templates.TemplateOptions{
		PageTitle: "View secret from your friend",
		Header:    &templates.ElementStyle{},
		Main: &templates.ElementStyle{
			CSSClass: "secret",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "secret",
		},
	}

	// Get secret by its key from the database.
	secret, err := a.Database.QueryGetSecretByKey(key)
	if err != nil {
		// Set the key to the secret, because the secret is not found (Secret struct has zero values).
		secret.Key = key

		// Set the template options.
		templateOptions.PageTitle = "Oops... Secret is not found"
		templateOptions.LogoVariant = "error"
		templateOptions.Component = pages.Secret(&secret, "not-found")

		// Render the secret page with 404 error.
		_ = templates.Layout(templateOptions).Render(r.Context(), w)

		return
	}

	// Check, if the secret is expired.
	if !helpers.DatetimeChecker(secret.ExpiresAt.Local(), time.Now().Local()) {
		// Set the template options.
		templateOptions.PageTitle = "Oops... Secret is expired"
		templateOptions.LogoVariant = "error"
		templateOptions.Component = pages.Secret(&secret, "expired")

		// Render the secret page with 400 error.
		_ = templates.Layout(templateOptions).Render(r.Context(), w)

		return
	}

	// Set the template options.
	templateOptions.Component = pages.Secret(&secret, "locked")

	// Render the secret page.
	_ = templates.Layout(templateOptions).Render(r.Context(), w)
}

// PageDashboardIndexHandler renders the dashboard index page (GET).
func (a *Application) PageDashboardIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Create template options.
	templateOptions := &templates.TemplateOptions{
		PageTitle: "Dashboard",
		Header: &templates.ElementStyle{
			IsHidden: true,
		},
		Main: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Component: pages.Dashboard(
			&templates.DashboardComponentOptions{
				Username: a.Config.MasterUsername,
			},
		),
	}

	// Render the dashboard index page.
	_ = templates.Layout(templateOptions).Render(r.Context(), w)
}

// PageDashboardAddSecretHandler renders the add secret page (GET).
func (a *Application) PageDashboardAddSecretHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Create template options.
	templateOptions := &templates.TemplateOptions{
		PageTitle: "Add secret",
		Header: &templates.ElementStyle{
			IsHidden: true,
		},
		Main: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Component: pages.Dashboard(
			&templates.DashboardComponentOptions{
				State: "add-secret",
			},
		),
	}

	// Render the dashboard add secret page.
	_ = templates.Layout(templateOptions).Render(r.Context(), w)
}

// PageDashboardShareSecretHandler renders the share secret page (GET).
func (a *Application) PageDashboardShareSecretHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get secret by its key from the database.
	secret, err := a.Database.QueryGetSecretByKey(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Set template data.
	shareURL := url.URL{
		Scheme: a.Config.DomainSchema,
		Host:   a.Config.Domain,
		Path:   fmt.Sprintf("get/%s", key),
	}

	// Set component options.
	componentOptions := &templates.DashboardComponentOptions{
		State:    "share-secret",
		ShareURL: shareURL.String(),
		Secret:   &secret,
	}

	// Check, if the URL has an 'access_code' parameter.
	if r.URL.Query().Get("access_code") != "" {
		componentOptions.Data = map[string]string{
			"AccessCode": r.URL.Query().Get("access_code"),
		}
	}

	// Create template options.
	templateOptions := &templates.TemplateOptions{
		PageTitle: "Share secret",
		Header: &templates.ElementStyle{
			IsHidden: true,
		},
		Main: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "dashboard",
		},
		Component: pages.Dashboard(componentOptions),
	}

	// Render the dashboard share secret page.
	_ = templates.Layout(templateOptions).Render(r.Context(), w)
}
