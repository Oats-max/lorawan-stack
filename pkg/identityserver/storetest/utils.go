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

package storetest

import (
	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func fieldMask(paths ...string) *pbtypes.FieldMask {
	return &pbtypes.FieldMask{Paths: ttnpb.ExcludeFields(
		paths,
		"ids",
		"created_at",
		"updated_at",
		"deleted_at",
	)}
}

func fieldMaskWith(mask *pbtypes.FieldMask, paths ...string) *pbtypes.FieldMask {
	return &pbtypes.FieldMask{Paths: append(mask.Paths, paths...)}
}
