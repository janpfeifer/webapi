// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package url

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/html/media"
)

// using following types:
// file.Blob
// media.MediaSource

// source idl files:
// url.idl

// transform files:
// url.go.md

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

// interface: URL
type URL struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URL) JSValue() js.Value {
	return _this.Value_JS
}

// URLFromJS is casting a js.Wrapper into URL.
func URLFromJS(value js.Wrapper) *URL {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URL{}
	ret.Value_JS = input
	return ret
}

func CreateObjectURL(blob *file.Blob) (_result string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("createObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func RevokeObjectURL(url string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("revokeObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_method.Invoke(_args[0:_end]...)
	return
}

func CreateObjectURL2(mediaSource *media.MediaSource) (_result string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("createObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := mediaSource.JSValue()
	_args[0] = _p0
	_end++
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func NewURL(url string, base *string) (_result *URL) {
	_klass := js.Global().Get("URL")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	if base != nil {
		_p1 := base
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *URL // javascript: URL _what_return_name
	)
	_converted = URLFromJS(_returned)
	_result = _converted
	return
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *URL) Href() string {
	var ret string
	value := _this.Value_JS.Get("href")
	ret = (value).String()
	return ret
}

// SetHref setting attribute 'href' with
// type string (idl: USVString).
func (_this *URL) SetHref(value string) {
	input := value
	_this.Value_JS.Set("href", input)
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *URL) Origin() string {
	var ret string
	value := _this.Value_JS.Get("origin")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: USVString).
func (_this *URL) Protocol() string {
	var ret string
	value := _this.Value_JS.Get("protocol")
	ret = (value).String()
	return ret
}

// SetProtocol setting attribute 'protocol' with
// type string (idl: USVString).
func (_this *URL) SetProtocol(value string) {
	input := value
	_this.Value_JS.Set("protocol", input)
}

// Username returning attribute 'username' with
// type string (idl: USVString).
func (_this *URL) Username() string {
	var ret string
	value := _this.Value_JS.Get("username")
	ret = (value).String()
	return ret
}

// SetUsername setting attribute 'username' with
// type string (idl: USVString).
func (_this *URL) SetUsername(value string) {
	input := value
	_this.Value_JS.Set("username", input)
}

// Password returning attribute 'password' with
// type string (idl: USVString).
func (_this *URL) Password() string {
	var ret string
	value := _this.Value_JS.Get("password")
	ret = (value).String()
	return ret
}

// SetPassword setting attribute 'password' with
// type string (idl: USVString).
func (_this *URL) SetPassword(value string) {
	input := value
	_this.Value_JS.Set("password", input)
}

// Host returning attribute 'host' with
// type string (idl: USVString).
func (_this *URL) Host() string {
	var ret string
	value := _this.Value_JS.Get("host")
	ret = (value).String()
	return ret
}

// SetHost setting attribute 'host' with
// type string (idl: USVString).
func (_this *URL) SetHost(value string) {
	input := value
	_this.Value_JS.Set("host", input)
}

// Hostname returning attribute 'hostname' with
// type string (idl: USVString).
func (_this *URL) Hostname() string {
	var ret string
	value := _this.Value_JS.Get("hostname")
	ret = (value).String()
	return ret
}

// SetHostname setting attribute 'hostname' with
// type string (idl: USVString).
func (_this *URL) SetHostname(value string) {
	input := value
	_this.Value_JS.Set("hostname", input)
}

// Port returning attribute 'port' with
// type string (idl: USVString).
func (_this *URL) Port() string {
	var ret string
	value := _this.Value_JS.Get("port")
	ret = (value).String()
	return ret
}

// SetPort setting attribute 'port' with
// type string (idl: USVString).
func (_this *URL) SetPort(value string) {
	input := value
	_this.Value_JS.Set("port", input)
}

// Pathname returning attribute 'pathname' with
// type string (idl: USVString).
func (_this *URL) Pathname() string {
	var ret string
	value := _this.Value_JS.Get("pathname")
	ret = (value).String()
	return ret
}

// SetPathname setting attribute 'pathname' with
// type string (idl: USVString).
func (_this *URL) SetPathname(value string) {
	input := value
	_this.Value_JS.Set("pathname", input)
}

// Search returning attribute 'search' with
// type string (idl: USVString).
func (_this *URL) Search() string {
	var ret string
	value := _this.Value_JS.Get("search")
	ret = (value).String()
	return ret
}

// SetSearch setting attribute 'search' with
// type string (idl: USVString).
func (_this *URL) SetSearch(value string) {
	input := value
	_this.Value_JS.Set("search", input)
}

// SearchParams returning attribute 'searchParams' with
// type URLSearchParams (idl: URLSearchParams).
func (_this *URL) SearchParams() *URLSearchParams {
	var ret *URLSearchParams
	value := _this.Value_JS.Get("searchParams")
	ret = URLSearchParamsFromJS(value)
	return ret
}

// Hash returning attribute 'hash' with
// type string (idl: USVString).
func (_this *URL) Hash() string {
	var ret string
	value := _this.Value_JS.Get("hash")
	ret = (value).String()
	return ret
}

// SetHash setting attribute 'hash' with
// type string (idl: USVString).
func (_this *URL) SetHash(value string) {
	input := value
	_this.Value_JS.Set("hash", input)
}

func (_this *URL) ToJSON() (_result string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("toJSON", _args[0:_end]...)
	var (
		_converted string // javascript: USVString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// interface: URLSearchParams
type URLSearchParams struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URLSearchParams) JSValue() js.Value {
	return _this.Value_JS
}

// URLSearchParamsFromJS is casting a js.Wrapper into URLSearchParams.
func URLSearchParamsFromJS(value js.Wrapper) *URLSearchParams {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URLSearchParams{}
	ret.Value_JS = input
	return ret
}

func (_this *URLSearchParams) Append(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("append", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Delete(name string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("delete", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Get(name string) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *string // javascript: USVString _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *URLSearchParams) GetAll(name string) (_result []string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getAll", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<USVString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *URLSearchParams) Has(name string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *URLSearchParams) Set(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Sort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("sort", _args[0:_end]...)
	return
}
