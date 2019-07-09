// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package fonts

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/css/cssom"
)

// using following types:
// cssom.CSSRule
// cssom.CSSStyleDeclaration

// source idl files:
// css-fonts.idl

// transform files:
// css-fonts.go.md

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

// interface: CSSFontFaceRule
type CSSFontFaceRule struct {
	cssom.CSSRule
}

// CSSFontFaceRuleFromJS is casting a js.Wrapper into CSSFontFaceRule.
func CSSFontFaceRuleFromJS(value js.Wrapper) *CSSFontFaceRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFaceRule{}
	ret.Value_JS = input
	return ret
}

// Style returning attribute 'style' with
// type cssom.CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSFontFaceRule) Style() *cssom.CSSStyleDeclaration {
	var ret *cssom.CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = cssom.CSSStyleDeclarationFromJS(value)
	return ret
}

// interface: CSSFontFeatureValuesMap
type CSSFontFeatureValuesMap struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontFeatureValuesMap) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontFeatureValuesMapFromJS is casting a js.Wrapper into CSSFontFeatureValuesMap.
func CSSFontFeatureValuesMapFromJS(value js.Wrapper) *CSSFontFeatureValuesMap {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesMap{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontFeatureValuesMap) Set(featureValueName string, values *Union) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := featureValueName
	_args[0] = _p0
	_end++
	_p1 := values.JSValue()
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

// interface: CSSFontFeatureValuesRule
type CSSFontFeatureValuesRule struct {
	cssom.CSSRule
}

// CSSFontFeatureValuesRuleFromJS is casting a js.Wrapper into CSSFontFeatureValuesRule.
func CSSFontFeatureValuesRuleFromJS(value js.Wrapper) *CSSFontFeatureValuesRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesRule{}
	ret.Value_JS = input
	return ret
}

// FontFamily returning attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontFeatureValuesRule) FontFamily() string {
	var ret string
	value := _this.Value_JS.Get("fontFamily")
	ret = (value).String()
	return ret
}

// SetFontFamily setting attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontFeatureValuesRule) SetFontFamily(value string) {
	input := value
	_this.Value_JS.Set("fontFamily", input)
}

// Annotation returning attribute 'annotation' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Annotation() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("annotation")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Ornaments returning attribute 'ornaments' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Ornaments() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("ornaments")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Stylistic returning attribute 'stylistic' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Stylistic() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("stylistic")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Swash returning attribute 'swash' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Swash() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("swash")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// CharacterVariant returning attribute 'characterVariant' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) CharacterVariant() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("characterVariant")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Styleset returning attribute 'styleset' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Styleset() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("styleset")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// interface: CSSFontPaletteValuesRule
type CSSFontPaletteValuesRule struct {
	cssom.CSSRule
}

// CSSFontPaletteValuesRuleFromJS is casting a js.Wrapper into CSSFontPaletteValuesRule.
func CSSFontPaletteValuesRuleFromJS(value js.Wrapper) *CSSFontPaletteValuesRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontPaletteValuesRule{}
	ret.Value_JS = input
	return ret
}

// FontFamily returning attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) FontFamily() string {
	var ret string
	value := _this.Value_JS.Get("fontFamily")
	ret = (value).String()
	return ret
}

// SetFontFamily setting attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) SetFontFamily(value string) {
	input := value
	_this.Value_JS.Set("fontFamily", input)
}

// BasePalette returning attribute 'basePalette' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) BasePalette() string {
	var ret string
	value := _this.Value_JS.Get("basePalette")
	ret = (value).String()
	return ret
}

// SetBasePalette setting attribute 'basePalette' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) SetBasePalette(value string) {
	input := value
	_this.Value_JS.Set("basePalette", input)
}
