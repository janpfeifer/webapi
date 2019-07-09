// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package backgroundtask

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// requestidlecallback.idl

// transform files:
// requestidlecallback.go.md

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

// callback: IdleRequestCallback
type IdleRequestCallbackFunc func(deadline *IdleDeadline)

// IdleRequestCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type IdleRequestCallback js.Func

func IdleRequestCallbackToJS(callback IdleRequestCallbackFunc) *IdleRequestCallback {
	if callback == nil {
		return nil
	}
	ret := IdleRequestCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *IdleDeadline // javascript: IdleDeadline deadline
		)
		_p0 = IdleDeadlineFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func IdleRequestCallbackFromJS(_value js.Value) IdleRequestCallbackFunc {
	return func(deadline *IdleDeadline) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := deadline.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: IdleRequestOptions
type IdleRequestOptions struct {
	Timeout uint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *IdleRequestOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Timeout
	out.Set("timeout", value0)
	return out
}

// IdleRequestOptionsFromJS is allocating a new
// IdleRequestOptions object and copy all values from
// input javascript object
func IdleRequestOptionsFromJS(value js.Wrapper) *IdleRequestOptions {
	input := value.JSValue()
	var out IdleRequestOptions
	var (
		value0 uint // javascript: unsigned long {timeout Timeout timeout}
	)
	value0 = (uint)((input.Get("timeout")).Int())
	out.Timeout = value0
	return &out
}

// interface: IdleDeadline
type IdleDeadline struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *IdleDeadline) JSValue() js.Value {
	return _this.Value_JS
}

// IdleDeadlineFromJS is casting a js.Wrapper into IdleDeadline.
func IdleDeadlineFromJS(value js.Wrapper) *IdleDeadline {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &IdleDeadline{}
	ret.Value_JS = input
	return ret
}

// DidTimeout returning attribute 'didTimeout' with
// type bool (idl: boolean).
func (_this *IdleDeadline) DidTimeout() bool {
	var ret bool
	value := _this.Value_JS.Get("didTimeout")
	ret = (value).Bool()
	return ret
}

func (_this *IdleDeadline) TimeRemaining() (_result float64) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("timeRemaining", _args[0:_end]...)
	var (
		_converted float64 // javascript: double _what_return_name
	)
	_converted = (_returned).Float()
	_result = _converted
	return
}
