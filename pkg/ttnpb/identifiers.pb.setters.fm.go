// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ApplicationIdentifiers) SetFields(src *ApplicationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_id":
			if len(subs) > 0 {
				return fmt.Errorf("'application_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationId = src.ApplicationId
			} else {
				var zero string
				dst.ApplicationId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClientIdentifiers) SetFields(src *ClientIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client_id":
			if len(subs) > 0 {
				return fmt.Errorf("'client_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ClientId = src.ClientId
			} else {
				var zero string
				dst.ClientId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceIdentifiers) SetFields(src *EndDeviceIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeviceId = src.DeviceId
			} else {
				var zero string
				dst.DeviceId = zero
			}
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEui = src.DevEui
			} else {
				dst.DevEui = nil
			}
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}
		case "dev_addr":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddr = src.DevAddr
			} else {
				dst.DevAddr = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GatewayIdentifiers) SetFields(src *GatewayIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_id":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayId = src.GatewayId
			} else {
				var zero string
				dst.GatewayId = zero
			}
		case "eui":
			if len(subs) > 0 {
				return fmt.Errorf("'eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Eui = src.Eui
			} else {
				dst.Eui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationIdentifiers) SetFields(src *OrganizationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_id":
			if len(subs) > 0 {
				return fmt.Errorf("'organization_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrganizationId = src.OrganizationId
			} else {
				var zero string
				dst.OrganizationId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserIdentifiers) SetFields(src *UserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_id":
			if len(subs) > 0 {
				return fmt.Errorf("'user_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UserId = src.UserId
			} else {
				var zero string
				dst.UserId = zero
			}
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationOrUserIdentifiers) SetFields(src *OrganizationOrUserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "organization_ids":
					_, srcTypeOk := src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds).OrganizationIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds).OrganizationIds
						} else if srcTypeOk {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_OrganizationIds{OrganizationIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcTypeOk := src.Ids.(*OrganizationOrUserIdentifiers_UserIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*OrganizationOrUserIdentifiers_UserIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_UserIds).UserIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_UserIds).UserIds
						} else if srcTypeOk {
							newDst = &UserIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_UserIds{UserIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EntityIdentifiers) SetFields(src *EntityIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "application_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_ApplicationIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_ApplicationIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_ApplicationIds).ApplicationIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_ApplicationIds).ApplicationIds
						} else if srcTypeOk {
							newDst = &ApplicationIdentifiers{}
							dst.Ids = &EntityIdentifiers_ApplicationIds{ApplicationIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "client_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_ClientIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_ClientIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ClientIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_ClientIds).ClientIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_ClientIds).ClientIds
						} else if srcTypeOk {
							newDst = &ClientIdentifiers{}
							dst.Ids = &EntityIdentifiers_ClientIds{ClientIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "device_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_DeviceIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_DeviceIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *EndDeviceIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_DeviceIds).DeviceIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_DeviceIds).DeviceIds
						} else if srcTypeOk {
							newDst = &EndDeviceIdentifiers{}
							dst.Ids = &EntityIdentifiers_DeviceIds{DeviceIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "gateway_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_GatewayIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_GatewayIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *GatewayIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_GatewayIds).GatewayIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_GatewayIds).GatewayIds
						} else if srcTypeOk {
							newDst = &GatewayIdentifiers{}
							dst.Ids = &EntityIdentifiers_GatewayIds{GatewayIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "organization_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_OrganizationIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_OrganizationIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_OrganizationIds).OrganizationIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_OrganizationIds).OrganizationIds
						} else if srcTypeOk {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &EntityIdentifiers_OrganizationIds{OrganizationIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcTypeOk := src.Ids.(*EntityIdentifiers_UserIds)
					srcValid := srcTypeOk || src.Ids == nil || len(oneofSubs) == 0
					if !srcValid {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Ids.(*EntityIdentifiers_UserIds)
					dstValid := dstTypeOk || dst.Ids == nil || len(oneofSubs) == 0
					if !dstValid {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if srcTypeOk {
							newSrc = src.Ids.(*EntityIdentifiers_UserIds).UserIds
						}
						if dstTypeOk {
							newDst = dst.Ids.(*EntityIdentifiers_UserIds).UserIds
						} else if srcTypeOk {
							newDst = &UserIdentifiers{}
							dst.Ids = &EntityIdentifiers_UserIds{UserIds: newDst}
						} else {
							dst.Ids = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceVersionIdentifiers) SetFields(src *EndDeviceVersionIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandId = src.BrandId
			} else {
				var zero string
				dst.BrandId = zero
			}
		case "model_id":
			if len(subs) > 0 {
				return fmt.Errorf("'model_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ModelId = src.ModelId
			} else {
				var zero string
				dst.ModelId = zero
			}
		case "hardware_version":
			if len(subs) > 0 {
				return fmt.Errorf("'hardware_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HardwareVersion = src.HardwareVersion
			} else {
				var zero string
				dst.HardwareVersion = zero
			}
		case "firmware_version":
			if len(subs) > 0 {
				return fmt.Errorf("'firmware_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FirmwareVersion = src.FirmwareVersion
			} else {
				var zero string
				dst.FirmwareVersion = zero
			}
		case "band_id":
			if len(subs) > 0 {
				return fmt.Errorf("'band_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BandId = src.BandId
			} else {
				var zero string
				dst.BandId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *NetworkIdentifiers) SetFields(src *NetworkIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetId = src.NetId
			} else {
				dst.NetId = nil
			}
		case "tenant_id":
			if len(subs) > 0 {
				return fmt.Errorf("'tenant_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TenantId = src.TenantId
			} else {
				var zero string
				dst.TenantId = zero
			}
		case "cluster_id":
			if len(subs) > 0 {
				return fmt.Errorf("'cluster_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ClusterId = src.ClusterId
			} else {
				var zero string
				dst.ClusterId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
