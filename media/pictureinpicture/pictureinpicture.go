// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package pictureinpicture

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget

// source idl files:
// picture-in-picture.idl

// transform files:
// picture-in-picture.go.md

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

// dictionary: EnterPictureInPictureEventInit
type EnterPictureInPictureEventInit struct {
	Bubbles                bool
	Cancelable             bool
	Composed               bool
	PictureInPictureWindow *PictureInPictureWindow
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EnterPictureInPictureEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.PictureInPictureWindow.JSValue()
	out.Set("pictureInPictureWindow", value3)
	return out
}

// EnterPictureInPictureEventInitFromJS is allocating a new
// EnterPictureInPictureEventInit object and copy all values from
// input javascript object
func EnterPictureInPictureEventInitFromJS(value js.Wrapper) *EnterPictureInPictureEventInit {
	input := value.JSValue()
	var out EnterPictureInPictureEventInit
	var (
		value0 bool                    // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                    // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                    // javascript: boolean {composed Composed composed}
		value3 *PictureInPictureWindow // javascript: PictureInPictureWindow {pictureInPictureWindow PictureInPictureWindow pictureInPictureWindow}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = PictureInPictureWindowFromJS(input.Get("pictureInPictureWindow"))
	out.PictureInPictureWindow = value3
	return &out
}

// interface: EnterPictureInPictureEvent
type EnterPictureInPictureEvent struct {
	domcore.Event
}

// EnterPictureInPictureEventFromJS is casting a js.Wrapper into EnterPictureInPictureEvent.
func EnterPictureInPictureEventFromJS(value js.Wrapper) *EnterPictureInPictureEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &EnterPictureInPictureEvent{}
	ret.Value_JS = input
	return ret
}

func NewEnterPictureInPictureEvent(_type string, eventInitDict *EnterPictureInPictureEventInit) (_result *EnterPictureInPictureEvent) {
	_klass := js.Global().Get("EnterPictureInPictureEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *EnterPictureInPictureEvent // javascript: EnterPictureInPictureEvent _what_return_name
	)
	_converted = EnterPictureInPictureEventFromJS(_returned)
	_result = _converted
	return
}

// PictureInPictureWindow returning attribute 'pictureInPictureWindow' with
// type PictureInPictureWindow (idl: PictureInPictureWindow).
func (_this *EnterPictureInPictureEvent) PictureInPictureWindow() *PictureInPictureWindow {
	var ret *PictureInPictureWindow
	value := _this.Value_JS.Get("pictureInPictureWindow")
	ret = PictureInPictureWindowFromJS(value)
	return ret
}

// interface: PictureInPictureWindow
type PictureInPictureWindow struct {
	domcore.EventTarget
}

// PictureInPictureWindowFromJS is casting a js.Wrapper into PictureInPictureWindow.
func PictureInPictureWindowFromJS(value js.Wrapper) *PictureInPictureWindow {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PictureInPictureWindow{}
	ret.Value_JS = input
	return ret
}

// Width returning attribute 'width' with
// type int (idl: long).
func (_this *PictureInPictureWindow) Width() int {
	var ret int
	value := _this.Value_JS.Get("width")
	ret = (value).Int()
	return ret
}

// Height returning attribute 'height' with
// type int (idl: long).
func (_this *PictureInPictureWindow) Height() int {
	var ret int
	value := _this.Value_JS.Get("height")
	ret = (value).Int()
	return ret
}

// Onresize returning attribute 'onresize' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *PictureInPictureWindow) Onresize() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onresize")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnresize setting attribute 'onresize' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *PictureInPictureWindow) SetOnresize(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onresize", input)
}
