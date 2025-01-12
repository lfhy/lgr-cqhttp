// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x88D_0.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// fetch group 0x88d_0 0x88d_14
type OidbSvcTrpcTcp0X88D struct {
	AppID   uint32                       `protobuf:"varint,1,opt"` // 537099973
	Config2 *OidbSvcTrpcTcp0X88D_Config2 `protobuf:"bytes,2,opt"`
	_       [0]func()
}

type OidbSvcTrpcTcp0X88D_Config2 struct {
	GroupUin  uint32         `protobuf:"varint,1,opt"`
	GroupInfo *D88DGroupInfo `protobuf:"bytes,2,opt"`
	_         [0]func()
}

type D88DGroupHeadPortrait struct {
	_ [0]func()
}

type D88DGroupExInfoOnly struct {
	_ [0]func()
}

type OidbSvcTrpcTcp0X88D_Response struct {
	Info *RspGroupInfo `protobuf:"bytes,1,opt"`
	_    [0]func()
}

type RspGroupInfo struct {
	GroupUin  proto.Option[uint64] `protobuf:"varint,1,opt"`
	Result    proto.Option[uint32] `protobuf:"varint,2,opt"`
	GroupInfo *D88DGroupInfoResp   `protobuf:"bytes,3,opt"`
	_         [0]func()
}

type D88DGroupInfo struct {
	GroupOwner        proto.Option[bool] `protobuf:"varint,1,opt"`
	GroupCreateTime   proto.Option[bool] `protobuf:"varint,2,opt"`
	GroupFlag         proto.Option[bool] `protobuf:"varint,3,opt"`
	GroupFlagExt      proto.Option[bool] `protobuf:"varint,4,opt"`
	GroupMemberMaxNum proto.Option[bool] `protobuf:"varint,5,opt"`
	GroupMemberNum    proto.Option[bool] `protobuf:"varint,6,opt"`
	GroupOption       proto.Option[bool] `protobuf:"varint,7,opt"`
	GroupClassExt     proto.Option[bool] `protobuf:"varint,8,opt"`
	GroupSpecialClass proto.Option[bool] `protobuf:"varint,9,opt"`
	GroupLevel        proto.Option[bool] `protobuf:"varint,10,opt"`
	GroupFace         proto.Option[bool] `protobuf:"varint,11,opt"`
	GroupDefaultPage  proto.Option[bool] `protobuf:"varint,12,opt"`
	// optional bool GroupInfoSeq = 13;
	GroupRoamingTime proto.Option[bool]   `protobuf:"varint,14,opt"`
	GroupName        proto.Option[string] `protobuf:"bytes,15,opt"`
	// optional string  GroupMemo = 16;
	GroupFingerMemo proto.Option[string] `protobuf:"bytes,17,opt"`
	GroupClassText  proto.Option[string] `protobuf:"bytes,18,opt"`
	// repeated bool GroupAllianceCode = 19;
	// optional bool GroupExtraAadmNum = 20;
	GroupUin         proto.Option[bool]   `protobuf:"varint,21,opt"`
	GroupCurMsgSeq   proto.Option[bool]   `protobuf:"varint,22,opt"`
	GroupLastMsgTime proto.Option[bool]   `protobuf:"varint,23,opt"`
	GroupQuestion    proto.Option[string] `protobuf:"bytes,24,opt"`
	GroupAnswer      proto.Option[string] `protobuf:"bytes,25,opt"`
	// optional bool GroupVisitorMaxNum = 26;
	// optional bool GroupVisitorCurNum = 27;
	// optional bool LevelNameSeq = 28;
	// optional bool GroupAdminMaxNum = 29;
	// optional bool GroupAioSkinTimestamp = 30;
	// optional bool GroupBoardSkinTimestamp = 31;
	// optional string  GroupAioSkinUrl = 32;
	// optional string  GroupBoardSkinUrl = 33;
	// optional bool GroupCoverSkinTimestamp = 34;
	// optional string  GroupCoverSkinUrl = 35;
	GroupGrade proto.Option[bool] `protobuf:"varint,36,opt"`
	// optional bool ActiveMemberNum = 37;
	CertificationType   proto.Option[bool]   `protobuf:"varint,38,opt"`
	CertificationText   proto.Option[string] `protobuf:"bytes,39,opt"`
	GroupRichFingerMemo proto.Option[string] `protobuf:"bytes,40,opt"`
	// repeated D88DTagRecord tagRecord = 41;
	// optional D88DGroupGeoInfo groupGeoInfo = 42;
	HeadPortraitSeq   proto.Option[bool]     `protobuf:"varint,43,opt"`
	MsgHeadPortrait   *D88DGroupHeadPortrait `protobuf:"bytes,44,opt"`
	ShutupTimestamp   proto.Option[bool]     `protobuf:"varint,45,opt"`
	ShutupTimestampMe proto.Option[bool]     `protobuf:"varint,46,opt"`
	CreateSourceFlag  proto.Option[bool]     `protobuf:"varint,47,opt"`
	// optional bool CmduinMsgSeq = 48;
	// optional bool CmduinJoinTime = 49;
	// optional bool CmduinUinFlag = 50;
	// optional bool CmduinFlagEx = 51;
	// optional bool CmduinNewMobileFlag = 52;
	// optional bool CmduinReadMsgSeq = 53;
	// optional bool CmduinLastMsgTime = 54;
	GroupTypeFlag    proto.Option[bool] `protobuf:"varint,55,opt"`
	AppPrivilegeFlag proto.Option[bool] `protobuf:"varint,56,opt"`
	// optional D88DGroupExInfoOnly StGroupExInfo = 57;
	GroupSecLevel proto.Option[bool] `protobuf:"varint,58,opt"`
	// optional bool GroupSecLevelInfo = 59;
	// optional bool CmduinPrivilege = 60;
	PoidInfo proto.Option[string] `protobuf:"bytes,61,opt"`
	// optional bool CmduinFlagEx2 = 62;
	// optional bool ConfUin = 63;
	// optional bool ConfMaxMsgSeq = 64;
	// optional bool ConfToGroupTime = 65;
	// optional bool PasswordRedbagTime = 66;
	SubscriptionUin proto.Option[bool] `protobuf:"varint,67,opt"`
	// optional bool MemberListChangeSeq = 68;
	// optional bool MembercardSeq = 69;
	// optional bool RootId = 70;
	// optional bool ParentId = 71;
	// optional bool TeamSeq = 72;
	// optional bool HistoryMsgBeginTime = 73;
	// optional bool InviteNoAuthNumLimit = 74;
	// optional bool CmduinHistoryMsgSeq = 75;
	// optional bool CmduinJoinMsgSeq = 76;
	GroupFlagext3 proto.Option[bool] `protobuf:"varint,77,opt"`
	// optional bool GroupOpenAppid = 78;
	IsConfGroup           proto.Option[bool] `protobuf:"varint,79,opt"`
	IsModifyConfGroupFace proto.Option[bool] `protobuf:"varint,80,opt"`
	IsModifyConfGroupName proto.Option[bool] `protobuf:"varint,81,opt"`
	NoFingerOpenFlag      proto.Option[bool] `protobuf:"varint,82,opt"`
	NoCodeFingerOpenFlag  proto.Option[bool] `protobuf:"varint,83,opt"`
	_                     [0]func()
}

type D88DGroupInfoResp struct {
	GroupOwner        string `protobuf:"bytes,1,opt"`
	GroupCreateTime   uint32 `protobuf:"varint,2,opt"`
	GroupMemberMaxNum uint32 `protobuf:"varint,5,opt"`
	GroupMemberNum    uint32 `protobuf:"varint,6,opt"`
	GroupLevel        uint32 `protobuf:"varint,10,opt"`
	GroupName         string `protobuf:"bytes,15,opt"`
	GroupMemo         string `protobuf:"bytes,16,opt"`
	GroupUin          uint32 `protobuf:"varint,21,opt"`
	GroupCurMsgSeq    uint32 `protobuf:"varint,22,opt"`
	GroupLastMsgTime  uint32 `protobuf:"varint,23,opt"`
	_                 [0]func()
}
