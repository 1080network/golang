// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/organization.proto

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

	currencyv1 "github.com/1080network/golang/serviceprovider/proto/common/enums/currencyv1"

	organizationcategoryv1 "github.com/1080network/golang/serviceprovider/proto/common/enums/organizationcategoryv1"
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

	_ = organizationcategoryv1.OrganizationCategory(0)
)

// Validate checks the field values on Organization with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Organization) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Organization with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrganizationMultiError, or
// nil if none found.
func (m *Organization) ValidateAll() error {
	return m.validate(true)
}

func (m *Organization) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetOrganizationKey()); l < 30 || l > 50 {
		err := OrganizationValidationError{
			field:  "OrganizationKey",
			reason: "value length must be between 30 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetCreated()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrganizationValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrganizationValidationError{
					field:  "Created",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrganizationValidationError{
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
				errors = append(errors, OrganizationValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrganizationValidationError{
					field:  "Updated",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrganizationValidationError{
				field:  "Updated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetPartnerKey()); l < 30 || l > 50 {
		err := OrganizationValidationError{
			field:  "PartnerKey",
			reason: "value length must be between 30 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrganizationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrganizationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrganizationValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if !_Organization_DomesticAchRoutingNumber_Pattern.MatchString(m.GetDomesticAchRoutingNumber()) {
		err := OrganizationValidationError{
			field:  "DomesticAchRoutingNumber",
			reason: "value does not match regex pattern \"^|\\\\w{9}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Organization_InternationalAchRoutingNumber_Pattern.MatchString(m.GetInternationalAchRoutingNumber()) {
		err := OrganizationValidationError{
			field:  "InternationalAchRoutingNumber",
			reason: "value does not match regex pattern \"^|\\\\w{9}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Organization_WireRoutingNumber_Pattern.MatchString(m.GetWireRoutingNumber()) {
		err := OrganizationValidationError{
			field:  "WireRoutingNumber",
			reason: "value does not match regex pattern \"^|\\\\w{9}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Organization_SwiftRoutingNumber_Pattern.MatchString(m.GetSwiftRoutingNumber()) {
		err := OrganizationValidationError{
			field:  "SwiftRoutingNumber",
			reason: "value does not match regex pattern \"^|\\\\w{9}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Organization_BankAccountNumber_Pattern.MatchString(m.GetBankAccountNumber()) {
		err := OrganizationValidationError{
			field:  "BankAccountNumber",
			reason: "value does not match regex pattern \"^|\\\\w{14}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrganizationMultiError(errors)
	}

	return nil
}

// OrganizationMultiError is an error wrapping multiple validation errors
// returned by Organization.ValidateAll() if the designated constraints aren't met.
type OrganizationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationMultiError) AllErrors() []error { return m }

// OrganizationValidationError is the validation error returned by
// Organization.Validate if the designated constraints aren't met.
type OrganizationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationValidationError) ErrorName() string { return "OrganizationValidationError" }

// Error satisfies the builtin error interface
func (e OrganizationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganization.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationValidationError{}

var _Organization_DomesticAchRoutingNumber_Pattern = regexp.MustCompile("^|\\w{9}$")

var _Organization_InternationalAchRoutingNumber_Pattern = regexp.MustCompile("^|\\w{9}$")

var _Organization_WireRoutingNumber_Pattern = regexp.MustCompile("^|\\w{9}$")

var _Organization_SwiftRoutingNumber_Pattern = regexp.MustCompile("^|\\w{9}$")

var _Organization_BankAccountNumber_Pattern = regexp.MustCompile("^|\\w{14}$")

// Validate checks the field values on LegacyConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LegacyConfiguration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LegacyConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LegacyConfigurationMultiError, or nil if none found.
func (m *LegacyConfiguration) ValidateAll() error {
	return m.validate(true)
}

func (m *LegacyConfiguration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInterchangePercentage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "InterchangePercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "InterchangePercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInterchangePercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "InterchangePercentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDisputeRatePercentage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeRatePercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeRatePercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisputeRatePercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "DisputeRatePercentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Currency

	if all {
		switch v := interface{}(m.GetDisputeFee()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisputeFee()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "DisputeFee",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDisputeManagementCost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeManagementCost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "DisputeManagementCost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisputeManagementCost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "DisputeManagementCost",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNetworkMembershipFee()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkMembershipFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkMembershipFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNetworkMembershipFee()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "NetworkMembershipFee",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNetworkTransactionFee()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkTransactionFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkTransactionFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNetworkTransactionFee()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "NetworkTransactionFee",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNetworkReportingFee()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkReportingFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LegacyConfigurationValidationError{
					field:  "NetworkReportingFee",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNetworkReportingFee()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LegacyConfigurationValidationError{
				field:  "NetworkReportingFee",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LegacyConfigurationMultiError(errors)
	}

	return nil
}

// LegacyConfigurationMultiError is an error wrapping multiple validation
// errors returned by LegacyConfiguration.ValidateAll() if the designated
// constraints aren't met.
type LegacyConfigurationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LegacyConfigurationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LegacyConfigurationMultiError) AllErrors() []error { return m }

// LegacyConfigurationValidationError is the validation error returned by
// LegacyConfiguration.Validate if the designated constraints aren't met.
type LegacyConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LegacyConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LegacyConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LegacyConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LegacyConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LegacyConfigurationValidationError) ErrorName() string {
	return "LegacyConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e LegacyConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLegacyConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LegacyConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LegacyConfigurationValidationError{}
