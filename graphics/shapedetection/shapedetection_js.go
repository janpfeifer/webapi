// Code generated by webidl-bind. DO NOT EDIT.

package shapedetection

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/geometry"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// geometry.DOMRectReadOnly
// javascript.FrozenArray
// javascript.PromiseFinally

// source idl files:
// promises.idl
// shape-detection-api.idl

// transform files:
// promises.go.md
// shape-detection-api.go.md

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

// enum: BarcodeFormat
type BarcodeFormat int

const (
	AztecBarcodeFormat BarcodeFormat = iota
	Code128BarcodeFormat
	Code39BarcodeFormat
	Code93BarcodeFormat
	CodabarBarcodeFormat
	DataMatrixBarcodeFormat
	Ean13BarcodeFormat
	Ean8BarcodeFormat
	ItfBarcodeFormat
	Pdf417BarcodeFormat
	QrCodeBarcodeFormat
	UnknownBarcodeFormat
	UpcABarcodeFormat
	UpcEBarcodeFormat
)

var barcodeFormatToWasmTable = []string{
	"aztec", "code_128", "code_39", "code_93", "codabar", "data_matrix", "ean_13", "ean_8", "itf", "pdf417", "qr_code", "unknown", "upc_a", "upc_e",
}

var barcodeFormatFromWasmTable = map[string]BarcodeFormat{
	"aztec": AztecBarcodeFormat, "code_128": Code128BarcodeFormat, "code_39": Code39BarcodeFormat, "code_93": Code93BarcodeFormat, "codabar": CodabarBarcodeFormat, "data_matrix": DataMatrixBarcodeFormat, "ean_13": Ean13BarcodeFormat, "ean_8": Ean8BarcodeFormat, "itf": ItfBarcodeFormat, "pdf417": Pdf417BarcodeFormat, "qr_code": QrCodeBarcodeFormat, "unknown": UnknownBarcodeFormat, "upc_a": UpcABarcodeFormat, "upc_e": UpcEBarcodeFormat,
}

// JSValue is converting this enum into a java object
func (this *BarcodeFormat) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this BarcodeFormat) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(barcodeFormatToWasmTable) {
		return barcodeFormatToWasmTable[idx]
	}
	panic("unknown input value")
}

// BarcodeFormatFromJS is converting a javascript value into
// a BarcodeFormat enum value.
func BarcodeFormatFromJS(value js.Value) BarcodeFormat {
	key := value.String()
	conv, ok := barcodeFormatFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: LandmarkType
type LandmarkType int

const (
	MouthLandmarkType LandmarkType = iota
	EyeLandmarkType
	NoseLandmarkType
)

var landmarkTypeToWasmTable = []string{
	"mouth", "eye", "nose",
}

var landmarkTypeFromWasmTable = map[string]LandmarkType{
	"mouth": MouthLandmarkType, "eye": EyeLandmarkType, "nose": NoseLandmarkType,
}

// JSValue is converting this enum into a java object
func (this *LandmarkType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this LandmarkType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(landmarkTypeToWasmTable) {
		return landmarkTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// LandmarkTypeFromJS is converting a javascript value into
// a LandmarkType enum value.
func LandmarkTypeFromJS(value js.Value) LandmarkType {
	key := value.String()
	conv, ok := landmarkTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseSequenceBarcodeFormatOnFulfilledFunc func(value []BarcodeFormat)

// PromiseSequenceBarcodeFormatOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceBarcodeFormatOnFulfilled js.Func

func PromiseSequenceBarcodeFormatOnFulfilledToJS(callback PromiseSequenceBarcodeFormatOnFulfilledFunc) *PromiseSequenceBarcodeFormatOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceBarcodeFormatOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []BarcodeFormat // javascript: sequence<BarcodeFormat> value
		)
		__length0 := args[0].Length()
		__array0 := make([]BarcodeFormat, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 BarcodeFormat
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = BarcodeFormatFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceBarcodeFormatOnFulfilledFromJS(_value js.Value) PromiseSequenceBarcodeFormatOnFulfilledFunc {
	return func(value []BarcodeFormat) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(value))
		for __idx0, __seq_in0 := range value {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseSequenceBarcodeFormatOnRejectedFunc func(reason js.Value)

// PromiseSequenceBarcodeFormatOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceBarcodeFormatOnRejected js.Func

func PromiseSequenceBarcodeFormatOnRejectedToJS(callback PromiseSequenceBarcodeFormatOnRejectedFunc) *PromiseSequenceBarcodeFormatOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceBarcodeFormatOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceBarcodeFormatOnRejectedFromJS(_value js.Value) PromiseSequenceBarcodeFormatOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromiseSequenceDetectedBarcodeOnFulfilledFunc func(value []*DetectedBarcode)

// PromiseSequenceDetectedBarcodeOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDetectedBarcodeOnFulfilled js.Func

func PromiseSequenceDetectedBarcodeOnFulfilledToJS(callback PromiseSequenceDetectedBarcodeOnFulfilledFunc) *PromiseSequenceDetectedBarcodeOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDetectedBarcodeOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*DetectedBarcode // javascript: sequence<DetectedBarcode> value
		)
		__length0 := args[0].Length()
		__array0 := make([]*DetectedBarcode, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *DetectedBarcode
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = DetectedBarcodeFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceDetectedBarcodeOnFulfilledFromJS(_value js.Value) PromiseSequenceDetectedBarcodeOnFulfilledFunc {
	return func(value []*DetectedBarcode) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(value))
		for __idx0, __seq_in0 := range value {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseSequenceDetectedBarcodeOnRejectedFunc func(reason js.Value)

// PromiseSequenceDetectedBarcodeOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDetectedBarcodeOnRejected js.Func

func PromiseSequenceDetectedBarcodeOnRejectedToJS(callback PromiseSequenceDetectedBarcodeOnRejectedFunc) *PromiseSequenceDetectedBarcodeOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDetectedBarcodeOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceDetectedBarcodeOnRejectedFromJS(_value js.Value) PromiseSequenceDetectedBarcodeOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromiseSequenceDetectedFaceOnFulfilledFunc func(value []*DetectedFace)

// PromiseSequenceDetectedFaceOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDetectedFaceOnFulfilled js.Func

func PromiseSequenceDetectedFaceOnFulfilledToJS(callback PromiseSequenceDetectedFaceOnFulfilledFunc) *PromiseSequenceDetectedFaceOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDetectedFaceOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*DetectedFace // javascript: sequence<DetectedFace> value
		)
		__length0 := args[0].Length()
		__array0 := make([]*DetectedFace, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *DetectedFace
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = DetectedFaceFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceDetectedFaceOnFulfilledFromJS(_value js.Value) PromiseSequenceDetectedFaceOnFulfilledFunc {
	return func(value []*DetectedFace) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(value))
		for __idx0, __seq_in0 := range value {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseSequenceDetectedFaceOnRejectedFunc func(reason js.Value)

// PromiseSequenceDetectedFaceOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseSequenceDetectedFaceOnRejected js.Func

func PromiseSequenceDetectedFaceOnRejectedToJS(callback PromiseSequenceDetectedFaceOnRejectedFunc) *PromiseSequenceDetectedFaceOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseSequenceDetectedFaceOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseSequenceDetectedFaceOnRejectedFromJS(_value js.Value) PromiseSequenceDetectedFaceOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: BarcodeDetectorOptions
type BarcodeDetectorOptions struct {
	Formats []BarcodeFormat
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BarcodeDetectorOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Formats))
	for __idx0, __seq_in0 := range _this.Formats {
		__seq_out0 := __seq_in0.JSValue()
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("formats", value0)
	return out
}

// BarcodeDetectorOptionsFromJS is allocating a new
// BarcodeDetectorOptions object and copy all values from
// input javascript object
func BarcodeDetectorOptionsFromJS(value js.Wrapper) *BarcodeDetectorOptions {
	input := value.JSValue()
	var out BarcodeDetectorOptions
	var (
		value0 []BarcodeFormat // javascript: sequence<BarcodeFormat> {formats Formats formats}
	)
	__length0 := input.Get("formats").Length()
	__array0 := make([]BarcodeFormat, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 BarcodeFormat
		__seq_in0 := input.Get("formats").Index(__idx0)
		__seq_out0 = BarcodeFormatFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Formats = value0
	return &out
}

// dictionary: FaceDetectorOptions
type FaceDetectorOptions struct {
	MaxDetectedFaces int
	FastMode         bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FaceDetectorOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.MaxDetectedFaces
	out.Set("maxDetectedFaces", value0)
	value1 := _this.FastMode
	out.Set("fastMode", value1)
	return out
}

// FaceDetectorOptionsFromJS is allocating a new
// FaceDetectorOptions object and copy all values from
// input javascript object
func FaceDetectorOptionsFromJS(value js.Wrapper) *FaceDetectorOptions {
	input := value.JSValue()
	var out FaceDetectorOptions
	var (
		value0 int  // javascript: unsigned short {maxDetectedFaces MaxDetectedFaces maxDetectedFaces}
		value1 bool // javascript: boolean {fastMode FastMode fastMode}
	)
	value0 = (input.Get("maxDetectedFaces")).Int()
	out.MaxDetectedFaces = value0
	value1 = (input.Get("fastMode")).Bool()
	out.FastMode = value1
	return &out
}

// dictionary: Landmark
type Landmark struct {
	Locations *javascript.FrozenArray
	Type      LandmarkType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Landmark) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Locations.JSValue()
	out.Set("locations", value0)
	value1 := _this.Type.JSValue()
	out.Set("type", value1)
	return out
}

// LandmarkFromJS is allocating a new
// Landmark object and copy all values from
// input javascript object
func LandmarkFromJS(value js.Wrapper) *Landmark {
	input := value.JSValue()
	var out Landmark
	var (
		value0 *javascript.FrozenArray // javascript: FrozenArray {locations Locations locations}
		value1 LandmarkType            // javascript: LandmarkType {type Type _type}
	)
	value0 = javascript.FrozenArrayFromJS(input.Get("locations"))
	out.Locations = value0
	value1 = LandmarkTypeFromJS(input.Get("type"))
	out.Type = value1
	return &out
}

// interface: BarcodeDetector
type BarcodeDetector struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *BarcodeDetector) JSValue() js.Value {
	return _this.Value_JS
}

// BarcodeDetectorFromJS is casting a js.Wrapper into BarcodeDetector.
func BarcodeDetectorFromJS(value js.Wrapper) *BarcodeDetector {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BarcodeDetector{}
	ret.Value_JS = input
	return ret
}

func GetSupportedFormats() (_result *PromiseSequenceBarcodeFormat) {
	_klass := js.Global().Get("BarcodeDetector")
	_method := _klass.Get("getSupportedFormats")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted *PromiseSequenceBarcodeFormat // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceBarcodeFormatFromJS(_returned)
	_result = _converted
	return
}

func NewBarcodeDetector(barcodeDetectorOptions *BarcodeDetectorOptions) (_result *BarcodeDetector) {
	_klass := js.Global().Get("BarcodeDetector")
	var (
		_args [1]interface{}
		_end  int
	)
	if barcodeDetectorOptions != nil {
		_p0 := barcodeDetectorOptions.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *BarcodeDetector // javascript: BarcodeDetector _what_return_name
	)
	_converted = BarcodeDetectorFromJS(_returned)
	_result = _converted
	return
}

func (_this *BarcodeDetector) Detect(image *Union) (_result *PromiseSequenceDetectedBarcode) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("detect", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedBarcode // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedBarcodeFromJS(_returned)
	_result = _converted
	return
}

// interface: DetectedBarcode
type DetectedBarcode struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DetectedBarcode) JSValue() js.Value {
	return _this.Value_JS
}

// DetectedBarcodeFromJS is casting a js.Wrapper into DetectedBarcode.
func DetectedBarcodeFromJS(value js.Wrapper) *DetectedBarcode {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DetectedBarcode{}
	ret.Value_JS = input
	return ret
}

// BoundingBox returning attribute 'boundingBox' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *DetectedBarcode) BoundingBox() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("boundingBox")
	ret = geometry.DOMRectReadOnlyFromJS(value)
	return ret
}

// RawValue returning attribute 'rawValue' with
// type string (idl: DOMString).
func (_this *DetectedBarcode) RawValue() string {
	var ret string
	value := _this.Value_JS.Get("rawValue")
	ret = (value).String()
	return ret
}

// Format returning attribute 'format' with
// type BarcodeFormat (idl: BarcodeFormat).
func (_this *DetectedBarcode) Format() BarcodeFormat {
	var ret BarcodeFormat
	value := _this.Value_JS.Get("format")
	ret = BarcodeFormatFromJS(value)
	return ret
}

// CornerPoints returning attribute 'cornerPoints' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *DetectedBarcode) CornerPoints() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("cornerPoints")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// interface: DetectedFace
type DetectedFace struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DetectedFace) JSValue() js.Value {
	return _this.Value_JS
}

// DetectedFaceFromJS is casting a js.Wrapper into DetectedFace.
func DetectedFaceFromJS(value js.Wrapper) *DetectedFace {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DetectedFace{}
	ret.Value_JS = input
	return ret
}

// BoundingBox returning attribute 'boundingBox' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *DetectedFace) BoundingBox() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("boundingBox")
	ret = geometry.DOMRectReadOnlyFromJS(value)
	return ret
}

// Landmarks returning attribute 'landmarks' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *DetectedFace) Landmarks() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("landmarks")
	if value.Type() != js.TypeNull {
		ret = javascript.FrozenArrayFromJS(value)
	}
	return ret
}

// interface: FaceDetector
type FaceDetector struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FaceDetector) JSValue() js.Value {
	return _this.Value_JS
}

// FaceDetectorFromJS is casting a js.Wrapper into FaceDetector.
func FaceDetectorFromJS(value js.Wrapper) *FaceDetector {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FaceDetector{}
	ret.Value_JS = input
	return ret
}

func NewFaceDetector(faceDetectorOptions *FaceDetectorOptions) (_result *FaceDetector) {
	_klass := js.Global().Get("FaceDetector")
	var (
		_args [1]interface{}
		_end  int
	)
	if faceDetectorOptions != nil {
		_p0 := faceDetectorOptions.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FaceDetector // javascript: FaceDetector _what_return_name
	)
	_converted = FaceDetectorFromJS(_returned)
	_result = _converted
	return
}

func (_this *FaceDetector) Detect(image *Union) (_result *PromiseSequenceDetectedFace) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("detect", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedFace // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedFaceFromJS(_returned)
	_result = _converted
	return
}

// interface: Promise
type PromiseSequenceBarcodeFormat struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseSequenceBarcodeFormat) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseSequenceBarcodeFormatFromJS is casting a js.Wrapper into PromiseSequenceBarcodeFormat.
func PromiseSequenceBarcodeFormatFromJS(value js.Wrapper) *PromiseSequenceBarcodeFormat {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseSequenceBarcodeFormat{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseSequenceBarcodeFormat) Then(onFulfilled *PromiseSequenceBarcodeFormatOnFulfilled, onRejected *PromiseSequenceBarcodeFormatOnRejected) (_result *PromiseSequenceBarcodeFormat) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseSequenceBarcodeFormat // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceBarcodeFormatFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceBarcodeFormat) Catch(onRejected *PromiseSequenceBarcodeFormatOnRejected) (_result *PromiseSequenceBarcodeFormat) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseSequenceBarcodeFormat // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceBarcodeFormatFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceBarcodeFormat) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseSequenceBarcodeFormat) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseSequenceBarcodeFormat // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceBarcodeFormatFromJS(_returned)
	_result = _converted
	return
}

// interface: Promise
type PromiseSequenceDetectedBarcode struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseSequenceDetectedBarcode) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseSequenceDetectedBarcodeFromJS is casting a js.Wrapper into PromiseSequenceDetectedBarcode.
func PromiseSequenceDetectedBarcodeFromJS(value js.Wrapper) *PromiseSequenceDetectedBarcode {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseSequenceDetectedBarcode{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseSequenceDetectedBarcode) Then(onFulfilled *PromiseSequenceDetectedBarcodeOnFulfilled, onRejected *PromiseSequenceDetectedBarcodeOnRejected) (_result *PromiseSequenceDetectedBarcode) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedBarcode // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedBarcodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDetectedBarcode) Catch(onRejected *PromiseSequenceDetectedBarcodeOnRejected) (_result *PromiseSequenceDetectedBarcode) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedBarcode // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedBarcodeFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDetectedBarcode) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseSequenceDetectedBarcode) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedBarcode // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedBarcodeFromJS(_returned)
	_result = _converted
	return
}

// interface: Promise
type PromiseSequenceDetectedFace struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseSequenceDetectedFace) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseSequenceDetectedFaceFromJS is casting a js.Wrapper into PromiseSequenceDetectedFace.
func PromiseSequenceDetectedFaceFromJS(value js.Wrapper) *PromiseSequenceDetectedFace {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseSequenceDetectedFace{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseSequenceDetectedFace) Then(onFulfilled *PromiseSequenceDetectedFaceOnFulfilled, onRejected *PromiseSequenceDetectedFaceOnRejected) (_result *PromiseSequenceDetectedFace) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedFace // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedFaceFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDetectedFace) Catch(onRejected *PromiseSequenceDetectedFaceOnRejected) (_result *PromiseSequenceDetectedFace) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedFace // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedFaceFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseSequenceDetectedFace) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseSequenceDetectedFace) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseSequenceDetectedFace // javascript: Promise _what_return_name
	)
	_converted = PromiseSequenceDetectedFaceFromJS(_returned)
	_result = _converted
	return
}
