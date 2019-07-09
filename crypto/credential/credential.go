// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package credential

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.AbortSignal
// javascript.Promise

// source idl files:
// credential-management.idl

// transform files:
// credential-management.go.md

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

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

// enum: CredentialMediationRequirement
type CredentialMediationRequirement int

const (
	SilentCredentialMediationRequirement CredentialMediationRequirement = iota
	OptionalCredentialMediationRequirement
	RequiredCredentialMediationRequirement
)

var credentialMediationRequirementToWasmTable = []string{
	"silent", "optional", "required",
}

var credentialMediationRequirementFromWasmTable = map[string]CredentialMediationRequirement{
	"silent": SilentCredentialMediationRequirement, "optional": OptionalCredentialMediationRequirement, "required": RequiredCredentialMediationRequirement,
}

// JSValue is converting this enum into a java object
func (this *CredentialMediationRequirement) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this CredentialMediationRequirement) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(credentialMediationRequirementToWasmTable) {
		return credentialMediationRequirementToWasmTable[idx]
	}
	panic("unknown input value")
}

// CredentialMediationRequirementFromJS is converting a javascript value into
// a CredentialMediationRequirement enum value.
func CredentialMediationRequirementFromJS(value js.Value) CredentialMediationRequirement {
	key := value.String()
	conv, ok := credentialMediationRequirementFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: CredentialCreationOptions
type CredentialCreationOptions struct {
	Signal    *domcore.AbortSignal
	Password  *Union
	Federated *FederatedCredentialInit
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CredentialCreationOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Signal.JSValue()
	out.Set("signal", value0)
	value1 := _this.Password.JSValue()
	out.Set("password", value1)
	value2 := _this.Federated.JSValue()
	out.Set("federated", value2)
	return out
}

// CredentialCreationOptionsFromJS is allocating a new
// CredentialCreationOptions object and copy all values from
// input javascript object
func CredentialCreationOptionsFromJS(value js.Wrapper) *CredentialCreationOptions {
	input := value.JSValue()
	var out CredentialCreationOptions
	var (
		value0 *domcore.AbortSignal     // javascript: AbortSignal {signal Signal signal}
		value1 *Union                   // javascript: Union {password Password password}
		value2 *FederatedCredentialInit // javascript: FederatedCredentialInit {federated Federated federated}
	)
	value0 = domcore.AbortSignalFromJS(input.Get("signal"))
	out.Signal = value0
	value1 = UnionFromJS(input.Get("password"))
	out.Password = value1
	value2 = FederatedCredentialInitFromJS(input.Get("federated"))
	out.Federated = value2
	return &out
}

// dictionary: CredentialData
type CredentialData struct {
	Id string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CredentialData) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Id
	out.Set("id", value0)
	return out
}

// CredentialDataFromJS is allocating a new
// CredentialData object and copy all values from
// input javascript object
func CredentialDataFromJS(value js.Wrapper) *CredentialData {
	input := value.JSValue()
	var out CredentialData
	var (
		value0 string // javascript: USVString {id Id id}
	)
	value0 = (input.Get("id")).String()
	out.Id = value0
	return &out
}

// dictionary: CredentialRequestOptions
type CredentialRequestOptions struct {
	Mediation CredentialMediationRequirement
	Signal    *domcore.AbortSignal
	Password  bool
	Federated *FederatedCredentialRequestOptions
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CredentialRequestOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Mediation.JSValue()
	out.Set("mediation", value0)
	value1 := _this.Signal.JSValue()
	out.Set("signal", value1)
	value2 := _this.Password
	out.Set("password", value2)
	value3 := _this.Federated.JSValue()
	out.Set("federated", value3)
	return out
}

// CredentialRequestOptionsFromJS is allocating a new
// CredentialRequestOptions object and copy all values from
// input javascript object
func CredentialRequestOptionsFromJS(value js.Wrapper) *CredentialRequestOptions {
	input := value.JSValue()
	var out CredentialRequestOptions
	var (
		value0 CredentialMediationRequirement     // javascript: CredentialMediationRequirement {mediation Mediation mediation}
		value1 *domcore.AbortSignal               // javascript: AbortSignal {signal Signal signal}
		value2 bool                               // javascript: boolean {password Password password}
		value3 *FederatedCredentialRequestOptions // javascript: FederatedCredentialRequestOptions {federated Federated federated}
	)
	value0 = CredentialMediationRequirementFromJS(input.Get("mediation"))
	out.Mediation = value0
	value1 = domcore.AbortSignalFromJS(input.Get("signal"))
	out.Signal = value1
	value2 = (input.Get("password")).Bool()
	out.Password = value2
	value3 = FederatedCredentialRequestOptionsFromJS(input.Get("federated"))
	out.Federated = value3
	return &out
}

// dictionary: FederatedCredentialInit
type FederatedCredentialInit struct {
	Id       string
	Name     string
	IconURL  string
	Origin   string
	Provider string
	Protocol string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FederatedCredentialInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Id
	out.Set("id", value0)
	value1 := _this.Name
	out.Set("name", value1)
	value2 := _this.IconURL
	out.Set("iconURL", value2)
	value3 := _this.Origin
	out.Set("origin", value3)
	value4 := _this.Provider
	out.Set("provider", value4)
	value5 := _this.Protocol
	out.Set("protocol", value5)
	return out
}

// FederatedCredentialInitFromJS is allocating a new
// FederatedCredentialInit object and copy all values from
// input javascript object
func FederatedCredentialInitFromJS(value js.Wrapper) *FederatedCredentialInit {
	input := value.JSValue()
	var out FederatedCredentialInit
	var (
		value0 string // javascript: USVString {id Id id}
		value1 string // javascript: USVString {name Name name}
		value2 string // javascript: USVString {iconURL IconURL iconURL}
		value3 string // javascript: USVString {origin Origin origin}
		value4 string // javascript: USVString {provider Provider provider}
		value5 string // javascript: DOMString {protocol Protocol protocol}
	)
	value0 = (input.Get("id")).String()
	out.Id = value0
	value1 = (input.Get("name")).String()
	out.Name = value1
	value2 = (input.Get("iconURL")).String()
	out.IconURL = value2
	value3 = (input.Get("origin")).String()
	out.Origin = value3
	value4 = (input.Get("provider")).String()
	out.Provider = value4
	value5 = (input.Get("protocol")).String()
	out.Protocol = value5
	return &out
}

// dictionary: FederatedCredentialRequestOptions
type FederatedCredentialRequestOptions struct {
	Providers []string
	Protocols []string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FederatedCredentialRequestOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Providers))
	for __idx0, __seq_in0 := range _this.Providers {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("providers", value0)
	value1 := js.Global().Get("Array").New(len(_this.Protocols))
	for __idx1, __seq_in1 := range _this.Protocols {
		__seq_out1 := __seq_in1
		value1.SetIndex(__idx1, __seq_out1)
	}
	out.Set("protocols", value1)
	return out
}

// FederatedCredentialRequestOptionsFromJS is allocating a new
// FederatedCredentialRequestOptions object and copy all values from
// input javascript object
func FederatedCredentialRequestOptionsFromJS(value js.Wrapper) *FederatedCredentialRequestOptions {
	input := value.JSValue()
	var out FederatedCredentialRequestOptions
	var (
		value0 []string // javascript: sequence<USVString> {providers Providers providers}
		value1 []string // javascript: sequence<DOMString> {protocols Protocols protocols}
	)
	__length0 := input.Get("providers").Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := input.Get("providers").Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Providers = value0
	__length1 := input.Get("protocols").Length()
	__array1 := make([]string, __length1, __length1)
	for __idx1 := 0; __idx1 < __length1; __idx1++ {
		var __seq_out1 string
		__seq_in1 := input.Get("protocols").Index(__idx1)
		__seq_out1 = (__seq_in1).String()
		__array1[__idx1] = __seq_out1
	}
	value1 = __array1
	out.Protocols = value1
	return &out
}

// dictionary: PasswordCredentialData
type PasswordCredentialData struct {
	Id       string
	Name     string
	IconURL  string
	Origin   string
	Password string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PasswordCredentialData) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Id
	out.Set("id", value0)
	value1 := _this.Name
	out.Set("name", value1)
	value2 := _this.IconURL
	out.Set("iconURL", value2)
	value3 := _this.Origin
	out.Set("origin", value3)
	value4 := _this.Password
	out.Set("password", value4)
	return out
}

// PasswordCredentialDataFromJS is allocating a new
// PasswordCredentialData object and copy all values from
// input javascript object
func PasswordCredentialDataFromJS(value js.Wrapper) *PasswordCredentialData {
	input := value.JSValue()
	var out PasswordCredentialData
	var (
		value0 string // javascript: USVString {id Id id}
		value1 string // javascript: USVString {name Name name}
		value2 string // javascript: USVString {iconURL IconURL iconURL}
		value3 string // javascript: USVString {origin Origin origin}
		value4 string // javascript: USVString {password Password password}
	)
	value0 = (input.Get("id")).String()
	out.Id = value0
	value1 = (input.Get("name")).String()
	out.Name = value1
	value2 = (input.Get("iconURL")).String()
	out.IconURL = value2
	value3 = (input.Get("origin")).String()
	out.Origin = value3
	value4 = (input.Get("password")).String()
	out.Password = value4
	return &out
}

// interface: Credential
type Credential struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Credential) JSValue() js.Value {
	return _this.Value_JS
}

// CredentialFromJS is casting a js.Wrapper into Credential.
func CredentialFromJS(value js.Wrapper) *Credential {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Credential{}
	ret.Value_JS = input
	return ret
}

// Id returning attribute 'id' with
// type string (idl: USVString).
func (_this *Credential) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Credential) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

// interface: CredentialsContainer
type CredentialsContainer struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CredentialsContainer) JSValue() js.Value {
	return _this.Value_JS
}

// CredentialsContainerFromJS is casting a js.Wrapper into CredentialsContainer.
func CredentialsContainerFromJS(value js.Wrapper) *CredentialsContainer {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CredentialsContainer{}
	ret.Value_JS = input
	return ret
}

func (_this *CredentialsContainer) Get(options *CredentialRequestOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CredentialsContainer) Store(credential *Credential) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := credential.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("store", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CredentialsContainer) Create(options *CredentialCreationOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("create", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CredentialsContainer) PreventSilentAccess() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("preventSilentAccess", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: FederatedCredential
type FederatedCredential struct {
	Credential
}

// FederatedCredentialFromJS is casting a js.Wrapper into FederatedCredential.
func FederatedCredentialFromJS(value js.Wrapper) *FederatedCredential {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FederatedCredential{}
	ret.Value_JS = input
	return ret
}

func NewFederatedCredential(data *FederatedCredentialInit) (_result *FederatedCredential) {
	_klass := js.Global().Get("FederatedCredential")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FederatedCredential // javascript: FederatedCredential _what_return_name
	)
	_converted = FederatedCredentialFromJS(_returned)
	_result = _converted
	return
}

// Provider returning attribute 'provider' with
// type string (idl: USVString).
func (_this *FederatedCredential) Provider() string {
	var ret string
	value := _this.Value_JS.Get("provider")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: DOMString).
func (_this *FederatedCredential) Protocol() *string {
	var ret *string
	value := _this.Value_JS.Get("protocol")
	if value.Type() != js.TypeNull {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// Name returning attribute 'name' with
// type string (idl: USVString).
func (_this *FederatedCredential) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// IconURL returning attribute 'iconURL' with
// type string (idl: USVString).
func (_this *FederatedCredential) IconURL() string {
	var ret string
	value := _this.Value_JS.Get("iconURL")
	ret = (value).String()
	return ret
}

// interface: PasswordCredential
type PasswordCredential struct {
	Credential
}

// PasswordCredentialFromJS is casting a js.Wrapper into PasswordCredential.
func PasswordCredentialFromJS(value js.Wrapper) *PasswordCredential {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PasswordCredential{}
	ret.Value_JS = input
	return ret
}

func NewPasswordCredential(data *PasswordCredentialData) (_result *PasswordCredential) {
	_klass := js.Global().Get("PasswordCredential")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *PasswordCredential // javascript: PasswordCredential _what_return_name
	)
	_converted = PasswordCredentialFromJS(_returned)
	_result = _converted
	return
}

// Password returning attribute 'password' with
// type string (idl: USVString).
func (_this *PasswordCredential) Password() string {
	var ret string
	value := _this.Value_JS.Get("password")
	ret = (value).String()
	return ret
}

// Name returning attribute 'name' with
// type string (idl: USVString).
func (_this *PasswordCredential) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// IconURL returning attribute 'iconURL' with
// type string (idl: USVString).
func (_this *PasswordCredential) IconURL() string {
	var ret string
	value := _this.Value_JS.Get("iconURL")
	ret = (value).String()
	return ret
}
