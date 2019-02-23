// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package appmanifest

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventInit
// javascript.Promise

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

// enum: AppBannerPromptOutcome
type AppBannerPromptOutcome int

const (
	AcceptedAppBannerPromptOutcome AppBannerPromptOutcome = iota
	DismissedAppBannerPromptOutcome
)

var appBannerPromptOutcomeToWasmTable = []string{
	"accepted", "dismissed",
}

var appBannerPromptOutcomeFromWasmTable = map[string]AppBannerPromptOutcome{
	"accepted": AcceptedAppBannerPromptOutcome, "dismissed": DismissedAppBannerPromptOutcome,
}

// JSValue is converting this enum into a java object
func (this *AppBannerPromptOutcome) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AppBannerPromptOutcome) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(appBannerPromptOutcomeToWasmTable) {
		return appBannerPromptOutcomeToWasmTable[idx]
	}
	panic("unknown input value")
}

// AppBannerPromptOutcomeFromJS is converting a javascript value into
// a AppBannerPromptOutcome enum value.
func AppBannerPromptOutcomeFromJS(value js.Value) AppBannerPromptOutcome {
	key := value.String()
	conv, ok := appBannerPromptOutcomeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: ImageResource
type ImageResource struct {
	Src      string
	Sizes    string
	Type     string
	Purpose  string
	Platform string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ImageResource) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Src
	out.Set("src", value0)
	value1 := _this.Sizes
	out.Set("sizes", value1)
	value2 := _this.Type
	out.Set("type", value2)
	value3 := _this.Purpose
	out.Set("purpose", value3)
	value4 := _this.Platform
	out.Set("platform", value4)
	return out
}

// ImageResourceFromJS is allocating a new
// ImageResource object and copy all values from
// input javascript object
func ImageResourceFromJS(value js.Wrapper) *ImageResource {
	input := value.JSValue()
	var out ImageResource
	var (
		value0 string // javascript: USVString {src Src src}
		value1 string // javascript: DOMString {sizes Sizes sizes}
		value2 string // javascript: USVString {type Type _type}
		value3 string // javascript: USVString {purpose Purpose purpose}
		value4 string // javascript: USVString {platform Platform platform}
	)
	value0 = (input.Get("src")).String()
	out.Src = value0
	value1 = (input.Get("sizes")).String()
	out.Sizes = value1
	value2 = (input.Get("type")).String()
	out.Type = value2
	value3 = (input.Get("purpose")).String()
	out.Purpose = value3
	value4 = (input.Get("platform")).String()
	out.Platform = value4
	return &out
}

// dictionary: PromptResponseObject
type PromptResponseObject struct {
	UserChoice AppBannerPromptOutcome
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PromptResponseObject) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.UserChoice.JSValue()
	out.Set("userChoice", value0)
	return out
}

// PromptResponseObjectFromJS is allocating a new
// PromptResponseObject object and copy all values from
// input javascript object
func PromptResponseObjectFromJS(value js.Wrapper) *PromptResponseObject {
	input := value.JSValue()
	var out PromptResponseObject
	var (
		value0 AppBannerPromptOutcome // javascript: AppBannerPromptOutcome {userChoice UserChoice userChoice}
	)
	value0 = AppBannerPromptOutcomeFromJS(input.Get("userChoice"))
	out.UserChoice = value0
	return &out
}

// interface: BeforeInstallPromptEvent
type BeforeInstallPromptEvent struct {
	domcore.Event
}

// BeforeInstallPromptEventFromJS is casting a js.Wrapper into BeforeInstallPromptEvent.
func BeforeInstallPromptEventFromJS(value js.Wrapper) *BeforeInstallPromptEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BeforeInstallPromptEvent{}
	ret.Value_JS = input
	return ret
}

func NewBeforeInstallPromptEvent(_type string, eventInitDict *domcore.EventInit) (_result *BeforeInstallPromptEvent) {
	_klass := js.Global().Get("BeforeInstallPromptEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *BeforeInstallPromptEvent // javascript: BeforeInstallPromptEvent _what_return_name
	)
	_converted = BeforeInstallPromptEventFromJS(_returned)
	_result = _converted
	return
}

func (_this *BeforeInstallPromptEvent) Prompt() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("prompt", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}
