// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/pod/pod.proto

package pod

import (
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

type PodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PodNamespace  string     `protobuf:"bytes,2,opt,name=pod_namespace,json=podNamespace,proto3" json:"pod_namespace,omitempty"`
	PodName       string     `protobuf:"bytes,3,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	PodTeamId     string     `protobuf:"bytes,4,opt,name=pod_team_id,json=podTeamId,proto3" json:"pod_team_id,omitempty"`
	PodCpuMax     float32    `protobuf:"fixed32,5,opt,name=pod_cpu_max,json=podCpuMax,proto3" json:"pod_cpu_max,omitempty"`
	PodReplicas   int32      `protobuf:"varint,6,opt,name=pod_replicas,json=podReplicas,proto3" json:"pod_replicas,omitempty"`
	PodMemoryMax  float32    `protobuf:"fixed32,7,opt,name=pod_memory_max,json=podMemoryMax,proto3" json:"pod_memory_max,omitempty"`
	PodPorts      []*PodPort `protobuf:"bytes,8,rep,name=pod_ports,json=podPorts,proto3" json:"pod_ports,omitempty"`
	PodEnv        []*PodEnv  `protobuf:"bytes,9,rep,name=pod_env,json=podEnv,proto3" json:"pod_env,omitempty"`
	PodPullPolicy string     `protobuf:"bytes,10,opt,name=pod_pull_policy,json=podPullPolicy,proto3" json:"pod_pull_policy,omitempty"`
	PodRestart    string     `protobuf:"bytes,11,opt,name=pod_restart,json=podRestart,proto3" json:"pod_restart,omitempty"`
	PodType       string     `protobuf:"bytes,12,opt,name=pod_type,json=podType,proto3" json:"pod_type,omitempty"`
	PodImage      string     `protobuf:"bytes,13,opt,name=pod_image,json=podImage,proto3" json:"pod_image,omitempty"`
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{0}
}

func (x *PodInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PodInfo) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

func (x *PodInfo) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *PodInfo) GetPodTeamId() string {
	if x != nil {
		return x.PodTeamId
	}
	return ""
}

func (x *PodInfo) GetPodCpuMax() float32 {
	if x != nil {
		return x.PodCpuMax
	}
	return 0
}

func (x *PodInfo) GetPodReplicas() int32 {
	if x != nil {
		return x.PodReplicas
	}
	return 0
}

func (x *PodInfo) GetPodMemoryMax() float32 {
	if x != nil {
		return x.PodMemoryMax
	}
	return 0
}

func (x *PodInfo) GetPodPorts() []*PodPort {
	if x != nil {
		return x.PodPorts
	}
	return nil
}

func (x *PodInfo) GetPodEnv() []*PodEnv {
	if x != nil {
		return x.PodEnv
	}
	return nil
}

func (x *PodInfo) GetPodPullPolicy() string {
	if x != nil {
		return x.PodPullPolicy
	}
	return ""
}

func (x *PodInfo) GetPodRestart() string {
	if x != nil {
		return x.PodRestart
	}
	return ""
}

func (x *PodInfo) GetPodType() string {
	if x != nil {
		return x.PodType
	}
	return ""
}

func (x *PodInfo) GetPodImage() string {
	if x != nil {
		return x.PodImage
	}
	return ""
}

type PodPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodId         int64  `protobuf:"varint,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	ContainerPort int32  `protobuf:"varint,2,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
	Protocol      string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *PodPort) Reset() {
	*x = PodPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodPort) ProtoMessage() {}

func (x *PodPort) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodPort.ProtoReflect.Descriptor instead.
func (*PodPort) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{1}
}

func (x *PodPort) GetPodId() int64 {
	if x != nil {
		return x.PodId
	}
	return 0
}

func (x *PodPort) GetContainerPort() int32 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

func (x *PodPort) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type PodEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodId    int64  `protobuf:"varint,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	EnvKey   string `protobuf:"bytes,2,opt,name=env_key,json=envKey,proto3" json:"env_key,omitempty"`
	EnvValue string `protobuf:"bytes,3,opt,name=env_value,json=envValue,proto3" json:"env_value,omitempty"`
}

func (x *PodEnv) Reset() {
	*x = PodEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodEnv) ProtoMessage() {}

func (x *PodEnv) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodEnv.ProtoReflect.Descriptor instead.
func (*PodEnv) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{2}
}

func (x *PodEnv) GetPodId() int64 {
	if x != nil {
		return x.PodId
	}
	return 0
}

func (x *PodEnv) GetEnvKey() string {
	if x != nil {
		return x.EnvKey
	}
	return ""
}

func (x *PodEnv) GetEnvValue() string {
	if x != nil {
		return x.EnvValue
	}
	return ""
}

type PodID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PodID) Reset() {
	*x = PodID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodID) ProtoMessage() {}

func (x *PodID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodID.ProtoReflect.Descriptor instead.
func (*PodID) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{3}
}

func (x *PodID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type FindAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAll) Reset() {
	*x = FindAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAll) ProtoMessage() {}

func (x *FindAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAll.ProtoReflect.Descriptor instead.
func (*FindAll) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{5}
}

type AllPod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodInfo []*PodInfo `protobuf:"bytes,1,rep,name=pod_info,json=podInfo,proto3" json:"pod_info,omitempty"`
}

func (x *AllPod) Reset() {
	*x = AllPod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pod_pod_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPod) ProtoMessage() {}

func (x *AllPod) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pod_pod_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPod.ProtoReflect.Descriptor instead.
func (*AllPod) Descriptor() ([]byte, []int) {
	return file_proto_pod_pod_proto_rawDescGZIP(), []int{6}
}

func (x *AllPod) GetPodInfo() []*PodInfo {
	if x != nil {
		return x.PodInfo
	}
	return nil
}

var File_proto_pod_pod_proto protoreflect.FileDescriptor

var file_proto_pod_pod_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x6f, 0x64, 0x22, 0xb4, 0x03, 0x0a, 0x07, 0x50,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x64,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x63, 0x70,
	0x75, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x64,
	0x43, 0x70, 0x75, 0x4d, 0x61, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x64, 0x5f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x64,
	0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x70, 0x6f, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x61, 0x78, 0x12,
	0x29, 0x0a, 0x09, 0x70, 0x6f, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x08, 0x70, 0x6f, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x70, 0x6f,
	0x64, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6f,
	0x64, 0x2e, 0x50, 0x6f, 0x64, 0x45, 0x6e, 0x76, 0x52, 0x06, 0x70, 0x6f, 0x64, 0x45, 0x6e, 0x76,
	0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x64, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6f, 0x64, 0x50, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x55, 0x0a, 0x06, 0x50, 0x6f, 0x64, 0x45, 0x6e, 0x76,
	0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x76, 0x4b, 0x65, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x17, 0x0a,
	0x05, 0x50, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x09, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x22,
	0x31, 0x0a, 0x06, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f,
	0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0xda, 0x01, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64,
	0x12, 0x0a, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x70,
	0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0a, 0x2e, 0x70,
	0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50,
	0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x1a, 0x0b, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x64, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x64, 0x3b, 0x70,
	0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pod_pod_proto_rawDescOnce sync.Once
	file_proto_pod_pod_proto_rawDescData = file_proto_pod_pod_proto_rawDesc
)

func file_proto_pod_pod_proto_rawDescGZIP() []byte {
	file_proto_pod_pod_proto_rawDescOnce.Do(func() {
		file_proto_pod_pod_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pod_pod_proto_rawDescData)
	})
	return file_proto_pod_pod_proto_rawDescData
}

var file_proto_pod_pod_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_pod_pod_proto_goTypes = []interface{}{
	(*PodInfo)(nil),  // 0: pod.PodInfo
	(*PodPort)(nil),  // 1: pod.PodPort
	(*PodEnv)(nil),   // 2: pod.PodEnv
	(*PodID)(nil),    // 3: pod.PodID
	(*Response)(nil), // 4: pod.Response
	(*FindAll)(nil),  // 5: pod.FindAll
	(*AllPod)(nil),   // 6: pod.AllPod
}
var file_proto_pod_pod_proto_depIdxs = []int32{
	1, // 0: pod.PodInfo.pod_ports:type_name -> pod.PodPort
	2, // 1: pod.PodInfo.pod_env:type_name -> pod.PodEnv
	0, // 2: pod.AllPod.pod_info:type_name -> pod.PodInfo
	0, // 3: pod.Pod.AddPod:input_type -> pod.PodInfo
	3, // 4: pod.Pod.DeletePod:input_type -> pod.PodID
	3, // 5: pod.Pod.FindPodByID:input_type -> pod.PodID
	0, // 6: pod.Pod.UpdatePod:input_type -> pod.PodInfo
	5, // 7: pod.Pod.FindAllPod:input_type -> pod.FindAll
	4, // 8: pod.Pod.AddPod:output_type -> pod.Response
	4, // 9: pod.Pod.DeletePod:output_type -> pod.Response
	0, // 10: pod.Pod.FindPodByID:output_type -> pod.PodInfo
	4, // 11: pod.Pod.UpdatePod:output_type -> pod.Response
	6, // 12: pod.Pod.FindAllPod:output_type -> pod.AllPod
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_pod_pod_proto_init() }
func file_proto_pod_pod_proto_init() {
	if File_proto_pod_pod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pod_pod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfo); i {
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
		file_proto_pod_pod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodPort); i {
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
		file_proto_pod_pod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodEnv); i {
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
		file_proto_pod_pod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodID); i {
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
		file_proto_pod_pod_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_pod_pod_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAll); i {
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
		file_proto_pod_pod_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPod); i {
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
			RawDescriptor: file_proto_pod_pod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pod_pod_proto_goTypes,
		DependencyIndexes: file_proto_pod_pod_proto_depIdxs,
		MessageInfos:      file_proto_pod_pod_proto_msgTypes,
	}.Build()
	File_proto_pod_pod_proto = out.File
	file_proto_pod_pod_proto_rawDesc = nil
	file_proto_pod_pod_proto_goTypes = nil
	file_proto_pod_pod_proto_depIdxs = nil
}
