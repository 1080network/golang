// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: serviceprovider/serviceprovider/v1/service_provider.proto

package serviceproviderv1

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

	accounttypev1 "mica/proto/common/enums/accounttypev1"

	currencyv1 "mica/proto/common/enums/currencyv1"
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

	_ = accounttypev1.AccountType(0)

	_ = currencyv1.Currency(0)
)

// Validate checks the field values on ServiceProviderAccount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServiceProviderAccount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceProviderAccount with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceProviderAccountMultiError, or nil if none found.
func (m *ServiceProviderAccount) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceProviderAccount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetServiceProviderAccountKey()); l < 30 || l > 50 {
		err := ServiceProviderAccountValidationError{
			field:  "ServiceProviderAccountKey",
			reason: "value length must be between 30 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetVersion() < 0 {
		err := ServiceProviderAccountValidationError{
			field:  "Version",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceProviderAccountValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceProviderAccountValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceProviderAccountValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceProviderAccountValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceProviderAccountValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceProviderAccountValidationError{
				field:  "Updated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AccountType

	// no validation rules for Currency

	if len(errors) > 0 {
		return ServiceProviderAccountMultiError(errors)
	}

	return nil
}

// ServiceProviderAccountMultiError is an error wrapping multiple validation
// errors returned by ServiceProviderAccount.ValidateAll() if the designated
// constraints aren't met.
type ServiceProviderAccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceProviderAccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceProviderAccountMultiError) AllErrors() []error { return m }

// ServiceProviderAccountValidationError is the validation error returned by
// ServiceProviderAccount.Validate if the designated constraints aren't met.
type ServiceProviderAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceProviderAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceProviderAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceProviderAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceProviderAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceProviderAccountValidationError) ErrorName() string {
	return "ServiceProviderAccountValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceProviderAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceProviderAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceProviderAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceProviderAccountValidationError{}

// Validate checks the field values on ServiceProvider with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServiceProvider) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceProvider with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceProviderMultiError, or nil if none found.
func (m *ServiceProvider) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceProvider) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetServiceProviderKey()); l < 30 || l > 50 {
		err := ServiceProviderValidationError{
			field:  "ServiceProviderKey",
			reason: "value length must be between 30 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetVersion() < 0 {
		err := ServiceProviderValidationError{
			field:  "Version",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceProviderValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceProviderValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceProviderValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceProviderValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceProviderValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceProviderValidationError{
				field:  "Updated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for ServiceProviderUrl

	for idx, item := range m.GetServiceProviderAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServiceProviderValidationError{
						field:  fmt.Sprintf("ServiceProviderAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServiceProviderValidationError{
						field:  fmt.Sprintf("ServiceProviderAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceProviderValidationError{
					field:  fmt.Sprintf("ServiceProviderAccounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServiceProviderMultiError(errors)
	}

	return nil
}

// ServiceProviderMultiError is an error wrapping multiple validation errors
// returned by ServiceProvider.ValidateAll() if the designated constraints
// aren't met.
type ServiceProviderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceProviderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceProviderMultiError) AllErrors() []error { return m }

// ServiceProviderValidationError is the validation error returned by
// ServiceProvider.Validate if the designated constraints aren't met.
type ServiceProviderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceProviderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceProviderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceProviderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceProviderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceProviderValidationError) ErrorName() string { return "ServiceProviderValidationError" }

// Error satisfies the builtin error interface
func (e ServiceProviderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceProvider.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceProviderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceProviderValidationError{}

// Validate checks the field values on GetServiceProviderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetServiceProviderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetServiceProviderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetServiceProviderRequestMultiError, or nil if none found.
func (m *GetServiceProviderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetServiceProviderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetServiceProviderRequestMultiError(errors)
	}

	return nil
}

// GetServiceProviderRequestMultiError is an error wrapping multiple validation
// errors returned by GetServiceProviderRequest.ValidateAll() if the
// designated constraints aren't met.
type GetServiceProviderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetServiceProviderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetServiceProviderRequestMultiError) AllErrors() []error { return m }

// GetServiceProviderRequestValidationError is the validation error returned by
// GetServiceProviderRequest.Validate if the designated constraints aren't met.
type GetServiceProviderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetServiceProviderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetServiceProviderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetServiceProviderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetServiceProviderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetServiceProviderRequestValidationError) ErrorName() string {
	return "GetServiceProviderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetServiceProviderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetServiceProviderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetServiceProviderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetServiceProviderRequestValidationError{}

// Validate checks the field values on GetServiceProviderResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetServiceProviderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetServiceProviderResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetServiceProviderResponseMultiError, or nil if none found.
func (m *GetServiceProviderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetServiceProviderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetServiceProviderResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetServiceProviderResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetServiceProviderResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServiceProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetServiceProviderResponseValidationError{
					field:  "ServiceProvider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetServiceProviderResponseValidationError{
					field:  "ServiceProvider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServiceProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetServiceProviderResponseValidationError{
				field:  "ServiceProvider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetServiceProviderResponseMultiError(errors)
	}

	return nil
}

// GetServiceProviderResponseMultiError is an error wrapping multiple
// validation errors returned by GetServiceProviderResponse.ValidateAll() if
// the designated constraints aren't met.
type GetServiceProviderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetServiceProviderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetServiceProviderResponseMultiError) AllErrors() []error { return m }

// GetServiceProviderResponseValidationError is the validation error returned
// by GetServiceProviderResponse.Validate if the designated constraints aren't met.
type GetServiceProviderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetServiceProviderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetServiceProviderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetServiceProviderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetServiceProviderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetServiceProviderResponseValidationError) ErrorName() string {
	return "GetServiceProviderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetServiceProviderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetServiceProviderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetServiceProviderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetServiceProviderResponseValidationError{}
