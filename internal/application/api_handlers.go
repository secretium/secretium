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
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Parse the form data.
	if err := r.ParseForm(); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Form data", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Get form values.
	name := r.FormValue("name")
	value := r.FormValue("value")
	accessCode := r.FormValue("access_code")
	expiresAt := r.FormValue("expires_at")
	isExpireAfterFirstUnlock := r.FormValue("is_expire_after_first_unlock") == "on"

	// Check, if the form values are valid.
	if err := helpers.ValidateAddSecretForm(name, value, accessCode); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(err).Render(r.Context(), w)

		return
	}

	// Get current date and time.
	createdAt := time.Now()

	// Create a new hashed access code string with salt.
	accessCodeHashed := helpers.HashString(
		fmt.Sprintf("%d", createdAt.Unix()),
		a.Config.SecretKey, accessCode,
	)

	// Shorten the access code to 16 characters.
	key := helpers.ShortenString(accessCodeHashed, 16)

	// Encrypt the secret value.
	valueEncrypted, err := helpers.EncryptString(a.Config.SecretKey, value)
	if err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the encrypt secret error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Encrypt secret", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Parse the 'expires_at' datetime.
	expiresAtDuration, err := helpers.ExpiresDatetimeSwitcher(createdAt, expiresAt)
	if err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the encrypt secret error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Expires datetime", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Create a new secret record.
	secret := &database.Secret{
		CreatedAt:                createdAt,
		ExpiresAt:                expiresAtDuration,
		Name:                     name,
		AccessCode:               accessCodeHashed,
		Key:                      key,
		Value:                    valueEncrypted,
		IsExpireAfterFirstUnlock: isExpireAfterFirstUnlock,
	}

	// Add the record to the database.
	if err := a.Database.QueryAddSecret(secret); err != nil {
		// Send a 500 bad request response.
		w.WriteHeader(http.StatusInternalServerError)

		// Render the encrypt secret error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Add secret", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Redirect to the share secret page.
	w.Header().Set("HX-Location", fmt.Sprintf("/dashboard/share/%s", secret.Key))
}

// APIUnlockSecretHandler renders the unlocked secret block (POST).
func (a *Application) APIUnlockSecretHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

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

	// Create a new hashed access code string with salt.
	accessCodeHashed := helpers.HashString(
		fmt.Sprintf("%d", secret.CreatedAt.Unix()),
		a.Config.SecretKey, accessCode,
	)

	// Check, if the access code is correct.
	if accessCodeHashed != secret.AccessCode {
		// Send a 400 not found response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the secret page with 400 error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Access code", Message: messages.ErrSecretAccessCodeNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Decrypt the secret value.
	decryptedValue, err := helpers.DecryptString(a.Config.SecretKey, secret.Value)
	if err != nil {
		// Send a 400 not found response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the secret page with 400 error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Decrypt secret", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Set component options.
	secret.Value = decryptedValue

	// Render the secret page.
	_ = pages.Secret(&secret, "unlocked").Render(r.Context(), w)
}

// APIRenewSecretExpiresAtFieldByKeyHandler renews a secret 'expires_at' field by its key from the database (PATCH).
func (a *Application) APIRenewSecretExpiresAtFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateExpiresAtFieldByKey(key, time.Now().Add(time.Hour*24).Local()); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the HX-Trigger header (to trigger a re-render by htmx).
	w.Header().Set("HX-Trigger", "getActiveSecrets, getExpiredSecrets")
}

// APIRestoreSecretAccessCodeFieldByKeyHandler restores a secret 'access_code' field by its key from the database (PATCH).
func (a *Application) APIRestoreSecretAccessCodeFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Get the secret record by its key from the database.
	secret, err := a.Database.QueryGetSecretByKey(key)
	if err != nil {
		helpers.WrapHTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	// Set the access code to a new random value (size 8).
	accessCode, err := helpers.GenerateAccessCode(8)
	if err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Create a new hashed access code string with salt.
	accessCodeHashed := helpers.HashString(
		fmt.Sprintf("%d", secret.CreatedAt.Unix()),
		a.Config.SecretKey, accessCode,
	)

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateAccessCodeFieldByKey(key, accessCodeHashed); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Render the restore access code block.
	_ = components.DashboardRestoreAccessCode(accessCode).Render(r.Context(), w)
}

// APIExpireSecretExpiresAtFieldByKeyHandler expires a secret 'expires_at' field by its key from the database (PATCH).
func (a *Application) APIExpireSecretExpiresAtFieldByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Patch the record by its key from the database.
	if err := a.Database.QueryUpdateExpiresAtFieldByKey(key, time.Now().Add(time.Second*1).Local()); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

// APIDeleteSecretByKeyHandler deletes a secret by its key from the database (DELETE).
func (a *Application) APIDeleteSecretByKeyHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Delete the record by its key from the database.
	if err := a.Database.QueryDeleteSecretByKey(key); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the HX-Trigger header (to trigger a re-render by htmx).
	w.Header().Set("HX-Trigger", "getActiveSecrets, getExpiredSecrets")
}

// APIDashboardActiveSecretsHandler renders the active secrets block (GET).
func (a *Application) APIDashboardActiveSecretsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Get all active secrets.
	secrets, err := a.Database.QueryGetActiveSecrets()
	if err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Render the active secrets block.
	_ = components.ActiveSecrets(secrets).Render(r.Context(), w)
}

// APIDashboardExpiredSecretsHandler renders the expired secrets block.
func (a *Application) APIDashboardExpiredSecretsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Extract the token expires datetime from the JWT.
	token, err := helpers.ExtractJWT(r.Header.Get("Authorization"), a.Config.SecretKey)
	if err != nil {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTClaimsNotValid)
		return
	}

	// Check, if the token is expired.
	if token.ExpiresAt < time.Now().Unix() {
		// Send a 403 forbidden response.
		helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTExpired)
		return
	}

	// Get all expired secrets.
	secrets, err := a.Database.QueryGetExpiredSecrets()
	if err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Render the active secrets block.
	_ = components.ExpiredSecrets(secrets).Render(r.Context(), w)
}

// APIUserLoginHandler logs in the user (POST).
func (a *Application) APIUserLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Parse the form data.
	if err := r.ParseForm(); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Form data", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}

	// Get form values.
	username := r.FormValue("username")
	masterPassword := r.FormValue("master_password")

	// Check, if the form values are valid.
	if err := helpers.ValidateUserSignInForm(username, masterPassword); err != nil {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(err).Render(r.Context(), w)

		return
	}

	// Authenticate the admin.
	if a.Config.MasterUsername != username || a.Config.MasterPassword != masterPassword {
		// Send a 401 unauthorized response.
		w.WriteHeader(http.StatusUnauthorized)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Form data", Message: messages.ErrFormLoginUserCredentialsNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Generate JWT.
	jwt, err := helpers.GenerateJWT(a.Config.SecretKey)
	if err != nil {
		// Send a 401 unauthorized response.
		w.WriteHeader(http.StatusUnauthorized)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "JWT", Message: err.Error()},
			},
		).Render(r.Context(), w)

		return
	}
	// Save JWT to local storage (on the client) and redirect to the dashboard page.
	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"jwtSaveToLocalStorage": %q}`, jwt))
	w.Header().Set("HX-Location", "/dashboard")
}

// APIUserLogoutHandler logs out the user (GET).
func (a *Application) APIUserLogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Check, if the request has a 'HX-Request' header.
	if !helpers.IsRequestValid(r) {
		// Send a 400 bad request response.
		w.WriteHeader(http.StatusBadRequest)

		// Render the form validation error.
		_ = components.FormValidationError(
			[]*messages.ErrorField{
				{Name: "Request", Message: messages.ErrHTMXHeaderNotValid},
			},
		).Render(r.Context(), w)

		return
	}

	// Remove JWT from local storage (on the client) and redirect to the index page.
	w.Header().Set("HX-Trigger", `{"jwtRemoveFromLocalStorage": ""}`)
	w.Header().Set("HX-Redirect", "/")
}
