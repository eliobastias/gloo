// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/kube/pod.proto

package kube

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
func (m *Pod) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Pod)
	if !ok {
		that2, ok := that.(Pod)
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

	if len(m.GetExtraLabels()) != len(target.GetExtraLabels()) {
		return false
	}
	for k, v := range m.GetExtraLabels() {

		if strings.Compare(v, target.GetExtraLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetExtraAnnotations()) != len(target.GetExtraAnnotations()) {
		return false
	}
	for k, v := range m.GetExtraAnnotations() {

		if strings.Compare(v, target.GetExtraAnnotations()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if len(m.GetImagePullSecrets()) != len(target.GetImagePullSecrets()) {
		return false
	}
	for idx, v := range m.GetImagePullSecrets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetImagePullSecrets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetImagePullSecrets()[idx]) {
				return false
			}
		}

	}

	if len(m.GetNodeSelector()) != len(target.GetNodeSelector()) {
		return false
	}
	for k, v := range m.GetNodeSelector() {

		if strings.Compare(v, target.GetNodeSelector()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetAffinity()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAffinity()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAffinity(), target.GetAffinity()) {
			return false
		}
	}

	if len(m.GetTolerations()) != len(target.GetTolerations()) {
		return false
	}
	for idx, v := range m.GetTolerations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTolerations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTolerations()[idx]) {
				return false
			}
		}

	}

	return true
}
