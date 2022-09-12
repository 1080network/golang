// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/address.proto

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

	countryv1 "mica/proto/common/enums/countryv1"

	regionv1 "mica/proto/common/enums/regionv1"
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

	_ = countryv1.Country(0)

	_ = regionv1.Region(0)
)

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AddressMultiError, or nil if none found.
func (m *Address) ValidateAll() error {
	return m.validate(true)
}

func (m *Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Locality

	// no validation rules for Region

	// no validation rules for PostalCode

	// no validation rules for Country

	if len(errors) > 0 {
		return AddressMultiError(errors)
	}

	return nil
}

// AddressMultiError is an error wrapping multiple validation errors returned
// by Address.ValidateAll() if the designated constraints aren't met.
type AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressMultiError) AllErrors() []error { return m }

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}
