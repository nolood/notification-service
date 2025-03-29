// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: template.proto

package templatev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTemplateRequest) Reset() {
	*x = AddTemplateRequest{}
	mi := &file_template_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTemplateRequest) ProtoMessage() {}

func (x *AddTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTemplateRequest.ProtoReflect.Descriptor instead.
func (*AddTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{0}
}

func (x *AddTemplateRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AddTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTemplateResponse) Reset() {
	*x = AddTemplateResponse{}
	mi := &file_template_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTemplateResponse) ProtoMessage() {}

func (x *AddTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTemplateResponse.ProtoReflect.Descriptor instead.
func (*AddTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{1}
}

func (x *AddTemplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddTemplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTemplatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTemplatesRequest) Reset() {
	*x = GetTemplatesRequest{}
	mi := &file_template_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplatesRequest) ProtoMessage() {}

func (x *GetTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplatesRequest.ProtoReflect.Descriptor instead.
func (*GetTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{2}
}

func (x *GetTemplatesRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetTemplatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Templates     []*Template            `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTemplatesResponse) Reset() {
	*x = GetTemplatesResponse{}
	mi := &file_template_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplatesResponse) ProtoMessage() {}

func (x *GetTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetTemplatesResponse) GetTemplates() []*Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

type Template struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Template) Reset() {
	*x = Template{}
	mi := &file_template_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{4}
}

func (x *Template) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTemplateRequest) Reset() {
	*x = UpdateTemplateRequest{}
	mi := &file_template_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateRequest) ProtoMessage() {}

func (x *UpdateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTemplateRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UpdateTemplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTemplateResponse) Reset() {
	*x = UpdateTemplateResponse{}
	mi := &file_template_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateResponse) ProtoMessage() {}

func (x *UpdateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateResponse.ProtoReflect.Descriptor instead.
func (*UpdateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTemplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTemplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTemplateResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DeleteTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	mi := &file_template_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTemplateRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *DeleteTemplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTemplateResponse) Reset() {
	*x = DeleteTemplateResponse{}
	mi := &file_template_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateResponse) ProtoMessage() {}

func (x *DeleteTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTemplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_template_proto protoreflect.FileDescriptor

const file_template_proto_rawDesc = "" +
	"\n" +
	"\x0etemplate.proto\x12\btemplate\"Y\n" +
	"\x12AddTemplateRequest\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"9\n" +
	"\x13AddTemplateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\",\n" +
	"\x13GetTemplatesRequest\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\"H\n" +
	"\x14GetTemplatesResponse\x120\n" +
	"\ttemplates\x18\x01 \x03(\v2\x12.template.TemplateR\ttemplates\"H\n" +
	"\bTemplate\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"l\n" +
	"\x15UpdateTemplateRequest\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\"V\n" +
	"\x16UpdateTemplateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\">\n" +
	"\x15DeleteTemplateRequest\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\"(\n" +
	"\x16DeleteTemplateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xd6\x02\n" +
	"\x0fTemplateService\x12J\n" +
	"\vAddTemplate\x12\x1c.template.AddTemplateRequest\x1a\x1d.template.AddTemplateResponse\x12M\n" +
	"\fGetTemplates\x12\x1d.template.GetTemplatesRequest\x1a\x1e.template.GetTemplatesResponse\x12S\n" +
	"\x0eUpdateTemplate\x12\x1f.template.UpdateTemplateRequest\x1a .template.UpdateTemplateResponse\x12S\n" +
	"\x0eDeleteTemplate\x12\x1f.template.DeleteTemplateRequest\x1a .template.DeleteTemplateResponseB\x1fZ\x1dnolood.template.v1;templatev1b\x06proto3"

var (
	file_template_proto_rawDescOnce sync.Once
	file_template_proto_rawDescData []byte
)

func file_template_proto_rawDescGZIP() []byte {
	file_template_proto_rawDescOnce.Do(func() {
		file_template_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_template_proto_rawDesc), len(file_template_proto_rawDesc)))
	})
	return file_template_proto_rawDescData
}

var file_template_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_template_proto_goTypes = []any{
	(*AddTemplateRequest)(nil),     // 0: template.AddTemplateRequest
	(*AddTemplateResponse)(nil),    // 1: template.AddTemplateResponse
	(*GetTemplatesRequest)(nil),    // 2: template.GetTemplatesRequest
	(*GetTemplatesResponse)(nil),   // 3: template.GetTemplatesResponse
	(*Template)(nil),               // 4: template.Template
	(*UpdateTemplateRequest)(nil),  // 5: template.UpdateTemplateRequest
	(*UpdateTemplateResponse)(nil), // 6: template.UpdateTemplateResponse
	(*DeleteTemplateRequest)(nil),  // 7: template.DeleteTemplateRequest
	(*DeleteTemplateResponse)(nil), // 8: template.DeleteTemplateResponse
}
var file_template_proto_depIdxs = []int32{
	4, // 0: template.GetTemplatesResponse.templates:type_name -> template.Template
	0, // 1: template.TemplateService.AddTemplate:input_type -> template.AddTemplateRequest
	2, // 2: template.TemplateService.GetTemplates:input_type -> template.GetTemplatesRequest
	5, // 3: template.TemplateService.UpdateTemplate:input_type -> template.UpdateTemplateRequest
	7, // 4: template.TemplateService.DeleteTemplate:input_type -> template.DeleteTemplateRequest
	1, // 5: template.TemplateService.AddTemplate:output_type -> template.AddTemplateResponse
	3, // 6: template.TemplateService.GetTemplates:output_type -> template.GetTemplatesResponse
	6, // 7: template.TemplateService.UpdateTemplate:output_type -> template.UpdateTemplateResponse
	8, // 8: template.TemplateService.DeleteTemplate:output_type -> template.DeleteTemplateResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_template_proto_init() }
func file_template_proto_init() {
	if File_template_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_template_proto_rawDesc), len(file_template_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_proto_goTypes,
		DependencyIndexes: file_template_proto_depIdxs,
		MessageInfos:      file_template_proto_msgTypes,
	}.Build()
	File_template_proto = out.File
	file_template_proto_goTypes = nil
	file_template_proto_depIdxs = nil
}
