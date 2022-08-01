// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

package ttnpb

import (
	"context"

	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (p MACPayload) ValidateContext(context.Context) error {
	if h := p.GetFHdr(); h == nil || types.MustDevAddr(h.DevAddr).OrZero().IsZero() {
		return errMissing("DevAddr")
	}
	return p.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (p JoinRequestPayload) ValidateContext(context.Context) error {
	if devEUI := types.MustEUI64(p.DevEui).OrZero(); devEUI.IsZero() {
		return errMissing("DevEUI")
	}
	return p.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GatewayUplinkMessage) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *DownlinkQueueRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
