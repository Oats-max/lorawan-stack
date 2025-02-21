// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _rights_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on Rights with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Rights) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = RightsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "rights":

			for idx, item := range m.GetRights() {
				_, _ = idx, item

				if _, ok := Right_name[int32(item)]; !ok {
					return RightsValidationError{
						field:  fmt.Sprintf("rights[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		default:
			return RightsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// RightsValidationError is the validation error returned by
// Rights.ValidateFields if the designated constraints aren't met.
type RightsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RightsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RightsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RightsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RightsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RightsValidationError) ErrorName() string { return "RightsValidationError" }

// Error satisfies the builtin error interface
func (e RightsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRights.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RightsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RightsValidationError{}

// ValidateFields checks the field values on APIKey with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *APIKey) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = APIKeyFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id":
			// no validation rules for Id
		case "key":
			// no validation rules for Key
		case "name":

			if utf8.RuneCountInString(m.GetName()) > 50 {
				return APIKeyValidationError{
					field:  "name",
					reason: "value length must be at most 50 runes",
				}
			}

		case "rights":

			for idx, item := range m.GetRights() {
				_, _ = idx, item

				if _, ok := Right_name[int32(item)]; !ok {
					return APIKeyValidationError{
						field:  fmt.Sprintf("rights[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		case "created_at":

			if v, ok := interface{}(m.GetCreatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return APIKeyValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "updated_at":

			if v, ok := interface{}(m.GetUpdatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return APIKeyValidationError{
						field:  "updated_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "expires_at":

			if t := m.GetExpiresAt(); t != nil {
				ts, err := types.TimestampFromProto(t)
				if err != nil {
					return APIKeyValidationError{
						field:  "expires_at",
						reason: "value is not a valid timestamp",
						cause:  err,
					}
				}

				now := time.Now()

				if ts.Sub(now) <= 0 {
					return APIKeyValidationError{
						field:  "expires_at",
						reason: "value must be greater than now",
					}
				}

			}

		default:
			return APIKeyValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// APIKeyValidationError is the validation error returned by
// APIKey.ValidateFields if the designated constraints aren't met.
type APIKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e APIKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e APIKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e APIKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e APIKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e APIKeyValidationError) ErrorName() string { return "APIKeyValidationError" }

// Error satisfies the builtin error interface
func (e APIKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAPIKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = APIKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = APIKeyValidationError{}

// ValidateFields checks the field values on APIKeys with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *APIKeys) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = APIKeysFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "api_keys":

			for idx, item := range m.GetApiKeys() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return APIKeysValidationError{
							field:  fmt.Sprintf("api_keys[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return APIKeysValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// APIKeysValidationError is the validation error returned by
// APIKeys.ValidateFields if the designated constraints aren't met.
type APIKeysValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e APIKeysValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e APIKeysValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e APIKeysValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e APIKeysValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e APIKeysValidationError) ErrorName() string { return "APIKeysValidationError" }

// Error satisfies the builtin error interface
func (e APIKeysValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAPIKeys.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = APIKeysValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = APIKeysValidationError{}

// ValidateFields checks the field values on Collaborator with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Collaborator) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CollaboratorFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if m.GetIds() == nil {
				return CollaboratorValidationError{
					field:  "ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CollaboratorValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rights":

			for idx, item := range m.GetRights() {
				_, _ = idx, item

				if _, ok := Right_name[int32(item)]; !ok {
					return CollaboratorValidationError{
						field:  fmt.Sprintf("rights[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		default:
			return CollaboratorValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CollaboratorValidationError is the validation error returned by
// Collaborator.ValidateFields if the designated constraints aren't met.
type CollaboratorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollaboratorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollaboratorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollaboratorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollaboratorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollaboratorValidationError) ErrorName() string { return "CollaboratorValidationError" }

// Error satisfies the builtin error interface
func (e CollaboratorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollaborator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollaboratorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollaboratorValidationError{}

// ValidateFields checks the field values on GetCollaboratorResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCollaboratorResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetCollaboratorResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(m.GetIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetCollaboratorResponseValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rights":

		default:
			return GetCollaboratorResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetCollaboratorResponseValidationError is the validation error returned by
// GetCollaboratorResponse.ValidateFields if the designated constraints aren't met.
type GetCollaboratorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCollaboratorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCollaboratorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCollaboratorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCollaboratorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCollaboratorResponseValidationError) ErrorName() string {
	return "GetCollaboratorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCollaboratorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCollaboratorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCollaboratorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCollaboratorResponseValidationError{}

// ValidateFields checks the field values on Collaborators with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Collaborators) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CollaboratorsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "collaborators":

			for idx, item := range m.GetCollaborators() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return CollaboratorsValidationError{
							field:  fmt.Sprintf("collaborators[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return CollaboratorsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CollaboratorsValidationError is the validation error returned by
// Collaborators.ValidateFields if the designated constraints aren't met.
type CollaboratorsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollaboratorsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollaboratorsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollaboratorsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollaboratorsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollaboratorsValidationError) ErrorName() string { return "CollaboratorsValidationError" }

// Error satisfies the builtin error interface
func (e CollaboratorsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollaborators.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollaboratorsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollaboratorsValidationError{}
