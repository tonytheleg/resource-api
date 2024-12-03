// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.3
// source: api/resources/v1/resources_service.proto

package resources

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResourceType:
	//
	//	*CreateResourceRequest_RhelHost
	//	*CreateResourceRequest_K8SCluster
	ResourceType isCreateResourceRequest_ResourceType `protobuf_oneof:"resource_type"`
}

func (x *CreateResourceRequest) Reset() {
	*x = CreateResourceRequest{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRequest) ProtoMessage() {}

func (x *CreateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{0}
}

func (m *CreateResourceRequest) GetResourceType() isCreateResourceRequest_ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (x *CreateResourceRequest) GetRhelHost() *RhelHost {
	if x, ok := x.GetResourceType().(*CreateResourceRequest_RhelHost); ok {
		return x.RhelHost
	}
	return nil
}

func (x *CreateResourceRequest) GetK8SCluster() *K8SCluster {
	if x, ok := x.GetResourceType().(*CreateResourceRequest_K8SCluster); ok {
		return x.K8SCluster
	}
	return nil
}

type isCreateResourceRequest_ResourceType interface {
	isCreateResourceRequest_ResourceType()
}

type CreateResourceRequest_RhelHost struct {
	RhelHost *RhelHost `protobuf:"bytes,1,opt,name=rhel_host,proto3,oneof"`
}

type CreateResourceRequest_K8SCluster struct {
	K8SCluster *K8SCluster `protobuf:"bytes,2,opt,name=k8s_cluster,proto3,oneof"`
}

func (*CreateResourceRequest_RhelHost) isCreateResourceRequest_ResourceType() {}

func (*CreateResourceRequest_K8SCluster) isCreateResourceRequest_ResourceType() {}

type CreateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceResponse) Reset() {
	*x = CreateResourceResponse{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceResponse) ProtoMessage() {}

func (x *CreateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{1}
}

type UpdateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResourceType:
	//
	//	*UpdateResourceRequest_RhelHost
	//	*UpdateResourceRequest_K8SCluster
	ResourceType isUpdateResourceRequest_ResourceType `protobuf_oneof:"resource_type"`
}

func (x *UpdateResourceRequest) Reset() {
	*x = UpdateResourceRequest{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRequest) ProtoMessage() {}

func (x *UpdateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{2}
}

func (m *UpdateResourceRequest) GetResourceType() isUpdateResourceRequest_ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (x *UpdateResourceRequest) GetRhelHost() *RhelHost {
	if x, ok := x.GetResourceType().(*UpdateResourceRequest_RhelHost); ok {
		return x.RhelHost
	}
	return nil
}

func (x *UpdateResourceRequest) GetK8SCluster() *K8SCluster {
	if x, ok := x.GetResourceType().(*UpdateResourceRequest_K8SCluster); ok {
		return x.K8SCluster
	}
	return nil
}

type isUpdateResourceRequest_ResourceType interface {
	isUpdateResourceRequest_ResourceType()
}

type UpdateResourceRequest_RhelHost struct {
	RhelHost *RhelHost `protobuf:"bytes,1,opt,name=rhel_host,proto3,oneof"`
}

type UpdateResourceRequest_K8SCluster struct {
	K8SCluster *K8SCluster `protobuf:"bytes,2,opt,name=k8s_cluster,proto3,oneof"`
}

func (*UpdateResourceRequest_RhelHost) isUpdateResourceRequest_ResourceType() {}

func (*UpdateResourceRequest_K8SCluster) isUpdateResourceRequest_ResourceType() {}

type UpdateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResourceResponse) Reset() {
	*x = UpdateResourceResponse{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceResponse) ProtoMessage() {}

func (x *UpdateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceResponse.ProtoReflect.Descriptor instead.
func (*UpdateResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{3}
}

type DeleteResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResourceType:
	//
	//	*DeleteResourceRequest_RhelHost
	//	*DeleteResourceRequest_K8SCluster
	ResourceType isDeleteResourceRequest_ResourceType `protobuf_oneof:"resource_type"`
}

func (x *DeleteResourceRequest) Reset() {
	*x = DeleteResourceRequest{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRequest) ProtoMessage() {}

func (x *DeleteResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{4}
}

func (m *DeleteResourceRequest) GetResourceType() isDeleteResourceRequest_ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (x *DeleteResourceRequest) GetRhelHost() *RhelHost {
	if x, ok := x.GetResourceType().(*DeleteResourceRequest_RhelHost); ok {
		return x.RhelHost
	}
	return nil
}

func (x *DeleteResourceRequest) GetK8SCluster() *K8SCluster {
	if x, ok := x.GetResourceType().(*DeleteResourceRequest_K8SCluster); ok {
		return x.K8SCluster
	}
	return nil
}

type isDeleteResourceRequest_ResourceType interface {
	isDeleteResourceRequest_ResourceType()
}

type DeleteResourceRequest_RhelHost struct {
	RhelHost *RhelHost `protobuf:"bytes,1,opt,name=rhel_host,proto3,oneof"`
}

type DeleteResourceRequest_K8SCluster struct {
	K8SCluster *K8SCluster `protobuf:"bytes,2,opt,name=k8s_cluster,proto3,oneof"`
}

func (*DeleteResourceRequest_RhelHost) isDeleteResourceRequest_ResourceType() {}

func (*DeleteResourceRequest_K8SCluster) isDeleteResourceRequest_ResourceType() {}

type DeleteResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResourceResponse) Reset() {
	*x = DeleteResourceResponse{}
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceResponse) ProtoMessage() {}

func (x *DeleteResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_v1_resources_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_resources_v1_resources_service_proto_rawDescGZIP(), []int{5}
}

var File_api_resources_v1_resources_service_proto protoreflect.FileDescriptor

var file_api_resources_v1_resources_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09,
	0x72, 0x68, 0x65, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x72, 0x68, 0x65, 0x6c, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x68, 0x65, 0x6c, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x72, 0x68, 0x65, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x3c, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0f, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x18,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x68, 0x65, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x72, 0x68, 0x65, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6b, 0x38,
	0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x88, 0x03, 0x0a, 0x15, 0x4b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x23, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a,
	0x1a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x2a, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x37,
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50,
	0x01, 0x5a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_resources_v1_resources_service_proto_rawDescOnce sync.Once
	file_api_resources_v1_resources_service_proto_rawDescData = file_api_resources_v1_resources_service_proto_rawDesc
)

func file_api_resources_v1_resources_service_proto_rawDescGZIP() []byte {
	file_api_resources_v1_resources_service_proto_rawDescOnce.Do(func() {
		file_api_resources_v1_resources_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_resources_v1_resources_service_proto_rawDescData)
	})
	return file_api_resources_v1_resources_service_proto_rawDescData
}

var file_api_resources_v1_resources_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_resources_v1_resources_service_proto_goTypes = []any{
	(*CreateResourceRequest)(nil),  // 0: resources.v1.CreateResourceRequest
	(*CreateResourceResponse)(nil), // 1: resources.v1.CreateResourceResponse
	(*UpdateResourceRequest)(nil),  // 2: resources.v1.UpdateResourceRequest
	(*UpdateResourceResponse)(nil), // 3: resources.v1.UpdateResourceResponse
	(*DeleteResourceRequest)(nil),  // 4: resources.v1.DeleteResourceRequest
	(*DeleteResourceResponse)(nil), // 5: resources.v1.DeleteResourceResponse
	(*RhelHost)(nil),               // 6: resources.v1.RhelHost
	(*K8SCluster)(nil),             // 7: resources.v1.K8sCluster
}
var file_api_resources_v1_resources_service_proto_depIdxs = []int32{
	6, // 0: resources.v1.CreateResourceRequest.rhel_host:type_name -> resources.v1.RhelHost
	7, // 1: resources.v1.CreateResourceRequest.k8s_cluster:type_name -> resources.v1.K8sCluster
	6, // 2: resources.v1.UpdateResourceRequest.rhel_host:type_name -> resources.v1.RhelHost
	7, // 3: resources.v1.UpdateResourceRequest.k8s_cluster:type_name -> resources.v1.K8sCluster
	6, // 4: resources.v1.DeleteResourceRequest.rhel_host:type_name -> resources.v1.RhelHost
	7, // 5: resources.v1.DeleteResourceRequest.k8s_cluster:type_name -> resources.v1.K8sCluster
	0, // 6: resources.v1.KesselResourceService.CreateResource:input_type -> resources.v1.CreateResourceRequest
	2, // 7: resources.v1.KesselResourceService.UpdateResource:input_type -> resources.v1.UpdateResourceRequest
	4, // 8: resources.v1.KesselResourceService.DeleteResource:input_type -> resources.v1.DeleteResourceRequest
	1, // 9: resources.v1.KesselResourceService.CreateResource:output_type -> resources.v1.CreateResourceResponse
	3, // 10: resources.v1.KesselResourceService.UpdateResource:output_type -> resources.v1.UpdateResourceResponse
	5, // 11: resources.v1.KesselResourceService.DeleteResource:output_type -> resources.v1.DeleteResourceResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_resources_v1_resources_service_proto_init() }
func file_api_resources_v1_resources_service_proto_init() {
	if File_api_resources_v1_resources_service_proto != nil {
		return
	}
	file_api_resources_v1_common_proto_init()
	file_api_resources_v1_resource_data_proto_init()
	file_api_resources_v1_resource_types_proto_init()
	file_api_resources_v1_resources_service_proto_msgTypes[0].OneofWrappers = []any{
		(*CreateResourceRequest_RhelHost)(nil),
		(*CreateResourceRequest_K8SCluster)(nil),
	}
	file_api_resources_v1_resources_service_proto_msgTypes[2].OneofWrappers = []any{
		(*UpdateResourceRequest_RhelHost)(nil),
		(*UpdateResourceRequest_K8SCluster)(nil),
	}
	file_api_resources_v1_resources_service_proto_msgTypes[4].OneofWrappers = []any{
		(*DeleteResourceRequest_RhelHost)(nil),
		(*DeleteResourceRequest_K8SCluster)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_resources_v1_resources_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_resources_v1_resources_service_proto_goTypes,
		DependencyIndexes: file_api_resources_v1_resources_service_proto_depIdxs,
		MessageInfos:      file_api_resources_v1_resources_service_proto_msgTypes,
	}.Build()
	File_api_resources_v1_resources_service_proto = out.File
	file_api_resources_v1_resources_service_proto_rawDesc = nil
	file_api_resources_v1_resources_service_proto_goTypes = nil
	file_api_resources_v1_resources_service_proto_depIdxs = nil
}