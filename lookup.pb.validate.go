// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: lookup.proto

package api

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

// Validate checks the field values on Lookup with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Lookup) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Lookup with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LookupMultiError, or nil if none found.
func (m *Lookup) ValidateAll() error {
	return m.validate(true)
}

func (m *Lookup) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupMultiError(errors)
	}

	return nil
}

// LookupMultiError is an error wrapping multiple validation errors returned by
// Lookup.ValidateAll() if the designated constraints aren't met.
type LookupMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupMultiError) AllErrors() []error { return m }

// LookupValidationError is the validation error returned by Lookup.Validate if
// the designated constraints aren't met.
type LookupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupValidationError) ErrorName() string { return "LookupValidationError" }

// Error satisfies the builtin error interface
func (e LookupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookup.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupValidationError{}

// Validate checks the field values on LookupField with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LookupField) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupField with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LookupFieldMultiError, or
// nil if none found.
func (m *LookupField) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupField) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	switch m.Value.(type) {

	case *LookupField_StringValue:

		if all {
			switch v := interface{}(m.GetStringValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "StringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "StringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStringValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "StringValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedStringValue:

		if all {
			switch v := interface{}(m.GetRepeatedStringValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedStringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedStringValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedStringValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedStringValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_IntValue:

		if all {
			switch v := interface{}(m.GetIntValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "IntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "IntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetIntValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "IntValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedIntValue:

		if all {
			switch v := interface{}(m.GetRepeatedIntValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedIntValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedIntValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_UintValue:

		if all {
			switch v := interface{}(m.GetUintValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "UintValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "UintValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUintValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "UintValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedUintValue:

		if all {
			switch v := interface{}(m.GetRepeatedUintValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedUintValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedUintValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedUintValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedUintValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_FloatValue:

		if all {
			switch v := interface{}(m.GetFloatValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "FloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "FloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFloatValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "FloatValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedFloatValue:

		if all {
			switch v := interface{}(m.GetRepeatedFloatValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedFloatValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedFloatValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_BoolValue:

		if all {
			switch v := interface{}(m.GetBoolValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "BoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "BoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBoolValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "BoolValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedBoolValue:

		if all {
			switch v := interface{}(m.GetRepeatedBoolValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedBoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedBoolValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedBoolValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedBoolValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RangeIntValue:

		if all {
			switch v := interface{}(m.GetRangeIntValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RangeIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RangeIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRangeIntValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RangeIntValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedRangeIntValue:

		if all {
			switch v := interface{}(m.GetRepeatedRangeIntValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedRangeIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedRangeIntValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedRangeIntValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedRangeIntValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RangeFloatValue:

		if all {
			switch v := interface{}(m.GetRangeFloatValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RangeFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RangeFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRangeFloatValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RangeFloatValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedRangeFloatValue:

		if all {
			switch v := interface{}(m.GetRepeatedRangeFloatValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedRangeFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedRangeFloatValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedRangeFloatValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedRangeFloatValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_GeoPointValue:

		if all {
			switch v := interface{}(m.GetGeoPointValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "GeoPointValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "GeoPointValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGeoPointValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "GeoPointValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedGeoPointValue:

		if all {
			switch v := interface{}(m.GetRepeatedGeoPointValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedGeoPointValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedGeoPointValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedGeoPointValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedGeoPointValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_GeoRectValue:

		if all {
			switch v := interface{}(m.GetGeoRectValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "GeoRectValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "GeoRectValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGeoRectValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "GeoRectValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LookupField_RepeatedGeoRectValue:

		if all {
			switch v := interface{}(m.GetRepeatedGeoRectValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedGeoRectValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupFieldValidationError{
						field:  "RepeatedGeoRectValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRepeatedGeoRectValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupFieldValidationError{
					field:  "RepeatedGeoRectValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupFieldMultiError(errors)
	}

	return nil
}

// LookupFieldMultiError is an error wrapping multiple validation errors
// returned by LookupField.ValidateAll() if the designated constraints aren't met.
type LookupFieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupFieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupFieldMultiError) AllErrors() []error { return m }

// LookupFieldValidationError is the validation error returned by
// LookupField.Validate if the designated constraints aren't met.
type LookupFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupFieldValidationError) ErrorName() string { return "LookupFieldValidationError" }

// Error satisfies the builtin error interface
func (e LookupFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupFieldValidationError{}

// Validate checks the field values on LookupSegmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSegmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSegmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSegmentRequestMultiError, or nil if none found.
func (m *LookupSegmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSegmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Index

	if all {
		switch v := interface{}(m.GetLookup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LookupSegmentRequestValidationError{
					field:  "Lookup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LookupSegmentRequestValidationError{
					field:  "Lookup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLookup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupSegmentRequestValidationError{
				field:  "Lookup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LookupSegmentRequestMultiError(errors)
	}

	return nil
}

// LookupSegmentRequestMultiError is an error wrapping multiple validation
// errors returned by LookupSegmentRequest.ValidateAll() if the designated
// constraints aren't met.
type LookupSegmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSegmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSegmentRequestMultiError) AllErrors() []error { return m }

// LookupSegmentRequestValidationError is the validation error returned by
// LookupSegmentRequest.Validate if the designated constraints aren't met.
type LookupSegmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSegmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSegmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSegmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSegmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSegmentRequestValidationError) ErrorName() string {
	return "LookupSegmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSegmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSegmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSegmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSegmentRequestValidationError{}

// Validate checks the field values on LookupSegmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSegmentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSegmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSegmentResponseMultiError, or nil if none found.
func (m *LookupSegmentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSegmentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupSegmentResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupSegmentResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupSegmentResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupSegmentResponseMultiError(errors)
	}

	return nil
}

// LookupSegmentResponseMultiError is an error wrapping multiple validation
// errors returned by LookupSegmentResponse.ValidateAll() if the designated
// constraints aren't met.
type LookupSegmentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSegmentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSegmentResponseMultiError) AllErrors() []error { return m }

// LookupSegmentResponseValidationError is the validation error returned by
// LookupSegmentResponse.Validate if the designated constraints aren't met.
type LookupSegmentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSegmentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSegmentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSegmentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSegmentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSegmentResponseValidationError) ErrorName() string {
	return "LookupSegmentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSegmentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSegmentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSegmentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSegmentResponseValidationError{}

// Validate checks the field values on LookupSegmentKeyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSegmentKeyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSegmentKeyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSegmentKeyRequestMultiError, or nil if none found.
func (m *LookupSegmentKeyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSegmentKeyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Index

	if all {
		switch v := interface{}(m.GetLookup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LookupSegmentKeyRequestValidationError{
					field:  "Lookup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LookupSegmentKeyRequestValidationError{
					field:  "Lookup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLookup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupSegmentKeyRequestValidationError{
				field:  "Lookup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LookupSegmentKeyRequestMultiError(errors)
	}

	return nil
}

// LookupSegmentKeyRequestMultiError is an error wrapping multiple validation
// errors returned by LookupSegmentKeyRequest.ValidateAll() if the designated
// constraints aren't met.
type LookupSegmentKeyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSegmentKeyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSegmentKeyRequestMultiError) AllErrors() []error { return m }

// LookupSegmentKeyRequestValidationError is the validation error returned by
// LookupSegmentKeyRequest.Validate if the designated constraints aren't met.
type LookupSegmentKeyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSegmentKeyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSegmentKeyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSegmentKeyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSegmentKeyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSegmentKeyRequestValidationError) ErrorName() string {
	return "LookupSegmentKeyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSegmentKeyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSegmentKeyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSegmentKeyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSegmentKeyRequestValidationError{}

// Validate checks the field values on LookupSegmentKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupSegmentKeyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupSegmentKeyResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupSegmentKeyResponseMultiError, or nil if none found.
func (m *LookupSegmentKeyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupSegmentKeyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupSegmentKeyResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupSegmentKeyResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupSegmentKeyResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupSegmentKeyResponseMultiError(errors)
	}

	return nil
}

// LookupSegmentKeyResponseMultiError is an error wrapping multiple validation
// errors returned by LookupSegmentKeyResponse.ValidateAll() if the designated
// constraints aren't met.
type LookupSegmentKeyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupSegmentKeyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupSegmentKeyResponseMultiError) AllErrors() []error { return m }

// LookupSegmentKeyResponseValidationError is the validation error returned by
// LookupSegmentKeyResponse.Validate if the designated constraints aren't met.
type LookupSegmentKeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupSegmentKeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupSegmentKeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupSegmentKeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupSegmentKeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupSegmentKeyResponseValidationError) ErrorName() string {
	return "LookupSegmentKeyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LookupSegmentKeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupSegmentKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupSegmentKeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupSegmentKeyResponseValidationError{}