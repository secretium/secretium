package helpers

import (
	"fmt"

	"github.com/secretium/secretium/internal/constants"
	"github.com/secretium/secretium/internal/messages"
)

// ErrorField represents an error field.
type ErrorField struct {
	Name, Message string
}

// ValidateAddSecretForm returns nil if the given add secret form values are valid.
func ValidateAddSecretForm(name, value string) (errorFields []*messages.ErrorField) {
	// Check if the name is empty or not valid (length should be greater than 3 and less than 32).
	if name == "" ||
		len(name) < constants.ConstFormAddSecretNameMinLength ||
		len(name) > constants.ConstFormAddSecretNameMaxLength {
		// Append error field.
		errorFields = append(
			errorFields,
			&messages.ErrorField{
				Name: "Name",
				Message: fmt.Sprintf(
					messages.ErrFormAddSecretNameLengthNotValid,
					constants.ConstFormAddSecretNameMinLength,
					constants.ConstFormAddSecretNameMaxLength,
				),
			},
		)
	}

	// Check if the value is empty or not valid (length should be greater or equal to 1).
	if value == "" || len(value) < constants.ConstFormAddSecretValueMinLength {
		// Append error field.
		errorFields = append(
			errorFields,
			&messages.ErrorField{
				Name: "Value",
				Message: fmt.Sprintf(
					messages.ErrFormAddSecretValueLengthNotValid,
					constants.ConstFormAddSecretValueMinLength,
				),
			},
		)
	}

	return errorFields
}

// ValidateViewSecretForm returns nil if the given view secret form access code is valid.
func ValidateViewSecretForm(accessCode string) (errorFields []*messages.ErrorField) {
	// Check if the access code is empty or not valid (length should be greater than 6 and less than 32).
	if accessCode == "" ||
		len(accessCode) < constants.ConstFormAddSecretAccessCodeMinLength ||
		len(accessCode) > constants.ConstFormAddSecretAccessCodeMaxLength {
		// Append error field.
		errorFields = append(
			errorFields,
			&messages.ErrorField{
				Name: "Access code",
				Message: fmt.Sprintf(
					messages.ErrFormAddSecretAccessCodeLengthNotValid,
					constants.ConstFormAddSecretAccessCodeMinLength,
					constants.ConstFormAddSecretAccessCodeMaxLength,
				),
			},
		)
	}

	return errorFields
}

// ValidateUserSignInForm returns nil if the given user sign in form values are valid.
func ValidateUserSignInForm(username, masterPassword string) (errorFields []*messages.ErrorField) {
	// Check if the username is empty or not valid (length should be greater than 4 and less than 16).
	if username == "" ||
		len(username) < constants.ConstConfigMasterUsernameMinLength ||
		len(username) > constants.ConstConfigMasterUsernameMaxLength {
		// Append error field.
		errorFields = append(
			errorFields,
			&messages.ErrorField{
				Name: "Username",
				Message: fmt.Sprintf(
					messages.ErrConfigMasterUsernameLengthNotValid,
					constants.ConstConfigMasterUsernameMinLength,
					constants.ConstConfigMasterUsernameMaxLength,
				),
			},
		)
	}

	// Check if the master password is empty or not valid (length should be greater than 8 and less than 16).
	if masterPassword == "" ||
		len(masterPassword) < constants.ConstConfigMasterPasswordMinLength ||
		len(masterPassword) > constants.ConstConfigMasterPasswordMaxLength {
		// Append error field.
		errorFields = append(
			errorFields,
			&messages.ErrorField{
				Name: "Master password",
				Message: fmt.Sprintf(
					messages.ErrConfigMasterPasswordLengthNotValid,
					constants.ConstConfigMasterPasswordMinLength,
					constants.ConstConfigMasterPasswordMaxLength,
				),
			},
		)
	}

	return errorFields
}
