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

package notification

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/asgardeo/thunder/internal/notification/common"
	"github.com/asgardeo/thunder/internal/system/error/serviceerror"
)

// validateNotificationSender validates the notification sender data.
func validateNotificationSender(sender common.NotificationSenderDTO) *serviceerror.ServiceError {
	if sender.Name == "" {
		return &ErrorInvalidSenderName
	}

	switch sender.Type {
	case common.NotificationSenderTypeMessage:
		return validateMessageNotificationSender(sender)
	default:
		return &ErrorInvalidSenderType
	}
}

// validateMessageNotificationSender validates a message notification sender.
func validateMessageNotificationSender(sender common.NotificationSenderDTO) *serviceerror.ServiceError {
	if sender.Provider == "" {
		return &ErrorInvalidProvider
	}
	if sender.Provider != common.MessageProviderTypeTwilio &&
		sender.Provider != common.MessageProviderTypeVonage &&
		sender.Provider != common.MessageProviderTypeCustom {
		return &ErrorInvalidProvider
	}

	if err := validateMessageNotificationSenderProperties(sender); err != nil {
		svcErr := ErrorInvalidRequestFormat
		svcErr.ErrorDescription = err.Error()
		return &svcErr
	}

	return nil
}

// validateMessageNotificationSenderProperties validates the properties of a message notification sender.
func validateMessageNotificationSenderProperties(sender common.NotificationSenderDTO) error {
	if len(sender.Properties) == 0 {
		return errors.New("message notification sender properties cannot be empty")
	}

	switch sender.Provider {
	case common.MessageProviderTypeTwilio:
		return validateTwilioProperties(sender.Properties)
	case common.MessageProviderTypeVonage:
		return validateVonageProperties(sender.Properties)
	case common.MessageProviderTypeCustom:
		return validateCustomProperties(sender.Properties)
	default:
		return errors.New("unsupported message notification sender")
	}
}

// validateTwilioProperties validates the message notification sender properties for a Twilio client.
func validateTwilioProperties(properties []common.SenderProperty) error {
	requiredProps := map[string]bool{
		"account_sid": false,
		"auth_token":  false,
		"sender_id":   false,
	}
	err := validateSenderProperties(properties, requiredProps)
	if err != nil {
		return err
	}

	// Validate the account SID format
	sIDRegex := `^AC[0-9a-fA-F]{32}$`
	sid := ""
	for _, prop := range properties {
		if prop.Name == common.TwilioPropKeyAccountSID {
			sid = prop.Value
			break
		}
	}
	matched, err := regexp.MatchString(sIDRegex, sid)
	if err != nil {
		return fmt.Errorf("failed to validate Twilio account SID: %w", err)
	}
	if !matched {
		return errors.New("invalid Twilio account SID format")
	}

	return nil
}

// validateVonageProperties validates the message notification sender properties for a Vonage client.
func validateVonageProperties(properties []common.SenderProperty) error {
	requiredProps := map[string]bool{
		"api_key":    false,
		"api_secret": false,
		"sender_id":  false,
	}
	return validateSenderProperties(properties, requiredProps)
}

// validateCustomProperties validates the message notification sender properties for a custom client.
func validateCustomProperties(properties []common.SenderProperty) error {
	url := ""
	httpMethod := ""
	contentType := ""
	for _, prop := range properties {
		if prop.Name == "" {
			return errors.New("properties must have non-empty name")
		}
		switch prop.Name {
		case common.CustomPropKeyURL:
			url = prop.Value
		case common.CustomPropKeyHTTPMethod:
			httpMethod = strings.ToUpper(prop.Value)
		case common.CustomPropKeyContentType:
			contentType = strings.ToUpper(prop.Value)
		}
	}
	if url == "" {
		return errors.New("custom provider must have a URL property")
	}
	if httpMethod != "" && httpMethod != http.MethodGet &&
		httpMethod != http.MethodPost && httpMethod != http.MethodPut {
		return errors.New("custom provider must have a valid HTTP method")
	}
	if contentType != "" && contentType != "JSON" && contentType != "FORM" {
		return errors.New("custom provider must have a valid content type (JSON or FORM)")
	}

	return nil
}

// validateSenderProperties validates the properties for a notification sender.
func validateSenderProperties(properties []common.SenderProperty, requiredProperties map[string]bool) error {
	for _, prop := range properties {
		if prop.Name == "" {
			return errors.New("properties must have non-empty name")
		}
		if _, exists := requiredProperties[prop.Name]; exists {
			requiredProperties[prop.Name] = true
		}
	}

	// Check if all required properties are present
	for key, found := range requiredProperties {
		if !found {
			return errors.New("required property missing for the provider: " + key)
		}
	}
	return nil
}
