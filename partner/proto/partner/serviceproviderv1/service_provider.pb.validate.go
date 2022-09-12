// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: partner/serviceprovider/v1/service_provider.proto

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

// Validate checks the field values on SearchServiceProviderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchServiceProviderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchServiceProviderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchServiceProviderRequestMultiError, or nil if none found.
func (m *SearchServiceProviderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchServiceProviderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SearchServiceProviderRequestMultiError(errors)
	}

	return nil
}

// SearchServiceProviderRequestMultiError is an error wrapping multiple
// validation errors returned by SearchServiceProviderRequest.ValidateAll() if
// the designated constraints aren't met.
type SearchServiceProviderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchServiceProviderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchServiceProviderRequestMultiError) AllErrors() []error { return m }

// SearchServiceProviderRequestValidationError is the validation error returned
// by SearchServiceProviderRequest.Validate if the designated constraints
// aren't met.
type SearchServiceProviderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchServiceProviderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchServiceProviderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchServiceProviderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchServiceProviderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchServiceProviderRequestValidationError) ErrorName() string {
	return "SearchServiceProviderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchServiceProviderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchServiceProviderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchServiceProviderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchServiceProviderRequestValidationError{}

// Validate checks the field values on SearchServiceProviderResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchServiceProviderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchServiceProviderResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// SearchServiceProviderResponseMultiError, or nil if none found.
func (m *SearchServiceProviderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchServiceProviderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SearchServiceProviderResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SearchServiceProviderResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SearchServiceProviderResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetServiceProviders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchServiceProviderResponseValidationError{
						field:  fmt.Sprintf("ServiceProviders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchServiceProviderResponseValidationError{
						field:  fmt.Sprintf("ServiceProviders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchServiceProviderResponseValidationError{
					field:  fmt.Sprintf("ServiceProviders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SearchServiceProviderResponseMultiError(errors)
	}

	return nil
}

// SearchServiceProviderResponseMultiError is an error wrapping multiple
// validation errors returned by SearchServiceProviderResponse.ValidateAll()
// if the designated constraints aren't met.
type SearchServiceProviderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchServiceProviderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchServiceProviderResponseMultiError) AllErrors() []error { return m }

// SearchServiceProviderResponseValidationError is the validation error
// returned by SearchServiceProviderResponse.Validate if the designated
// constraints aren't met.
type SearchServiceProviderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchServiceProviderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchServiceProviderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchServiceProviderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchServiceProviderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchServiceProviderResponseValidationError) ErrorName() string {
	return "SearchServiceProviderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SearchServiceProviderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchServiceProviderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchServiceProviderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchServiceProviderResponseValidationError{}
