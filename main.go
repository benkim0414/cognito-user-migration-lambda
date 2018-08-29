package main

import (
	"errors"
)

const (
	// CognitoTriggerSourceUserMigrationAuthentication triggers a user migration at the time of sign in.
	CognitoTriggerSourceUserMigrationAuthentication = "UserMigration_Authentication"
	// CognitoTriggerSourceUserMigrationForgotPassword triggers a user migration during forgot-password flow.
	CognitoTriggerSourceUserMigrationForgotPassword = "UserMigration_ForgotPassword"
)

var (
	// ErrBadPassword is thrown if the password entered by the user for sign-in is wrong.
	ErrBadPassword = errors.New("Bad password")
	// ErrBadTriggerSource is thrown when the specified trigger source is not supported.
	ErrBadTriggerSource = errors.New("Bad trigger source")
)
