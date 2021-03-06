// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/gateway_resources.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteTable) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteTable)
	if !ok {
		that2, ok := that.(RouteTable)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListGatewaysRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGatewaysRequest)
	if !ok {
		that2, ok := that.(ListGatewaysRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListGatewaysResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListGatewaysResponse)
	if !ok {
		that2, ok := that.(ListGatewaysResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetGateways()) != len(target.GetGateways()) {
		return false
	}
	for idx, v := range m.GetGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetGateways()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GetGatewayYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetGatewayYamlRequest)
	if !ok {
		that2, ok := that.(GetGatewayYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGatewayRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGatewayRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGatewayRef(), target.GetGatewayRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetGatewayYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetGatewayYamlResponse)
	if !ok {
		that2, ok := that.(GetGatewayYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListVirtualServicesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListVirtualServicesRequest)
	if !ok {
		that2, ok := that.(ListVirtualServicesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListVirtualServicesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListVirtualServicesResponse)
	if !ok {
		that2, ok := that.(ListVirtualServicesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetVirtualServices()) != len(target.GetVirtualServices()) {
		return false
	}
	for idx, v := range m.GetVirtualServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualServices()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GetVirtualServiceYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceYamlRequest)
	if !ok {
		that2, ok := that.(GetVirtualServiceYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetVirtualServiceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualServiceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualServiceRef(), target.GetVirtualServiceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetVirtualServiceYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetVirtualServiceYamlResponse)
	if !ok {
		that2, ok := that.(GetVirtualServiceYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListRouteTablesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListRouteTablesRequest)
	if !ok {
		that2, ok := that.(ListRouteTablesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListRouteTablesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListRouteTablesResponse)
	if !ok {
		that2, ok := that.(ListRouteTablesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetRouteTables()) != len(target.GetRouteTables()) {
		return false
	}
	for idx, v := range m.GetRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRouteTables()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GetRouteTableYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetRouteTableYamlRequest)
	if !ok {
		that2, ok := that.(GetRouteTableYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRouteTableRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteTableRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteTableRef(), target.GetRouteTableRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetRouteTableYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetRouteTableYamlResponse)
	if !ok {
		that2, ok := that.(GetRouteTableYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}
