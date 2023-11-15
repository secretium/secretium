package messages

const (

	/*
		Configuration error messages.
	*/

	// ErrConfigSecretKeyEmpty is returned when the secret key is empty.
	ErrConfigSecretKeyEmpty string = "secret key is empty"

	// ErrConfigSecretKeyLengthNotValid is returned when the secret key is not valid.
	ErrConfigSecretKeyLengthNotValid string = "secret key is not valid (length should be greater or equal to %d)"

	// ErrConfigMasterUsernameEmpty is returned when the master username is empty.
	ErrConfigMasterUsernameEmpty string = "master username is empty"

	// ErrConfigMasterUsernameLengthNotValid is returned when the master username is not valid.
	ErrConfigMasterUsernameLengthNotValid string = "master username is not valid (length should be greater than %d and less than %d)"

	// ErrConfigMasterPasswordEmpty is returned when the master password is empty.
	ErrConfigMasterPasswordEmpty string = "master password is empty"

	// ErrConfigMasterPasswordLengthNotValid is returned when the master password is not valid.
	ErrConfigMasterPasswordLengthNotValid string = "master password is not valid (length should be greater than %d and less than %d)"

	// ErrConfigDomainNotValid is returned when the domain has an invalid format.
	ErrConfigDomainNotValid string = "domain URL is not valid"

	// ErrConfigDomainSchemaNotValid is returned when the HTTP schema of the domain is not valid.
	ErrConfigDomainSchemaNotValid string = "domain HTTP schema is not valid"

	// ErrConfigServerPortNotValid is returned when the server port is not valid.
	ErrConfigServerPortNotValid string = "server port is not valid"

	// ErrConfigServerTimezoneNotValid is returned when the server timezone is not valid.
	ErrConfigServerTimezoneNotValid string = "server timezone is not valid"

	// ErrConfigServerReadTimeoutNotValid is returned when the server read timeout is not valid.
	ErrConfigServerReadTimeoutNotValid string = "server read timeout is not valid"

	// ErrConfigServerWriteTimeoutNotValid is returned when the server write timeout is not valid.
	ErrConfigServerWriteTimeoutNotValid string = "server write timeout is not valid"

	/*
		HTMX error messages.
	*/

	// ErrHTMXHeaderNotValid is returned when the 'HX-Request' header is missing or not valid.
	ErrHTMXHeaderNotValid string = "missing 'HX-Request' header or its value is not valid"

	/*
		Secret error messages.
	*/

	// ErrSecretKeyEmptyOrNotFound is returned when the secret key is empty or not found.
	ErrSecretKeyEmptyOrNotFound string = "secret key is empty or not found"

	// ErrSecretKeyLengthNotValid is returned when the secret key is not valid.
	ErrSecretKeyLengthNotValid string = "secret key is not valid (length should be strictly %d)"

	// ErrSecretExpiresAtNotValid is returned when the secret expires at datetime is not valid.
	ErrSecretExpiresAtNotValid string = "secret expires at datetime is not valid"

	// ErrSecretAccessCodeNotValid is returned when the secret access code is not valid.
	ErrSecretAccessCodeNotValid string = "secret access code is not valid"

	/*
		Form error messages.
	*/

	// ErrFormDataNotValid is returned when the form data is not valid.
	ErrFormDataNotValid string = "form data is not valid"

	// ErrFormSignInUserCredentialsNotValid is returned when the sign in user credentials are not valid.
	ErrFormLoginUserCredentialsNotValid string = "master username or password are empty or not valid"

	// ErrFormAddSecretNameLengthNotValid is returned when the secret name is not valid.
	ErrFormAddSecretNameLengthNotValid string = "secret name is not valid (length should be greater than %d and less than %d)"

	// ErrFormAddSecretValueLengthNotValid is returned when the secret value is not valid.
	ErrFormAddSecretValueLengthNotValid string = "secret value is not valid (length should be greater or equal to %d)"

	// ErrFormAddSecretAccessCodeLengthNotValid is returned when the secret access code is not valid.
	ErrFormAddSecretAccessCodeLengthNotValid string = "secret access code is not valid (length should be greater than %d and less than %d)"

	/*
		Session error messages.
	*/

	// ErrSessionUserNotAuthenticated is returned when the user is not authenticated.
	ErrSessionUserNotAuthenticated string = "user is not authenticated"
)

// ErrorField represents an error field.
type ErrorField struct {
	Name, Message string
}
