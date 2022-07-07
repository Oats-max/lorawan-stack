// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.3
// - protoc             v3.9.1
// source: lorawan-stack/api/organization.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the Organization message to JSON.
func (x *Organization) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
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
	if x.AdministrativeContact != nil || s.HasField("administrative_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("administrative_contact")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.AdministrativeContact)
	}
	if x.TechnicalContact != nil || s.HasField("technical_contact") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("technical_contact")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.TechnicalContact)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Organization to JSON.
func (x *Organization) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Organization message from JSON.
func (x *Organization) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
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
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
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
		case "administrative_contact", "administrativeContact":
			s.AddField("administrative_contact")
			if s.ReadNil() {
				x.AdministrativeContact = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.AdministrativeContact = &v
		case "technical_contact", "technicalContact":
			s.AddField("technical_contact")
			if s.ReadNil() {
				x.TechnicalContact = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.TechnicalContact = &v
		}
	})
}

// UnmarshalJSON unmarshals the Organization from JSON.
func (x *Organization) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Organizations message to JSON.
func (x *Organizations) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Organizations) > 0 || s.HasField("organizations") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organizations")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Organizations {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("organizations"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Organizations to JSON.
func (x *Organizations) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Organizations message from JSON.
func (x *Organizations) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organizations":
			s.AddField("organizations")
			if s.ReadNil() {
				x.Organizations = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Organizations = append(x.Organizations, nil)
					return
				}
				v := &Organization{}
				v.UnmarshalProtoJSON(s.WithField("organizations", false))
				if s.Err() != nil {
					return
				}
				x.Organizations = append(x.Organizations, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Organizations from JSON.
func (x *Organizations) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateOrganizationRequest message to JSON.
func (x *CreateOrganizationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Organization != nil || s.HasField("organization") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization")
		x.Organization.MarshalProtoJSON(s.WithField("organization"))
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Collaborator)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CreateOrganizationRequest to JSON.
func (x *CreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateOrganizationRequest message from JSON.
func (x *CreateOrganizationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization":
			if s.ReadNil() {
				x.Organization = nil
				return
			}
			x.Organization = &Organization{}
			x.Organization.UnmarshalProtoJSON(s.WithField("organization", true))
		case "collaborator":
			s.AddField("collaborator")
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Collaborator = &v
		}
	})
}

// UnmarshalJSON unmarshals the CreateOrganizationRequest from JSON.
func (x *CreateOrganizationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateOrganizationRequest message to JSON.
func (x *UpdateOrganizationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Organization != nil || s.HasField("organization") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization")
		x.Organization.MarshalProtoJSON(s.WithField("organization"))
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

// MarshalJSON marshals the UpdateOrganizationRequest to JSON.
func (x *UpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateOrganizationRequest message from JSON.
func (x *UpdateOrganizationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization":
			if s.ReadNil() {
				x.Organization = nil
				return
			}
			x.Organization = &Organization{}
			x.Organization.UnmarshalProtoJSON(s.WithField("organization", true))
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

// UnmarshalJSON unmarshals the UpdateOrganizationRequest from JSON.
func (x *UpdateOrganizationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CreateOrganizationAPIKeyRequest message to JSON.
func (x *CreateOrganizationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
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

// MarshalJSON marshals the CreateOrganizationAPIKeyRequest to JSON.
func (x *CreateOrganizationAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CreateOrganizationAPIKeyRequest message from JSON.
func (x *CreateOrganizationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			if s.ReadNil() {
				x.OrganizationIds = nil
				return
			}
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
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

// UnmarshalJSON unmarshals the CreateOrganizationAPIKeyRequest from JSON.
func (x *CreateOrganizationAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the UpdateOrganizationAPIKeyRequest message to JSON.
func (x *UpdateOrganizationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
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

// MarshalJSON marshals the UpdateOrganizationAPIKeyRequest to JSON.
func (x *UpdateOrganizationAPIKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the UpdateOrganizationAPIKeyRequest message from JSON.
func (x *UpdateOrganizationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			if s.ReadNil() {
				x.OrganizationIds = nil
				return
			}
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
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

// UnmarshalJSON unmarshals the UpdateOrganizationAPIKeyRequest from JSON.
func (x *UpdateOrganizationAPIKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SetOrganizationCollaboratorRequest message to JSON.
func (x *SetOrganizationCollaboratorRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		x.Collaborator.MarshalProtoJSON(s.WithField("collaborator"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SetOrganizationCollaboratorRequest to JSON.
func (x *SetOrganizationCollaboratorRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SetOrganizationCollaboratorRequest message from JSON.
func (x *SetOrganizationCollaboratorRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			if s.ReadNil() {
				x.OrganizationIds = nil
				return
			}
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
		case "collaborator":
			if s.ReadNil() {
				x.Collaborator = nil
				return
			}
			x.Collaborator = &Collaborator{}
			x.Collaborator.UnmarshalProtoJSON(s.WithField("collaborator", true))
		}
	})
}

// UnmarshalJSON unmarshals the SetOrganizationCollaboratorRequest from JSON.
func (x *SetOrganizationCollaboratorRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
