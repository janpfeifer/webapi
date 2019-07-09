// Code generated by webidl-bind. DO NOT EDIT.

package fontmetrics

import "syscall/js"

import (
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// javascript.FrozenArray

// source idl files:
// font-metrics-api.idl

// transform files:
// font-metrics-api.go.md

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

// interface: Baseline
type Baseline struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Baseline) JSValue() js.Value {
	return _this.Value_JS
}

// BaselineFromJS is casting a js.Wrapper into Baseline.
func BaselineFromJS(value js.Wrapper) *Baseline {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Baseline{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *Baseline) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Value returning attribute 'value' with
// type float64 (idl: double).
func (_this *Baseline) Value() float64 {
	var ret float64
	value := _this.Value_JS.Get("value")
	ret = (value).Float()
	return ret
}

// interface: Font
type Font struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Font) JSValue() js.Value {
	return _this.Value_JS
}

// FontFromJS is casting a js.Wrapper into Font.
func FontFromJS(value js.Wrapper) *Font {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Font{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *Font) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// GlyphsRendered returning attribute 'glyphsRendered' with
// type uint (idl: unsigned long).
func (_this *Font) GlyphsRendered() uint {
	var ret uint
	value := _this.Value_JS.Get("glyphsRendered")
	ret = (uint)((value).Int())
	return ret
}

// interface: FontMetrics
type FontMetrics struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FontMetrics) JSValue() js.Value {
	return _this.Value_JS
}

// FontMetricsFromJS is casting a js.Wrapper into FontMetrics.
func FontMetricsFromJS(value js.Wrapper) *FontMetrics {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FontMetrics{}
	ret.Value_JS = input
	return ret
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *FontMetrics) Width() float64 {
	var ret float64
	value := _this.Value_JS.Get("width")
	ret = (value).Float()
	return ret
}

// Advances returning attribute 'advances' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *FontMetrics) Advances() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("advances")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// BoundingBoxLeft returning attribute 'boundingBoxLeft' with
// type float64 (idl: double).
func (_this *FontMetrics) BoundingBoxLeft() float64 {
	var ret float64
	value := _this.Value_JS.Get("boundingBoxLeft")
	ret = (value).Float()
	return ret
}

// BoundingBoxRight returning attribute 'boundingBoxRight' with
// type float64 (idl: double).
func (_this *FontMetrics) BoundingBoxRight() float64 {
	var ret float64
	value := _this.Value_JS.Get("boundingBoxRight")
	ret = (value).Float()
	return ret
}

// Height returning attribute 'height' with
// type float64 (idl: double).
func (_this *FontMetrics) Height() float64 {
	var ret float64
	value := _this.Value_JS.Get("height")
	ret = (value).Float()
	return ret
}

// EmHeightAscent returning attribute 'emHeightAscent' with
// type float64 (idl: double).
func (_this *FontMetrics) EmHeightAscent() float64 {
	var ret float64
	value := _this.Value_JS.Get("emHeightAscent")
	ret = (value).Float()
	return ret
}

// EmHeightDescent returning attribute 'emHeightDescent' with
// type float64 (idl: double).
func (_this *FontMetrics) EmHeightDescent() float64 {
	var ret float64
	value := _this.Value_JS.Get("emHeightDescent")
	ret = (value).Float()
	return ret
}

// BoundingBoxAscent returning attribute 'boundingBoxAscent' with
// type float64 (idl: double).
func (_this *FontMetrics) BoundingBoxAscent() float64 {
	var ret float64
	value := _this.Value_JS.Get("boundingBoxAscent")
	ret = (value).Float()
	return ret
}

// BoundingBoxDescent returning attribute 'boundingBoxDescent' with
// type float64 (idl: double).
func (_this *FontMetrics) BoundingBoxDescent() float64 {
	var ret float64
	value := _this.Value_JS.Get("boundingBoxDescent")
	ret = (value).Float()
	return ret
}

// FontBoundingBoxAscent returning attribute 'fontBoundingBoxAscent' with
// type float64 (idl: double).
func (_this *FontMetrics) FontBoundingBoxAscent() float64 {
	var ret float64
	value := _this.Value_JS.Get("fontBoundingBoxAscent")
	ret = (value).Float()
	return ret
}

// FontBoundingBoxDescent returning attribute 'fontBoundingBoxDescent' with
// type float64 (idl: double).
func (_this *FontMetrics) FontBoundingBoxDescent() float64 {
	var ret float64
	value := _this.Value_JS.Get("fontBoundingBoxDescent")
	ret = (value).Float()
	return ret
}

// DominantBaseline returning attribute 'dominantBaseline' with
// type Baseline (idl: Baseline).
func (_this *FontMetrics) DominantBaseline() *Baseline {
	var ret *Baseline
	value := _this.Value_JS.Get("dominantBaseline")
	ret = BaselineFromJS(value)
	return ret
}

// Baselines returning attribute 'baselines' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *FontMetrics) Baselines() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("baselines")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Fonts returning attribute 'fonts' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *FontMetrics) Fonts() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("fonts")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}
