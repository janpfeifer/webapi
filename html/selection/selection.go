// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package selection

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom"
)

// using following types:
// dom.Node
// dom.Range

// source idl files:
// selection-api.idl

// transform files:
// selection-api.go.md

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

// interface: Selection
type Selection struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Selection) JSValue() js.Value {
	return _this.Value_JS
}

// SelectionFromJS is casting a js.Wrapper into Selection.
func SelectionFromJS(value js.Wrapper) *Selection {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Selection{}
	ret.Value_JS = input
	return ret
}

// AnchorNode returning attribute 'anchorNode' with
// type dom.Node (idl: Node).
func (_this *Selection) AnchorNode() *dom.Node {
	var ret *dom.Node
	value := _this.Value_JS.Get("anchorNode")
	if value.Type() != js.TypeNull {
		ret = dom.NodeFromJS(value)
	}
	return ret
}

// AnchorOffset returning attribute 'anchorOffset' with
// type uint (idl: unsigned long).
func (_this *Selection) AnchorOffset() uint {
	var ret uint
	value := _this.Value_JS.Get("anchorOffset")
	ret = (uint)((value).Int())
	return ret
}

// FocusNode returning attribute 'focusNode' with
// type dom.Node (idl: Node).
func (_this *Selection) FocusNode() *dom.Node {
	var ret *dom.Node
	value := _this.Value_JS.Get("focusNode")
	if value.Type() != js.TypeNull {
		ret = dom.NodeFromJS(value)
	}
	return ret
}

// FocusOffset returning attribute 'focusOffset' with
// type uint (idl: unsigned long).
func (_this *Selection) FocusOffset() uint {
	var ret uint
	value := _this.Value_JS.Get("focusOffset")
	ret = (uint)((value).Int())
	return ret
}

// IsCollapsed returning attribute 'isCollapsed' with
// type bool (idl: boolean).
func (_this *Selection) IsCollapsed() bool {
	var ret bool
	value := _this.Value_JS.Get("isCollapsed")
	ret = (value).Bool()
	return ret
}

// RangeCount returning attribute 'rangeCount' with
// type uint (idl: unsigned long).
func (_this *Selection) RangeCount() uint {
	var ret uint
	value := _this.Value_JS.Get("rangeCount")
	ret = (uint)((value).Int())
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Selection) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

func (_this *Selection) GetRangeAt(index uint) (_result *dom.Range) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getRangeAt", _args[0:_end]...)
	var (
		_converted *dom.Range // javascript: Range _what_return_name
	)
	_converted = dom.RangeFromJS(_returned)
	_result = _converted
	return
}

func (_this *Selection) AddRange(_range *dom.Range) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _range.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("addRange", _args[0:_end]...)
	return
}

func (_this *Selection) RemoveRange(_range *dom.Range) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _range.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("removeRange", _args[0:_end]...)
	return
}

func (_this *Selection) RemoveAllRanges() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("removeAllRanges", _args[0:_end]...)
	return
}

func (_this *Selection) Empty() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("empty", _args[0:_end]...)
	return
}

func (_this *Selection) Collapse(node *dom.Node, offset *uint) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	if offset != nil {
		_p1 := offset
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("collapse", _args[0:_end]...)
	return
}

func (_this *Selection) SetPosition(node *dom.Node, offset *uint) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	if offset != nil {
		_p1 := offset
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("setPosition", _args[0:_end]...)
	return
}

func (_this *Selection) CollapseToStart() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("collapseToStart", _args[0:_end]...)
	return
}

func (_this *Selection) CollapseToEnd() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("collapseToEnd", _args[0:_end]...)
	return
}

func (_this *Selection) Extend(node *dom.Node, offset *uint) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	if offset != nil {
		_p1 := offset
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("extend", _args[0:_end]...)
	return
}

func (_this *Selection) SetBaseAndExtent(anchorNode *dom.Node, anchorOffset uint, focusNode *dom.Node, focusOffset uint) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := anchorNode.JSValue()
	_args[0] = _p0
	_end++
	_p1 := anchorOffset
	_args[1] = _p1
	_end++
	_p2 := focusNode.JSValue()
	_args[2] = _p2
	_end++
	_p3 := focusOffset
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("setBaseAndExtent", _args[0:_end]...)
	return
}

func (_this *Selection) SelectAllChildren(node *dom.Node) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("selectAllChildren", _args[0:_end]...)
	return
}

func (_this *Selection) DeleteFromDocument() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("deleteFromDocument", _args[0:_end]...)
	return
}

func (_this *Selection) ContainsNode(node *dom.Node, allowPartialContainment *bool) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	if allowPartialContainment != nil {
		_p1 := allowPartialContainment
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("containsNode", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}
