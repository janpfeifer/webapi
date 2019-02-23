// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package datatransfer

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/file/entries"
	"github.com/gowebapi/webapi/html/draganddrop"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// dom.Element
// draganddrop.FunctionStringCallback
// entries.FileSystemEntry
// file.File
// file.FileList
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

// interface: DataTransfer
type DataTransfer struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransfer) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferFromJS is casting a js.Wrapper into DataTransfer.
func DataTransferFromJS(value js.Wrapper) *DataTransfer {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransfer{}
	ret.Value_JS = input
	return ret
}

func NewDataTransfer() (_result *DataTransfer) {
	_klass := js.Global().Get("DataTransfer")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DataTransfer // javascript: DataTransfer _what_return_name
	)
	_converted = DataTransferFromJS(_returned)
	_result = _converted
	return
}

// DropEffect returning attribute 'dropEffect' with
// type string (idl: DOMString).
func (_this *DataTransfer) DropEffect() string {
	var ret string
	value := _this.Value_JS.Get("dropEffect")
	ret = (value).String()
	return ret
}

// SetDropEffect setting attribute 'dropEffect' with
// type string (idl: DOMString).
func (_this *DataTransfer) SetDropEffect(value string) {
	input := value
	_this.Value_JS.Set("dropEffect", input)
}

// EffectAllowed returning attribute 'effectAllowed' with
// type string (idl: DOMString).
func (_this *DataTransfer) EffectAllowed() string {
	var ret string
	value := _this.Value_JS.Get("effectAllowed")
	ret = (value).String()
	return ret
}

// SetEffectAllowed setting attribute 'effectAllowed' with
// type string (idl: DOMString).
func (_this *DataTransfer) SetEffectAllowed(value string) {
	input := value
	_this.Value_JS.Set("effectAllowed", input)
}

// Items returning attribute 'items' with
// type DataTransferItemList (idl: DataTransferItemList).
func (_this *DataTransfer) Items() *DataTransferItemList {
	var ret *DataTransferItemList
	value := _this.Value_JS.Get("items")
	ret = DataTransferItemListFromJS(value)
	return ret
}

// Types returning attribute 'types' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *DataTransfer) Types() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("types")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Files returning attribute 'files' with
// type file.FileList (idl: FileList).
func (_this *DataTransfer) Files() *file.FileList {
	var ret *file.FileList
	value := _this.Value_JS.Get("files")
	ret = file.FileListFromJS(value)
	return ret
}

func (_this *DataTransfer) SetDragImage(image *dom.Element, x int, y int) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := x
	_args[1] = _p1
	_end++
	_p2 := y
	_args[2] = _p2
	_end++
	_this.Value_JS.Call("setDragImage", _args[0:_end]...)
	return
}

func (_this *DataTransfer) GetData(format string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := format
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getData", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *DataTransfer) SetData(format string, data string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := format
	_args[0] = _p0
	_end++
	_p1 := data
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("setData", _args[0:_end]...)
	return
}

func (_this *DataTransfer) ClearData(format *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	if format != nil {
		_p0 := format
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("clearData", _args[0:_end]...)
	return
}

// interface: DataTransferItem
type DataTransferItem struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransferItem) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferItemFromJS is casting a js.Wrapper into DataTransferItem.
func DataTransferItemFromJS(value js.Wrapper) *DataTransferItem {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransferItem{}
	ret.Value_JS = input
	return ret
}

// Kind returning attribute 'kind' with
// type string (idl: DOMString).
func (_this *DataTransferItem) Kind() string {
	var ret string
	value := _this.Value_JS.Get("kind")
	ret = (value).String()
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *DataTransferItem) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

func (_this *DataTransferItem) GetAsString(callback *draganddrop.FunctionStringCallback) {
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
	_this.Value_JS.Call("getAsString", _args[0:_end]...)
	return
}

func (_this *DataTransferItem) GetAsFile() (_result *file.File) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getAsFile", _args[0:_end]...)
	var (
		_converted *file.File // javascript: File _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = file.FileFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *DataTransferItem) WebkitGetAsEntry() (_result *entries.FileSystemEntry) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("webkitGetAsEntry", _args[0:_end]...)
	var (
		_converted *entries.FileSystemEntry // javascript: FileSystemEntry _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = entries.FileSystemEntryFromJS(_returned)
	}
	_result = _converted
	return
}

// interface: DataTransferItemList
type DataTransferItemList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransferItemList) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferItemListFromJS is casting a js.Wrapper into DataTransferItemList.
func DataTransferItemListFromJS(value js.Wrapper) *DataTransferItemList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransferItemList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *DataTransferItemList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *DataTransferItemList) Add(data string, _type string) (_result *DataTransferItem) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_p1 := _type
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("add", _args[0:_end]...)
	var (
		_converted *DataTransferItem // javascript: DataTransferItem _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = DataTransferItemFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *DataTransferItemList) Add2(data *file.File) (_result *DataTransferItem) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("add", _args[0:_end]...)
	var (
		_converted *DataTransferItem // javascript: DataTransferItem _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = DataTransferItemFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *DataTransferItemList) Remove(index uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("remove", _args[0:_end]...)
	return
}

func (_this *DataTransferItemList) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("clear", _args[0:_end]...)
	return
}
