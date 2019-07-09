// Code generated by webidl-bind. DO NOT EDIT.

package lock

import "syscall/js"

import (
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// javascript.Promise

// source idl files:
// keyboard-lock.idl

// transform files:
// keyboard-lock.go.md

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

// interface: Keyboard
type Keyboard struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Keyboard) JSValue() js.Value {
	return _this.Value_JS
}

// KeyboardFromJS is casting a js.Wrapper into Keyboard.
func KeyboardFromJS(value js.Wrapper) *Keyboard {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Keyboard{}
	ret.Value_JS = input
	return ret
}

func (_this *Keyboard) Lock(keyCodes []string) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if keyCodes != nil {
		_p0 := js.Global().Get("Array").New(len(keyCodes))
		for __idx0, __seq_in0 := range keyCodes {
			__seq_out0 := __seq_in0
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("lock", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Keyboard) Unlock() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("unlock", _args[0:_end]...)
	return
}
