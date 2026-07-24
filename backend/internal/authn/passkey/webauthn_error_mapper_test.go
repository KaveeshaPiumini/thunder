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
	"errors"
	"fmt"
	"testing"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/stretchr/testify/assert"

	tidcommon "github.com/thunder-id/thunderid/pkg/thunderidengine/common"
)

// TestMapWebAuthnError_LibraryCategories verifies that every go-webauthn protocol error
// category is wrapped into its dedicated internal ServiceError code.
func TestMapWebAuthnError_LibraryCategories(t *testing.T) {
	fallback := ErrorInvalidSignature

	tests := []struct {
		name     string
		libErr   *protocol.Error
		expected tidcommon.ServiceError
	}{
		{"invalid_request", protocol.ErrBadRequest, ErrorWebAuthnBadRequest},
		{"policy_restriction", protocol.ErrPolicyRestriction, ErrorWebAuthnPolicyRestriction},
		{"challenge_mismatch", protocol.ErrChallengeMismatch, ErrorInvalidChallenge},
		{"parse_error", protocol.ErrParsingData, ErrorWebAuthnResponseParseError},
		{"auth_data", protocol.ErrAuthData, ErrorWebAuthnAuthenticatorDataInvalid},
		{"verification_error", protocol.ErrVerification, ErrorWebAuthnVerificationFailed},
		{"attestation_error", protocol.ErrAttestation, ErrorInvalidAttestationResponse},
		{"invalid_attestation", protocol.ErrInvalidAttestation, ErrorInvalidAttestationResponse},
		{"invalid_attestation_format", protocol.ErrAttestationFormat, ErrorInvalidAttestationResponse},
		{"invalid_certificate", protocol.ErrAttestationCertificate, ErrorWebAuthnInvalidCertificate},
		{"invalid_signature", protocol.ErrAssertionSignature, ErrorInvalidSignature},
		{"invalid_key_type", protocol.ErrUnsupportedKey, ErrorWebAuthnUnsupportedKeyType},
		{"unsupported_key_algorithm", protocol.ErrUnsupportedAlgorithm, ErrorWebAuthnUnsupportedAlgorithm},
		{"invalid_metadata", protocol.ErrMetadata, ErrorWebAuthnMetadataError},
		{"spec_unimplemented", protocol.ErrNotSpecImplemented, ErrorWebAuthnNotImplemented},
		{"not_implemented", protocol.ErrNotImplemented, ErrorWebAuthnNotImplemented},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mapWebAuthnError(tc.libErr, fallback)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expected.Code, result.Code)
			assert.Equal(t, tc.expected.Type, result.Type)
			// The internal error must never carry the raw library message.
			assert.NotEqual(t, tc.libErr.Details, result.ErrorDescription.DefaultValue)
		})
	}
}

// TestMapWebAuthnError_WithDetailsAndInfo ensures the mapping is driven by the error type and
// is not affected by library-supplied details or debug info that vary between library versions.
func TestMapWebAuthnError_WithDetailsAndInfo(t *testing.T) {
	libErr := protocol.ErrVerification.
		WithDetails("some internal library detail").
		WithInfo("stack-like debug info")

	result := mapWebAuthnError(libErr, ErrorInvalidSignature)

	assert.NotNil(t, result)
	assert.Equal(t, ErrorWebAuthnVerificationFailed.Code, result.Code)
	assert.NotContains(t, result.ErrorDescription.DefaultValue, "internal library detail")
}

// TestMapWebAuthnError_UnknownCredential verifies the dedicated unknown-credential wrapper type
// is mapped to the credential-not-found error rather than the fallback.
func TestMapWebAuthnError_UnknownCredential(t *testing.T) {
	libErr := &protocol.ErrorUnknownCredential{Err: protocol.ErrBadRequest}

	result := mapWebAuthnError(libErr, ErrorInvalidSignature)

	assert.NotNil(t, result)
	assert.Equal(t, ErrorCredentialNotFound.Code, result.Code)
}

// TestMapWebAuthnError_WrappedProtocolError verifies that a protocol error wrapped with
// fmt.Errorf is still recognized through errors.As.
func TestMapWebAuthnError_WrappedProtocolError(t *testing.T) {
	wrapped := fmt.Errorf("failed to validate assertion: %w", protocol.ErrChallengeMismatch)

	result := mapWebAuthnError(wrapped, ErrorInvalidSignature)

	assert.NotNil(t, result)
	assert.Equal(t, ErrorInvalidChallenge.Code, result.Code)
}

// TestMapWebAuthnError_UnknownType falls back for a protocol error whose type is not recognized.
func TestMapWebAuthnError_UnknownType(t *testing.T) {
	libErr := &protocol.Error{Type: "some_future_error_type", Details: "unexpected"}

	result := mapWebAuthnError(libErr, ErrorInvalidAttestationResponse)

	assert.NotNil(t, result)
	assert.Equal(t, ErrorInvalidAttestationResponse.Code, result.Code)
}

// TestMapWebAuthnError_NonProtocolError falls back for a plain error that is not a protocol error.
func TestMapWebAuthnError_NonProtocolError(t *testing.T) {
	result := mapWebAuthnError(errors.New("plain error"), ErrorInvalidSignature)

	assert.NotNil(t, result)
	assert.Equal(t, ErrorInvalidSignature.Code, result.Code)
}

// TestMapWebAuthnError_Nil returns nil when there is no error.
func TestMapWebAuthnError_Nil(t *testing.T) {
	assert.Nil(t, mapWebAuthnError(nil, ErrorInvalidSignature))
}
