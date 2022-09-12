// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/payment_token.proto

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

	currencyv1 "mica/proto/common/enums/currencyv1"

	paymenttokentypev1 "mica/proto/common/enums/paymenttokentypev1"
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

	_ = currencyv1.Currency(0)

	_ = paymenttokentypev1.PaymentTokenType(0)
)

// Validate checks the field values on CommonPaymentToken with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonPaymentToken) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonPaymentToken with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonPaymentTokenMultiError, or nil if none found.
func (m *CommonPaymentToken) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonPaymentToken) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PaymentTokenType

	// no validation rules for PaymentToken

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetCreated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonPaymentTokenValidationError{
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
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonPaymentTokenValidationError{
				field:  "Updated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpirationDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "ExpirationDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonPaymentTokenValidationError{
					field:  "ExpirationDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpirationDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonPaymentTokenValidationError{
				field:  "ExpirationDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Currency

	if len(errors) > 0 {
		return CommonPaymentTokenMultiError(errors)
	}

	return nil
}

// CommonPaymentTokenMultiError is an error wrapping multiple validation errors
// returned by CommonPaymentToken.ValidateAll() if the designated constraints
// aren't met.
type CommonPaymentTokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonPaymentTokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonPaymentTokenMultiError) AllErrors() []error { return m }

// CommonPaymentTokenValidationError is the validation error returned by
// CommonPaymentToken.Validate if the designated constraints aren't met.
type CommonPaymentTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonPaymentTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonPaymentTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonPaymentTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonPaymentTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonPaymentTokenValidationError) ErrorName() string {
	return "CommonPaymentTokenValidationError"
}

// Error satisfies the builtin error interface
func (e CommonPaymentTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonPaymentToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonPaymentTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonPaymentTokenValidationError{}
