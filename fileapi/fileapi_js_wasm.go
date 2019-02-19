// Code generated by webidlgenerator. DO NOT EDIT.

package fileapi

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.DOMException
// domcore.EventHandler
// domcore.EventTarget
// javascript.ArrayBuffer

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

// enum: EndingType
type EndingType int

const (
	TransparentEndingType EndingType = iota
	NativeEndingType
)

var endingTypeToWasmTable = []string{
	"transparent", "native",
}

var endingTypeFromWasmTable = map[string]EndingType{
	"transparent": TransparentEndingType, "native": NativeEndingType,
}

// JSValue is converting this enum into a java object
func (this *EndingType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this EndingType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(endingTypeToWasmTable) {
		return endingTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// EndingTypeFromJS is converting a javascript value into
// a EndingType enum value.
func EndingTypeFromJS(value js.Value) EndingType {
	key := value.String()
	conv, ok := endingTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: BlobCallback
type BlobCallbackFunc func(blob *Blob)

// BlobCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type BlobCallback js.Func

func BlobCallbackToJS(callback BlobCallbackFunc) *BlobCallback {
	if callback == nil {
		return nil
	}
	ret := BlobCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Blob // javascript: Blob blob
		)
		if args[0].Type() != js.TypeNull {
			_p0 = BlobFromJS(args[0])
		}
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func BlobCallbackFromJS(_value js.Value) BlobCallbackFunc {
	return func(blob *Blob) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := blob.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: BlobPropertyBag
type BlobPropertyBag struct {
	Type    string
	Endings EndingType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BlobPropertyBag) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type
	out.Set("type", value0)
	value1 := _this.Endings.JSValue()
	out.Set("endings", value1)
	return out
}

// BlobPropertyBagFromJS is allocating a new
// BlobPropertyBag object and copy all values from
// input javascript object
func BlobPropertyBagFromJS(value js.Wrapper) *BlobPropertyBag {
	input := value.JSValue()
	var out BlobPropertyBag
	var (
		value0 string     // javascript: DOMString {type Type _type}
		value1 EndingType // javascript: EndingType {endings Endings endings}
	)
	value0 = (input.Get("type")).String()
	out.Type = value0
	value1 = EndingTypeFromJS(input.Get("endings"))
	out.Endings = value1
	return &out
}

// dictionary: FilePropertyBag
type FilePropertyBag struct {
	LastModified int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FilePropertyBag) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.LastModified
	out.Set("lastModified", value0)
	return out
}

// FilePropertyBagFromJS is allocating a new
// FilePropertyBag object and copy all values from
// input javascript object
func FilePropertyBagFromJS(value js.Wrapper) *FilePropertyBag {
	input := value.JSValue()
	var out FilePropertyBag
	var (
		value0 int // javascript: long long {lastModified LastModified lastModified}
	)
	value0 = (input.Get("lastModified")).Int()
	out.LastModified = value0
	return &out
}

// interface: Blob
type Blob struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Blob) JSValue() js.Value {
	return _this.Value_JS
}

// BlobFromJS is casting a js.Wrapper into Blob.
func BlobFromJS(value js.Wrapper) *Blob {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Blob{}
	ret.Value_JS = input
	return ret
}

func NewBlob(blobParts []*Union, options *BlobPropertyBag) (_result *Blob) {
	_klass := js.Global().Get("Blob")
	var (
		_args [2]interface{}
		_end  int
	)
	if blobParts != nil {
		_p0 := js.Global().Get("Array").New(len(blobParts))
		for __idx, __seq_in := range blobParts {
			__seq_out := __seq_in.JSValue()
			_p0.SetIndex(__idx, __seq_out)
		}
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Blob // javascript: Blob _what_return_name
	)
	_converted = BlobFromJS(_returned)
	_result = _converted
	return
}

// Size returning attribute 'size' with
// type int (idl: unsigned long long).
func (_this *Blob) Size() int {
	var ret int
	value := _this.Value_JS.Get("size")
	ret = (value).Int()
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Blob) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

func (_this *Blob) Slice(start *int, end *int, contentType *string) (_result *Blob) {
	var (
		_args [3]interface{}
		_end  int
	)
	if start != nil {
		_p0 := start
		_args[0] = _p0
		_end++
	}
	if end != nil {
		_p1 := end
		_args[1] = _p1
		_end++
	}
	if contentType != nil {
		_p2 := contentType
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("slice", _args[0:_end]...)
	var (
		_converted *Blob // javascript: Blob _what_return_name
	)
	_converted = BlobFromJS(_returned)
	_result = _converted
	return
}

// interface: File
type File struct {
	Blob
}

// FileFromJS is casting a js.Wrapper into File.
func FileFromJS(value js.Wrapper) *File {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &File{}
	ret.Value_JS = input
	return ret
}

func NewFile(fileBits []*Union, fileName string, options *FilePropertyBag) (_result *File) {
	_klass := js.Global().Get("File")
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(fileBits))
	for __idx, __seq_in := range fileBits {
		__seq_out := __seq_in.JSValue()
		_p0.SetIndex(__idx, __seq_out)
	}
	_args[0] = _p0
	_end++
	_p1 := fileName
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *File // javascript: File _what_return_name
	)
	_converted = FileFromJS(_returned)
	_result = _converted
	return
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *File) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// LastModified returning attribute 'lastModified' with
// type int (idl: long long).
func (_this *File) LastModified() int {
	var ret int
	value := _this.Value_JS.Get("lastModified")
	ret = (value).Int()
	return ret
}

// interface: FileList
type FileList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileList) JSValue() js.Value {
	return _this.Value_JS
}

// FileListFromJS is casting a js.Wrapper into FileList.
func FileListFromJS(value js.Wrapper) *FileList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *FileList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *FileList) Item(index uint) (_result *File) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *File // javascript: File _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = FileFromJS(_returned)
	}
	_result = _converted
	return
}

// interface: FileReader
type FileReader struct {
	domcore.EventTarget
}

// FileReaderFromJS is casting a js.Wrapper into FileReader.
func FileReaderFromJS(value js.Wrapper) *FileReader {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileReader{}
	ret.Value_JS = input
	return ret
}

const EMPTY_FileReader int = 0
const LOADING_FileReader int = 1
const DONE_FileReader int = 2

func NewFileReader() (_result *FileReader) {
	_klass := js.Global().Get("FileReader")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FileReader // javascript: FileReader _what_return_name
	)
	_converted = FileReaderFromJS(_returned)
	_result = _converted
	return
}

// ReadyState returning attribute 'readyState' with
// type int (idl: unsigned short).
func (_this *FileReader) ReadyState() int {
	var ret int
	value := _this.Value_JS.Get("readyState")
	ret = (value).Int()
	return ret
}

// Result returning attribute 'result' with
// type Union (idl: Union).
func (_this *FileReader) Result() *Union {
	var ret *Union
	value := _this.Value_JS.Get("result")
	if value.Type() != js.TypeNull {
		ret = UnionFromJS(value)
	}
	return ret
}

// Error returning attribute 'error' with
// type domcore.DOMException (idl: DOMException).
func (_this *FileReader) Error() *domcore.DOMException {
	var ret *domcore.DOMException
	value := _this.Value_JS.Get("error")
	if value.Type() != js.TypeNull {
		ret = domcore.DOMExceptionFromJS(value)
	}
	return ret
}

// Onloadstart returning attribute 'onloadstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onloadstart() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadstart")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnloadstart setting attribute 'onloadstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnloadstart(value *domcore.EventHandler) {
	var __callback3 js.Value
	if value != nil {
		__callback3 = (*value).Value
	} else {
		__callback3 = js.Null()
	}
	input := __callback3
	_this.Value_JS.Set("onloadstart", input)
}

// Onprogress returning attribute 'onprogress' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onprogress() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onprogress")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnprogress setting attribute 'onprogress' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnprogress(value *domcore.EventHandler) {
	var __callback4 js.Value
	if value != nil {
		__callback4 = (*value).Value
	} else {
		__callback4 = js.Null()
	}
	input := __callback4
	_this.Value_JS.Set("onprogress", input)
}

// Onload returning attribute 'onload' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onload() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onload")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnload setting attribute 'onload' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnload(value *domcore.EventHandler) {
	var __callback5 js.Value
	if value != nil {
		__callback5 = (*value).Value
	} else {
		__callback5 = js.Null()
	}
	input := __callback5
	_this.Value_JS.Set("onload", input)
}

// Onabort returning attribute 'onabort' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onabort() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onabort")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnabort setting attribute 'onabort' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnabort(value *domcore.EventHandler) {
	var __callback6 js.Value
	if value != nil {
		__callback6 = (*value).Value
	} else {
		__callback6 = js.Null()
	}
	input := __callback6
	_this.Value_JS.Set("onabort", input)
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnerror(value *domcore.EventHandler) {
	var __callback7 js.Value
	if value != nil {
		__callback7 = (*value).Value
	} else {
		__callback7 = js.Null()
	}
	input := __callback7
	_this.Value_JS.Set("onerror", input)
}

// Onloadend returning attribute 'onloadend' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) Onloadend() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadend")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnloadend setting attribute 'onloadend' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) SetOnloadend(value *domcore.EventHandler) {
	var __callback8 js.Value
	if value != nil {
		__callback8 = (*value).Value
	} else {
		__callback8 = js.Null()
	}
	input := __callback8
	_this.Value_JS.Set("onloadend", input)
}

func (_this *FileReader) ReadAsArrayBuffer(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsArrayBuffer", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsBinaryString(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsBinaryString", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsText(blob *Blob, encoding *string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	if encoding != nil {
		_p1 := encoding
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("readAsText", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsDataURL(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsDataURL", _args[0:_end]...)
	return
}

func (_this *FileReader) Abort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("abort", _args[0:_end]...)
	return
}

// interface: FileReaderSync
type FileReaderSync struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileReaderSync) JSValue() js.Value {
	return _this.Value_JS
}

// FileReaderSyncFromJS is casting a js.Wrapper into FileReaderSync.
func FileReaderSyncFromJS(value js.Wrapper) *FileReaderSync {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileReaderSync{}
	ret.Value_JS = input
	return ret
}

func NewFileReaderSync() (_result *FileReaderSync) {
	_klass := js.Global().Get("FileReaderSync")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FileReaderSync // javascript: FileReaderSync _what_return_name
	)
	_converted = FileReaderSyncFromJS(_returned)
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsArrayBuffer(blob *Blob) (_result *javascript.ArrayBuffer) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsArrayBuffer", _args[0:_end]...)
	var (
		_converted *javascript.ArrayBuffer // javascript: ArrayBuffer _what_return_name
	)
	_converted = javascript.ArrayBufferFromJS(_returned)
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsBinaryString(blob *Blob) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsBinaryString", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsText(blob *Blob, encoding *string) (_result string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	if encoding != nil {
		_p1 := encoding
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("readAsText", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsDataURL(blob *Blob) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsDataURL", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}
