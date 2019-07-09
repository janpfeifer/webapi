// Code generated by webidl-bind. DO NOT EDIT.

package remoteplayback

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget
// javascript.Promise

// source idl files:
// remote-playback.idl

// transform files:
// remote-playback.go.md

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

// enum: RemotePlaybackState
type RemotePlaybackState int

const (
	ConnectingRemotePlaybackState RemotePlaybackState = iota
	ConnectedRemotePlaybackState
	DisconnectedRemotePlaybackState
)

var remotePlaybackStateToWasmTable = []string{
	"connecting", "connected", "disconnected",
}

var remotePlaybackStateFromWasmTable = map[string]RemotePlaybackState{
	"connecting": ConnectingRemotePlaybackState, "connected": ConnectedRemotePlaybackState, "disconnected": DisconnectedRemotePlaybackState,
}

// JSValue is converting this enum into a java object
func (this *RemotePlaybackState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RemotePlaybackState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(remotePlaybackStateToWasmTable) {
		return remotePlaybackStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// RemotePlaybackStateFromJS is converting a javascript value into
// a RemotePlaybackState enum value.
func RemotePlaybackStateFromJS(value js.Value) RemotePlaybackState {
	key := value.String()
	conv, ok := remotePlaybackStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: RemotePlaybackAvailabilityCallback
type RemotePlaybackAvailabilityCallbackFunc func(available bool)

// RemotePlaybackAvailabilityCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type RemotePlaybackAvailabilityCallback js.Func

func RemotePlaybackAvailabilityCallbackToJS(callback RemotePlaybackAvailabilityCallbackFunc) *RemotePlaybackAvailabilityCallback {
	if callback == nil {
		return nil
	}
	ret := RemotePlaybackAvailabilityCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 bool // javascript: boolean available
		)
		_p0 = (args[0]).Bool()
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func RemotePlaybackAvailabilityCallbackFromJS(_value js.Value) RemotePlaybackAvailabilityCallbackFunc {
	return func(available bool) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := available
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// interface: RemotePlayback
type RemotePlayback struct {
	domcore.EventTarget
}

// RemotePlaybackFromJS is casting a js.Wrapper into RemotePlayback.
func RemotePlaybackFromJS(value js.Wrapper) *RemotePlayback {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &RemotePlayback{}
	ret.Value_JS = input
	return ret
}

// State returning attribute 'state' with
// type RemotePlaybackState (idl: RemotePlaybackState).
func (_this *RemotePlayback) State() RemotePlaybackState {
	var ret RemotePlaybackState
	value := _this.Value_JS.Get("state")
	ret = RemotePlaybackStateFromJS(value)
	return ret
}

// Onconnecting returning attribute 'onconnecting' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) Onconnecting() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnecting")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnecting setting attribute 'onconnecting' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) SetOnconnecting(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onconnecting", input)
}

// Onconnect returning attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) Onconnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnect")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnect setting attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) SetOnconnect(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onconnect", input)
}

// Ondisconnect returning attribute 'ondisconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) Ondisconnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ondisconnect")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOndisconnect setting attribute 'ondisconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *RemotePlayback) SetOndisconnect(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("ondisconnect", input)
}

func (_this *RemotePlayback) WatchAvailability(callback *RemotePlaybackAvailabilityCallback) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("watchAvailability", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *RemotePlayback) CancelWatchAvailability(id *int) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if id != nil {
		_p0 := id
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("cancelWatchAvailability", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *RemotePlayback) Prompt() (_result *javascript.Promise) {
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
