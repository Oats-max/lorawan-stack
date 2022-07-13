// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.3
// - protoc             v3.9.1
// source: lorawan-stack/api/gatewayserver.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GatewayUp message to JSON.
func (x *GatewayUp) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.UplinkMessages) > 0 || s.HasField("uplink_messages") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("uplink_messages")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.UplinkMessages {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("uplink_messages"))
		}
		s.WriteArrayEnd()
	}
	if x.GatewayStatus != nil || s.HasField("gateway_status") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_status")
		x.GatewayStatus.MarshalProtoJSON(s.WithField("gateway_status"))
	}
	if x.TxAcknowledgment != nil || s.HasField("tx_acknowledgment") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tx_acknowledgment")
		x.TxAcknowledgment.MarshalProtoJSON(s.WithField("tx_acknowledgment"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GatewayUp to JSON.
func (x *GatewayUp) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GatewayUp message from JSON.
func (x *GatewayUp) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "uplink_messages", "uplinkMessages":
			s.AddField("uplink_messages")
			if s.ReadNil() {
				x.UplinkMessages = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.UplinkMessages = append(x.UplinkMessages, nil)
					return
				}
				v := &UplinkMessage{}
				v.UnmarshalProtoJSON(s.WithField("uplink_messages", false))
				if s.Err() != nil {
					return
				}
				x.UplinkMessages = append(x.UplinkMessages, v)
			})
		case "gateway_status", "gatewayStatus":
			if s.ReadNil() {
				x.GatewayStatus = nil
				return
			}
			x.GatewayStatus = &GatewayStatus{}
			x.GatewayStatus.UnmarshalProtoJSON(s.WithField("gateway_status", true))
		case "tx_acknowledgment", "txAcknowledgment":
			if s.ReadNil() {
				x.TxAcknowledgment = nil
				return
			}
			x.TxAcknowledgment = &TxAcknowledgment{}
			x.TxAcknowledgment.UnmarshalProtoJSON(s.WithField("tx_acknowledgment", true))
		}
	})
}

// UnmarshalJSON unmarshals the GatewayUp from JSON.
func (x *GatewayUp) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GatewayDown message to JSON.
func (x *GatewayDown) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.DownlinkMessage != nil || s.HasField("downlink_message") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("downlink_message")
		x.DownlinkMessage.MarshalProtoJSON(s.WithField("downlink_message"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GatewayDown to JSON.
func (x *GatewayDown) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GatewayDown message from JSON.
func (x *GatewayDown) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "downlink_message", "downlinkMessage":
			if s.ReadNil() {
				x.DownlinkMessage = nil
				return
			}
			x.DownlinkMessage = &DownlinkMessage{}
			x.DownlinkMessage.UnmarshalProtoJSON(s.WithField("downlink_message", true))
		}
	})
}

// UnmarshalJSON unmarshals the GatewayDown from JSON.
func (x *GatewayDown) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
