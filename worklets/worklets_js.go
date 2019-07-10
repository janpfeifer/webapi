// Code generated by webidl-bind. DO NOT EDIT.

package worklets

import "syscall/js"

import (
	"github.com/gowebapi/webapi/fetch"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// fetch.RequestCredentials
// javascript.PromiseVoid

// source idl files:
// worklets.idl

// transform files:
// worklets.go.md

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

// dictionary: WorkletOptions
type WorkletOptions struct {
	Credentials fetch.RequestCredentials
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *WorkletOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Credentials.JSValue()
	out.Set("credentials", value0)
	return out
}

// WorkletOptionsFromJS is allocating a new
// WorkletOptions object and copy all values from
// input javascript object
func WorkletOptionsFromJS(value js.Wrapper) *WorkletOptions {
	input := value.JSValue()
	var out WorkletOptions
	var (
		value0 fetch.RequestCredentials // javascript: RequestCredentials {credentials Credentials credentials}
	)
	value0 = fetch.RequestCredentialsFromJS(input.Get("credentials"))
	out.Credentials = value0
	return &out
}

// interface: Worklet
type Worklet struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Worklet) JSValue() js.Value {
	return _this.Value_JS
}

// WorkletFromJS is casting a js.Wrapper into Worklet.
func WorkletFromJS(value js.Wrapper) *Worklet {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Worklet{}
	ret.Value_JS = input
	return ret
}

func (_this *Worklet) AddModule(moduleURL string, options *WorkletOptions) (_result *javascript.PromiseVoid) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := moduleURL
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("addModule", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

// interface: WorkletGlobalScope
type WorkletGlobalScope struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WorkletGlobalScope) JSValue() js.Value {
	return _this.Value_JS
}

// WorkletGlobalScopeFromJS is casting a js.Wrapper into WorkletGlobalScope.
func WorkletGlobalScopeFromJS(value js.Wrapper) *WorkletGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WorkletGlobalScope{}
	ret.Value_JS = input
	return ret
}
