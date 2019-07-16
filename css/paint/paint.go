// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package paint

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/geometry"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/webidl"
	"github.com/gowebapi/webapi/worklets"
)

// using following types:
// canvas.CanvasFillRule
// canvas.CanvasGradient
// canvas.CanvasLineCap
// canvas.CanvasLineJoin
// canvas.CanvasPattern
// canvas.ImageSmoothingQuality
// canvas.Path2D
// geometry.DOMMatrix
// geometry.DOMMatrix2DInit
// webidl.VoidFunction
// worklets.WorkletGlobalScope

// source idl files:
// css-paint-api.idl

// transform files:
// css-paint-api.go.md

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

// dictionary: PaintRenderingContext2DSettings
type PaintRenderingContext2DSettings struct {
	Alpha bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PaintRenderingContext2DSettings) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Alpha
	out.Set("alpha", value0)
	return out
}

// PaintRenderingContext2DSettingsFromJS is allocating a new
// PaintRenderingContext2DSettings object and copy all values from
// input javascript object
func PaintRenderingContext2DSettingsFromJS(value js.Wrapper) *PaintRenderingContext2DSettings {
	input := value.JSValue()
	var out PaintRenderingContext2DSettings
	var (
		value0 bool // javascript: boolean {alpha Alpha alpha}
	)
	value0 = (input.Get("alpha")).Bool()
	out.Alpha = value0
	return &out
}

// class: PaintRenderingContext2D
type PaintRenderingContext2D struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PaintRenderingContext2D) JSValue() js.Value {
	return _this.Value_JS
}

// PaintRenderingContext2DFromJS is casting a js.Wrapper into PaintRenderingContext2D.
func PaintRenderingContext2DFromJS(value js.Wrapper) *PaintRenderingContext2D {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PaintRenderingContext2D{}
	ret.Value_JS = input
	return ret
}

// GlobalAlpha returning attribute 'globalAlpha' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) GlobalAlpha() float64 {
	var ret float64
	value := _this.Value_JS.Get("globalAlpha")
	ret = (value).Float()
	return ret
}

// SetGlobalAlpha setting attribute 'globalAlpha' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetGlobalAlpha(value float64) {
	input := value
	_this.Value_JS.Set("globalAlpha", input)
}

// GlobalCompositeOperation returning attribute 'globalCompositeOperation' with
// type string (idl: DOMString).
func (_this *PaintRenderingContext2D) GlobalCompositeOperation() string {
	var ret string
	value := _this.Value_JS.Get("globalCompositeOperation")
	ret = (value).String()
	return ret
}

// SetGlobalCompositeOperation setting attribute 'globalCompositeOperation' with
// type string (idl: DOMString).
func (_this *PaintRenderingContext2D) SetGlobalCompositeOperation(value string) {
	input := value
	_this.Value_JS.Set("globalCompositeOperation", input)
}

// ImageSmoothingEnabled returning attribute 'imageSmoothingEnabled' with
// type bool (idl: boolean).
func (_this *PaintRenderingContext2D) ImageSmoothingEnabled() bool {
	var ret bool
	value := _this.Value_JS.Get("imageSmoothingEnabled")
	ret = (value).Bool()
	return ret
}

// SetImageSmoothingEnabled setting attribute 'imageSmoothingEnabled' with
// type bool (idl: boolean).
func (_this *PaintRenderingContext2D) SetImageSmoothingEnabled(value bool) {
	input := value
	_this.Value_JS.Set("imageSmoothingEnabled", input)
}

// ImageSmoothingQuality returning attribute 'imageSmoothingQuality' with
// type canvas.ImageSmoothingQuality (idl: ImageSmoothingQuality).
func (_this *PaintRenderingContext2D) ImageSmoothingQuality() canvas.ImageSmoothingQuality {
	var ret canvas.ImageSmoothingQuality
	value := _this.Value_JS.Get("imageSmoothingQuality")
	ret = canvas.ImageSmoothingQualityFromJS(value)
	return ret
}

// SetImageSmoothingQuality setting attribute 'imageSmoothingQuality' with
// type canvas.ImageSmoothingQuality (idl: ImageSmoothingQuality).
func (_this *PaintRenderingContext2D) SetImageSmoothingQuality(value canvas.ImageSmoothingQuality) {
	input := value.JSValue()
	_this.Value_JS.Set("imageSmoothingQuality", input)
}

// StrokeStyle returning attribute 'strokeStyle' with
// type Union (idl: Union).
func (_this *PaintRenderingContext2D) StrokeStyle() *Union {
	var ret *Union
	value := _this.Value_JS.Get("strokeStyle")
	ret = UnionFromJS(value)
	return ret
}

// SetStrokeStyle setting attribute 'strokeStyle' with
// type Union (idl: Union).
func (_this *PaintRenderingContext2D) SetStrokeStyle(value *Union) {
	input := value.JSValue()
	_this.Value_JS.Set("strokeStyle", input)
}

// FillStyle returning attribute 'fillStyle' with
// type Union (idl: Union).
func (_this *PaintRenderingContext2D) FillStyle() *Union {
	var ret *Union
	value := _this.Value_JS.Get("fillStyle")
	ret = UnionFromJS(value)
	return ret
}

// SetFillStyle setting attribute 'fillStyle' with
// type Union (idl: Union).
func (_this *PaintRenderingContext2D) SetFillStyle(value *Union) {
	input := value.JSValue()
	_this.Value_JS.Set("fillStyle", input)
}

// ShadowOffsetX returning attribute 'shadowOffsetX' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) ShadowOffsetX() float64 {
	var ret float64
	value := _this.Value_JS.Get("shadowOffsetX")
	ret = (value).Float()
	return ret
}

// SetShadowOffsetX setting attribute 'shadowOffsetX' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetShadowOffsetX(value float64) {
	input := value
	_this.Value_JS.Set("shadowOffsetX", input)
}

// ShadowOffsetY returning attribute 'shadowOffsetY' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) ShadowOffsetY() float64 {
	var ret float64
	value := _this.Value_JS.Get("shadowOffsetY")
	ret = (value).Float()
	return ret
}

// SetShadowOffsetY setting attribute 'shadowOffsetY' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetShadowOffsetY(value float64) {
	input := value
	_this.Value_JS.Set("shadowOffsetY", input)
}

// ShadowBlur returning attribute 'shadowBlur' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) ShadowBlur() float64 {
	var ret float64
	value := _this.Value_JS.Get("shadowBlur")
	ret = (value).Float()
	return ret
}

// SetShadowBlur setting attribute 'shadowBlur' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetShadowBlur(value float64) {
	input := value
	_this.Value_JS.Set("shadowBlur", input)
}

// ShadowColor returning attribute 'shadowColor' with
// type string (idl: DOMString).
func (_this *PaintRenderingContext2D) ShadowColor() string {
	var ret string
	value := _this.Value_JS.Get("shadowColor")
	ret = (value).String()
	return ret
}

// SetShadowColor setting attribute 'shadowColor' with
// type string (idl: DOMString).
func (_this *PaintRenderingContext2D) SetShadowColor(value string) {
	input := value
	_this.Value_JS.Set("shadowColor", input)
}

// LineWidth returning attribute 'lineWidth' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) LineWidth() float64 {
	var ret float64
	value := _this.Value_JS.Get("lineWidth")
	ret = (value).Float()
	return ret
}

// SetLineWidth setting attribute 'lineWidth' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetLineWidth(value float64) {
	input := value
	_this.Value_JS.Set("lineWidth", input)
}

// LineCap returning attribute 'lineCap' with
// type canvas.CanvasLineCap (idl: CanvasLineCap).
func (_this *PaintRenderingContext2D) LineCap() canvas.CanvasLineCap {
	var ret canvas.CanvasLineCap
	value := _this.Value_JS.Get("lineCap")
	ret = canvas.CanvasLineCapFromJS(value)
	return ret
}

// SetLineCap setting attribute 'lineCap' with
// type canvas.CanvasLineCap (idl: CanvasLineCap).
func (_this *PaintRenderingContext2D) SetLineCap(value canvas.CanvasLineCap) {
	input := value.JSValue()
	_this.Value_JS.Set("lineCap", input)
}

// LineJoin returning attribute 'lineJoin' with
// type canvas.CanvasLineJoin (idl: CanvasLineJoin).
func (_this *PaintRenderingContext2D) LineJoin() canvas.CanvasLineJoin {
	var ret canvas.CanvasLineJoin
	value := _this.Value_JS.Get("lineJoin")
	ret = canvas.CanvasLineJoinFromJS(value)
	return ret
}

// SetLineJoin setting attribute 'lineJoin' with
// type canvas.CanvasLineJoin (idl: CanvasLineJoin).
func (_this *PaintRenderingContext2D) SetLineJoin(value canvas.CanvasLineJoin) {
	input := value.JSValue()
	_this.Value_JS.Set("lineJoin", input)
}

// MiterLimit returning attribute 'miterLimit' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) MiterLimit() float64 {
	var ret float64
	value := _this.Value_JS.Get("miterLimit")
	ret = (value).Float()
	return ret
}

// SetMiterLimit setting attribute 'miterLimit' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetMiterLimit(value float64) {
	input := value
	_this.Value_JS.Set("miterLimit", input)
}

// LineDashOffset returning attribute 'lineDashOffset' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) LineDashOffset() float64 {
	var ret float64
	value := _this.Value_JS.Get("lineDashOffset")
	ret = (value).Float()
	return ret
}

// SetLineDashOffset setting attribute 'lineDashOffset' with
// type float64 (idl: unrestricted double).
func (_this *PaintRenderingContext2D) SetLineDashOffset(value float64) {
	input := value
	_this.Value_JS.Set("lineDashOffset", input)
}

func (_this *PaintRenderingContext2D) Save() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("save", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Restore() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("restore", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Scale(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("scale", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Rotate(angle float64) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := angle
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("rotate", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Translate(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("translate", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := a
	_args[0] = _p0
	_end++
	_p1 := b
	_args[1] = _p1
	_end++
	_p2 := c
	_args[2] = _p2
	_end++
	_p3 := d
	_args[3] = _p3
	_end++
	_p4 := e
	_args[4] = _p4
	_end++
	_p5 := f
	_args[5] = _p5
	_end++
	_this.Value_JS.Call("transform", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) GetTransform() (_result *geometry.DOMMatrix) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getTransform", _args[0:_end]...)
	var (
		_converted *geometry.DOMMatrix // javascript: DOMMatrix _what_return_name
	)
	_converted = geometry.DOMMatrixFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := a
	_args[0] = _p0
	_end++
	_p1 := b
	_args[1] = _p1
	_end++
	_p2 := c
	_args[2] = _p2
	_end++
	_p3 := d
	_args[3] = _p3
	_end++
	_p4 := e
	_args[4] = _p4
	_end++
	_p5 := f
	_args[5] = _p5
	_end++
	_this.Value_JS.Call("setTransform", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) SetTransform2(transform *geometry.DOMMatrix2DInit) {
	var (
		_args [1]interface{}
		_end  int
	)
	if transform != nil {
		_p0 := transform.JSValue()
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("setTransform", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) ResetTransform() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("resetTransform", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (_result *canvas.CanvasGradient) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := x0
	_args[0] = _p0
	_end++
	_p1 := y0
	_args[1] = _p1
	_end++
	_p2 := x1
	_args[2] = _p2
	_end++
	_p3 := y1
	_args[3] = _p3
	_end++
	_returned := _this.Value_JS.Call("createLinearGradient", _args[0:_end]...)
	var (
		_converted *canvas.CanvasGradient // javascript: CanvasGradient _what_return_name
	)
	_converted = canvas.CanvasGradientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (_result *canvas.CanvasGradient) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := x0
	_args[0] = _p0
	_end++
	_p1 := y0
	_args[1] = _p1
	_end++
	_p2 := r0
	_args[2] = _p2
	_end++
	_p3 := x1
	_args[3] = _p3
	_end++
	_p4 := y1
	_args[4] = _p4
	_end++
	_p5 := r1
	_args[5] = _p5
	_end++
	_returned := _this.Value_JS.Call("createRadialGradient", _args[0:_end]...)
	var (
		_converted *canvas.CanvasGradient // javascript: CanvasGradient _what_return_name
	)
	_converted = canvas.CanvasGradientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) CreatePattern(image *Union, repetition string) (_result *canvas.CanvasPattern) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := repetition
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("createPattern", _args[0:_end]...)
	var (
		_converted *canvas.CanvasPattern // javascript: CanvasPattern _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = canvas.CanvasPatternFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := w
	_args[2] = _p2
	_end++
	_p3 := h
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("clearRect", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := w
	_args[2] = _p2
	_end++
	_p3 := h
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("fillRect", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := w
	_args[2] = _p2
	_end++
	_p3 := h
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("strokeRect", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) BeginPath() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("beginPath", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Fill(fillRule *canvas.CanvasFillRule) {
	var (
		_args [1]interface{}
		_end  int
	)
	if fillRule != nil {
		_p0 := fillRule.JSValue()
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("fill", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Fill2(path *canvas.Path2D, fillRule *canvas.CanvasFillRule) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := path.JSValue()
	_args[0] = _p0
	_end++
	if fillRule != nil {
		_p1 := fillRule.JSValue()
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("fill", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Stroke() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("stroke", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Stroke2(path *canvas.Path2D) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := path.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("stroke", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Clip(fillRule *canvas.CanvasFillRule) {
	var (
		_args [1]interface{}
		_end  int
	)
	if fillRule != nil {
		_p0 := fillRule.JSValue()
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("clip", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Clip2(path *canvas.Path2D, fillRule *canvas.CanvasFillRule) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := path.JSValue()
	_args[0] = _p0
	_end++
	if fillRule != nil {
		_p1 := fillRule.JSValue()
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("clip", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) IsPointInPath(x float64, y float64, fillRule *canvas.CanvasFillRule) (_result bool) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	if fillRule != nil {
		_p2 := fillRule.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("isPointInPath", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) IsPointInPath2(path *canvas.Path2D, x float64, y float64, fillRule *canvas.CanvasFillRule) (_result bool) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := path.JSValue()
	_args[0] = _p0
	_end++
	_p1 := x
	_args[1] = _p1
	_end++
	_p2 := y
	_args[2] = _p2
	_end++
	if fillRule != nil {
		_p3 := fillRule.JSValue()
		_args[3] = _p3
		_end++
	}
	_returned := _this.Value_JS.Call("isPointInPath", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) IsPointInStroke(x float64, y float64) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("isPointInStroke", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) IsPointInStroke2(path *canvas.Path2D, x float64, y float64) (_result bool) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := path.JSValue()
	_args[0] = _p0
	_end++
	_p1 := x
	_args[1] = _p1
	_end++
	_p2 := y
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("isPointInStroke", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) DrawImage(image *Union, dx float64, dy float64) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := dx
	_args[1] = _p1
	_end++
	_p2 := dy
	_args[2] = _p2
	_end++
	_this.Value_JS.Call("drawImage", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) DrawImage2(image *Union, dx float64, dy float64, dw float64, dh float64) {
	var (
		_args [5]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := dx
	_args[1] = _p1
	_end++
	_p2 := dy
	_args[2] = _p2
	_end++
	_p3 := dw
	_args[3] = _p3
	_end++
	_p4 := dh
	_args[4] = _p4
	_end++
	_this.Value_JS.Call("drawImage", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) DrawImage3(image *Union, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) {
	var (
		_args [9]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := sx
	_args[1] = _p1
	_end++
	_p2 := sy
	_args[2] = _p2
	_end++
	_p3 := sw
	_args[3] = _p3
	_end++
	_p4 := sh
	_args[4] = _p4
	_end++
	_p5 := dx
	_args[5] = _p5
	_end++
	_p6 := dy
	_args[6] = _p6
	_end++
	_p7 := dw
	_args[7] = _p7
	_end++
	_p8 := dh
	_args[8] = _p8
	_end++
	_this.Value_JS.Call("drawImage", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) SetLineDash(segments js.Value) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := segments
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("setLineDash", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) GetLineDash() (_result js.Value) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getLineDash", _args[0:_end]...)
	var (
		_converted js.Value // javascript: typed-array _what_return_name
	)
	_converted = _returned
	_result = _converted
	return
}

func (_this *PaintRenderingContext2D) ClosePath() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("closePath", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) MoveTo(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("moveTo", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) LineTo(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("lineTo", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := cpx
	_args[0] = _p0
	_end++
	_p1 := cpy
	_args[1] = _p1
	_end++
	_p2 := x
	_args[2] = _p2
	_end++
	_p3 := y
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("quadraticCurveTo", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := cp1x
	_args[0] = _p0
	_end++
	_p1 := cp1y
	_args[1] = _p1
	_end++
	_p2 := cp2x
	_args[2] = _p2
	_end++
	_p3 := cp2y
	_args[3] = _p3
	_end++
	_p4 := x
	_args[4] = _p4
	_end++
	_p5 := y
	_args[5] = _p5
	_end++
	_this.Value_JS.Call("bezierCurveTo", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	var (
		_args [5]interface{}
		_end  int
	)
	_p0 := x1
	_args[0] = _p0
	_end++
	_p1 := y1
	_args[1] = _p1
	_end++
	_p2 := x2
	_args[2] = _p2
	_end++
	_p3 := y2
	_args[3] = _p3
	_end++
	_p4 := radius
	_args[4] = _p4
	_end++
	_this.Value_JS.Call("arcTo", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Rect(x float64, y float64, w float64, h float64) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := w
	_args[2] = _p2
	_end++
	_p3 := h
	_args[3] = _p3
	_end++
	_this.Value_JS.Call("rect", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, anticlockwise *bool) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := radius
	_args[2] = _p2
	_end++
	_p3 := startAngle
	_args[3] = _p3
	_end++
	_p4 := endAngle
	_args[4] = _p4
	_end++
	if anticlockwise != nil {
		_p5 := anticlockwise
		_args[5] = _p5
		_end++
	}
	_this.Value_JS.Call("arc", _args[0:_end]...)
	return
}

func (_this *PaintRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, anticlockwise *bool) {
	var (
		_args [8]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_p2 := radiusX
	_args[2] = _p2
	_end++
	_p3 := radiusY
	_args[3] = _p3
	_end++
	_p4 := rotation
	_args[4] = _p4
	_end++
	_p5 := startAngle
	_args[5] = _p5
	_end++
	_p6 := endAngle
	_args[6] = _p6
	_end++
	if anticlockwise != nil {
		_p7 := anticlockwise
		_args[7] = _p7
		_end++
	}
	_this.Value_JS.Call("ellipse", _args[0:_end]...)
	return
}

// class: PaintSize
type PaintSize struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PaintSize) JSValue() js.Value {
	return _this.Value_JS
}

// PaintSizeFromJS is casting a js.Wrapper into PaintSize.
func PaintSizeFromJS(value js.Wrapper) *PaintSize {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PaintSize{}
	ret.Value_JS = input
	return ret
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *PaintSize) Width() float64 {
	var ret float64
	value := _this.Value_JS.Get("width")
	ret = (value).Float()
	return ret
}

// Height returning attribute 'height' with
// type float64 (idl: double).
func (_this *PaintSize) Height() float64 {
	var ret float64
	value := _this.Value_JS.Get("height")
	ret = (value).Float()
	return ret
}

// class: PaintWorkletGlobalScope
type PaintWorkletGlobalScope struct {
	worklets.WorkletGlobalScope
}

// PaintWorkletGlobalScopeFromJS is casting a js.Wrapper into PaintWorkletGlobalScope.
func PaintWorkletGlobalScopeFromJS(value js.Wrapper) *PaintWorkletGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PaintWorkletGlobalScope{}
	ret.Value_JS = input
	return ret
}

// DevicePixelRatio returning attribute 'devicePixelRatio' with
// type float64 (idl: unrestricted double).
func (_this *PaintWorkletGlobalScope) DevicePixelRatio() float64 {
	var ret float64
	value := _this.Value_JS.Get("devicePixelRatio")
	ret = (value).Float()
	return ret
}

func (_this *PaintWorkletGlobalScope) RegisterPaint(name string, paintCtor *webidl.VoidFunction) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++

	var __callback1 js.Value
	if paintCtor != nil {
		__callback1 = (*paintCtor).Value
	} else {
		__callback1 = js.Null()
	}
	_p1 := __callback1
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("registerPaint", _args[0:_end]...)
	return
}
