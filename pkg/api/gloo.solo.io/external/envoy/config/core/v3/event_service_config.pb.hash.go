// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/event_service_config.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *EventServiceConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3.EventServiceConfig")); err != nil {
		return 0, err
	}

	switch m.ConfigSourceSpecifier.(type) {

	case *EventServiceConfig_GrpcService:

		if h, ok := interface{}(m.GetGrpcService()).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetGrpcService(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}