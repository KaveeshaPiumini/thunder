/*
 * Copyright (c) 2026, WSO2 LLC. (https://www.wso2.com).
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

package passkey

import (
	tidcommon "github.com/thunder-id/thunderid/pkg/thunderidengine/common"
)

// Client errors for passkey authentication service

var (
	// ErrorEmptyUserIdentifier is returned when both userID and username are empty.
	ErrorEmptyUserIdentifier = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1001",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_user_identifier",
			DefaultValue: "Empty user identifier",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_user_identifier_description",
			DefaultValue: "Either user ID or username must be provided",
		},
	}
	// ErrorEmptyRelyingPartyID is returned when the relying party ID is empty.
	ErrorEmptyRelyingPartyID = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1002",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_relying_party_id",
			DefaultValue: "Empty relying party ID",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_relying_party_id_description",
			DefaultValue: "The relying party ID is required",
		},
	}
	// ErrorEmptyCredentialID is returned when the credential ID is empty.
	ErrorEmptyCredentialID = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1003",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_credential_id",
			DefaultValue: "Empty credential ID",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_credential_id_description",
			DefaultValue: "The credential ID is required",
		},
	}
	// ErrorEmptyCredentialType is returned when the credential type is empty.
	ErrorEmptyCredentialType = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1004",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_credential_type",
			DefaultValue: "Empty credential type",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_credential_type_description",
			DefaultValue: "The credential type is required",
		},
	}
	// ErrorInvalidAuthenticatorResponse is returned when the authenticator response is invalid.
	ErrorInvalidAuthenticatorResponse = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1005",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_authenticator_response",
			DefaultValue: "Invalid authenticator response",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key: "error.passkeyservice.invalid_authenticator_response_description",
			DefaultValue: "The authenticator response is missing required fields " +
				"(clientDataJSON, authenticatorData, or signature)",
		},
	}
	// ErrorEmptySessionToken is returned when the session token is empty.
	ErrorEmptySessionToken = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1006",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_session_token",
			DefaultValue: "Empty session token",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.empty_session_token_description",
			DefaultValue: "The session token is required",
		},
	}
	// ErrorInvalidFinishData is returned when the finish data is nil.
	ErrorInvalidFinishData = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1007",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_finish_data",
			DefaultValue: "Invalid finish data",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_finish_data_description",
			DefaultValue: "The finish data cannot be null",
		},
	}
	// ErrorInvalidChallenge is returned when the challenge validation fails.
	ErrorInvalidChallenge = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1008",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_challenge",
			DefaultValue: "Invalid challenge",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_challenge_description",
			DefaultValue: "The challenge in the response does not match the expected challenge",
		},
	}
	// ErrorInvalidSignature is returned when signature verification fails.
	ErrorInvalidSignature = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1009",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_signature",
			DefaultValue: "Invalid signature",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_signature_description",
			DefaultValue: "The signature verification failed",
		},
	}
	// ErrorCredentialNotFound is returned when the credential is not found.
	ErrorCredentialNotFound = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1010",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.credential_not_found",
			DefaultValue: "Passkey credential not found",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.credential_not_found_description",
			DefaultValue: "The specified credential was not found for the user",
		},
	}
	// ErrorInvalidAttestationResponse is returned when the attestation response is invalid.
	ErrorInvalidAttestationResponse = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1011",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_attestation_response",
			DefaultValue: "Invalid attestation response",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_attestation_response_description",
			DefaultValue: "The attestation response is missing required fields (clientDataJSON or attestationObject)",
		},
	}
	// ErrorUserNotFound is returned when the user is not found.
	ErrorUserNotFound = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1012",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.user_not_found",
			DefaultValue: "User not found",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.user_not_found_description",
			DefaultValue: "The specified user was not found",
		},
	}
	// ErrorInvalidSessionToken is returned when the session token is invalid.
	ErrorInvalidSessionToken = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1013",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_session_token",
			DefaultValue: "Invalid session token",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.invalid_session_token_description",
			DefaultValue: "The session token is invalid or malformed",
		},
	}
	// ErrorSessionExpired is returned when the session has expired.
	ErrorSessionExpired = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1014",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.session_expired",
			DefaultValue: "Session expired",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.session_expired_description",
			DefaultValue: "The session has expired. Please start a new session",
		},
	}
	// ErrorNoCredentialsFound is returned when no credentials are found for the user.
	ErrorNoCredentialsFound = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1015",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.no_credentials_found",
			DefaultValue: "No credentials found",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.no_credentials_found_description",
			DefaultValue: "No credentials found for the user. Please register a credential first",
		},
	}
)

// WebAuthn library error codes.
//
// These errors wrap the failure categories reported by the underlying go-webauthn library
// (protocol.Error.Type) into stable internal codes, so that library-specific details are
// never returned to external clients and a change of the underlying library does not alter
// the API error contract. mapWebAuthnError performs the translation.

var (
	// ErrorWebAuthnBadRequest wraps the library "invalid_request" error category.
	ErrorWebAuthnBadRequest = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1016",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_bad_request",
			DefaultValue: "Invalid WebAuthn request",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_bad_request_description",
			DefaultValue: "The WebAuthn request data could not be processed",
		},
	}
	// ErrorWebAuthnPolicyRestriction wraps the library "policy_restriction" error category.
	ErrorWebAuthnPolicyRestriction = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1017",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_policy_restriction",
			DefaultValue: "Policy restriction",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_policy_restriction_description",
			DefaultValue: "The operation was prevented by a policy restriction",
		},
	}
	// ErrorWebAuthnResponseParseError wraps the library "parse_error" error category.
	ErrorWebAuthnResponseParseError = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1018",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_response_parse_error",
			DefaultValue: "Malformed authenticator response",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_response_parse_error_description",
			DefaultValue: "The authenticator response could not be parsed",
		},
	}
	// ErrorWebAuthnAuthenticatorDataInvalid wraps the library "auth_data" error category.
	ErrorWebAuthnAuthenticatorDataInvalid = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1019",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_authenticator_data_invalid",
			DefaultValue: "Invalid authenticator data",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_authenticator_data_invalid_description",
			DefaultValue: "The authenticator data failed verification",
		},
	}
	// ErrorWebAuthnVerificationFailed wraps the library "verification_error" error category.
	ErrorWebAuthnVerificationFailed = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1020",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_verification_failed",
			DefaultValue: "Verification failed",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_verification_failed_description",
			DefaultValue: "The authenticator response failed verification",
		},
	}
	// ErrorWebAuthnInvalidCertificate wraps the library "invalid_certificate" error category.
	ErrorWebAuthnInvalidCertificate = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1021",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_invalid_certificate",
			DefaultValue: "Invalid attestation certificate",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_invalid_certificate_description",
			DefaultValue: "The attestation certificate is invalid",
		},
	}
	// ErrorWebAuthnUnsupportedKeyType wraps the library "invalid_key_type" error category.
	ErrorWebAuthnUnsupportedKeyType = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1022",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_unsupported_key_type",
			DefaultValue: "Unsupported key type",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_unsupported_key_type_description",
			DefaultValue: "The credential public key type is not supported",
		},
	}
	// ErrorWebAuthnUnsupportedAlgorithm wraps the library "unsupported_key_algorithm" error category.
	ErrorWebAuthnUnsupportedAlgorithm = tidcommon.ServiceError{
		Type: tidcommon.ClientErrorType,
		Code: "PSK-1023",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_unsupported_algorithm",
			DefaultValue: "Unsupported key algorithm",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_unsupported_algorithm_description",
			DefaultValue: "The credential public key algorithm is not supported",
		},
	}
	// ErrorWebAuthnMetadataError wraps the library "invalid_metadata" error category. This is a
	// server-side error because attestation metadata validation depends on relying party configuration.
	ErrorWebAuthnMetadataError = tidcommon.ServiceError{
		Type: tidcommon.ServerErrorType,
		Code: "PSK-1024",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_metadata_error",
			DefaultValue: "Attestation metadata error",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_metadata_error_description",
			DefaultValue: "Failed to validate the attestation metadata",
		},
	}
	// ErrorWebAuthnNotImplemented wraps the library "not_implemented" and "spec_unimplemented" error
	// categories. This is a server-side error because it signals a limitation of the library or spec.
	ErrorWebAuthnNotImplemented = tidcommon.ServiceError{
		Type: tidcommon.ServerErrorType,
		Code: "PSK-1025",
		Error: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_not_implemented",
			DefaultValue: "Unsupported WebAuthn operation",
		},
		ErrorDescription: tidcommon.I18nMessage{
			Key:          "error.passkeyservice.webauthn_not_implemented_description",
			DefaultValue: "The requested WebAuthn operation is not supported",
		},
	}
)
