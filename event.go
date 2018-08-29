package main

import "github.com/aws/aws-lambda-go/events"

// CognitoEventUserPoolsMigrateUser is sent by AWS Cognito User Pools when a user does not exist
// in the user pool at the time of sign-in with a password, or in the forgot-password flow,
// allowing the Lambda to migrate users.
type CognitoEventUserPoolsMigrateUser struct {
	events.CognitoEventUserPoolsHeader
	Request  CognitoEventUserPoolsMigrateUserRequest  `json:"request"`
	Response CognitoEventUserPoolsMigrateUserResponse `json:"response"`
}

// CognitoEventUserPoolsMigrateUserRequest contains the request portion of a MigrateUser Event
type CognitoEventUserPoolsMigrateUserRequest struct {
	Password string `json:"password"`
}

// CognitoEventUserPoolsMigrateUserResponse contains the response portion of a MigrateUser Event
type CognitoEventUserPoolsMigrateUserResponse struct {
	UserAttributes         map[string]string `json:"userAttributes"`
	FinalUserStatus        string            `json:"finalUserStatus"`
	MessageAction          string            `json:"messageAction"`
	DesiredDeliveryMediums []string          `json:"desiredDeliveryMediums"`
	ForceAliasCreation     bool              `json:"forceAliasCreation"`
}
