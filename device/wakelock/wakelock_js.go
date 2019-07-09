// Code generated by webidl-bind. DO NOT EDIT.

package wakelock

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget

// source idl files:
// wake-lock.idl

// transform files:
// wake-lock.go.md

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

// enum: WakeLockType
type WakeLockType int

const (
	ScreenWakeLockType WakeLockType = iota
	SystemWakeLockType
)

var wakeLockTypeToWasmTable = []string{
	"screen", "system",
}

var wakeLockTypeFromWasmTable = map[string]WakeLockType{
	"screen": ScreenWakeLockType, "system": SystemWakeLockType,
}

// JSValue is converting this enum into a java object
func (this *WakeLockType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this WakeLockType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(wakeLockTypeToWasmTable) {
		return wakeLockTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// WakeLockTypeFromJS is converting a javascript value into
// a WakeLockType enum value.
func WakeLockTypeFromJS(value js.Value) WakeLockType {
	key := value.String()
	conv, ok := wakeLockTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// interface: WakeLock
type WakeLock struct {
	domcore.EventTarget
}

// WakeLockFromJS is casting a js.Wrapper into WakeLock.
func WakeLockFromJS(value js.Wrapper) *WakeLock {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WakeLock{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type WakeLockType (idl: WakeLockType).
func (_this *WakeLock) Type() WakeLockType {
	var ret WakeLockType
	value := _this.Value_JS.Get("type")
	ret = WakeLockTypeFromJS(value)
	return ret
}

// Active returning attribute 'active' with
// type bool (idl: boolean).
func (_this *WakeLock) Active() bool {
	var ret bool
	value := _this.Value_JS.Get("active")
	ret = (value).Bool()
	return ret
}

// Onactivechange returning attribute 'onactivechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WakeLock) Onactivechange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onactivechange")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnactivechange setting attribute 'onactivechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WakeLock) SetOnactivechange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onactivechange", input)
}

func (_this *WakeLock) CreateRequest() (_result *WakeLockRequest) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("createRequest", _args[0:_end]...)
	var (
		_converted *WakeLockRequest // javascript: WakeLockRequest _what_return_name
	)
	_converted = WakeLockRequestFromJS(_returned)
	_result = _converted
	return
}

// interface: WakeLockRequest
type WakeLockRequest struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WakeLockRequest) JSValue() js.Value {
	return _this.Value_JS
}

// WakeLockRequestFromJS is casting a js.Wrapper into WakeLockRequest.
func WakeLockRequestFromJS(value js.Wrapper) *WakeLockRequest {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WakeLockRequest{}
	ret.Value_JS = input
	return ret
}

func (_this *WakeLockRequest) Cancel() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("cancel", _args[0:_end]...)
	return
}
