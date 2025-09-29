/*
 * Copyright (c) 2025, WSO2 LLC. (https://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package oauth

import "github.com/asgardeo/thunder/internal/system/error/serviceerror"

// Client errors for OAuth authentication.
var (
	// ErrorEmptyIdpID is the error when the IDP identifier is empty.
	ErrorEmptyIdpID = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1001",
		Error:            "IDP id is empty",
		ErrorDescription: "The identity provider id cannot be empty",
	}
	// ErrorInvalidIDP is the error when the retrieved IDP is invalid.
	ErrorInvalidIDP = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1002",
		Error:            "Invalid identity provider",
		ErrorDescription: "The retrieved identity provider is invalid or empty",
	}
	// ErrorEmptyAuthorizationCode is the error when the authorization code is empty.
	ErrorEmptyAuthorizationCode = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1003",
		Error:            "Empty authorization code",
		ErrorDescription: "The authorization code cannot be empty",
	}
	// ErrorEmptyAccessToken is the error when the access token is empty.
	ErrorEmptyAccessToken = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1004",
		Error:            "Empty access token",
		ErrorDescription: "The access token cannot be empty",
	}
	// ErrorClientErrorWhileRetrievingIDP is the error when there is a client error while retrieving the IDP.
	ErrorClientErrorWhileRetrievingIDP = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1005",
		Error:            "Failed to retrieve identity provider",
		ErrorDescription: "A client error occurred while retrieving the identity provider configuration",
	}
	// ErrorEmptySubClaim is the error when the sub claim is empty.
	ErrorEmptySubClaim = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1006",
		Error:            "Empty sub claim",
		ErrorDescription: "The sub claim cannot be empty",
	}
	// ErrorUserNotFound is the error when the user is not found.
	ErrorUserNotFound = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1007",
		Error:            "User not found",
		ErrorDescription: "No user found for the provided sub claim",
	}
	// ErrorClientErrorWhileRetrievingUser is the error when there is a client error while retrieving the user.
	ErrorClientErrorWhileRetrievingUser = serviceerror.ServiceError{
		Type:             serviceerror.ClientErrorType,
		Code:             "AUTH-OAUTH-1008",
		Error:            "Failed to retrieve user",
		ErrorDescription: "A client error occurred while retrieving the internal user",
	}
)

// Server errors for OAuth authentication.
var (
	// ErrorUnexpectedServerError is a generic error for unexpected server errors.
	ErrorUnexpectedServerError = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5000",
		Error:            "Something went wrong",
		ErrorDescription: "An unexpected error occurred while processing the request",
	}
	// ErrorInvalidIDPConfig is the error when the IDP configuration is invalid or incomplete.
	ErrorInvalidIDPConfig = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5001",
		Error:            "Invalid IDP configuration",
		ErrorDescription: "The configuration for the specified identity provider is invalid or incomplete",
	}
	// ErrorDuringTokenExchange is the error when there is an error during the token exchange process.
	ErrorDuringTokenExchange = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5002",
		Error:            "Error during token exchange",
		ErrorDescription: "An error occurred while exchanging the authorization code for token",
	}
	// ErrorInvalidTokenResponse is the error when the token response is invalid.
	ErrorInvalidTokenResponse = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5003",
		Error:            "Invalid token response",
		ErrorDescription: "The token response received from the identity provider is invalid",
	}
	// ErrorFetchingUserInfo is the error when there is an error fetching user information.
	ErrorFetchingUserInfo = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5004",
		Error:            "Error fetching user information",
		ErrorDescription: "An error occurred while fetching user information from the identity provider",
	}
	// ErrorServerErrorWhileRetrievingUser is the error when there is a server error while retrieving the user.
	ErrorServerErrorWhileRetrievingUser = serviceerror.ServiceError{
		Type:             serviceerror.ServerErrorType,
		Code:             "AUTH-OAUTH-5005",
		Error:            "Failed to retrieve user",
		ErrorDescription: "A server error occurred while retrieving the internal user",
	}
)

// customServiceError creates a new service error based on an existing error with custom description.
func customServiceError(svcError serviceerror.ServiceError, errorDesc string) *serviceerror.ServiceError {
	err := &serviceerror.ServiceError{
		Type:             svcError.Type,
		Code:             svcError.Code,
		Error:            svcError.Error,
		ErrorDescription: svcError.ErrorDescription,
	}
	if errorDesc != "" {
		err.ErrorDescription = errorDesc
	}
	return err
}
