// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.3
// - protoc             v3.9.1
// source: lorawan-stack/api/user.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the User message to JSON.
func (x *User) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.CreatedAt)
		}
	}
	if x.UpdatedAt != nil || s.HasField("updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		if x.UpdatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.UpdatedAt)
		}
	}
	if x.DeletedAt != nil || s.HasField("deleted_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted_at")
		if x.DeletedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.DeletedAt)
		}
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if x.Description != "" || s.HasField("description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description")
		s.WriteString(x.Description)
	}
	if x.Attributes != nil || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.PrimaryEmailAddress != "" || s.HasField("primary_email_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("primary_email_address")
		s.WriteString(x.PrimaryEmailAddress)
	}
	if x.PrimaryEmailAddressValidatedAt != nil || s.HasField("primary_email_address_validated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("primary_email_address_validated_at")
		if x.PrimaryEmailAddressValidatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.PrimaryEmailAddressValidatedAt)
		}
	}
	if x.Password != "" || s.HasField("password") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("password")
		s.WriteString(x.Password)
	}
	if x.PasswordUpdatedAt != nil || s.HasField("password_updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("password_updated_at")
		if x.PasswordUpdatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.PasswordUpdatedAt)
		}
	}
	if x.RequirePasswordUpdate || s.HasField("require_password_update") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("require_password_update")
		s.WriteBool(x.RequirePasswordUpdate)
	}
	if x.State != 0 || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		x.State.MarshalProtoJSON(s)
	}
	if x.StateDescription != "" || s.HasField("state_description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state_description")
		s.WriteString(x.StateDescription)
	}
	if x.Admin || s.HasField("admin") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("admin")
		s.WriteBool(x.Admin)
	}
	if x.TemporaryPassword != "" || s.HasField("temporary_password") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password")
		s.WriteString(x.TemporaryPassword)
	}
	if x.TemporaryPasswordCreatedAt != nil || s.HasField("temporary_password_created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password_created_at")
		if x.TemporaryPasswordCreatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.TemporaryPasswordCreatedAt)
		}
	}
	if x.TemporaryPasswordExpiresAt != nil || s.HasField("temporary_password_expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password_expires_at")
		if x.TemporaryPasswordExpiresAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.TemporaryPasswordExpiresAt)
		}
	}
	if x.ProfilePicture != nil || s.HasField("profile_picture") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("profile_picture")
		// NOTE: Picture does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ProfilePicture)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the User to JSON.
func (x *User) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the User message from JSON.
func (x *User) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "created_at", "createdAt":
			s.AddField("created_at")
			if s.ReadNil() {
				x.CreatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			if s.ReadNil() {
				x.UpdatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = v
		case "deleted_at", "deletedAt":
			s.AddField("deleted_at")
			if s.ReadNil() {
				x.DeletedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.DeletedAt = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "description":
			s.AddField("description")
			x.Description = s.ReadString()
		case "attributes":
			s.AddField("attributes")
			if s.ReadNil() {
				x.Attributes = nil
				return
			}
			x.Attributes = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.Attributes[key] = s.ReadString()
			})
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			if s.ReadNil() {
				x.ContactInfo = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "primary_email_address", "primaryEmailAddress":
			s.AddField("primary_email_address")
			x.PrimaryEmailAddress = s.ReadString()
		case "primary_email_address_validated_at", "primaryEmailAddressValidatedAt":
			s.AddField("primary_email_address_validated_at")
			if s.ReadNil() {
				x.PrimaryEmailAddressValidatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.PrimaryEmailAddressValidatedAt = v
		case "password":
			s.AddField("password")
			x.Password = s.ReadString()
		case "password_updated_at", "passwordUpdatedAt":
			s.AddField("password_updated_at")
			if s.ReadNil() {
				x.PasswordUpdatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.PasswordUpdatedAt = v
		case "require_password_update", "requirePasswordUpdate":
			s.AddField("require_password_update")
			x.RequirePasswordUpdate = s.ReadBool()
		case "state":
			s.AddField("state")
			x.State.UnmarshalProtoJSON(s)
		case "state_description", "stateDescription":
			s.AddField("state_description")
			x.StateDescription = s.ReadString()
		case "admin":
			s.AddField("admin")
			x.Admin = s.ReadBool()
		case "temporary_password", "temporaryPassword":
			s.AddField("temporary_password")
			x.TemporaryPassword = s.ReadString()
		case "temporary_password_created_at", "temporaryPasswordCreatedAt":
			s.AddField("temporary_password_created_at")
			if s.ReadNil() {
				x.TemporaryPasswordCreatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.TemporaryPasswordCreatedAt = v
		case "temporary_password_expires_at", "temporaryPasswordExpiresAt":
			s.AddField("temporary_password_expires_at")
			if s.ReadNil() {
				x.TemporaryPasswordExpiresAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.TemporaryPasswordExpiresAt = v
		case "profile_picture", "profilePicture":
			s.AddField("profile_picture")
			if s.ReadNil() {
				x.ProfilePicture = nil
				return
			}
			// NOTE: Picture does not seem to implement UnmarshalProtoJSON.
			var v Picture
			gogo.UnmarshalMessage(s, &v)
			x.ProfilePicture = &v
		}
	})
}

// UnmarshalJSON unmarshals the User from JSON.
func (x *User) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Users message to JSON.
func (x *Users) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Users) > 0 || s.HasField("users") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("users")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Users {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("users"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Users to JSON.
func (x *Users) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Users message from JSON.
func (x *Users) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "users":
			s.AddField("users")
			if s.ReadNil() {
				x.Users = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Users = append(x.Users, nil)
					return
				}
				v := &User{}
				v.UnmarshalProtoJSON(s.WithField("users", false))
				if s.Err() != nil {
					return
				}
				x.Users = append(x.Users, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Users from JSON.
func (x *Users) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateUserRequest message to JSON.
func (x *CreateUserRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.User != nil || s.HasField("user") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user")
		x.User.MarshalProtoJSON(s.WithField("user"))
	}
	if x.InvitationToken != "" || s.HasField("invitation_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("invitation_token")
		s.WriteString(x.InvitationToken)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateUserRequest to JSON.
func (x *CreateUserRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateUserRequest message from JSON.
func (x *CreateUserRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user":
			if s.ReadNil() {
				x.User = nil
				return
			}
			x.User = &User{}
			x.User.UnmarshalProtoJSON(s.WithField("user", true))
		case "invitation_token", "invitationToken":
			s.AddField("invitation_token")
			x.InvitationToken = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the CreateUserRequest from JSON.
func (x *CreateUserRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateUserRequest message to JSON.
func (x *UpdateUserRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.User != nil || s.HasField("user") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user")
		x.User.MarshalProtoJSON(s.WithField("user"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the UpdateUserRequest to JSON.
func (x *UpdateUserRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateUserRequest message from JSON.
func (x *UpdateUserRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user":
			if s.ReadNil() {
				x.User = nil
				return
			}
			x.User = &User{}
			x.User.UnmarshalProtoJSON(s.WithField("user", true))
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// UnmarshalJSON unmarshals the UpdateUserRequest from JSON.
func (x *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateUserAPIKeyRequest message to JSON.
func (x *CreateUserAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.UserIds != nil || s.HasField("user_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.UserIds)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateUserAPIKeyRequest to JSON.
func (x *CreateUserAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateUserAPIKeyRequest message from JSON.
func (x *CreateUserAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			if s.ReadNil() {
				x.UserIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIds = &v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			if s.ReadNil() {
				x.ExpiresAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// UnmarshalJSON unmarshals the CreateUserAPIKeyRequest from JSON.
func (x *CreateUserAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateUserAPIKeyRequest message to JSON.
func (x *UpdateUserAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.UserIds != nil || s.HasField("user_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.UserIds)
	}
	if x.ApiKey != nil || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the UpdateUserAPIKeyRequest to JSON.
func (x *UpdateUserAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateUserAPIKeyRequest message from JSON.
func (x *UpdateUserAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			if s.ReadNil() {
				x.UserIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIds = &v
		case "api_key", "apiKey":
			if s.ReadNil() {
				x.ApiKey = nil
				return
			}
			x.ApiKey = &APIKey{}
			x.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// UnmarshalJSON unmarshals the UpdateUserAPIKeyRequest from JSON.
func (x *UpdateUserAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
