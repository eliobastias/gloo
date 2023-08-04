// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/filters/stages.proto

package filters

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// List of filter stages which can be selected for a HTTP filter.
type FilterStage_Stage int32

const (
	FilterStage_FaultStage     FilterStage_Stage = 0
	FilterStage_CorsStage      FilterStage_Stage = 1
	FilterStage_WafStage       FilterStage_Stage = 2
	FilterStage_AuthNStage     FilterStage_Stage = 3
	FilterStage_AuthZStage     FilterStage_Stage = 4
	FilterStage_RateLimitStage FilterStage_Stage = 5
	FilterStage_AcceptedStage  FilterStage_Stage = 6
	FilterStage_OutAuthStage   FilterStage_Stage = 7
	FilterStage_RouteStage     FilterStage_Stage = 8
)

// Enum value maps for FilterStage_Stage.
var (
	FilterStage_Stage_name = map[int32]string{
		0: "FaultStage",
		1: "CorsStage",
		2: "WafStage",
		3: "AuthNStage",
		4: "AuthZStage",
		5: "RateLimitStage",
		6: "AcceptedStage",
		7: "OutAuthStage",
		8: "RouteStage",
	}
	FilterStage_Stage_value = map[string]int32{
		"FaultStage":     0,
		"CorsStage":      1,
		"WafStage":       2,
		"AuthNStage":     3,
		"AuthZStage":     4,
		"RateLimitStage": 5,
		"AcceptedStage":  6,
		"OutAuthStage":   7,
		"RouteStage":     8,
	}
)

func (x FilterStage_Stage) Enum() *FilterStage_Stage {
	p := new(FilterStage_Stage)
	*p = x
	return p
}

func (x FilterStage_Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterStage_Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes[0].Descriptor()
}

func (FilterStage_Stage) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes[0]
}

func (x FilterStage_Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterStage_Stage.Descriptor instead.
func (FilterStage_Stage) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescGZIP(), []int{0, 0}
}

// Desired placement of the HTTP filter relative to the stage. The default is `During`.
type FilterStage_Predicate int32

const (
	FilterStage_During FilterStage_Predicate = 0
	FilterStage_Before FilterStage_Predicate = 1
	FilterStage_After  FilterStage_Predicate = 2
)

// Enum value maps for FilterStage_Predicate.
var (
	FilterStage_Predicate_name = map[int32]string{
		0: "During",
		1: "Before",
		2: "After",
	}
	FilterStage_Predicate_value = map[string]int32{
		"During": 0,
		"Before": 1,
		"After":  2,
	}
)

func (x FilterStage_Predicate) Enum() *FilterStage_Predicate {
	p := new(FilterStage_Predicate)
	*p = x
	return p
}

func (x FilterStage_Predicate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterStage_Predicate) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes[1].Descriptor()
}

func (FilterStage_Predicate) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes[1]
}

func (x FilterStage_Predicate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterStage_Predicate.Descriptor instead.
func (FilterStage_Predicate) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescGZIP(), []int{0, 1}
}

// FilterStage allows configuration of where in a filter chain a given HTTP filter is inserted,
// relative to one of the pre-defined stages.
type FilterStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stage of the filter chain in which the selected filter should be added.
	Stage FilterStage_Stage `protobuf:"varint,1,opt,name=stage,proto3,enum=filters.gloo.solo.io.FilterStage_Stage" json:"stage,omitempty"`
	// How this filter should be placed relative to the stage.
	Predicate FilterStage_Predicate `protobuf:"varint,2,opt,name=predicate,proto3,enum=filters.gloo.solo.io.FilterStage_Predicate" json:"predicate,omitempty"`
}

func (x *FilterStage) Reset() {
	*x = FilterStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterStage) ProtoMessage() {}

func (x *FilterStage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterStage.ProtoReflect.Descriptor instead.
func (*FilterStage) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescGZIP(), []int{0}
}

func (x *FilterStage) GetStage() FilterStage_Stage {
	if x != nil {
		return x.Stage
	}
	return FilterStage_FaultStage
}

func (x *FilterStage) GetPredicate() FilterStage_Predicate {
	if x != nil {
		return x.Predicate
	}
	return FilterStage_During
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x09,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x72, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x57, 0x61, 0x66, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x03, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x10, 0x08, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x10, 0x02, 0x42, 0x46, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_goTypes = []interface{}{
	(FilterStage_Stage)(0),     // 0: filters.gloo.solo.io.FilterStage.Stage
	(FilterStage_Predicate)(0), // 1: filters.gloo.solo.io.FilterStage.Predicate
	(*FilterStage)(nil),        // 2: filters.gloo.solo.io.FilterStage
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_depIdxs = []int32{
	0, // 0: filters.gloo.solo.io.FilterStage.stage:type_name -> filters.gloo.solo.io.FilterStage.Stage
	1, // 1: filters.gloo.solo.io.FilterStage.predicate:type_name -> filters.gloo.solo.io.FilterStage.Predicate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterStage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_filters_stages_proto_depIdxs = nil
}
