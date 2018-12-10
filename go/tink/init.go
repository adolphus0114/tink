// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Package tink defines interfaces for the crypto primitives that Tink supports.
package tink

import (
	"github.com/google/tink/go/internal"
	tinkpb "github.com/google/tink/proto/tink_go_proto"
)

// keysetHandle is used by package insecure and package testkeysethandle (via package internal)
// to create KeysetHandle from cleartext key material.
func keysetHandle(ks *tinkpb.Keyset) *KeysetHandle {
	return &KeysetHandle{ks}
}

func init() {
	internal.KeysetHandle = keysetHandle
}
