// Code generated by webidlgenerator. DO NOT EDIT.

// +build !js

package notification

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget

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

// enum: NotificationPermission
type PermissionMode int

const (
	Default PermissionMode = iota
	Denied
	Granted
)

var notificationPermissionToWasmTable = []string{
	"default", "denied", "granted",
}

var notificationPermissionFromWasmTable = map[string]PermissionMode{
	"default": Default, "denied": Denied, "granted": Granted,
}

// JSValue is converting this enum into a java object
func (this *PermissionMode) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PermissionMode) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(notificationPermissionToWasmTable) {
		return notificationPermissionToWasmTable[idx]
	}
	panic("unknown input value")
}

// PermissionModeFromJS is converting a javascript value into
// a PermissionMode enum value.
func PermissionModeFromJS(value js.Value) PermissionMode {
	key := value.String()
	conv, ok := notificationPermissionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: NotificationDirection
type Direction int

const (
	Auto Direction = iota
	LeftToRight
	RightToLeft
)

var notificationDirectionToWasmTable = []string{
	"auto", "ltr", "rtl",
}

var notificationDirectionFromWasmTable = map[string]Direction{
	"auto": Auto, "ltr": LeftToRight, "rtl": RightToLeft,
}

// JSValue is converting this enum into a java object
func (this *Direction) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Direction) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(notificationDirectionToWasmTable) {
		return notificationDirectionToWasmTable[idx]
	}
	panic("unknown input value")
}

// DirectionFromJS is converting a javascript value into
// a Direction enum value.
func DirectionFromJS(value js.Value) Direction {
	key := value.String()
	conv, ok := notificationDirectionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: NotificationPermissionCallback
type PermissionCallbackFunc func(permission PermissionMode)

// PermissionCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PermissionCallback js.Func

func PermissionCallbackToJS(callback PermissionCallbackFunc) *PermissionCallback {
	if callback == nil {
		return nil
	}
	ret := PermissionCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 PermissionMode // javascript: NotificationPermission permission
		)
		_p0 = PermissionModeFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PermissionCallbackFromJS(_value js.Value) PermissionCallbackFunc {
	return func(permission PermissionMode) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := permission.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: NotificationOptions
type Options struct {
	Dir  Direction
	Lang string
	Body string
	Tag  string
	Icon string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Options) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Dir.JSValue()
	out.Set("dir", value0)
	value1 := _this.Lang
	out.Set("lang", value1)
	value2 := _this.Body
	out.Set("body", value2)
	value3 := _this.Tag
	out.Set("tag", value3)
	value4 := _this.Icon
	out.Set("icon", value4)
	return out
}

// OptionsFromJS is allocating a new
// Options object and copy all values from
// input javascript object
func OptionsFromJS(value js.Wrapper) *Options {
	input := value.JSValue()
	var out Options
	var (
		value0 Direction // javascript: NotificationDirection {dir Dir dir}
		value1 string    // javascript: DOMString {lang Lang lang}
		value2 string    // javascript: DOMString {body Body body}
		value3 string    // javascript: DOMString {tag Tag tag}
		value4 string    // javascript: DOMString {icon Icon icon}
	)
	value0 = DirectionFromJS(input.Get("dir"))
	out.Dir = value0
	value1 = (input.Get("lang")).String()
	out.Lang = value1
	value2 = (input.Get("body")).String()
	out.Body = value2
	value3 = (input.Get("tag")).String()
	out.Tag = value3
	value4 = (input.Get("icon")).String()
	out.Icon = value4
	return &out
}

// interface: Notification
type Notification struct {
	domcore.EventTarget
}

// NotificationFromJS is casting a js.Wrapper into Notification.
func NotificationFromJS(value js.Wrapper) *Notification {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Notification{}
	ret.Value_JS = input
	return ret
}

// Permission returning attribute 'permission' with
// type PermissionMode (idl: NotificationPermission).
func Permission() PermissionMode {
	var ret PermissionMode
	_klass := js.Global().Get("Notification")
	value := _klass.Get("permission")
	ret = PermissionModeFromJS(value)
	return ret
}

func RequestPermission(callback *PermissionCallback) {
	_klass := js.Global().Get("Notification")
	_method := _klass.Get("requestPermission")
	var (
		_args [1]interface{}
		_end  int
	)
	if callback != nil {

		var __callback0 js.Value
		if callback != nil {
			__callback0 = (*callback).Value
		} else {
			__callback0 = js.Null()
		}
		_p0 := __callback0
		_args[0] = _p0
		_end++
	}
	_method.Invoke(_args[0:_end]...)
	return
}

func New(title string, options *Options) (_result *Notification) {
	_klass := js.Global().Get("Notification")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := title
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Notification // javascript: Notification _what_return_name
	)
	_converted = NotificationFromJS(_returned)
	_result = _converted
	return
}

// OnClick returning attribute 'onclick' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnClick() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclick")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnClick setting attribute 'onclick' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnClick(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onclick", input)
}

// OnShow returning attribute 'onshow' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnShow() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onshow")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnShow setting attribute 'onshow' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnShow(value *domcore.EventHandler) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("onshow", input)
}

// OnError returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnError() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnError setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnError(value *domcore.EventHandler) {
	var __callback2 js.Value
	if value != nil {
		__callback2 = (*value).Value
	} else {
		__callback2 = js.Null()
	}
	input := __callback2
	_this.Value_JS.Set("onerror", input)
}

// OnClose returning attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnClose() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclose")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnClose setting attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnClose(value *domcore.EventHandler) {
	var __callback3 js.Value
	if value != nil {
		__callback3 = (*value).Value
	} else {
		__callback3 = js.Null()
	}
	input := __callback3
	_this.Value_JS.Set("onclose", input)
}

// Title returning attribute 'title' with
// type string (idl: DOMString).
func (_this *Notification) Title() string {
	var ret string
	value := _this.Value_JS.Get("title")
	ret = (value).String()
	return ret
}

// Dir returning attribute 'dir' with
// type Direction (idl: NotificationDirection).
func (_this *Notification) Dir() Direction {
	var ret Direction
	value := _this.Value_JS.Get("dir")
	ret = DirectionFromJS(value)
	return ret
}

// Lang returning attribute 'lang' with
// type string (idl: DOMString).
func (_this *Notification) Lang() string {
	var ret string
	value := _this.Value_JS.Get("lang")
	ret = (value).String()
	return ret
}

// Body returning attribute 'body' with
// type string (idl: DOMString).
func (_this *Notification) Body() string {
	var ret string
	value := _this.Value_JS.Get("body")
	ret = (value).String()
	return ret
}

// Tag returning attribute 'tag' with
// type string (idl: DOMString).
func (_this *Notification) Tag() string {
	var ret string
	value := _this.Value_JS.Get("tag")
	ret = (value).String()
	return ret
}

// Icon returning attribute 'icon' with
// type string (idl: DOMString).
func (_this *Notification) Icon() string {
	var ret string
	value := _this.Value_JS.Get("icon")
	ret = (value).String()
	return ret
}

func (_this *Notification) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}
