// Code generated by webidl-bind. DO NOT EDIT.

package authentication

import "syscall/js"

import (
	"github.com/gowebapi/webapi/device/sensor"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// javascript.ArrayBuffer
// sensor.Coordinates

// source idl files:
// webauthn.idl

// transform files:
// webauthn.go.md

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// enum: AuthenticatorAttachment
type Attachment int

const (
	Platform Attachment = iota
	CrossPlatform
)

var authenticatorAttachmentToWasmTable = []string{
	"platform", "cross-platform",
}

var authenticatorAttachmentFromWasmTable = map[string]Attachment{
	"platform": Platform, "cross-platform": CrossPlatform,
}

// JSValue is converting this enum into a javascript object
func (this *Attachment) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Attachment) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(authenticatorAttachmentToWasmTable) {
		return authenticatorAttachmentToWasmTable[idx]
	}
	panic("unknown input value")
}

// AttachmentFromJS is converting a javascript value into
// a Attachment enum value.
func AttachmentFromJS(value js.Value) Attachment {
	key := value.String()
	conv, ok := authenticatorAttachmentFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: AttestationConveyancePreference
type AttestationConveyancePreference int

const (
	None AttestationConveyancePreference = iota
	Indirect
	Direct
)

var attestationConveyancePreferenceToWasmTable = []string{
	"none", "indirect", "direct",
}

var attestationConveyancePreferenceFromWasmTable = map[string]AttestationConveyancePreference{
	"none": None, "indirect": Indirect, "direct": Direct,
}

// JSValue is converting this enum into a javascript object
func (this *AttestationConveyancePreference) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AttestationConveyancePreference) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(attestationConveyancePreferenceToWasmTable) {
		return attestationConveyancePreferenceToWasmTable[idx]
	}
	panic("unknown input value")
}

// AttestationConveyancePreferenceFromJS is converting a javascript value into
// a AttestationConveyancePreference enum value.
func AttestationConveyancePreferenceFromJS(value js.Value) AttestationConveyancePreference {
	key := value.String()
	conv, ok := attestationConveyancePreferenceFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: TokenBindingStatus
type TokenBindingStatus int

const (
	PresentTokenBindingStatus TokenBindingStatus = iota
	SupportedTokenBindingStatus
)

var tokenBindingStatusToWasmTable = []string{
	"present", "supported",
}

var tokenBindingStatusFromWasmTable = map[string]TokenBindingStatus{
	"present": PresentTokenBindingStatus, "supported": SupportedTokenBindingStatus,
}

// JSValue is converting this enum into a javascript object
func (this *TokenBindingStatus) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this TokenBindingStatus) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(tokenBindingStatusToWasmTable) {
		return tokenBindingStatusToWasmTable[idx]
	}
	panic("unknown input value")
}

// TokenBindingStatusFromJS is converting a javascript value into
// a TokenBindingStatus enum value.
func TokenBindingStatusFromJS(value js.Value) TokenBindingStatus {
	key := value.String()
	conv, ok := tokenBindingStatusFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: AuthenticatorTransport
type Transport int

const (
	TransportUsb Transport = iota
	TransportNfc
	TransportBle
	TransportInternal
)

var authenticatorTransportToWasmTable = []string{
	"usb", "nfc", "ble", "internal",
}

var authenticatorTransportFromWasmTable = map[string]Transport{
	"usb": TransportUsb, "nfc": TransportNfc, "ble": TransportBle, "internal": TransportInternal,
}

// JSValue is converting this enum into a javascript object
func (this *Transport) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Transport) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(authenticatorTransportToWasmTable) {
		return authenticatorTransportToWasmTable[idx]
	}
	panic("unknown input value")
}

// TransportFromJS is converting a javascript value into
// a Transport enum value.
func TransportFromJS(value js.Value) Transport {
	key := value.String()
	conv, ok := authenticatorTransportFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: UserVerificationRequirement
type UserVerificationRequirement int

const (
	Required UserVerificationRequirement = iota
	Preferred
	Discouraged
)

var userVerificationRequirementToWasmTable = []string{
	"required", "preferred", "discouraged",
}

var userVerificationRequirementFromWasmTable = map[string]UserVerificationRequirement{
	"required": Required, "preferred": Preferred, "discouraged": Discouraged,
}

// JSValue is converting this enum into a javascript object
func (this *UserVerificationRequirement) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this UserVerificationRequirement) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(userVerificationRequirementToWasmTable) {
		return userVerificationRequirementToWasmTable[idx]
	}
	panic("unknown input value")
}

// UserVerificationRequirementFromJS is converting a javascript value into
// a UserVerificationRequirement enum value.
func UserVerificationRequirementFromJS(value js.Value) UserVerificationRequirement {
	key := value.String()
	conv, ok := userVerificationRequirementFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: AuthenticationExtensionsClientInputs
type AuthenticationExtensionsClientInputs struct {
	Appid         string
	TxAuthSimple  string
	TxAuthGeneric *TxAuthGenericArg
	AuthnSel      []*Union
	Exts          bool
	Uvi           bool
	Loc           bool
	Uvm           bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AuthenticationExtensionsClientInputs) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Appid
	out.Set("appid", value0)
	value1 := _this.TxAuthSimple
	out.Set("txAuthSimple", value1)
	value2 := _this.TxAuthGeneric.JSValue()
	out.Set("txAuthGeneric", value2)
	value3 := js.Global().Get("Array").New(len(_this.AuthnSel))
	for __idx3, __seq_in3 := range _this.AuthnSel {
		__seq_out3 := __seq_in3.JSValue()
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("authnSel", value3)
	value4 := _this.Exts
	out.Set("exts", value4)
	value5 := _this.Uvi
	out.Set("uvi", value5)
	value6 := _this.Loc
	out.Set("loc", value6)
	value7 := _this.Uvm
	out.Set("uvm", value7)
	return out
}

// AuthenticationExtensionsClientInputsFromJS is allocating a new
// AuthenticationExtensionsClientInputs object and copy all values from
// input javascript object
func AuthenticationExtensionsClientInputsFromJS(value js.Wrapper) *AuthenticationExtensionsClientInputs {
	input := value.JSValue()
	var out AuthenticationExtensionsClientInputs
	var (
		value0 string            // javascript: USVString {appid Appid appid}
		value1 string            // javascript: USVString {txAuthSimple TxAuthSimple txAuthSimple}
		value2 *TxAuthGenericArg // javascript: txAuthGenericArg {txAuthGeneric TxAuthGeneric txAuthGeneric}
		value3 []*Union          // javascript: sequence<Union> {authnSel AuthnSel authnSel}
		value4 bool              // javascript: boolean {exts Exts exts}
		value5 bool              // javascript: boolean {uvi Uvi uvi}
		value6 bool              // javascript: boolean {loc Loc loc}
		value7 bool              // javascript: boolean {uvm Uvm uvm}
	)
	value0 = (input.Get("appid")).String()
	out.Appid = value0
	value1 = (input.Get("txAuthSimple")).String()
	out.TxAuthSimple = value1
	value2 = TxAuthGenericArgFromJS(input.Get("txAuthGeneric"))
	out.TxAuthGeneric = value2
	__length3 := input.Get("authnSel").Length()
	__array3 := make([]*Union, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 *Union
		__seq_in3 := input.Get("authnSel").Index(__idx3)
		__seq_out3 = UnionFromJS(__seq_in3)
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.AuthnSel = value3
	value4 = (input.Get("exts")).Bool()
	out.Exts = value4
	value5 = (input.Get("uvi")).Bool()
	out.Uvi = value5
	value6 = (input.Get("loc")).Bool()
	out.Loc = value6
	value7 = (input.Get("uvm")).Bool()
	out.Uvm = value7
	return &out
}

// dictionary: AuthenticationExtensionsClientOutputs
type AuthenticationExtensionsClientOutputs struct {
	Appid         bool
	TxAuthSimple  string
	TxAuthGeneric *javascript.ArrayBuffer
	AuthnSel      bool
	Exts          []string
	Uvi           *javascript.ArrayBuffer
	Loc           *sensor.Coordinates
	Uvm           [][]uint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AuthenticationExtensionsClientOutputs) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Appid
	out.Set("appid", value0)
	value1 := _this.TxAuthSimple
	out.Set("txAuthSimple", value1)
	value2 := _this.TxAuthGeneric.JSValue()
	out.Set("txAuthGeneric", value2)
	value3 := _this.AuthnSel
	out.Set("authnSel", value3)
	value4 := js.Global().Get("Array").New(len(_this.Exts))
	for __idx4, __seq_in4 := range _this.Exts {
		__seq_out4 := __seq_in4
		value4.SetIndex(__idx4, __seq_out4)
	}
	out.Set("exts", value4)
	value5 := _this.Uvi.JSValue()
	out.Set("uvi", value5)
	value6 := _this.Loc.JSValue()
	out.Set("loc", value6)
	value7 := js.Global().Get("Array").New(len(_this.Uvm))
	for __idx7, __seq_in7 := range _this.Uvm {
		__seq_out7 := js.Global().Get("Array").New(len(__seq_in7))
		for __idx8, __seq_in8 := range __seq_in7 {
			__seq_out8 := __seq_in8
			__seq_out7.SetIndex(__idx8, __seq_out8)
		}
		value7.SetIndex(__idx7, __seq_out7)
	}
	out.Set("uvm", value7)
	return out
}

// AuthenticationExtensionsClientOutputsFromJS is allocating a new
// AuthenticationExtensionsClientOutputs object and copy all values from
// input javascript object
func AuthenticationExtensionsClientOutputsFromJS(value js.Wrapper) *AuthenticationExtensionsClientOutputs {
	input := value.JSValue()
	var out AuthenticationExtensionsClientOutputs
	var (
		value0 bool                    // javascript: boolean {appid Appid appid}
		value1 string                  // javascript: USVString {txAuthSimple TxAuthSimple txAuthSimple}
		value2 *javascript.ArrayBuffer // javascript: ArrayBuffer {txAuthGeneric TxAuthGeneric txAuthGeneric}
		value3 bool                    // javascript: boolean {authnSel AuthnSel authnSel}
		value4 []string                // javascript: sequence<USVString> {exts Exts exts}
		value5 *javascript.ArrayBuffer // javascript: ArrayBuffer {uvi Uvi uvi}
		value6 *sensor.Coordinates     // javascript: Coordinates {loc Loc loc}
		value7 [][]uint                // javascript: sequence<sequence<unsigned long>> {uvm Uvm uvm}
	)
	value0 = (input.Get("appid")).Bool()
	out.Appid = value0
	value1 = (input.Get("txAuthSimple")).String()
	out.TxAuthSimple = value1
	value2 = javascript.ArrayBufferFromJS(input.Get("txAuthGeneric"))
	out.TxAuthGeneric = value2
	value3 = (input.Get("authnSel")).Bool()
	out.AuthnSel = value3
	__length4 := input.Get("exts").Length()
	__array4 := make([]string, __length4, __length4)
	for __idx4 := 0; __idx4 < __length4; __idx4++ {
		var __seq_out4 string
		__seq_in4 := input.Get("exts").Index(__idx4)
		__seq_out4 = (__seq_in4).String()
		__array4[__idx4] = __seq_out4
	}
	value4 = __array4
	out.Exts = value4
	value5 = javascript.ArrayBufferFromJS(input.Get("uvi"))
	out.Uvi = value5
	value6 = sensor.CoordinatesFromJS(input.Get("loc"))
	out.Loc = value6
	__length7 := input.Get("uvm").Length()
	__array7 := make([][]uint, __length7, __length7)
	for __idx7 := 0; __idx7 < __length7; __idx7++ {
		var __seq_out7 []uint
		__seq_in7 := input.Get("uvm").Index(__idx7)
		__length8 := __seq_in7.Length()
		__array8 := make([]uint, __length8, __length8)
		for __idx8 := 0; __idx8 < __length8; __idx8++ {
			var __seq_out8 uint
			__seq_in8 := __seq_in7.Index(__idx8)
			__seq_out8 = (uint)((__seq_in8).Int())
			__array8[__idx8] = __seq_out8
		}
		__seq_out7 = __array8
		__array7[__idx7] = __seq_out7
	}
	value7 = __array7
	out.Uvm = value7
	return &out
}

// dictionary: authenticatorBiometricPerfBounds
type AuthenticatorBiometricPerfBounds struct {
	FAR float32
	FRR float32
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AuthenticatorBiometricPerfBounds) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.FAR
	out.Set("FAR", value0)
	value1 := _this.FRR
	out.Set("FRR", value1)
	return out
}

// AuthenticatorBiometricPerfBoundsFromJS is allocating a new
// AuthenticatorBiometricPerfBounds object and copy all values from
// input javascript object
func AuthenticatorBiometricPerfBoundsFromJS(value js.Wrapper) *AuthenticatorBiometricPerfBounds {
	input := value.JSValue()
	var out AuthenticatorBiometricPerfBounds
	var (
		value0 float32 // javascript: float {FAR FAR fAR}
		value1 float32 // javascript: float {FRR FRR fRR}
	)
	value0 = (float32)((input.Get("FAR")).Float())
	out.FAR = value0
	value1 = (float32)((input.Get("FRR")).Float())
	out.FRR = value1
	return &out
}

// dictionary: CollectedClientData
type CollectedClientData struct {
	Type         string
	Challenge    string
	Origin       string
	TokenBinding *TokenBinding
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CollectedClientData) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type
	out.Set("type", value0)
	value1 := _this.Challenge
	out.Set("challenge", value1)
	value2 := _this.Origin
	out.Set("origin", value2)
	value3 := _this.TokenBinding.JSValue()
	out.Set("tokenBinding", value3)
	return out
}

// CollectedClientDataFromJS is allocating a new
// CollectedClientData object and copy all values from
// input javascript object
func CollectedClientDataFromJS(value js.Wrapper) *CollectedClientData {
	input := value.JSValue()
	var out CollectedClientData
	var (
		value0 string        // javascript: DOMString {type Type _type}
		value1 string        // javascript: DOMString {challenge Challenge challenge}
		value2 string        // javascript: DOMString {origin Origin origin}
		value3 *TokenBinding // javascript: TokenBinding {tokenBinding TokenBinding tokenBinding}
	)
	value0 = (input.Get("type")).String()
	out.Type = value0
	value1 = (input.Get("challenge")).String()
	out.Challenge = value1
	value2 = (input.Get("origin")).String()
	out.Origin = value2
	value3 = TokenBindingFromJS(input.Get("tokenBinding"))
	out.TokenBinding = value3
	return &out
}

// dictionary: AuthenticatorSelectionCriteria
type SelectionCriteria struct {
	AuthenticatorAttachment Attachment
	RequireResidentKey      bool
	UserVerification        UserVerificationRequirement
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *SelectionCriteria) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.AuthenticatorAttachment.JSValue()
	out.Set("authenticatorAttachment", value0)
	value1 := _this.RequireResidentKey
	out.Set("requireResidentKey", value1)
	value2 := _this.UserVerification.JSValue()
	out.Set("userVerification", value2)
	return out
}

// SelectionCriteriaFromJS is allocating a new
// SelectionCriteria object and copy all values from
// input javascript object
func SelectionCriteriaFromJS(value js.Wrapper) *SelectionCriteria {
	input := value.JSValue()
	var out SelectionCriteria
	var (
		value0 Attachment                  // javascript: AuthenticatorAttachment {authenticatorAttachment AuthenticatorAttachment authenticatorAttachment}
		value1 bool                        // javascript: boolean {requireResidentKey RequireResidentKey requireResidentKey}
		value2 UserVerificationRequirement // javascript: UserVerificationRequirement {userVerification UserVerification userVerification}
	)
	value0 = AttachmentFromJS(input.Get("authenticatorAttachment"))
	out.AuthenticatorAttachment = value0
	value1 = (input.Get("requireResidentKey")).Bool()
	out.RequireResidentKey = value1
	value2 = UserVerificationRequirementFromJS(input.Get("userVerification"))
	out.UserVerification = value2
	return &out
}

// dictionary: TokenBinding
type TokenBinding struct {
	Status TokenBindingStatus
	Id     string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *TokenBinding) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Status.JSValue()
	out.Set("status", value0)
	value1 := _this.Id
	out.Set("id", value1)
	return out
}

// TokenBindingFromJS is allocating a new
// TokenBinding object and copy all values from
// input javascript object
func TokenBindingFromJS(value js.Wrapper) *TokenBinding {
	input := value.JSValue()
	var out TokenBinding
	var (
		value0 TokenBindingStatus // javascript: TokenBindingStatus {status Status status}
		value1 string             // javascript: DOMString {id Id id}
	)
	value0 = TokenBindingStatusFromJS(input.Get("status"))
	out.Status = value0
	value1 = (input.Get("id")).String()
	out.Id = value1
	return &out
}

// dictionary: txAuthGenericArg
type TxAuthGenericArg struct {
	ContentType string
	Content     *javascript.ArrayBuffer
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *TxAuthGenericArg) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.ContentType
	out.Set("contentType", value0)
	value1 := _this.Content.JSValue()
	out.Set("content", value1)
	return out
}

// TxAuthGenericArgFromJS is allocating a new
// TxAuthGenericArg object and copy all values from
// input javascript object
func TxAuthGenericArgFromJS(value js.Wrapper) *TxAuthGenericArg {
	input := value.JSValue()
	var out TxAuthGenericArg
	var (
		value0 string                  // javascript: USVString {contentType ContentType contentType}
		value1 *javascript.ArrayBuffer // javascript: ArrayBuffer {content Content content}
	)
	value0 = (input.Get("contentType")).String()
	out.ContentType = value0
	value1 = javascript.ArrayBufferFromJS(input.Get("content"))
	out.Content = value1
	return &out
}

// class: AuthenticatorAssertionResponse
type AssertionResponse struct {
	Response
}

// AssertionResponseFromJS is casting a js.Wrapper into AssertionResponse.
func AssertionResponseFromJS(value js.Wrapper) *AssertionResponse {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AssertionResponse{}
	ret.Value_JS = input
	return ret
}

// AuthenticatorData returning attribute 'authenticatorData' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *AssertionResponse) AuthenticatorData() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("authenticatorData")
	ret = javascript.ArrayBufferFromJS(value)
	return ret
}

// Signature returning attribute 'signature' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *AssertionResponse) Signature() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("signature")
	ret = javascript.ArrayBufferFromJS(value)
	return ret
}

// UserHandle returning attribute 'userHandle' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *AssertionResponse) UserHandle() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("userHandle")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = javascript.ArrayBufferFromJS(value)
	}
	return ret
}

// class: AuthenticatorAttestationResponse
type AttestationResponse struct {
	Response
}

// AttestationResponseFromJS is casting a js.Wrapper into AttestationResponse.
func AttestationResponseFromJS(value js.Wrapper) *AttestationResponse {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AttestationResponse{}
	ret.Value_JS = input
	return ret
}

// AttestationObject returning attribute 'attestationObject' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *AttestationResponse) AttestationObject() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("attestationObject")
	ret = javascript.ArrayBufferFromJS(value)
	return ret
}

// class: AuthenticatorResponse
type Response struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Response) JSValue() js.Value {
	return _this.Value_JS
}

// ResponseFromJS is casting a js.Wrapper into Response.
func ResponseFromJS(value js.Wrapper) *Response {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Response{}
	ret.Value_JS = input
	return ret
}

// ClientDataJSON returning attribute 'clientDataJSON' with
// type javascript.ArrayBuffer (idl: ArrayBuffer).
func (_this *Response) ClientDataJSON() *javascript.ArrayBuffer {
	var ret *javascript.ArrayBuffer
	value := _this.Value_JS.Get("clientDataJSON")
	ret = javascript.ArrayBufferFromJS(value)
	return ret
}
