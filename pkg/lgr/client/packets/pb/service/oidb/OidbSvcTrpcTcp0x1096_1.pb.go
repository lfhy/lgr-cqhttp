// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x1096_1.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// SetGroupAdmin
type OidbSvcTrpcTcp0X1096_1 struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	Uid      string `protobuf:"bytes,2,opt"`
	IsAdmin  bool   `protobuf:"varint,3,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X1096_1Response struct {
	Success proto.Option[string] `protobuf:"bytes,1,opt"`
	_       [0]func()
}
