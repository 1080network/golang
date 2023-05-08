// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: micashared/common/v1/area.proto

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

// Validate checks the field values on Circle with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Circle) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Circle with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CircleMultiError, or nil if none found.
func (m *Circle) ValidateAll() error {
	return m.validate(true)
}

func (m *Circle) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetLatitude(); val < -90 || val > 90 {
		err := CircleValidationError{
			field:  "Latitude",
			reason: "value must be inside range [-90, 90]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLongitude(); val < -180 || val > 180 {
		err := CircleValidationError{
			field:  "Longitude",
			reason: "value must be inside range [-180, 180]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetRadius(); val < 0 || val > 10000000 {
		err := CircleValidationError{
			field:  "Radius",
			reason: "value must be inside range [0, 10000000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CircleMultiError(errors)
	}

	return nil
}

// CircleMultiError is an error wrapping multiple validation errors returned by
// Circle.ValidateAll() if the designated constraints aren't met.
type CircleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircleMultiError) AllErrors() []error { return m }

// CircleValidationError is the validation error returned by Circle.Validate if
// the designated constraints aren't met.
type CircleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircleValidationError) ErrorName() string { return "CircleValidationError" }

// Error satisfies the builtin error interface
func (e CircleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircleValidationError{}