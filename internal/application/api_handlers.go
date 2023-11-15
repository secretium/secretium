package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/database"
	"github.com/secretium/secretium/internal/helpers"
	"github.com/secretium/secretium/internal/messages"
	"github.com/secretium/secretium/internal/templates"
	"github.com/secretium/secretium/internal/templates/components"
	"github.com/secretium/secretium/internal/templates/pages"
)

// APIAddSecretHandler adds a new secret to the database (POST).
func (a *Application) APIAddSecretHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the form data.
	if err := r.ParseForm(); err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Form data", Message: err.Error()},
				},
			),
			messages.ErrFormDataNotValid,
		)
		return
	}

	// Get form values.
	name := r.FormValue("name")
	value := r.FormValue("value")
	expiresAt := r.FormValue("expires_at")
	isExpireAfterFirstUnlock := r.FormValue("is_expire_after_first_unlock") == "on"

	// Check, if the form values are valid.
	if err := helpers.ValidateAddSecretForm(name, value); err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(err),
			messages.ErrFormDataNotValid,
		)
		return
	}

	// Get current date and time.
	createdAt := time.Now()

	// Create a new hashed access code string with salt and trim it to 8 characters.
	accessCodeHashed := helpers.HashString(8, fmt.Sprintf("%d", createdAt.Unix()), a.Config.SecretKey)

	// Create a new hashed key string with salt and trim it to 16 characters.
	keyHashed := helpers.HashString(16, accessCodeHashed, a.Config.SecretKey)

	// Encrypt the secret value.
	valueEncrypted, err := helpers.EncryptString(a.Config.SecretKey, value)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Encrypt secret", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Encrypt the access code value.
	accessCodeEncrypted, err := helpers.EncryptString(a.Config.SecretKey, accessCodeHashed)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Encrypt access code", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Parse the 'expires_at' datetime.
	expiresAtDuration, err := helpers.ExpiresDatetimeSwitcher(createdAt, expiresAt)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Expires datetime", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Create a new secret record.
	secret := &database.Secret{
		CreatedAt:                createdAt,
		ExpiresAt:                expiresAtDuration,
		Name:                     name,
		AccessCode:               accessCodeEncrypted,
		Key:                      keyHashed,
		Value:                    valueEncrypted,
		IsExpireAfterFirstUnlock: isExpireAfterFirstUnlock,
	}

	// Add the record to the database.
	if err := a.Database.QueryAddSecret(secret); err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Expires datetime", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Redirect to the share secret page.
	w.Header().Set("HX-Location", fmt.Sprintf("/dashboard/share/%s?access_code=%s", secret.Key, accessCodeHashed))
}

// APIUnlockSecretHandler renders the unlocked secret block (POST).
func (a *Application) APIUnlockSecretHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Parse the form data.
	if err := r.ParseForm(); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the encrypt secret error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Form data", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Get key and access code from the form inputs.
	accessCode := r.FormValue("access_code")

	// Check, if the form values are valid.
	if err := helpers.ValidateViewSecretForm(accessCode); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(err).Render(r.Context(), w)

		return
	}

	// Create template options.
	templateOptions := &templates.TemplateOptions{
		Header: &templates.ElementStyle{},
		Main: &templates.ElementStyle{
			CSSClass: "secret",
		},
		Footer: &templates.ElementStyle{
			CSSClass: "secret",
		},
	}

	// Add the secret to the database.
	secret, err := a.Database.QueryGetSecretByKey(key)
	if err != nil {
		// Send a 404 not found response.
		w.WriteHeader(http.StatusNotFound)

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

	// Check, if expiration date is in the future.
	if secret.ExpiresAt.Before(time.Now()) {
		// Send a 400 not found response.
		w.WriteHeader(http.StatusBadRequest)

		// Set the template options.
		templateOptions.PageTitle = "Oops... Secret is expired"
		templateOptions.LogoVariant = "error"
		templateOptions.Component = pages.Secret(&secret, "expired")

		// Render the secret page with 400 error.
		_ = templates.Layout(templateOptions).Render(r.Context(), w)

		return
	}

	// Decrypt the access code value.
	accessCodeDecrypted, err := helpers.DecryptString(a.Config.SecretKey, secret.AccessCode)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Decrypt access code", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Check, if the entered access code is equal to the decrypted access code.
	if accessCode != accessCodeDecrypted {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Access code", Message: messages.ErrSecretAccessCodeNotValid},
				},
			),
			messages.ErrSecretAccessCodeNotValid,
		)
		return
	}

	// Decrypt the secret value.
	decryptedValue, err := helpers.DecryptString(a.Config.SecretKey, secret.Value)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Decrypt secret", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Set component options.
	secret.Value = decryptedValue

	// Render the secret page.
	_ = pages.Secret(&secret, "unlocked").Render(r.Context(), w)
}

// APIRenewSecretExpiresAtFieldByKeyHandler renews a secret 'expires_at' field by its key from the database (PATCH).
func (a *Application) APIRenewSecretExpiresAtFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateExpiresAtFieldByKey(key, time.Now().Add(time.Hour*24).Local()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the HX-Trigger header (to trigger a re-render by htmx).
	w.Header().Set("HX-Trigger", "getActiveSecrets, getExpiredSecrets")
}

// APIRestoreSecretAccessCodeFieldByKeyHandler restores a secret 'access_code' field by its key from the database (PATCH).
func (a *Application) APIRestoreSecretAccessCodeFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the secret record by its key from the database.
	secret, err := a.Database.QueryGetSecretByKey(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Create a new hashed access code string with salt.
	accessCodeHashed := helpers.HashString(8, fmt.Sprintf("%d", secret.CreatedAt.Unix()), a.Config.SecretKey)

	// Encrypt the access code value.
	accessCodeEncrypted, err := helpers.EncryptString(a.Config.SecretKey, accessCodeHashed)
	if err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Encrypt access code", Message: err.Error()},
				},
			),
			err.Error(),
		)
		return
	}

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateAccessCodeFieldByKey(key, accessCodeEncrypted); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the restore access code block.
	_ = components.DashboardRestoreAccessCode(accessCodeHashed).Render(r.Context(), w)
}

// APIExpireSecretExpiresAtFieldByKeyHandler expires a secret 'expires_at' field by its key from the database (PATCH).
func (a *Application) APIExpireSecretExpiresAtFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateExpiresAtFieldByKey(key, time.Now().Add(time.Second*1).Local()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// APIDeleteSecretByKeyHandler deletes a secret by its key from the database (DELETE).
func (a *Application) APIDeleteSecretByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete the record by its key from the database.
	if err := a.Database.QueryDeleteSecretByKey(key); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the HX-Trigger header (to trigger a re-render by htmx).
	w.Header().Set("HX-Trigger", "getActiveSecrets, getExpiredSecrets")
}

// APIDashboardActiveSecretsHandler renders the active secrets block (GET).
func (a *Application) APIDashboardActiveSecretsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get all active secrets.
	secrets, err := a.Database.QueryGetActiveSecrets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the active secrets block.
	_ = components.ActiveSecrets(secrets).Render(r.Context(), w)
}

// APIDashboardExpiredSecretsHandler renders the expired secrets block.
func (a *Application) APIDashboardExpiredSecretsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get all expired secrets.
	secrets, err := a.Database.QueryGetExpiredSecrets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the active secrets block.
	_ = components.ExpiredSecrets(secrets).Render(r.Context(), w)
}

// APIUserLoginHandler logs in the user (POST).
func (a *Application) APIUserLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the form data.
	if err := r.ParseForm(); err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Form data", Message: err.Error()},
				},
			),
			messages.ErrFormDataNotValid,
		)
		return
	}

	// Get form values.
	username := r.FormValue("username")
	masterPassword := r.FormValue("master_password")

	// Check, if the form values are valid.
	if err := helpers.ValidateUserSignInForm(username, masterPassword); err != nil {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusBadRequest,
			components.FormValidationError(err),
			messages.ErrFormLoginUserCredentialsNotValid,
		)
		return
	}

	// Authenticate the admin.
	if a.Config.MasterUsername != username || a.Config.MasterPassword != masterPassword {
		// Wrap the error with template.
		helpers.WrapHTTPError(
			w, r, http.StatusUnauthorized,
			components.FormValidationError(
				[]*messages.ErrorField{
					{Name: "Form data", Message: messages.ErrFormLoginUserCredentialsNotValid},
				},
			),
			messages.ErrSessionUserNotAuthenticated,
		)
		return
	}

	// Set session.
	a.Session.Manager.Put(r.Context(), "authenticated", true)

	// Redirect to the dashboard page.
	w.Header().Set("HX-Redirect", "/dashboard")
}

// APIUserLogoutHandler logs out the user (GET).
func (a *Application) APIUserLogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Remove session.
	a.Session.Manager.Remove(r.Context(), "authenticated")

	// Redirect to the index page.
	w.Header().Set("HX-Redirect", "/")
}
