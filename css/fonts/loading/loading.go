// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package loading

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.FrozenArray
// javascript.Promise

// source idl files:
// css-font-loading.idl

// transform files:
// css-font-loading.go.md

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

// enum: FontFaceLoadStatus
type FontFaceLoadStatus int

const (
	UnloadedFontFaceLoadStatus FontFaceLoadStatus = iota
	LoadingFontFaceLoadStatus
	LoadedFontFaceLoadStatus
	ErrorFontFaceLoadStatus
)

var fontFaceLoadStatusToWasmTable = []string{
	"unloaded", "loading", "loaded", "error",
}

var fontFaceLoadStatusFromWasmTable = map[string]FontFaceLoadStatus{
	"unloaded": UnloadedFontFaceLoadStatus, "loading": LoadingFontFaceLoadStatus, "loaded": LoadedFontFaceLoadStatus, "error": ErrorFontFaceLoadStatus,
}

// JSValue is converting this enum into a java object
func (this *FontFaceLoadStatus) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this FontFaceLoadStatus) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(fontFaceLoadStatusToWasmTable) {
		return fontFaceLoadStatusToWasmTable[idx]
	}
	panic("unknown input value")
}

// FontFaceLoadStatusFromJS is converting a javascript value into
// a FontFaceLoadStatus enum value.
func FontFaceLoadStatusFromJS(value js.Value) FontFaceLoadStatus {
	key := value.String()
	conv, ok := fontFaceLoadStatusFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: FontFaceSetLoadStatus
type FontFaceSetLoadStatus int

const (
	LoadingFontFaceSetLoadStatus FontFaceSetLoadStatus = iota
	LoadedFontFaceSetLoadStatus
)

var fontFaceSetLoadStatusToWasmTable = []string{
	"loading", "loaded",
}

var fontFaceSetLoadStatusFromWasmTable = map[string]FontFaceSetLoadStatus{
	"loading": LoadingFontFaceSetLoadStatus, "loaded": LoadedFontFaceSetLoadStatus,
}

// JSValue is converting this enum into a java object
func (this *FontFaceSetLoadStatus) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this FontFaceSetLoadStatus) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(fontFaceSetLoadStatusToWasmTable) {
		return fontFaceSetLoadStatusToWasmTable[idx]
	}
	panic("unknown input value")
}

// FontFaceSetLoadStatusFromJS is converting a javascript value into
// a FontFaceSetLoadStatus enum value.
func FontFaceSetLoadStatusFromJS(value js.Value) FontFaceSetLoadStatus {
	key := value.String()
	conv, ok := fontFaceSetLoadStatusFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: FontFaceDescriptors
type FontFaceDescriptors struct {
	Style             string
	Weight            string
	Stretch           string
	UnicodeRange      string
	Variant           string
	FeatureSettings   string
	VariationSettings string
	Display           string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FontFaceDescriptors) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Style
	out.Set("style", value0)
	value1 := _this.Weight
	out.Set("weight", value1)
	value2 := _this.Stretch
	out.Set("stretch", value2)
	value3 := _this.UnicodeRange
	out.Set("unicodeRange", value3)
	value4 := _this.Variant
	out.Set("variant", value4)
	value5 := _this.FeatureSettings
	out.Set("featureSettings", value5)
	value6 := _this.VariationSettings
	out.Set("variationSettings", value6)
	value7 := _this.Display
	out.Set("display", value7)
	return out
}

// FontFaceDescriptorsFromJS is allocating a new
// FontFaceDescriptors object and copy all values from
// input javascript object
func FontFaceDescriptorsFromJS(value js.Wrapper) *FontFaceDescriptors {
	input := value.JSValue()
	var out FontFaceDescriptors
	var (
		value0 string // javascript: DOMString {style Style style}
		value1 string // javascript: DOMString {weight Weight weight}
		value2 string // javascript: DOMString {stretch Stretch stretch}
		value3 string // javascript: DOMString {unicodeRange UnicodeRange unicodeRange}
		value4 string // javascript: DOMString {variant Variant variant}
		value5 string // javascript: DOMString {featureSettings FeatureSettings featureSettings}
		value6 string // javascript: DOMString {variationSettings VariationSettings variationSettings}
		value7 string // javascript: DOMString {display Display display}
	)
	value0 = (input.Get("style")).String()
	out.Style = value0
	value1 = (input.Get("weight")).String()
	out.Weight = value1
	value2 = (input.Get("stretch")).String()
	out.Stretch = value2
	value3 = (input.Get("unicodeRange")).String()
	out.UnicodeRange = value3
	value4 = (input.Get("variant")).String()
	out.Variant = value4
	value5 = (input.Get("featureSettings")).String()
	out.FeatureSettings = value5
	value6 = (input.Get("variationSettings")).String()
	out.VariationSettings = value6
	value7 = (input.Get("display")).String()
	out.Display = value7
	return &out
}

// dictionary: FontFaceSetLoadEventInit
type FontFaceSetLoadEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Fontfaces  []*FontFace
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FontFaceSetLoadEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := js.Global().Get("Array").New(len(_this.Fontfaces))
	for __idx3, __seq_in3 := range _this.Fontfaces {
		__seq_out3 := __seq_in3.JSValue()
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("fontfaces", value3)
	return out
}

// FontFaceSetLoadEventInitFromJS is allocating a new
// FontFaceSetLoadEventInit object and copy all values from
// input javascript object
func FontFaceSetLoadEventInitFromJS(value js.Wrapper) *FontFaceSetLoadEventInit {
	input := value.JSValue()
	var out FontFaceSetLoadEventInit
	var (
		value0 bool        // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool        // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool        // javascript: boolean {composed Composed composed}
		value3 []*FontFace // javascript: sequence<FontFace> {fontfaces Fontfaces fontfaces}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	__length3 := input.Get("fontfaces").Length()
	__array3 := make([]*FontFace, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 *FontFace
		__seq_in3 := input.Get("fontfaces").Index(__idx3)
		__seq_out3 = FontFaceFromJS(__seq_in3)
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.Fontfaces = value3
	return &out
}

// interface: FontFace
type FontFace struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FontFace) JSValue() js.Value {
	return _this.Value_JS
}

// FontFaceFromJS is casting a js.Wrapper into FontFace.
func FontFaceFromJS(value js.Wrapper) *FontFace {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FontFace{}
	ret.Value_JS = input
	return ret
}

func NewFontFace(family string, source *Union, descriptors *FontFaceDescriptors) (_result *FontFace) {
	_klass := js.Global().Get("FontFace")
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := family
	_args[0] = _p0
	_end++
	_p1 := source.JSValue()
	_args[1] = _p1
	_end++
	if descriptors != nil {
		_p2 := descriptors.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FontFace // javascript: FontFace _what_return_name
	)
	_converted = FontFaceFromJS(_returned)
	_result = _converted
	return
}

// Family returning attribute 'family' with
// type string (idl: DOMString).
func (_this *FontFace) Family() string {
	var ret string
	value := _this.Value_JS.Get("family")
	ret = (value).String()
	return ret
}

// SetFamily setting attribute 'family' with
// type string (idl: DOMString).
func (_this *FontFace) SetFamily(value string) {
	input := value
	_this.Value_JS.Set("family", input)
}

// Style returning attribute 'style' with
// type string (idl: DOMString).
func (_this *FontFace) Style() string {
	var ret string
	value := _this.Value_JS.Get("style")
	ret = (value).String()
	return ret
}

// SetStyle setting attribute 'style' with
// type string (idl: DOMString).
func (_this *FontFace) SetStyle(value string) {
	input := value
	_this.Value_JS.Set("style", input)
}

// Weight returning attribute 'weight' with
// type string (idl: DOMString).
func (_this *FontFace) Weight() string {
	var ret string
	value := _this.Value_JS.Get("weight")
	ret = (value).String()
	return ret
}

// SetWeight setting attribute 'weight' with
// type string (idl: DOMString).
func (_this *FontFace) SetWeight(value string) {
	input := value
	_this.Value_JS.Set("weight", input)
}

// Stretch returning attribute 'stretch' with
// type string (idl: DOMString).
func (_this *FontFace) Stretch() string {
	var ret string
	value := _this.Value_JS.Get("stretch")
	ret = (value).String()
	return ret
}

// SetStretch setting attribute 'stretch' with
// type string (idl: DOMString).
func (_this *FontFace) SetStretch(value string) {
	input := value
	_this.Value_JS.Set("stretch", input)
}

// UnicodeRange returning attribute 'unicodeRange' with
// type string (idl: DOMString).
func (_this *FontFace) UnicodeRange() string {
	var ret string
	value := _this.Value_JS.Get("unicodeRange")
	ret = (value).String()
	return ret
}

// SetUnicodeRange setting attribute 'unicodeRange' with
// type string (idl: DOMString).
func (_this *FontFace) SetUnicodeRange(value string) {
	input := value
	_this.Value_JS.Set("unicodeRange", input)
}

// Variant returning attribute 'variant' with
// type string (idl: DOMString).
func (_this *FontFace) Variant() string {
	var ret string
	value := _this.Value_JS.Get("variant")
	ret = (value).String()
	return ret
}

// SetVariant setting attribute 'variant' with
// type string (idl: DOMString).
func (_this *FontFace) SetVariant(value string) {
	input := value
	_this.Value_JS.Set("variant", input)
}

// FeatureSettings returning attribute 'featureSettings' with
// type string (idl: DOMString).
func (_this *FontFace) FeatureSettings() string {
	var ret string
	value := _this.Value_JS.Get("featureSettings")
	ret = (value).String()
	return ret
}

// SetFeatureSettings setting attribute 'featureSettings' with
// type string (idl: DOMString).
func (_this *FontFace) SetFeatureSettings(value string) {
	input := value
	_this.Value_JS.Set("featureSettings", input)
}

// VariationSettings returning attribute 'variationSettings' with
// type string (idl: DOMString).
func (_this *FontFace) VariationSettings() string {
	var ret string
	value := _this.Value_JS.Get("variationSettings")
	ret = (value).String()
	return ret
}

// SetVariationSettings setting attribute 'variationSettings' with
// type string (idl: DOMString).
func (_this *FontFace) SetVariationSettings(value string) {
	input := value
	_this.Value_JS.Set("variationSettings", input)
}

// Display returning attribute 'display' with
// type string (idl: DOMString).
func (_this *FontFace) Display() string {
	var ret string
	value := _this.Value_JS.Get("display")
	ret = (value).String()
	return ret
}

// SetDisplay setting attribute 'display' with
// type string (idl: DOMString).
func (_this *FontFace) SetDisplay(value string) {
	input := value
	_this.Value_JS.Set("display", input)
}

// Status returning attribute 'status' with
// type FontFaceLoadStatus (idl: FontFaceLoadStatus).
func (_this *FontFace) Status() FontFaceLoadStatus {
	var ret FontFaceLoadStatus
	value := _this.Value_JS.Get("status")
	ret = FontFaceLoadStatusFromJS(value)
	return ret
}

// Loaded returning attribute 'loaded' with
// type javascript.Promise (idl: Promise).
func (_this *FontFace) Loaded() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("loaded")
	ret = javascript.PromiseFromJS(value)
	return ret
}

func (_this *FontFace) Load() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("load", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: FontFaceSet
type FontFaceSet struct {
	domcore.EventTarget
}

// FontFaceSetFromJS is casting a js.Wrapper into FontFaceSet.
func FontFaceSetFromJS(value js.Wrapper) *FontFaceSet {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FontFaceSet{}
	ret.Value_JS = input
	return ret
}

func NewFontFaceSet(initialFaces []*FontFace) (_result *FontFaceSet) {
	_klass := js.Global().Get("FontFaceSet")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(initialFaces))
	for __idx0, __seq_in0 := range initialFaces {
		__seq_out0 := __seq_in0.JSValue()
		_p0.SetIndex(__idx0, __seq_out0)
	}
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FontFaceSet // javascript: FontFaceSet _what_return_name
	)
	_converted = FontFaceSetFromJS(_returned)
	_result = _converted
	return
}

// Onloading returning attribute 'onloading' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) Onloading() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloading")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnloading setting attribute 'onloading' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) SetOnloading(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onloading", input)
}

// Onloadingdone returning attribute 'onloadingdone' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) Onloadingdone() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadingdone")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnloadingdone setting attribute 'onloadingdone' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) SetOnloadingdone(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onloadingdone", input)
}

// Onloadingerror returning attribute 'onloadingerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) Onloadingerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadingerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnloadingerror setting attribute 'onloadingerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FontFaceSet) SetOnloadingerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onloadingerror", input)
}

// Ready returning attribute 'ready' with
// type javascript.Promise (idl: Promise).
func (_this *FontFaceSet) Ready() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("ready")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// Status returning attribute 'status' with
// type FontFaceSetLoadStatus (idl: FontFaceSetLoadStatus).
func (_this *FontFaceSet) Status() FontFaceSetLoadStatus {
	var ret FontFaceSetLoadStatus
	value := _this.Value_JS.Get("status")
	ret = FontFaceSetLoadStatusFromJS(value)
	return ret
}

func (_this *FontFaceSet) Add(font *FontFace) (_result *FontFaceSet) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := font.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("add", _args[0:_end]...)
	var (
		_converted *FontFaceSet // javascript: FontFaceSet _what_return_name
	)
	_converted = FontFaceSetFromJS(_returned)
	_result = _converted
	return
}

func (_this *FontFaceSet) Delete(font *FontFace) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := font.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *FontFaceSet) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("clear", _args[0:_end]...)
	return
}

func (_this *FontFaceSet) Load(font string, text *string) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := font
	_args[0] = _p0
	_end++
	if text != nil {
		_p1 := text
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("load", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *FontFaceSet) Check(font string, text *string) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := font
	_args[0] = _p0
	_end++
	if text != nil {
		_p1 := text
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("check", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: FontFaceSetLoadEvent
type FontFaceSetLoadEvent struct {
	domcore.Event
}

// FontFaceSetLoadEventFromJS is casting a js.Wrapper into FontFaceSetLoadEvent.
func FontFaceSetLoadEventFromJS(value js.Wrapper) *FontFaceSetLoadEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FontFaceSetLoadEvent{}
	ret.Value_JS = input
	return ret
}

func NewFontFaceSetLoadEvent(_type string, eventInitDict *FontFaceSetLoadEventInit) (_result *FontFaceSetLoadEvent) {
	_klass := js.Global().Get("FontFaceSetLoadEvent")
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
		_converted *FontFaceSetLoadEvent // javascript: FontFaceSetLoadEvent _what_return_name
	)
	_converted = FontFaceSetLoadEventFromJS(_returned)
	_result = _converted
	return
}

// Fontfaces returning attribute 'fontfaces' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *FontFaceSetLoadEvent) Fontfaces() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("fontfaces")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// interface: FontFaceSource
type FontFaceSource struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FontFaceSource) JSValue() js.Value {
	return _this.Value_JS
}

// FontFaceSourceFromJS is casting a js.Wrapper into FontFaceSource.
func FontFaceSourceFromJS(value js.Wrapper) *FontFaceSource {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FontFaceSource{}
	ret.Value_JS = input
	return ret
}

// Fonts returning attribute 'fonts' with
// type FontFaceSet (idl: FontFaceSet).
func (_this *FontFaceSource) Fonts() *FontFaceSet {
	var ret *FontFaceSet
	value := _this.Value_JS.Get("fonts")
	ret = FontFaceSetFromJS(value)
	return ret
}
