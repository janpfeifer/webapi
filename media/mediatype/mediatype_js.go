// Code generated by webidl-bind. DO NOT EDIT.

package mediatype

import "syscall/js"

// using following types:

// source idl files:
// image-capture.idl

// transform files:
// image-capture.go.md

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

// dictionary: Point2D
type Point2D struct {
	X float64
	Y float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Point2D) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.X
	out.Set("x", value0)
	value1 := _this.Y
	out.Set("y", value1)
	return out
}

// Point2DFromJS is allocating a new
// Point2D object and copy all values from
// input javascript object
func Point2DFromJS(value js.Wrapper) *Point2D {
	input := value.JSValue()
	var out Point2D
	var (
		value0 float64 // javascript: double {x X x}
		value1 float64 // javascript: double {y Y y}
	)
	value0 = (input.Get("x")).Float()
	out.X = value0
	value1 = (input.Get("y")).Float()
	out.Y = value1
	return &out
}

// interface: MediaSettingsRange
type MediaSettingsRange struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaSettingsRange) JSValue() js.Value {
	return _this.Value_JS
}

// MediaSettingsRangeFromJS is casting a js.Wrapper into MediaSettingsRange.
func MediaSettingsRangeFromJS(value js.Wrapper) *MediaSettingsRange {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaSettingsRange{}
	ret.Value_JS = input
	return ret
}

// Max returning attribute 'max' with
// type float64 (idl: double).
func (_this *MediaSettingsRange) Max() float64 {
	var ret float64
	value := _this.Value_JS.Get("max")
	ret = (value).Float()
	return ret
}

// Min returning attribute 'min' with
// type float64 (idl: double).
func (_this *MediaSettingsRange) Min() float64 {
	var ret float64
	value := _this.Value_JS.Get("min")
	ret = (value).Float()
	return ret
}

// Step returning attribute 'step' with
// type float64 (idl: double).
func (_this *MediaSettingsRange) Step() float64 {
	var ret float64
	value := _this.Value_JS.Get("step")
	ret = (value).Float()
	return ret
}
