// Code generated by webidl-bind. DO NOT EDIT.

package boxtree

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// dom.Node
// javascript.FrozenArray

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

// interface: DeadFragmentInformation
type DeadFragmentInformation struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DeadFragmentInformation) JSValue() js.Value {
	return _this.Value_JS
}

// DeadFragmentInformationFromJS is casting a js.Wrapper into DeadFragmentInformation.
func DeadFragmentInformationFromJS(value js.Wrapper) *DeadFragmentInformation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DeadFragmentInformation{}
	ret.Value_JS = input
	return ret
}

// Node returning attribute 'node' with
// type dom.Node (idl: Node).
func (_this *DeadFragmentInformation) Node() *dom.Node {
	var ret *dom.Node
	value := _this.Value_JS.Get("node")
	ret = dom.NodeFromJS(value)
	return ret
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *DeadFragmentInformation) Width() float64 {
	var ret float64
	value := _this.Value_JS.Get("width")
	ret = (value).Float()
	return ret
}

// Height returning attribute 'height' with
// type float64 (idl: double).
func (_this *DeadFragmentInformation) Height() float64 {
	var ret float64
	value := _this.Value_JS.Get("height")
	ret = (value).Float()
	return ret
}

// Top returning attribute 'top' with
// type float64 (idl: double).
func (_this *DeadFragmentInformation) Top() float64 {
	var ret float64
	value := _this.Value_JS.Get("top")
	ret = (value).Float()
	return ret
}

// Left returning attribute 'left' with
// type float64 (idl: double).
func (_this *DeadFragmentInformation) Left() float64 {
	var ret float64
	value := _this.Value_JS.Get("left")
	ret = (value).Float()
	return ret
}

// IsOverflowed returning attribute 'isOverflowed' with
// type bool (idl: boolean).
func (_this *DeadFragmentInformation) IsOverflowed() bool {
	var ret bool
	value := _this.Value_JS.Get("isOverflowed")
	ret = (value).Bool()
	return ret
}

// Children returning attribute 'children' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *DeadFragmentInformation) Children() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("children")
	if value.Type() != js.TypeNull {
		ret = javascript.FrozenArrayFromJS(value)
	}
	return ret
}

// NextSibling returning attribute 'nextSibling' with
// type DeadFragmentInformation (idl: DeadFragmentInformation).
func (_this *DeadFragmentInformation) NextSibling() *DeadFragmentInformation {
	var ret *DeadFragmentInformation
	value := _this.Value_JS.Get("nextSibling")
	if value.Type() != js.TypeNull {
		ret = DeadFragmentInformationFromJS(value)
	}
	return ret
}

// PreviousSibling returning attribute 'previousSibling' with
// type DeadFragmentInformation (idl: DeadFragmentInformation).
func (_this *DeadFragmentInformation) PreviousSibling() *DeadFragmentInformation {
	var ret *DeadFragmentInformation
	value := _this.Value_JS.Get("previousSibling")
	if value.Type() != js.TypeNull {
		ret = DeadFragmentInformationFromJS(value)
	}
	return ret
}

// NextInBox returning attribute 'nextInBox' with
// type DeadFragmentInformation (idl: DeadFragmentInformation).
func (_this *DeadFragmentInformation) NextInBox() *DeadFragmentInformation {
	var ret *DeadFragmentInformation
	value := _this.Value_JS.Get("nextInBox")
	if value.Type() != js.TypeNull {
		ret = DeadFragmentInformationFromJS(value)
	}
	return ret
}

// PreviousInBox returning attribute 'previousInBox' with
// type DeadFragmentInformation (idl: DeadFragmentInformation).
func (_this *DeadFragmentInformation) PreviousInBox() *DeadFragmentInformation {
	var ret *DeadFragmentInformation
	value := _this.Value_JS.Get("previousInBox")
	if value.Type() != js.TypeNull {
		ret = DeadFragmentInformationFromJS(value)
	}
	return ret
}
