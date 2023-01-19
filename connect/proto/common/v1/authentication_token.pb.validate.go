// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/authentication_token.proto

package commonv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AuthenticationHeaderConfiguration with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *AuthenticationHeaderConfiguration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthenticationHeaderConfiguration
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// AuthenticationHeaderConfigurationMultiError, or nil if none found.
func (m *AuthenticationHeaderConfiguration) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthenticationHeaderConfiguration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetHeaderKey()); l < 1 || l > 20 {
		err := AuthenticationHeaderConfigurationValidationError{
			field:  "HeaderKey",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHeaderValuePrefix()) > 30 {
		err := AuthenticationHeaderConfigurationValidationError{
			field:  "HeaderValuePrefix",
			reason: "value length must be at most 30 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AuthenticationHeaderConfigurationMultiError(errors)
	}

	return nil
}

// AuthenticationHeaderConfigurationMultiError is an error wrapping multiple
// validation errors returned by
// AuthenticationHeaderConfiguration.ValidateAll() if the designated
// constraints aren't met.
type AuthenticationHeaderConfigurationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthenticationHeaderConfigurationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthenticationHeaderConfigurationMultiError) AllErrors() []error { return m }

// AuthenticationHeaderConfigurationValidationError is the validation error
// returned by AuthenticationHeaderConfiguration.Validate if the designated
// constraints aren't met.
type AuthenticationHeaderConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticationHeaderConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticationHeaderConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticationHeaderConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticationHeaderConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticationHeaderConfigurationValidationError) ErrorName() string {
	return "AuthenticationHeaderConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticationHeaderConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticationHeaderConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticationHeaderConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticationHeaderConfigurationValidationError{}

// Validate checks the field values on ApiTokenConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApiTokenConfiguration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApiTokenConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApiTokenConfigurationMultiError, or nil if none found.
func (m *ApiTokenConfiguration) ValidateAll() error {
	return m.validate(true)
}

func (m *ApiTokenConfiguration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHeaderConfiguration() == nil {
		err := ApiTokenConfigurationValidationError{
			field:  "HeaderConfiguration",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHeaderConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApiTokenConfigurationValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApiTokenConfigurationValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeaderConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiTokenConfigurationValidationError{
				field:  "HeaderConfiguration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetApiTokenValue()) < 1 {
		err := ApiTokenConfigurationValidationError{
			field:  "ApiTokenValue",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPemCaChain()) < 1 {
		err := ApiTokenConfigurationValidationError{
			field:  "PemCaChain",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ApiTokenConfigurationMultiError(errors)
	}

	return nil
}

// ApiTokenConfigurationMultiError is an error wrapping multiple validation
// errors returned by ApiTokenConfiguration.ValidateAll() if the designated
// constraints aren't met.
type ApiTokenConfigurationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApiTokenConfigurationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApiTokenConfigurationMultiError) AllErrors() []error { return m }

// ApiTokenConfigurationValidationError is the validation error returned by
// ApiTokenConfiguration.Validate if the designated constraints aren't met.
type ApiTokenConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiTokenConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiTokenConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiTokenConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiTokenConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiTokenConfigurationValidationError) ErrorName() string {
	return "ApiTokenConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e ApiTokenConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiTokenConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiTokenConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiTokenConfigurationValidationError{}

// Validate checks the field values on CreateApiTokenConfigurationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreateApiTokenConfigurationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateApiTokenConfigurationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateApiTokenConfigurationRequestMultiError, or nil if none found.
func (m *CreateApiTokenConfigurationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateApiTokenConfigurationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHeaderConfiguration() == nil {
		err := CreateApiTokenConfigurationRequestValidationError{
			field:  "HeaderConfiguration",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHeaderConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateApiTokenConfigurationRequestValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateApiTokenConfigurationRequestValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeaderConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateApiTokenConfigurationRequestValidationError{
				field:  "HeaderConfiguration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetApiTokenValue()) < 1 {
		err := CreateApiTokenConfigurationRequestValidationError{
			field:  "ApiTokenValue",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPemCaChain()) < 1 {
		err := CreateApiTokenConfigurationRequestValidationError{
			field:  "PemCaChain",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateApiTokenConfigurationRequestMultiError(errors)
	}

	return nil
}

// CreateApiTokenConfigurationRequestMultiError is an error wrapping multiple
// validation errors returned by
// CreateApiTokenConfigurationRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateApiTokenConfigurationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateApiTokenConfigurationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateApiTokenConfigurationRequestMultiError) AllErrors() []error { return m }

// CreateApiTokenConfigurationRequestValidationError is the validation error
// returned by CreateApiTokenConfigurationRequest.Validate if the designated
// constraints aren't met.
type CreateApiTokenConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApiTokenConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApiTokenConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApiTokenConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApiTokenConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApiTokenConfigurationRequestValidationError) ErrorName() string {
	return "CreateApiTokenConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateApiTokenConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApiTokenConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApiTokenConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApiTokenConfigurationRequestValidationError{}

// Validate checks the field values on CreateApiTokenConfigurationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreateApiTokenConfigurationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateApiTokenConfigurationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateApiTokenConfigurationResponseMultiError, or nil if none found.
func (m *CreateApiTokenConfigurationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateApiTokenConfigurationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateApiTokenConfigurationResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetVersion() < 0 {
		err := CreateApiTokenConfigurationResponseValidationError{
			field:  "Version",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateApiTokenConfigurationResponseMultiError(errors)
	}

	return nil
}

// CreateApiTokenConfigurationResponseMultiError is an error wrapping multiple
// validation errors returned by
// CreateApiTokenConfigurationResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateApiTokenConfigurationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateApiTokenConfigurationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateApiTokenConfigurationResponseMultiError) AllErrors() []error { return m }

// CreateApiTokenConfigurationResponseValidationError is the validation error
// returned by CreateApiTokenConfigurationResponse.Validate if the designated
// constraints aren't met.
type CreateApiTokenConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApiTokenConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApiTokenConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApiTokenConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApiTokenConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApiTokenConfigurationResponseValidationError) ErrorName() string {
	return "CreateApiTokenConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateApiTokenConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApiTokenConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApiTokenConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApiTokenConfigurationResponseValidationError{}

// Validate checks the field values on GetApiTokenConfigurationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetApiTokenConfigurationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApiTokenConfigurationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetApiTokenConfigurationRequestMultiError, or nil if none found.
func (m *GetApiTokenConfigurationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApiTokenConfigurationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetApiTokenConfigurationRequestMultiError(errors)
	}

	return nil
}

// GetApiTokenConfigurationRequestMultiError is an error wrapping multiple
// validation errors returned by GetApiTokenConfigurationRequest.ValidateAll()
// if the designated constraints aren't met.
type GetApiTokenConfigurationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApiTokenConfigurationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApiTokenConfigurationRequestMultiError) AllErrors() []error { return m }

// GetApiTokenConfigurationRequestValidationError is the validation error
// returned by GetApiTokenConfigurationRequest.Validate if the designated
// constraints aren't met.
type GetApiTokenConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApiTokenConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApiTokenConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApiTokenConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApiTokenConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApiTokenConfigurationRequestValidationError) ErrorName() string {
	return "GetApiTokenConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApiTokenConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApiTokenConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApiTokenConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApiTokenConfigurationRequestValidationError{}

// Validate checks the field values on GetApiTokenConfigurationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetApiTokenConfigurationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApiTokenConfigurationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetApiTokenConfigurationResponseMultiError, or nil if none found.
func (m *GetApiTokenConfigurationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApiTokenConfigurationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApiTokenConfigurationResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetApiTokenConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetApiTokenConfigurationResponseValidationError{
					field:  "ApiTokenConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetApiTokenConfigurationResponseValidationError{
					field:  "ApiTokenConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApiTokenConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetApiTokenConfigurationResponseValidationError{
				field:  "ApiTokenConfiguration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetApiTokenConfigurationResponseMultiError(errors)
	}

	return nil
}

// GetApiTokenConfigurationResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetApiTokenConfigurationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetApiTokenConfigurationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApiTokenConfigurationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApiTokenConfigurationResponseMultiError) AllErrors() []error { return m }

// GetApiTokenConfigurationResponseValidationError is the validation error
// returned by GetApiTokenConfigurationResponse.Validate if the designated
// constraints aren't met.
type GetApiTokenConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApiTokenConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApiTokenConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApiTokenConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApiTokenConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApiTokenConfigurationResponseValidationError) ErrorName() string {
	return "GetApiTokenConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetApiTokenConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApiTokenConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApiTokenConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApiTokenConfigurationResponseValidationError{}

// Validate checks the field values on UpdateApiTokenConfigurationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdateApiTokenConfigurationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateApiTokenConfigurationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateApiTokenConfigurationRequestMultiError, or nil if none found.
func (m *UpdateApiTokenConfigurationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateApiTokenConfigurationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHeaderConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApiTokenConfigurationRequestValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApiTokenConfigurationRequestValidationError{
					field:  "HeaderConfiguration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeaderConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApiTokenConfigurationRequestValidationError{
				field:  "HeaderConfiguration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ApiTokenValue

	if len(errors) > 0 {
		return UpdateApiTokenConfigurationRequestMultiError(errors)
	}

	return nil
}

// UpdateApiTokenConfigurationRequestMultiError is an error wrapping multiple
// validation errors returned by
// UpdateApiTokenConfigurationRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateApiTokenConfigurationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateApiTokenConfigurationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateApiTokenConfigurationRequestMultiError) AllErrors() []error { return m }

// UpdateApiTokenConfigurationRequestValidationError is the validation error
// returned by UpdateApiTokenConfigurationRequest.Validate if the designated
// constraints aren't met.
type UpdateApiTokenConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateApiTokenConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateApiTokenConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateApiTokenConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateApiTokenConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateApiTokenConfigurationRequestValidationError) ErrorName() string {
	return "UpdateApiTokenConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateApiTokenConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateApiTokenConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateApiTokenConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateApiTokenConfigurationRequestValidationError{}

// Validate checks the field values on UpdateApiTokenConfigurationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdateApiTokenConfigurationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateApiTokenConfigurationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateApiTokenConfigurationResponseMultiError, or nil if none found.
func (m *UpdateApiTokenConfigurationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateApiTokenConfigurationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApiTokenConfigurationResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApiTokenConfigurationResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateApiTokenConfigurationResponseMultiError(errors)
	}

	return nil
}

// UpdateApiTokenConfigurationResponseMultiError is an error wrapping multiple
// validation errors returned by
// UpdateApiTokenConfigurationResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateApiTokenConfigurationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateApiTokenConfigurationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateApiTokenConfigurationResponseMultiError) AllErrors() []error { return m }

// UpdateApiTokenConfigurationResponseValidationError is the validation error
// returned by UpdateApiTokenConfigurationResponse.Validate if the designated
// constraints aren't met.
type UpdateApiTokenConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateApiTokenConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateApiTokenConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateApiTokenConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateApiTokenConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateApiTokenConfigurationResponseValidationError) ErrorName() string {
	return "UpdateApiTokenConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateApiTokenConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateApiTokenConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateApiTokenConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateApiTokenConfigurationResponseValidationError{}
