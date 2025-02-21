// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lorawan

import (
	"bytes"
	"math"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/byteutil"
)

// AppendMHDR appends encoded msg to dst.
func AppendMHDR(dst []byte, msg *ttnpb.MHDR) ([]byte, error) {
	if msg.GetMType() > 7 {
		return nil, errExpectedLowerOrEqual("MType", 7)(msg.GetMType())
	}
	if msg.GetMajor() > 4 {
		return nil, errExpectedLowerOrEqual("Major", 4)(msg.GetMajor())
	}
	return append(dst, byte(msg.GetMType())<<5|byte(msg.GetMajor())), nil
}

// MarshalMHDR returns encoded msg.
func MarshalMHDR(msg *ttnpb.MHDR) ([]byte, error) {
	return AppendMHDR(make([]byte, 0, 1), msg)
}

// UnmarshalMHDR unmarshals b into msg.
func UnmarshalMHDR(b []byte, msg *ttnpb.MHDR) error {
	if len(b) != 1 {
		return errExpectedLengthEncodedEqual("MHDR", 1)(len(b))
	}
	v := b[0]
	msg.MType = ttnpb.MType(v >> 5)
	msg.Major = ttnpb.Major(v & 3)
	return nil
}

// AppendFCtrl appends encoded msg to dst.
func AppendFCtrl(dst []byte, msg *ttnpb.FCtrl, isUplink bool, fOptsLen uint8) ([]byte, error) {
	if fOptsLen > 15 {
		return nil, errExpectedLowerOrEqual("FOptsLen", 15)(fOptsLen)
	}
	b := fOptsLen
	if msg.GetAdr() {
		b |= 1 << 7
	}
	if msg.GetAck() {
		b |= 1 << 5
	}
	if isUplink {
		if msg.GetAdrAckReq() {
			b |= 1 << 6
		}
		if msg.GetClassB() {
			b |= 1 << 4
		}
	} else {
		if msg.GetFPending() {
			b |= 1 << 4
		}
	}
	return append(dst, b), nil
}

// UnmarshalFCtrl unmarshals b into msg.
func UnmarshalFCtrl(b []byte, msg *ttnpb.FCtrl, isUplink bool) error {
	if len(b) != 1 {
		return errExpectedLengthEncodedEqual("FCtrl", 1)(len(b))
	}
	v := b[0]
	msg.Adr = v&(1<<7) > 0
	msg.Ack = v&(1<<5) > 0
	if isUplink {
		msg.AdrAckReq = v&(1<<6) > 0
		msg.ClassB = v&(1<<4) > 0
	} else {
		msg.FPending = v&(1<<4) > 0
	}
	return nil
}

// AppendFHDR appends encoded msg to dst.
func AppendFHDR(dst []byte, msg *ttnpb.FHDR, isUplink bool) ([]byte, error) {
	if n := len(msg.DevAddr); n != 4 {
		return nil, errExpectedLengthEqual("DevAddr", 4)(n)
	}
	dst = appendReverse(dst, msg.DevAddr...)
	fOptsLen := uint8(len(msg.GetFOpts()))
	if fOptsLen > 15 {
		return nil, errExpectedLowerOrEqual("FOptsLen", 15)(fOptsLen)
	}
	dst, err := AppendFCtrl(dst, msg.GetFCtrl(), isUplink, fOptsLen)
	if err != nil {
		return nil, errFailedEncoding("FCtrl").WithCause(err)
	}
	dst = byteutil.AppendUint32(dst, msg.GetFCnt()&0xffff, 2)
	dst = append(dst, msg.GetFOpts()...)
	return dst, nil
}

// UnmarshalFHDR unmarshals b into msg.
func UnmarshalFHDR(b []byte, msg *ttnpb.FHDR, isUplink bool) error {
	n := len(b)
	if n < 7 || n > 23 {
		return errExpectedLengthEncodedBound("FHDR", 7, 23)(n)
	}
	msg.DevAddr = make([]byte, 4)
	copyReverse(msg.DevAddr, b[0:4])
	msg.FCtrl = &ttnpb.FCtrl{}
	if err := UnmarshalFCtrl(b[4:5], msg.FCtrl, isUplink); err != nil {
		return errFailedDecoding("FCtrl").WithCause(err)
	}
	msg.FCnt = byteutil.ParseUint32(b[5:7])
	if n == 7 {
		// No FOpts
		return nil
	}
	msg.FOpts = make([]byte, 0, n-7)
	msg.FOpts = append(msg.FOpts, b[7:n]...)
	return nil
}

// AppendMACPayload appends encoded msg to dst.
func AppendMACPayload(dst []byte, msg *ttnpb.MACPayload, isUplink bool) ([]byte, error) {
	dst, err := AppendFHDR(dst, msg.FHdr, isUplink)
	if err != nil {
		return nil, errFailedEncoding("FHDR").WithCause(err)
	}
	if msg.FPort > math.MaxUint8 {
		return nil, errExpectedLowerOrEqual("FPort", math.MaxUint8)(msg.FPort)
	}
	if len(msg.FrmPayload) > 0 || msg.FPort != 0 {
		dst = append(dst, byte(msg.FPort))
	}
	dst = append(dst, msg.FrmPayload...)
	return dst, nil
}

// MarshalMACPayload returns encoded msg.
func MarshalMACPayload(msg *ttnpb.MACPayload, isUplink bool) ([]byte, error) {
	return AppendMACPayload(make([]byte, 0, 1), msg, isUplink)
}

// UnmarshalMACPayload unmarshals b into msg.
func UnmarshalMACPayload(b []byte, msg *ttnpb.MACPayload, isUplink bool) error {
	n := uint8(len(b))
	if n < 7 {
		return errExpectedLengthHigherOrEqual("FHDR", 7)(n)
	}
	fOptsLen := b[4] & 0xf
	fhdrLen := fOptsLen + 7
	if n < fhdrLen {
		return errExpectedLengthHigherOrEqual("MACPayload", fhdrLen)(n)
	}
	msg.FHdr = &ttnpb.FHDR{}
	if err := UnmarshalFHDR(b[0:fhdrLen], msg.FHdr, isUplink); err != nil {
		return errFailedDecoding("FHDR").WithCause(err)
	}

	fPortIdx := fhdrLen
	if n > fPortIdx {
		msg.FPort = uint32(b[fPortIdx])

		frmPayloadIdx := fPortIdx + 1
		if n > frmPayloadIdx {
			msg.FrmPayload = make([]byte, 0, n-frmPayloadIdx)
			msg.FrmPayload = append(msg.FrmPayload, b[frmPayloadIdx:]...)
		}
	}
	return nil
}

// AppendDLSettings appends encoded msg to dst.
func AppendDLSettings(dst []byte, msg *ttnpb.DLSettings) ([]byte, error) {
	if msg.GetRx1DrOffset() > 7 {
		return nil, errExpectedLowerOrEqual("Rx1DROffset", 7)(msg.GetRx1DrOffset())
	}
	if msg.GetRx2Dr() > 15 {
		return nil, errExpectedLowerOrEqual("Rx2DR", 15)(msg.GetRx2Dr())
	}
	b := byte(msg.GetRx2Dr())
	b |= byte(msg.GetRx1DrOffset() << 4)
	if msg.GetOptNeg() {
		b |= (1 << 7)
	}
	return append(dst, b), nil
}

// MarshalDLSettings returns encoded msg.
func MarshalDLSettings(msg *ttnpb.DLSettings) ([]byte, error) {
	return AppendDLSettings(make([]byte, 0, 1), msg)
}

// UnmarshalDLSettings unmarshals b into msg.
func UnmarshalDLSettings(b []byte, msg *ttnpb.DLSettings) error {
	if len(b) != 1 {
		return errExpectedLengthEncodedEqual("DLSettings", 1)(len(b))
	}
	v := uint32(b[0])
	msg.OptNeg = (v >> 7) != 0
	msg.Rx1DrOffset = ttnpb.DataRateOffset((v >> 4) & 0x7)
	msg.Rx2Dr = ttnpb.DataRateIndex(v & 0xf)
	return nil
}

var errMaxCFListFrequency = unexpectedValue(
	errors.DefineInvalidArgument("max_cflist_frequency", "expected CFList frequency to be less or equal to 0xFFFFFF", valueKey),
)

// AppendCFList appends encoded msg to dst.
func AppendCFList(dst []byte, msg *ttnpb.CFList) ([]byte, error) {
	switch msg.Type {
	case 0:
		n := len(msg.Freq)
		if n > 5 {
			return nil, errExpectedLengthLowerOrEqual("CFListFreq", 5)(n)
		}
		for _, freq := range msg.Freq {
			if freq > byteutil.MaxUint24 {
				return nil, errMaxCFListFrequency(freq)
			}
			dst = byteutil.AppendUint32(dst, freq, 3)
		}
		// Fill remaining space with zeros.
		dst = append(dst, bytes.Repeat([]byte{0x0}, 15-n*3)...)
	case 1:
		n := uint(len(msg.ChMasks))
		if n > 96 {
			return nil, errExpectedLengthLowerOrEqual("CFListChMasks", 96)(n)
		}
		for i := uint(0); i < n; i += 8 {
			var b byte
			for j := uint(0); j < 8 && i+j < n; j++ {
				if msg.ChMasks[i+j] {
					b |= (1 << j)
				}
			}
			dst = append(dst, b)
		}
		// Fill remaining space with zeros.
		dst = append(dst, bytes.Repeat([]byte{0x0}, int(15-(n+7)/8))...)
	}
	dst = append(dst, byte(msg.Type))
	return dst, nil
}

// MarshalCFList returns encoded msg.
func MarshalCFList(msg *ttnpb.CFList) ([]byte, error) {
	return AppendCFList(make([]byte, 0, 16), msg)
}

// UnmarshalCFList unmarshals b into msg.
func UnmarshalCFList(b []byte, msg *ttnpb.CFList) error {
	n := len(b)
	if n != 16 {
		return errExpectedLengthEncodedEqual("CFList", 16)(n)
	}
	msg.Type = ttnpb.CFListType(b[15])
	switch msg.Type {
	case 0:
		msg.Freq = make([]uint32, 0, 5)
		for i := 0; i < 15; i += 3 {
			freq := byteutil.ParseUint32(b[i : i+3])
			if freq != 0 {
				msg.Freq = append(msg.Freq, freq)
			}
		}
	case 1:
		msg.ChMasks = make([]bool, 0, 96)
		for _, m := range b[:12] {
			msg.ChMasks = append(msg.ChMasks,
				m&1 > 0,
				m&(1<<1) > 0,
				m&(1<<2) > 0,
				m&(1<<3) > 0,
				m&(1<<4) > 0,
				m&(1<<5) > 0,
				m&(1<<6) > 0,
				m&(1<<7) > 0,
			)
		}
	default:
		return errUnknown("CFListType")(msg.Type.String())
	}
	return nil
}

// AppendJoinAcceptPayload appends encoded msg to dst.
func AppendJoinAcceptPayload(dst []byte, msg *ttnpb.JoinAcceptPayload) ([]byte, error) {
	if n := len(msg.JoinNonce); n != 3 {
		return nil, errExpectedLengthEqual("JoinNonce", 3)(n)
	}
	dst = appendReverse(dst, msg.JoinNonce...)
	if n := len(msg.NetId); n != 3 {
		return nil, errExpectedLengthEqual("NetID", 3)(n)
	}
	dst = appendReverse(dst, msg.NetId...)
	if n := len(msg.DevAddr); n != 4 {
		return nil, errExpectedLengthEqual("DevAddr", 4)(n)
	}
	dst = appendReverse(dst, msg.DevAddr...)
	dst, err := AppendDLSettings(dst, msg.DlSettings)
	if err != nil {
		return nil, errFailedEncoding("DLSettings").WithCause(err)
	}
	if msg.RxDelay > math.MaxUint8 {
		return nil, errExpectedLowerOrEqual("RxDelay", math.MaxUint8)(msg.RxDelay)
	}
	dst = append(dst, byte(msg.RxDelay))
	if msg.GetCfList() != nil {
		dst, err = AppendCFList(dst, msg.CfList)
		if err != nil {
			return nil, errFailedEncoding("CFList").WithCause(err)
		}
	}
	return dst, nil
}

// MarshalJoinAcceptPayload returns encoded msg.
func MarshalJoinAcceptPayload(msg *ttnpb.JoinAcceptPayload) ([]byte, error) {
	if msg.GetCfList() != nil {
		return AppendJoinAcceptPayload(make([]byte, 0, 28), msg)
	}
	return AppendJoinAcceptPayload(make([]byte, 0, 12), msg)
}

// UnmarshalJoinAcceptPayload unmarshals b into msg.
func UnmarshalJoinAcceptPayload(b []byte, msg *ttnpb.JoinAcceptPayload) error {
	n := len(b)
	if n != 12 && n != 28 {
		return errExpectedLengthEncodedTwoChoices("JoinAcceptPayload", 12, 28)(n)
	}
	msg.JoinNonce, msg.NetId, msg.DevAddr = make([]byte, 3), make([]byte, 3), make([]byte, 4)
	copyReverse(msg.JoinNonce, b[0:3])
	copyReverse(msg.NetId, b[3:6])
	copyReverse(msg.DevAddr, b[6:10])
	msg.DlSettings = &ttnpb.DLSettings{}
	if err := UnmarshalDLSettings(b[10:11], msg.DlSettings); err != nil {
		return errFailedDecoding("DLSettings").WithCause(err)
	}
	msg.RxDelay = ttnpb.RxDelay(b[11])

	if n == 12 {
		return nil
	}
	msg.CfList = &ttnpb.CFList{}
	if err := UnmarshalCFList(b[12:], msg.CfList); err != nil {
		return errFailedDecoding("CFList").WithCause(err)
	}
	return nil
}

// AppendJoinRequestPayload appends encoded msg to dst.
func AppendJoinRequestPayload(dst []byte, msg *ttnpb.JoinRequestPayload) ([]byte, error) {
	if n := len(msg.JoinEui); n != 8 {
		return nil, errExpectedLengthEqual("JoinEUI", 8)(n)
	}
	dst = appendReverse(dst, msg.JoinEui...)
	if n := len(msg.DevEui); n != 8 {
		return nil, errExpectedLengthEqual("DevEUI", 8)(n)
	}
	dst = appendReverse(dst, msg.DevEui...)
	if n := len(msg.DevNonce); n != 2 {
		return nil, errExpectedLengthEqual("DevNonce", 2)(n)
	}
	dst = appendReverse(dst, msg.DevNonce...)
	return dst, nil
}

// MarshalJoinRequestPayload returns encoded msg.
func MarshalJoinRequestPayload(msg *ttnpb.JoinRequestPayload) ([]byte, error) {
	return AppendJoinRequestPayload(make([]byte, 0, 18), msg)
}

// UnmarshalJoinRequestPayload unmarshals b into msg.
func UnmarshalJoinRequestPayload(b []byte, msg *ttnpb.JoinRequestPayload) error {
	if len(b) != 18 {
		return errExpectedLengthEncodedEqual("JoinRequestPayload", 18)(len(b))
	}
	msg.JoinEui, msg.DevEui, msg.DevNonce = make([]byte, 8), make([]byte, 8), make([]byte, 2)
	copyReverse(msg.JoinEui, b[0:8])
	copyReverse(msg.DevEui, b[8:16])
	copyReverse(msg.DevNonce, b[16:18])
	return nil
}

// AppendRejoinRequestPayload appends encoded msg to dst.
func AppendRejoinRequestPayload(dst []byte, msg *ttnpb.RejoinRequestPayload) ([]byte, error) {
	dst = append(dst, byte(msg.RejoinType))
	switch msg.RejoinType {
	case 0, 2:
		if n := len(msg.NetId); n != 3 {
			return nil, errExpectedLengthEqual("NetID", 3)(n)
		}
		dst = appendReverse(dst, msg.NetId...)
		if n := len(msg.DevEui); n != 8 {
			return nil, errExpectedLengthEqual("DevEUI", 8)(n)
		}
		dst = appendReverse(dst, msg.DevEui...)
	case 1:
		if n := len(msg.JoinEui); n != 8 {
			return nil, errExpectedLengthEqual("JoinEUI", 8)(n)
		}
		dst = appendReverse(dst, msg.JoinEui...)
		if n := len(msg.DevEui); n != 8 {
			return nil, errExpectedLengthEqual("DevEUI", 8)(n)
		}
		dst = appendReverse(dst, msg.DevEui...)
	default:
		return nil, errUnknown("RejoinType")(msg.RejoinType)
	}
	if msg.RejoinCnt > math.MaxUint16 {
		return nil, errExpectedLowerOrEqual("RJcount1", math.MaxUint16)(msg.RejoinCnt)
	}
	dst = byteutil.AppendUint32(dst, msg.RejoinCnt, 2)
	return dst, nil
}

// MarshalRejoinRequestPayload returns encoded msg.
func MarshalRejoinRequestPayload(msg *ttnpb.RejoinRequestPayload) ([]byte, error) {
	if msg.RejoinType == 1 {
		return AppendRejoinRequestPayload(make([]byte, 0, 19), msg)
	}
	return AppendRejoinRequestPayload(make([]byte, 0, 14), msg)
}

var errEncryptedJoinAcceptPayloadLength = unexpectedValue(
	errors.DefineInvalidArgument("encrypted_joinacceptpayload_length", "encrypted JoinAcceptPayload should have length of 16 or 32", valueKey),
)

// UnmarshalRejoinRequestPayload unmarshals b into msg.
func UnmarshalRejoinRequestPayload(b []byte, msg *ttnpb.RejoinRequestPayload) error {
	msg.RejoinType = ttnpb.RejoinRequestType(b[0])
	switch msg.RejoinType {
	case 0, 2:
		if len(b) != 14 {
			return errExpectedLengthEqual("RejoinRequestPayload", 14)(len(b))
		}
		msg.NetId, msg.DevEui = make([]byte, 3), make([]byte, 8)
		copyReverse(msg.NetId, b[1:4])
		copyReverse(msg.DevEui, b[4:12])
		msg.RejoinCnt = byteutil.ParseUint32(b[12:14])
	case 1:
		if len(b) != 19 {
			return errExpectedLengthEqual("RejoinRequestPayload", 19)(len(b))
		}
		msg.JoinEui, msg.DevEui = make([]byte, 8), make([]byte, 8)
		copyReverse(msg.JoinEui, b[1:9])
		copyReverse(msg.DevEui, b[9:17])
		msg.RejoinCnt = byteutil.ParseUint32(b[17:19])
	default:
		return errUnknown("RejoinType")(msg.RejoinType.String())
	}
	return nil
}

// AppendMessage appends encoded msg to dst.
func AppendMessage(dst []byte, msg *ttnpb.Message) ([]byte, error) {
	if msg.MHdr == nil {
		return nil, errMissing("MHDR")
	}
	dst, err := AppendMHDR(dst, msg.MHdr)
	if err != nil {
		return nil, errFailedEncoding("MHDR").WithCause(err)
	}
	switch msg.MHdr.MType {
	case ttnpb.MType_CONFIRMED_DOWN, ttnpb.MType_UNCONFIRMED_DOWN:
		pld := msg.GetMacPayload()
		if pld == nil {
			return nil, errMissing("MACPayload")
		}
		dst, err = AppendMACPayload(dst, pld, false)
		if err != nil {
			return nil, errFailedEncoding("downlink MACPayload").WithCause(err)
		}
	case ttnpb.MType_CONFIRMED_UP, ttnpb.MType_UNCONFIRMED_UP:
		pld := msg.GetMacPayload()
		if pld == nil {
			return nil, errMissing("MACPayload")
		}
		dst, err = AppendMACPayload(dst, pld, true)
		if err != nil {
			return nil, errFailedEncoding("uplink MACPayload").WithCause(err)
		}
	case ttnpb.MType_JOIN_REQUEST:
		pld := msg.GetJoinRequestPayload()
		if pld == nil {
			return nil, errMissing("JoinRequestPayload")
		}
		dst, err = AppendJoinRequestPayload(dst, pld)
		if err != nil {
			return nil, errFailedEncoding("JoinRequestPayload").WithCause(err)
		}
	case ttnpb.MType_REJOIN_REQUEST:
		pld := msg.GetRejoinRequestPayload()
		if pld == nil {
			return nil, errMissing("RejoinRequestPayload")
		}
		dst, err = AppendRejoinRequestPayload(dst, pld)
		if err != nil {
			return nil, errFailedEncoding("RejoinRequestPayload").WithCause(err)
		}
	case ttnpb.MType_JOIN_ACCEPT:
		pld := msg.GetJoinAcceptPayload()
		if pld == nil {
			return nil, errMissing("JoinAcceptPayload")
		}
		n := len(pld.Encrypted)
		if n != 16 && n != 32 {
			return nil, errEncryptedJoinAcceptPayloadLength(n)
		}
		dst = append(dst, pld.Encrypted...)
	default:
		return nil, errUnknown("MType")(msg.MHdr.MType.String())
	}
	dst = append(dst, msg.Mic...)
	return dst, nil
}

// MarshalMessage returns encoded msg.
func MarshalMessage(msg *ttnpb.Message) ([]byte, error) {
	if msg.MHdr == nil {
		return nil, errMissing("MHDR")
	}
	switch msg.MHdr.MType {
	case ttnpb.MType_CONFIRMED_DOWN, ttnpb.MType_UNCONFIRMED_DOWN, ttnpb.MType_CONFIRMED_UP, ttnpb.MType_UNCONFIRMED_UP:
		// MHDR(1) + Payload(up to 250) + MIC(4)
		return AppendMessage(make([]byte, 0, 255), msg)
	case ttnpb.MType_JOIN_REQUEST:
		// MHDR(1) + Payload(18) + MIC(4)
		return AppendMessage(make([]byte, 0, 23), msg)
	case ttnpb.MType_REJOIN_REQUEST:
		// MHDR(1) + Payload(14|19) + MIC(4)
		return AppendMessage(make([]byte, 0, 24), msg)
	case ttnpb.MType_JOIN_ACCEPT:
		// MHDR(1) + Encrypted payload(16|32)
		return AppendMessage(make([]byte, 0, 33), msg)
	default:
		return nil, errUnknown("MType")(msg.MHdr.MType.String())
	}
}

// UnmarshalMessage unmarshals b into msg.
func UnmarshalMessage(b []byte, msg *ttnpb.Message) error {
	n := len(b)
	if n == 0 {
		return errMissing("PHYPayload")
	}
	msg.MHdr = &ttnpb.MHDR{}
	if err := UnmarshalMHDR(b[0:1], msg.MHdr); err != nil {
		return errFailedDecoding("MHDR").WithCause(err)
	}
	switch msg.MHdr.MType {
	case ttnpb.MType_CONFIRMED_DOWN, ttnpb.MType_UNCONFIRMED_DOWN:
		if n < 12 {
			return errExpectedLengthHigherOrEqual("FHDR", 7)(n - 5)
		}
		pld := &ttnpb.MACPayload{}
		if err := UnmarshalMACPayload(b[1:n-4], pld, false); err != nil {
			return errFailedDecoding("MACPayload").WithCause(err)
		}
		msg.Payload = &ttnpb.Message_MacPayload{MacPayload: pld}
		msg.Mic = b[n-4:]
	case ttnpb.MType_CONFIRMED_UP, ttnpb.MType_UNCONFIRMED_UP:
		if n < 12 {
			return errExpectedLengthHigherOrEqual("FHDR", 7)(n - 5)
		}
		pld := &ttnpb.MACPayload{}
		if err := UnmarshalMACPayload(b[1:n-4], pld, true); err != nil {
			return errFailedDecoding("MACPayload").WithCause(err)
		}
		msg.Payload = &ttnpb.Message_MacPayload{MacPayload: pld}
		msg.Mic = b[n-4:]
	case ttnpb.MType_JOIN_REQUEST:
		if n != 23 {
			return errExpectedLengthEqual("JoinRequestPHYPayload", 23)(n)
		}
		pld := &ttnpb.JoinRequestPayload{}
		if err := UnmarshalJoinRequestPayload(b[1:19], pld); err != nil {
			return errFailedDecoding("JoinRequestPayload").WithCause(err)
		}
		msg.Payload = &ttnpb.Message_JoinRequestPayload{JoinRequestPayload: pld}
		msg.Mic = b[19:]
	case ttnpb.MType_REJOIN_REQUEST:
		if n < 2 {
			return errExpectedLengthTwoChoices("RejoinRequestPHYPayload", 19, 24)(n)
		}
		var micIdx int
		if b[1] == 1 {
			// MHDR(1) + Payload(19)
			micIdx = 20
		} else {
			// MHDR(1) + Payload(14)
			micIdx = 15
		}
		if n != micIdx+4 {
			return errExpectedLengthTwoChoices("RejoinRequestPHYPayload", 19, 24)(n)
		}
		pld := &ttnpb.RejoinRequestPayload{}
		if err := UnmarshalRejoinRequestPayload(b[1:micIdx], pld); err != nil {
			return errFailedDecoding("RejoinRequestPayload").WithCause(err)
		}
		msg.Payload = &ttnpb.Message_RejoinRequestPayload{RejoinRequestPayload: pld}
		msg.Mic = b[micIdx:]
	case ttnpb.MType_JOIN_ACCEPT:
		if n != 17 && n != 33 {
			return errExpectedLengthTwoChoices("JoinAcceptPHYPayload", 17, 33)(n)
		}
		msg.Payload = &ttnpb.Message_JoinAcceptPayload{JoinAcceptPayload: &ttnpb.JoinAcceptPayload{Encrypted: b[1:]}}
	default:
		return errUnknown("MType")(msg.MHdr.MType.String())
	}
	return nil
}

// GetUplinkMessageIdentifiers parses the PHYPayload and retrieves the EndDeviceIdentifiers (except DeviceID).
func GetUplinkMessageIdentifiers(phyPayload []byte) (*ttnpb.EndDeviceIdentifiers, error) {
	n := len(phyPayload)
	if n == 0 {
		return nil, errMissing("PHYPayload")
	}
	var mhdr ttnpb.MHDR
	if err := UnmarshalMHDR(phyPayload[0:1], &mhdr); err != nil {
		return nil, errFailedDecoding("MHDR").WithCause(err)
	}
	switch mhdr.MType {
	case ttnpb.MType_CONFIRMED_UP, ttnpb.MType_UNCONFIRMED_UP:
		if n < 12 {
			return nil, errExpectedLengthHigherOrEqual("FHDR", 7)(n - 5)
		}
		var devAddr types.DevAddr
		copyReverse(devAddr[:], phyPayload[1:5])
		return &ttnpb.EndDeviceIdentifiers{DevAddr: devAddr.Bytes()}, nil
	case ttnpb.MType_JOIN_REQUEST:
		if n != 23 {
			return nil, errExpectedLengthEqual("JoinRequestPHYPayload", 23)(n)
		}
		var joinEUI, devEUI types.EUI64
		copyReverse(joinEUI[:], phyPayload[1:9])
		copyReverse(devEUI[:], phyPayload[9:17])
		return &ttnpb.EndDeviceIdentifiers{DevEui: devEUI.Bytes(), JoinEui: joinEUI.Bytes()}, nil
	case ttnpb.MType_REJOIN_REQUEST:
		if n != 19 && n != 24 {
			return nil, errExpectedLengthTwoChoices("RejoinRequestPHYPayload", 19, 24)(n)
		}
		switch phyPayload[1] {
		case 0, 2:
			if n != 19 {
				return nil, errExpectedLengthEqual("RejoinRequestPHYPayload", 19)(n)
			}
			var devEUI types.EUI64
			copyReverse(devEUI[:], phyPayload[5:13])
			return &ttnpb.EndDeviceIdentifiers{DevEui: devEUI.Bytes()}, nil
		case 1:
			if n != 24 {
				return nil, errExpectedLengthEqual("RejoinRequestPHYPayload", 24)(n)
			}
			var joinEUI, devEUI types.EUI64
			copyReverse(joinEUI[:], phyPayload[2:10])
			copyReverse(devEUI[:], phyPayload[10:18])
			return &ttnpb.EndDeviceIdentifiers{DevEui: devEUI.Bytes(), JoinEui: joinEUI.Bytes()}, nil
		default:
			return nil, errUnknown("RejoinType")(phyPayload[1])
		}
	default:
		return nil, errUnknown("MType")(mhdr.MType.String())
	}
}
