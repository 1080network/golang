// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: partner/widget/v1/widget.proto

package instrumentv1

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

// Validate checks the field values on InitializeWidgetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InitializeWidgetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InitializeWidgetRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InitializeWidgetRequestMultiError, or nil if none found.
func (m *InitializeWidgetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InitializeWidgetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserDemographic()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InitializeWidgetRequestValidationError{
					field:  "UserDemographic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InitializeWidgetRequestValidationError{
					field:  "UserDemographic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserDemographic()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InitializeWidgetRequestValidationError{
				field:  "UserDemographic",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InitializeWidgetRequestMultiError(errors)
	}

	return nil
}

// InitializeWidgetRequestMultiError is an error wrapping multiple validation
// errors returned by InitializeWidgetRequest.ValidateAll() if the designated
// constraints aren't met.
type InitializeWidgetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InitializeWidgetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InitializeWidgetRequestMultiError) AllErrors() []error { return m }

// InitializeWidgetRequestValidationError is the validation error returned by
// InitializeWidgetRequest.Validate if the designated constraints aren't met.
type InitializeWidgetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InitializeWidgetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InitializeWidgetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InitializeWidgetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InitializeWidgetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InitializeWidgetRequestValidationError) ErrorName() string {
	return "InitializeWidgetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InitializeWidgetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInitializeWidgetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InitializeWidgetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InitializeWidgetRequestValidationError{}

// Validate checks the field values on InitializeWidgetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InitializeWidgetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InitializeWidgetResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InitializeWidgetResponseMultiError, or nil if none found.
func (m *InitializeWidgetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *InitializeWidgetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InitializeWidgetResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InitializeWidgetResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InitializeWidgetResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SessionKey

	if len(errors) > 0 {
		return InitializeWidgetResponseMultiError(errors)
	}

	return nil
}

// InitializeWidgetResponseMultiError is an error wrapping multiple validation
// errors returned by InitializeWidgetResponse.ValidateAll() if the designated
// constraints aren't met.
type InitializeWidgetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InitializeWidgetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InitializeWidgetResponseMultiError) AllErrors() []error { return m }

// InitializeWidgetResponseValidationError is the validation error returned by
// InitializeWidgetResponse.Validate if the designated constraints aren't met.
type InitializeWidgetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InitializeWidgetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InitializeWidgetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InitializeWidgetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InitializeWidgetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InitializeWidgetResponseValidationError) ErrorName() string {
	return "InitializeWidgetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InitializeWidgetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInitializeWidgetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InitializeWidgetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InitializeWidgetResponseValidationError{}
