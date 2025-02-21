// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var AuthInfoResponseFieldPathsNested = []string{
	"access_method",
	"access_method.api_key",
	"access_method.api_key.api_key",
	"access_method.api_key.api_key.created_at",
	"access_method.api_key.api_key.expires_at",
	"access_method.api_key.api_key.id",
	"access_method.api_key.api_key.key",
	"access_method.api_key.api_key.name",
	"access_method.api_key.api_key.rights",
	"access_method.api_key.api_key.updated_at",
	"access_method.api_key.entity_ids",
	"access_method.api_key.entity_ids.ids",
	"access_method.api_key.entity_ids.ids.application_ids",
	"access_method.api_key.entity_ids.ids.application_ids.application_id",
	"access_method.api_key.entity_ids.ids.client_ids",
	"access_method.api_key.entity_ids.ids.client_ids.client_id",
	"access_method.api_key.entity_ids.ids.device_ids",
	"access_method.api_key.entity_ids.ids.device_ids.application_ids",
	"access_method.api_key.entity_ids.ids.device_ids.application_ids.application_id",
	"access_method.api_key.entity_ids.ids.device_ids.dev_addr",
	"access_method.api_key.entity_ids.ids.device_ids.dev_eui",
	"access_method.api_key.entity_ids.ids.device_ids.device_id",
	"access_method.api_key.entity_ids.ids.device_ids.join_eui",
	"access_method.api_key.entity_ids.ids.gateway_ids",
	"access_method.api_key.entity_ids.ids.gateway_ids.eui",
	"access_method.api_key.entity_ids.ids.gateway_ids.gateway_id",
	"access_method.api_key.entity_ids.ids.organization_ids",
	"access_method.api_key.entity_ids.ids.organization_ids.organization_id",
	"access_method.api_key.entity_ids.ids.user_ids",
	"access_method.api_key.entity_ids.ids.user_ids.email",
	"access_method.api_key.entity_ids.ids.user_ids.user_id",
	"access_method.oauth_access_token",
	"access_method.oauth_access_token.access_token",
	"access_method.oauth_access_token.client_ids",
	"access_method.oauth_access_token.client_ids.client_id",
	"access_method.oauth_access_token.created_at",
	"access_method.oauth_access_token.expires_at",
	"access_method.oauth_access_token.id",
	"access_method.oauth_access_token.refresh_token",
	"access_method.oauth_access_token.rights",
	"access_method.oauth_access_token.user_ids",
	"access_method.oauth_access_token.user_ids.email",
	"access_method.oauth_access_token.user_ids.user_id",
	"access_method.oauth_access_token.user_session_id",
	"access_method.user_session",
	"access_method.user_session.created_at",
	"access_method.user_session.expires_at",
	"access_method.user_session.session_id",
	"access_method.user_session.session_secret",
	"access_method.user_session.updated_at",
	"access_method.user_session.user_ids",
	"access_method.user_session.user_ids.email",
	"access_method.user_session.user_ids.user_id",
	"is_admin",
	"universal_rights",
	"universal_rights.rights",
}

var AuthInfoResponseFieldPathsTopLevel = []string{
	"access_method",
	"is_admin",
	"universal_rights",
}
var GetIsConfigurationRequestFieldPathsNested []string
var GetIsConfigurationRequestFieldPathsTopLevel []string
var IsConfigurationFieldPathsNested = []string{
	"admin_rights",
	"admin_rights.all",
	"end_device_picture",
	"end_device_picture.disable_upload",
	"profile_picture",
	"profile_picture.disable_upload",
	"profile_picture.use_gravatar",
	"user_login",
	"user_login.disable_credentials_login",
	"user_registration",
	"user_registration.admin_approval",
	"user_registration.admin_approval.required",
	"user_registration.contact_info_validation",
	"user_registration.contact_info_validation.required",
	"user_registration.enabled",
	"user_registration.invitation",
	"user_registration.invitation.required",
	"user_registration.invitation.token_ttl",
	"user_registration.password_requirements",
	"user_registration.password_requirements.max_length",
	"user_registration.password_requirements.min_digits",
	"user_registration.password_requirements.min_length",
	"user_registration.password_requirements.min_special",
	"user_registration.password_requirements.min_uppercase",
	"user_rights",
	"user_rights.create_applications",
	"user_rights.create_clients",
	"user_rights.create_gateways",
	"user_rights.create_organizations",
}

var IsConfigurationFieldPathsTopLevel = []string{
	"admin_rights",
	"end_device_picture",
	"profile_picture",
	"user_login",
	"user_registration",
	"user_rights",
}
var GetIsConfigurationResponseFieldPathsNested = []string{
	"configuration",
	"configuration.admin_rights",
	"configuration.admin_rights.all",
	"configuration.end_device_picture",
	"configuration.end_device_picture.disable_upload",
	"configuration.profile_picture",
	"configuration.profile_picture.disable_upload",
	"configuration.profile_picture.use_gravatar",
	"configuration.user_login",
	"configuration.user_login.disable_credentials_login",
	"configuration.user_registration",
	"configuration.user_registration.admin_approval",
	"configuration.user_registration.admin_approval.required",
	"configuration.user_registration.contact_info_validation",
	"configuration.user_registration.contact_info_validation.required",
	"configuration.user_registration.enabled",
	"configuration.user_registration.invitation",
	"configuration.user_registration.invitation.required",
	"configuration.user_registration.invitation.token_ttl",
	"configuration.user_registration.password_requirements",
	"configuration.user_registration.password_requirements.max_length",
	"configuration.user_registration.password_requirements.min_digits",
	"configuration.user_registration.password_requirements.min_length",
	"configuration.user_registration.password_requirements.min_special",
	"configuration.user_registration.password_requirements.min_uppercase",
	"configuration.user_rights",
	"configuration.user_rights.create_applications",
	"configuration.user_rights.create_clients",
	"configuration.user_rights.create_gateways",
	"configuration.user_rights.create_organizations",
}

var GetIsConfigurationResponseFieldPathsTopLevel = []string{
	"configuration",
}
var AuthInfoResponse_APIKeyAccessFieldPathsNested = []string{
	"api_key",
	"api_key.created_at",
	"api_key.expires_at",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"api_key.updated_at",
	"entity_ids",
	"entity_ids.ids",
	"entity_ids.ids.application_ids",
	"entity_ids.ids.application_ids.application_id",
	"entity_ids.ids.client_ids",
	"entity_ids.ids.client_ids.client_id",
	"entity_ids.ids.device_ids",
	"entity_ids.ids.device_ids.application_ids",
	"entity_ids.ids.device_ids.application_ids.application_id",
	"entity_ids.ids.device_ids.dev_addr",
	"entity_ids.ids.device_ids.dev_eui",
	"entity_ids.ids.device_ids.device_id",
	"entity_ids.ids.device_ids.join_eui",
	"entity_ids.ids.gateway_ids",
	"entity_ids.ids.gateway_ids.eui",
	"entity_ids.ids.gateway_ids.gateway_id",
	"entity_ids.ids.organization_ids",
	"entity_ids.ids.organization_ids.organization_id",
	"entity_ids.ids.user_ids",
	"entity_ids.ids.user_ids.email",
	"entity_ids.ids.user_ids.user_id",
}

var AuthInfoResponse_APIKeyAccessFieldPathsTopLevel = []string{
	"api_key",
	"entity_ids",
}
var IsConfiguration_UserRegistrationFieldPathsNested = []string{
	"admin_approval",
	"admin_approval.required",
	"contact_info_validation",
	"contact_info_validation.required",
	"enabled",
	"invitation",
	"invitation.required",
	"invitation.token_ttl",
	"password_requirements",
	"password_requirements.max_length",
	"password_requirements.min_digits",
	"password_requirements.min_length",
	"password_requirements.min_special",
	"password_requirements.min_uppercase",
}

var IsConfiguration_UserRegistrationFieldPathsTopLevel = []string{
	"admin_approval",
	"contact_info_validation",
	"enabled",
	"invitation",
	"password_requirements",
}
var IsConfiguration_ProfilePictureFieldPathsNested = []string{
	"disable_upload",
	"use_gravatar",
}

var IsConfiguration_ProfilePictureFieldPathsTopLevel = []string{
	"disable_upload",
	"use_gravatar",
}
var IsConfiguration_EndDevicePictureFieldPathsNested = []string{
	"disable_upload",
}

var IsConfiguration_EndDevicePictureFieldPathsTopLevel = []string{
	"disable_upload",
}
var IsConfiguration_UserRightsFieldPathsNested = []string{
	"create_applications",
	"create_clients",
	"create_gateways",
	"create_organizations",
}

var IsConfiguration_UserRightsFieldPathsTopLevel = []string{
	"create_applications",
	"create_clients",
	"create_gateways",
	"create_organizations",
}
var IsConfiguration_UserLoginFieldPathsNested = []string{
	"disable_credentials_login",
}

var IsConfiguration_UserLoginFieldPathsTopLevel = []string{
	"disable_credentials_login",
}
var IsConfiguration_AdminRightsFieldPathsNested = []string{
	"all",
}

var IsConfiguration_AdminRightsFieldPathsTopLevel = []string{
	"all",
}
var IsConfiguration_UserRegistration_InvitationFieldPathsNested = []string{
	"required",
	"token_ttl",
}

var IsConfiguration_UserRegistration_InvitationFieldPathsTopLevel = []string{
	"required",
	"token_ttl",
}
var IsConfiguration_UserRegistration_ContactInfoValidationFieldPathsNested = []string{
	"required",
}

var IsConfiguration_UserRegistration_ContactInfoValidationFieldPathsTopLevel = []string{
	"required",
}
var IsConfiguration_UserRegistration_AdminApprovalFieldPathsNested = []string{
	"required",
}

var IsConfiguration_UserRegistration_AdminApprovalFieldPathsTopLevel = []string{
	"required",
}
var IsConfiguration_UserRegistration_PasswordRequirementsFieldPathsNested = []string{
	"max_length",
	"min_digits",
	"min_length",
	"min_special",
	"min_uppercase",
}

var IsConfiguration_UserRegistration_PasswordRequirementsFieldPathsTopLevel = []string{
	"max_length",
	"min_digits",
	"min_length",
	"min_special",
	"min_uppercase",
}
