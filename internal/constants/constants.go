package constants

const (

	/*
		Configuration constants.
	*/

	// ConstConfigSecretKeyMinLength is the minimum length of the secret key.
	ConstConfigSecretKeyMinLength int = 16

	// ConstConfigMasterUsernameMinLength is the minimum length of the master username.
	ConstConfigMasterUsernameMinLength int = 4

	// ConstConfigMasterUsernameMaxLength is the maximum length of the master username.
	ConstConfigMasterUsernameMaxLength int = 16

	// ConstConfigMasterPasswordMinLength is the minimum length of the master password.
	ConstConfigMasterPasswordMinLength int = 8

	// ConstConfigMasterPasswordMaxLength is the maximum length of the master password.
	ConstConfigMasterPasswordMaxLength int = 16

	// ConstConfigDomain is the domain URL.
	ConstConfigDomain string = "localhost"

	// ConstConfigDomainSchema is the domain HTTP schema.
	ConstConfigDomainSchema string = "http"

	// ConstConfigServerPort is the port the server is listening on.
	ConstConfigServerPort string = "8787"

	// ConstConfigServerReadTimeout is the server read timeout.
	ConstConfigServerReadTimeout string = "5"

	// ConstConfigServerWriteTimeout is the server write timeout.
	ConstConfigServerWriteTimeout string = "10"

	// ConstConfigServerTimezone is the server timezone.
	ConstConfigServerTimezone string = "Europe/Moscow"

	// ConstConfigSQLitePath is the path to the SQLite database.
	ConstConfigSQLitePath string = "secretium-data"

	/*
		Form constants.
	*/

	// ConstFormAddSecretNameMinLength is the minimum length of the secret name.
	ConstFormAddSecretNameMinLength int = 3

	// ConstFormAddSecretNameMaxLength is the maximum length of the secret name.
	ConstFormAddSecretNameMaxLength int = 32

	// ConstFormAddSecretValueMinLength is the minimum length of the secret value.
	ConstFormAddSecretValueMinLength int = 1

	// ConstFormAddSecretAccessCodeMinLength is the minimum length of the secret access code.
	ConstFormAddSecretAccessCodeMinLength int = 6

	// ConstFormAddSecretAccessCodeMaxLength is the maximum length of the secret access code.
	ConstFormAddSecretAccessCodeMaxLength int = 32
)
