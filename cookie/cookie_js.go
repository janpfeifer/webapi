// Code generated by webidl-bind. DO NOT EDIT.

package cookie

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// domcore.ExtendableEvent
// javascript.Promise

// source idl files:
// cookie-store.idl

// transform files:
// cookie-store.go.md

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

// enum: CookieMatchType
type CookieMatchType int

const (
	EqualsCookieMatchType CookieMatchType = iota
	StartsWithCookieMatchType
)

var cookieMatchTypeToWasmTable = []string{
	"equals", "starts-with",
}

var cookieMatchTypeFromWasmTable = map[string]CookieMatchType{
	"equals": EqualsCookieMatchType, "starts-with": StartsWithCookieMatchType,
}

// JSValue is converting this enum into a java object
func (this *CookieMatchType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this CookieMatchType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(cookieMatchTypeToWasmTable) {
		return cookieMatchTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// CookieMatchTypeFromJS is converting a javascript value into
// a CookieMatchType enum value.
func CookieMatchTypeFromJS(value js.Value) CookieMatchType {
	key := value.String()
	conv, ok := cookieMatchTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: CookieSameSite
type CookieSameSite int

const (
	StrictCookieSameSite CookieSameSite = iota
	LaxCookieSameSite
	UnrestrictedCookieSameSite
)

var cookieSameSiteToWasmTable = []string{
	"strict", "lax", "unrestricted",
}

var cookieSameSiteFromWasmTable = map[string]CookieSameSite{
	"strict": StrictCookieSameSite, "lax": LaxCookieSameSite, "unrestricted": UnrestrictedCookieSameSite,
}

// JSValue is converting this enum into a java object
func (this *CookieSameSite) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this CookieSameSite) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(cookieSameSiteToWasmTable) {
		return cookieSameSiteToWasmTable[idx]
	}
	panic("unknown input value")
}

// CookieSameSiteFromJS is converting a javascript value into
// a CookieSameSite enum value.
func CookieSameSiteFromJS(value js.Value) CookieSameSite {
	key := value.String()
	conv, ok := cookieSameSiteFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: CookieChangeEventInit
type CookieChangeEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Changed    []*CookieListItem
	Deleted    []*CookieListItem
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieChangeEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := js.Global().Get("Array").New(len(_this.Changed))
	for __idx3, __seq_in3 := range _this.Changed {
		__seq_out3 := __seq_in3.JSValue()
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("changed", value3)
	value4 := js.Global().Get("Array").New(len(_this.Deleted))
	for __idx4, __seq_in4 := range _this.Deleted {
		__seq_out4 := __seq_in4.JSValue()
		value4.SetIndex(__idx4, __seq_out4)
	}
	out.Set("deleted", value4)
	return out
}

// CookieChangeEventInitFromJS is allocating a new
// CookieChangeEventInit object and copy all values from
// input javascript object
func CookieChangeEventInitFromJS(value js.Wrapper) *CookieChangeEventInit {
	input := value.JSValue()
	var out CookieChangeEventInit
	var (
		value0 bool              // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool              // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool              // javascript: boolean {composed Composed composed}
		value3 []*CookieListItem // javascript: sequence<CookieListItem> {changed Changed changed}
		value4 []*CookieListItem // javascript: sequence<CookieListItem> {deleted Deleted deleted}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	__length3 := input.Get("changed").Length()
	__array3 := make([]*CookieListItem, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 *CookieListItem
		__seq_in3 := input.Get("changed").Index(__idx3)
		__seq_out3 = CookieListItemFromJS(__seq_in3)
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.Changed = value3
	__length4 := input.Get("deleted").Length()
	__array4 := make([]*CookieListItem, __length4, __length4)
	for __idx4 := 0; __idx4 < __length4; __idx4++ {
		var __seq_out4 *CookieListItem
		__seq_in4 := input.Get("deleted").Index(__idx4)
		__seq_out4 = CookieListItemFromJS(__seq_in4)
		__array4[__idx4] = __seq_out4
	}
	value4 = __array4
	out.Deleted = value4
	return &out
}

// dictionary: CookieListItem
type CookieListItem struct {
	Name     string
	Value    string
	Domain   *string
	Path     string
	Expires  *int
	Secure   bool
	SameSite CookieSameSite
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieListItem) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Name
	out.Set("name", value0)
	value1 := _this.Value
	out.Set("value", value1)
	value2 := _this.Domain
	out.Set("domain", value2)
	value3 := _this.Path
	out.Set("path", value3)
	value4 := _this.Expires
	out.Set("expires", value4)
	value5 := _this.Secure
	out.Set("secure", value5)
	value6 := _this.SameSite.JSValue()
	out.Set("sameSite", value6)
	return out
}

// CookieListItemFromJS is allocating a new
// CookieListItem object and copy all values from
// input javascript object
func CookieListItemFromJS(value js.Wrapper) *CookieListItem {
	input := value.JSValue()
	var out CookieListItem
	var (
		value0 string         // javascript: USVString {name Name name}
		value1 string         // javascript: USVString {value Value value}
		value2 *string        // javascript: USVString {domain Domain domain}
		value3 string         // javascript: USVString {path Path path}
		value4 *int           // javascript: unsigned long long {expires Expires expires}
		value5 bool           // javascript: boolean {secure Secure secure}
		value6 CookieSameSite // javascript: CookieSameSite {sameSite SameSite sameSite}
	)
	value0 = (input.Get("name")).String()
	out.Name = value0
	value1 = (input.Get("value")).String()
	out.Value = value1
	if input.Get("domain").Type() != js.TypeNull {
		__tmp := (input.Get("domain")).String()
		value2 = &__tmp
	}
	out.Domain = value2
	value3 = (input.Get("path")).String()
	out.Path = value3
	if input.Get("expires").Type() != js.TypeNull {
		__tmp := (input.Get("expires")).Int()
		value4 = &__tmp
	}
	out.Expires = value4
	value5 = (input.Get("secure")).Bool()
	out.Secure = value5
	value6 = CookieSameSiteFromJS(input.Get("sameSite"))
	out.SameSite = value6
	return &out
}

// dictionary: CookieStoreDeleteOptions
type CookieStoreDeleteOptions struct {
	Name   string
	Domain *string
	Path   string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieStoreDeleteOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Name
	out.Set("name", value0)
	value1 := _this.Domain
	out.Set("domain", value1)
	value2 := _this.Path
	out.Set("path", value2)
	return out
}

// CookieStoreDeleteOptionsFromJS is allocating a new
// CookieStoreDeleteOptions object and copy all values from
// input javascript object
func CookieStoreDeleteOptionsFromJS(value js.Wrapper) *CookieStoreDeleteOptions {
	input := value.JSValue()
	var out CookieStoreDeleteOptions
	var (
		value0 string  // javascript: USVString {name Name name}
		value1 *string // javascript: USVString {domain Domain domain}
		value2 string  // javascript: USVString {path Path path}
	)
	value0 = (input.Get("name")).String()
	out.Name = value0
	if input.Get("domain").Type() != js.TypeNull {
		__tmp := (input.Get("domain")).String()
		value1 = &__tmp
	}
	out.Domain = value1
	value2 = (input.Get("path")).String()
	out.Path = value2
	return &out
}

// dictionary: CookieStoreGetOptions
type CookieStoreGetOptions struct {
	Name      string
	Url       string
	MatchType CookieMatchType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieStoreGetOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Name
	out.Set("name", value0)
	value1 := _this.Url
	out.Set("url", value1)
	value2 := _this.MatchType.JSValue()
	out.Set("matchType", value2)
	return out
}

// CookieStoreGetOptionsFromJS is allocating a new
// CookieStoreGetOptions object and copy all values from
// input javascript object
func CookieStoreGetOptionsFromJS(value js.Wrapper) *CookieStoreGetOptions {
	input := value.JSValue()
	var out CookieStoreGetOptions
	var (
		value0 string          // javascript: USVString {name Name name}
		value1 string          // javascript: USVString {url Url url}
		value2 CookieMatchType // javascript: CookieMatchType {matchType MatchType matchType}
	)
	value0 = (input.Get("name")).String()
	out.Name = value0
	value1 = (input.Get("url")).String()
	out.Url = value1
	value2 = CookieMatchTypeFromJS(input.Get("matchType"))
	out.MatchType = value2
	return &out
}

// dictionary: CookieStoreSetExtraOptions
type CookieStoreSetExtraOptions struct {
	Expires  *int
	Domain   *string
	Path     string
	Secure   bool
	SameSite CookieSameSite
	Name     string
	Value    string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieStoreSetExtraOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Expires
	out.Set("expires", value0)
	value1 := _this.Domain
	out.Set("domain", value1)
	value2 := _this.Path
	out.Set("path", value2)
	value3 := _this.Secure
	out.Set("secure", value3)
	value4 := _this.SameSite.JSValue()
	out.Set("sameSite", value4)
	value5 := _this.Name
	out.Set("name", value5)
	value6 := _this.Value
	out.Set("value", value6)
	return out
}

// CookieStoreSetExtraOptionsFromJS is allocating a new
// CookieStoreSetExtraOptions object and copy all values from
// input javascript object
func CookieStoreSetExtraOptionsFromJS(value js.Wrapper) *CookieStoreSetExtraOptions {
	input := value.JSValue()
	var out CookieStoreSetExtraOptions
	var (
		value0 *int           // javascript: unsigned long long {expires Expires expires}
		value1 *string        // javascript: USVString {domain Domain domain}
		value2 string         // javascript: USVString {path Path path}
		value3 bool           // javascript: boolean {secure Secure secure}
		value4 CookieSameSite // javascript: CookieSameSite {sameSite SameSite sameSite}
		value5 string         // javascript: USVString {name Name name}
		value6 string         // javascript: USVString {value Value value}
	)
	if input.Get("expires").Type() != js.TypeNull {
		__tmp := (input.Get("expires")).Int()
		value0 = &__tmp
	}
	out.Expires = value0
	if input.Get("domain").Type() != js.TypeNull {
		__tmp := (input.Get("domain")).String()
		value1 = &__tmp
	}
	out.Domain = value1
	value2 = (input.Get("path")).String()
	out.Path = value2
	value3 = (input.Get("secure")).Bool()
	out.Secure = value3
	value4 = CookieSameSiteFromJS(input.Get("sameSite"))
	out.SameSite = value4
	value5 = (input.Get("name")).String()
	out.Name = value5
	value6 = (input.Get("value")).String()
	out.Value = value6
	return &out
}

// dictionary: CookieStoreSetOptions
type CookieStoreSetOptions struct {
	Expires  *int
	Domain   *string
	Path     string
	Secure   bool
	SameSite CookieSameSite
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CookieStoreSetOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Expires
	out.Set("expires", value0)
	value1 := _this.Domain
	out.Set("domain", value1)
	value2 := _this.Path
	out.Set("path", value2)
	value3 := _this.Secure
	out.Set("secure", value3)
	value4 := _this.SameSite.JSValue()
	out.Set("sameSite", value4)
	return out
}

// CookieStoreSetOptionsFromJS is allocating a new
// CookieStoreSetOptions object and copy all values from
// input javascript object
func CookieStoreSetOptionsFromJS(value js.Wrapper) *CookieStoreSetOptions {
	input := value.JSValue()
	var out CookieStoreSetOptions
	var (
		value0 *int           // javascript: unsigned long long {expires Expires expires}
		value1 *string        // javascript: USVString {domain Domain domain}
		value2 string         // javascript: USVString {path Path path}
		value3 bool           // javascript: boolean {secure Secure secure}
		value4 CookieSameSite // javascript: CookieSameSite {sameSite SameSite sameSite}
	)
	if input.Get("expires").Type() != js.TypeNull {
		__tmp := (input.Get("expires")).Int()
		value0 = &__tmp
	}
	out.Expires = value0
	if input.Get("domain").Type() != js.TypeNull {
		__tmp := (input.Get("domain")).String()
		value1 = &__tmp
	}
	out.Domain = value1
	value2 = (input.Get("path")).String()
	out.Path = value2
	value3 = (input.Get("secure")).Bool()
	out.Secure = value3
	value4 = CookieSameSiteFromJS(input.Get("sameSite"))
	out.SameSite = value4
	return &out
}

// dictionary: ExtendableCookieChangeEventInit
type ExtendableCookieChangeEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Changed    []*CookieListItem
	Deleted    []*CookieListItem
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ExtendableCookieChangeEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := js.Global().Get("Array").New(len(_this.Changed))
	for __idx3, __seq_in3 := range _this.Changed {
		__seq_out3 := __seq_in3.JSValue()
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("changed", value3)
	value4 := js.Global().Get("Array").New(len(_this.Deleted))
	for __idx4, __seq_in4 := range _this.Deleted {
		__seq_out4 := __seq_in4.JSValue()
		value4.SetIndex(__idx4, __seq_out4)
	}
	out.Set("deleted", value4)
	return out
}

// ExtendableCookieChangeEventInitFromJS is allocating a new
// ExtendableCookieChangeEventInit object and copy all values from
// input javascript object
func ExtendableCookieChangeEventInitFromJS(value js.Wrapper) *ExtendableCookieChangeEventInit {
	input := value.JSValue()
	var out ExtendableCookieChangeEventInit
	var (
		value0 bool              // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool              // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool              // javascript: boolean {composed Composed composed}
		value3 []*CookieListItem // javascript: sequence<CookieListItem> {changed Changed changed}
		value4 []*CookieListItem // javascript: sequence<CookieListItem> {deleted Deleted deleted}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	__length3 := input.Get("changed").Length()
	__array3 := make([]*CookieListItem, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 *CookieListItem
		__seq_in3 := input.Get("changed").Index(__idx3)
		__seq_out3 = CookieListItemFromJS(__seq_in3)
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.Changed = value3
	__length4 := input.Get("deleted").Length()
	__array4 := make([]*CookieListItem, __length4, __length4)
	for __idx4 := 0; __idx4 < __length4; __idx4++ {
		var __seq_out4 *CookieListItem
		__seq_in4 := input.Get("deleted").Index(__idx4)
		__seq_out4 = CookieListItemFromJS(__seq_in4)
		__array4[__idx4] = __seq_out4
	}
	value4 = __array4
	out.Deleted = value4
	return &out
}

// interface: CookieChangeEvent
type CookieChangeEvent struct {
	domcore.Event
}

// CookieChangeEventFromJS is casting a js.Wrapper into CookieChangeEvent.
func CookieChangeEventFromJS(value js.Wrapper) *CookieChangeEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CookieChangeEvent{}
	ret.Value_JS = input
	return ret
}

func NewCookieChangeEvent(_type string, eventInitDict *CookieChangeEventInit) (_result *CookieChangeEvent) {
	_klass := js.Global().Get("CookieChangeEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *CookieChangeEvent // javascript: CookieChangeEvent _what_return_name
	)
	_converted = CookieChangeEventFromJS(_returned)
	_result = _converted
	return
}

// Changed returning attribute 'changed' with
// type []CookieListItem (idl: sequence<CookieListItem>).
func (_this *CookieChangeEvent) Changed() []*CookieListItem {
	var ret []*CookieListItem
	value := _this.Value_JS.Get("changed")
	__length0 := value.Length()
	__array0 := make([]*CookieListItem, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *CookieListItem
		__seq_in0 := value.Index(__idx0)
		__seq_out0 = CookieListItemFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	ret = __array0
	return ret
}

// Deleted returning attribute 'deleted' with
// type []CookieListItem (idl: sequence<CookieListItem>).
func (_this *CookieChangeEvent) Deleted() []*CookieListItem {
	var ret []*CookieListItem
	value := _this.Value_JS.Get("deleted")
	__length0 := value.Length()
	__array0 := make([]*CookieListItem, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *CookieListItem
		__seq_in0 := value.Index(__idx0)
		__seq_out0 = CookieListItemFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	ret = __array0
	return ret
}

// interface: CookieStore
type CookieStore struct {
	domcore.EventTarget
}

// CookieStoreFromJS is casting a js.Wrapper into CookieStore.
func CookieStoreFromJS(value js.Wrapper) *CookieStore {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CookieStore{}
	ret.Value_JS = input
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *CookieStore) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *CookieStore) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

func (_this *CookieStore) Get(name string) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) Get2(options *CookieStoreGetOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) GetAll(name string) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getAll", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) GetAll2(options *CookieStoreGetOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("getAll", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) Set(name string, value string, options *CookieStoreSetOptions) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) Set2(options *CookieStoreSetExtraOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := options.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) Delete(name string) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) Delete2(options *CookieStoreDeleteOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := options.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) SubscribeToChanges(subscriptions []*CookieStoreGetOptions) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(subscriptions))
	for __idx0, __seq_in0 := range subscriptions {
		__seq_out0 := __seq_in0.JSValue()
		_p0.SetIndex(__idx0, __seq_out0)
	}
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("subscribeToChanges", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *CookieStore) GetChangeSubscriptions() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getChangeSubscriptions", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: ExtendableCookieChangeEvent
type ExtendableCookieChangeEvent struct {
	domcore.ExtendableEvent
}

// ExtendableCookieChangeEventFromJS is casting a js.Wrapper into ExtendableCookieChangeEvent.
func ExtendableCookieChangeEventFromJS(value js.Wrapper) *ExtendableCookieChangeEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ExtendableCookieChangeEvent{}
	ret.Value_JS = input
	return ret
}

func NewExtendableCookieChangeEvent(_type string, eventInitDict *ExtendableCookieChangeEventInit) (_result *ExtendableCookieChangeEvent) {
	_klass := js.Global().Get("ExtendableCookieChangeEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ExtendableCookieChangeEvent // javascript: ExtendableCookieChangeEvent _what_return_name
	)
	_converted = ExtendableCookieChangeEventFromJS(_returned)
	_result = _converted
	return
}

// Changed returning attribute 'changed' with
// type []CookieListItem (idl: sequence<CookieListItem>).
func (_this *ExtendableCookieChangeEvent) Changed() []*CookieListItem {
	var ret []*CookieListItem
	value := _this.Value_JS.Get("changed")
	__length0 := value.Length()
	__array0 := make([]*CookieListItem, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *CookieListItem
		__seq_in0 := value.Index(__idx0)
		__seq_out0 = CookieListItemFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	ret = __array0
	return ret
}

// Deleted returning attribute 'deleted' with
// type []CookieListItem (idl: sequence<CookieListItem>).
func (_this *ExtendableCookieChangeEvent) Deleted() []*CookieListItem {
	var ret []*CookieListItem
	value := _this.Value_JS.Get("deleted")
	__length0 := value.Length()
	__array0 := make([]*CookieListItem, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *CookieListItem
		__seq_in0 := value.Index(__idx0)
		__seq_out0 = CookieListItemFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	ret = __array0
	return ret
}
