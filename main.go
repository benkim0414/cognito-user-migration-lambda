package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
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

// Handler migrates the user with an existing password and suppresses the welcome message from Amazon Cognito.
func Handler(event CognitoEventUserPoolsMigrateUser) (CognitoEventUserPoolsMigrateUser, error) {
	switch event.TriggerSource {
	case CognitoTriggerSourceUserMigrationAuthentication:
		// Authenticate the user with existing user directory service
		user, err := authenticateUser(event.UserName, event.Request.Password)
		if err != nil {
			return CognitoEventUserPoolsMigrateUser{}, err
		}
		event.Response.UserAttributes = map[string]string{
			"email":          user.email,
			"email_verified": "true",
		}
		event.Response.FinalUserStatus = "CONFIRMED"
		event.Response.MessageAction = "SUPPRESS"
		return event, nil
	case CognitoTriggerSourceUserMigrationForgotPassword:
		// Lookup the user in existing user directory service
		user, err := lookupUser(event.UserName)
		if err != nil {
			return CognitoEventUserPoolsMigrateUser{}, err
		}
		event.Response.UserAttributes = map[string]string{
			"email":          user.email,
			"email_verified": "true",
		}
		event.Response.MessageAction = "SUPPRESS"
		return event, nil
	default:
		return CognitoEventUserPoolsMigrateUser{}, ErrBadTriggerSource
	}
}

func main() {
	lambda.Start(Handler)
}
