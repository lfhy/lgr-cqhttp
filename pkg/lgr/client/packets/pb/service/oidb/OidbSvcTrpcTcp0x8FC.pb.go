// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x8FC.proto

package oidb

// Rename Group Member
type OidbSvcTrpcTcp0X8FC struct {
	GroupUin uint32                   `protobuf:"varint,1,opt"`
	Body     *OidbSvcTrpcTcp0X8FCBody `protobuf:"bytes,3,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X8FCBody struct {
	TargetUid              string `protobuf:"bytes,1,opt"`
	SpecialTitle           string `protobuf:"bytes,5,opt"`
	SpecialTitleExpireTime int32  `protobuf:"varint,6,opt"`
	UinName                string `protobuf:"bytes,7,opt"`
	TargetName             string `protobuf:"bytes,8,opt"`
	_                      [0]func()
}

type OidbSvcTrpcTcp0X8FC_3Response struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	_        [0]func()
}
