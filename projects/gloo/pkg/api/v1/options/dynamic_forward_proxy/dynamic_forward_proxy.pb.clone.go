// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/dynamic_forward_proxy/dynamic_forward_proxy.proto

package dynamic_forward_proxy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *FilterConfig) Clone() proto.Message {
	var target *FilterConfig
	if m == nil {
		return target
	}
	target = &FilterConfig{}

	if h, ok := interface{}(m.GetDnsCacheConfig()).(clone.Cloner); ok {
		target.DnsCacheConfig = h.Clone().(*DnsCacheConfig)
	} else {
		target.DnsCacheConfig = proto.Clone(m.GetDnsCacheConfig()).(*DnsCacheConfig)
	}

	target.SaveUpstreamAddress = m.GetSaveUpstreamAddress()

	return target
}

// Clone function
func (m *DnsCacheCircuitBreakers) Clone() proto.Message {
	var target *DnsCacheCircuitBreakers
	if m == nil {
		return target
	}
	target = &DnsCacheCircuitBreakers{}

	if h, ok := interface{}(m.GetMaxPendingRequests()).(clone.Cloner); ok {
		target.MaxPendingRequests = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxPendingRequests = proto.Clone(m.GetMaxPendingRequests()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *DnsCacheConfig) Clone() proto.Message {
	var target *DnsCacheConfig
	if m == nil {
		return target
	}
	target = &DnsCacheConfig{}

	target.DnsLookupFamily = m.GetDnsLookupFamily()

	if h, ok := interface{}(m.GetDnsRefreshRate()).(clone.Cloner); ok {
		target.DnsRefreshRate = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsRefreshRate = proto.Clone(m.GetDnsRefreshRate()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetHostTtl()).(clone.Cloner); ok {
		target.HostTtl = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.HostTtl = proto.Clone(m.GetHostTtl()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxHosts()).(clone.Cloner); ok {
		target.MaxHosts = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxHosts = proto.Clone(m.GetMaxHosts()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetDnsFailureRefreshRate()).(clone.Cloner); ok {
		target.DnsFailureRefreshRate = h.Clone().(*RefreshRate)
	} else {
		target.DnsFailureRefreshRate = proto.Clone(m.GetDnsFailureRefreshRate()).(*RefreshRate)
	}

	if h, ok := interface{}(m.GetDnsCacheCircuitBreaker()).(clone.Cloner); ok {
		target.DnsCacheCircuitBreaker = h.Clone().(*DnsCacheCircuitBreakers)
	} else {
		target.DnsCacheCircuitBreaker = proto.Clone(m.GetDnsCacheCircuitBreaker()).(*DnsCacheCircuitBreakers)
	}

	if m.GetPreresolveHostnames() != nil {
		target.PreresolveHostnames = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.SocketAddress, len(m.GetPreresolveHostnames()))
		for idx, v := range m.GetPreresolveHostnames() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.PreresolveHostnames[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.SocketAddress)
			} else {
				target.PreresolveHostnames[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.SocketAddress)
			}

		}
	}

	if h, ok := interface{}(m.GetDnsQueryTimeout()).(clone.Cloner); ok {
		target.DnsQueryTimeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DnsQueryTimeout = proto.Clone(m.GetDnsQueryTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	switch m.DnsCacheType.(type) {

	case *DnsCacheConfig_CaresDns:

		if h, ok := interface{}(m.GetCaresDns()).(clone.Cloner); ok {
			target.DnsCacheType = &DnsCacheConfig_CaresDns{
				CaresDns: h.Clone().(*CaresDnsResolverConfig),
			}
		} else {
			target.DnsCacheType = &DnsCacheConfig_CaresDns{
				CaresDns: proto.Clone(m.GetCaresDns()).(*CaresDnsResolverConfig),
			}
		}

	case *DnsCacheConfig_AppleDns:

		if h, ok := interface{}(m.GetAppleDns()).(clone.Cloner); ok {
			target.DnsCacheType = &DnsCacheConfig_AppleDns{
				AppleDns: h.Clone().(*AppleDnsResolverConfig),
			}
		} else {
			target.DnsCacheType = &DnsCacheConfig_AppleDns{
				AppleDns: proto.Clone(m.GetAppleDns()).(*AppleDnsResolverConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *RefreshRate) Clone() proto.Message {
	var target *RefreshRate
	if m == nil {
		return target
	}
	target = &RefreshRate{}

	if h, ok := interface{}(m.GetBaseInterval()).(clone.Cloner); ok {
		target.BaseInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.BaseInterval = proto.Clone(m.GetBaseInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetMaxInterval()).(clone.Cloner); ok {
		target.MaxInterval = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.MaxInterval = proto.Clone(m.GetMaxInterval()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}

// Clone function
func (m *PerRouteConfig) Clone() proto.Message {
	var target *PerRouteConfig
	if m == nil {
		return target
	}
	target = &PerRouteConfig{}

	switch m.HostRewriteSpecifier.(type) {

	case *PerRouteConfig_HostRewrite:

		target.HostRewriteSpecifier = &PerRouteConfig_HostRewrite{
			HostRewrite: m.GetHostRewrite(),
		}

	case *PerRouteConfig_AutoHostRewriteHeader:

		target.HostRewriteSpecifier = &PerRouteConfig_AutoHostRewriteHeader{
			AutoHostRewriteHeader: m.GetAutoHostRewriteHeader(),
		}

	}

	return target
}

// Clone function
func (m *DnsResolverOptions) Clone() proto.Message {
	var target *DnsResolverOptions
	if m == nil {
		return target
	}
	target = &DnsResolverOptions{}

	target.UseTcpForDnsLookups = m.GetUseTcpForDnsLookups()

	target.NoDefaultSearchDomain = m.GetNoDefaultSearchDomain()

	return target
}

// Clone function
func (m *CaresDnsResolverConfig) Clone() proto.Message {
	var target *CaresDnsResolverConfig
	if m == nil {
		return target
	}
	target = &CaresDnsResolverConfig{}

	if m.GetResolvers() != nil {
		target.Resolvers = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.Address, len(m.GetResolvers()))
		for idx, v := range m.GetResolvers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolvers[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.Address)
			} else {
				target.Resolvers[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.Address)
			}

		}
	}

	if h, ok := interface{}(m.GetDnsResolverOptions()).(clone.Cloner); ok {
		target.DnsResolverOptions = h.Clone().(*DnsResolverOptions)
	} else {
		target.DnsResolverOptions = proto.Clone(m.GetDnsResolverOptions()).(*DnsResolverOptions)
	}

	return target
}

// Clone function
func (m *AppleDnsResolverConfig) Clone() proto.Message {
	var target *AppleDnsResolverConfig
	if m == nil {
		return target
	}
	target = &AppleDnsResolverConfig{}

	return target
}
