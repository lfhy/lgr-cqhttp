// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/login/ntlogin.proto

package login

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type SsoNTLoginEncryptedData struct {
	Sign    []byte `protobuf:"bytes,1,opt"`
	GcmCalc []byte `protobuf:"bytes,3,opt"`
	Type    int32  `protobuf:"varint,4,opt"`
}

type SsoNTLoginBase struct {
	Header *SsoNTLoginHeader `protobuf:"bytes,1,opt"`
	Body   []byte            `protobuf:"bytes,2,opt"`
}

type SsoNTLoginHeader struct {
	Uin     *SsoNTLoginUin     `protobuf:"bytes,1,opt"`
	System  *SsoNTLoginSystem  `protobuf:"bytes,2,opt"`
	Version *SsoNTLoginVersion `protobuf:"bytes,3,opt"`
	Error   *SsoNTLoginError   `protobuf:"bytes,4,opt"`
	Cookie  *SsoNTLoginCookie  `protobuf:"bytes,5,opt"`
	_       [0]func()
}

type SsoNTLoginResponse struct {
	Credentials *SsoNTLoginCredentials `protobuf:"bytes,1,opt"`
	Captcha     *SsoNTLoginCaptchaUrl  `protobuf:"bytes,2,opt"`
	Unusual     *SsoNTLoginUnusual     `protobuf:"bytes,3,opt"`
	Uid         *SsoNTLoginUid         `protobuf:"bytes,4,opt"`
	_           [0]func()
}

type SsoNTLoginEasyLogin struct {
	TempPassword []byte                   `protobuf:"bytes,1,opt"`
	Captcha      *SsoNTLoginCaptchaSubmit `protobuf:"bytes,2,opt"`
}

type SsoNTLoginCaptchaSubmit struct {
	Ticket  string `protobuf:"bytes,1,opt"`
	RandStr string `protobuf:"bytes,2,opt"`
	Aid     string `protobuf:"bytes,3,opt"`
	_       [0]func()
}

type SsoNTLoginCaptchaUrl struct {
	Url string `protobuf:"bytes,3,opt"`
	_   [0]func()
}

type SsoNTLoginCookie struct {
	Cookie proto.Option[string] `protobuf:"bytes,1,opt"`
	_      [0]func()
}

type SsoNTLoginCredentials struct {
	TempPassword []byte `protobuf:"bytes,3,opt"`
	Tgt          []byte `protobuf:"bytes,4,opt"`
	D2           []byte `protobuf:"bytes,5,opt"`
	D2Key        []byte `protobuf:"bytes,6,opt"`
}

type SsoNTLoginError struct {
	ErrorCode          uint32               `protobuf:"varint,1,opt"`
	Tag                string               `protobuf:"bytes,2,opt"`
	Message            string               `protobuf:"bytes,3,opt"`
	NewDeviceVerifyUrl proto.Option[string] `protobuf:"bytes,5,opt"`
	_                  [0]func()
}

type SsoNTLoginSystem struct {
	OS         proto.Option[string] `protobuf:"bytes,1,opt"`
	DeviceName proto.Option[string] `protobuf:"bytes,2,opt"`
	Type       int32                `protobuf:"varint,3,opt"`
	Guid       []byte               `protobuf:"bytes,4,opt"`
}

type SsoNTLoginUid struct {
	Uid string `protobuf:"bytes,2,opt"`
	_   [0]func()
}

type SsoNTLoginUin struct {
	Uin proto.Option[string] `protobuf:"bytes,1,opt"`
	_   [0]func()
}

type SsoNTLoginUnusual struct {
	Sig []byte `protobuf:"bytes,2,opt"`
}

type SsoNTLoginVersion struct {
	KernelVersion proto.Option[string] `protobuf:"bytes,1,opt"`
	AppId         int32                `protobuf:"varint,2,opt"`
	PackageName   proto.Option[string] `protobuf:"bytes,3,opt"`
	_             [0]func()
}
