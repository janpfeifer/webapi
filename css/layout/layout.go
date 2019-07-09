// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package layout

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/css/typedom"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/webidl"
	"github.com/gowebapi/webapi/worklets"
)

// using following types:
// javascript.FrozenArray
// javascript.Promise
// typedom.StylePropertyMapReadOnly
// webidl.VoidFunction
// worklets.WorkletGlobalScope

// source idl files:
// css-layout-api.idl

// transform files:
// css-layout-api.go.md

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

// enum: BlockFragmentationType
type BlockFragmentationType int

const (
	NoneBlockFragmentationType BlockFragmentationType = iota
	PageBlockFragmentationType
	ColumnBlockFragmentationType
	RegionBlockFragmentationType
)

var blockFragmentationTypeToWasmTable = []string{
	"none", "page", "column", "region",
}

var blockFragmentationTypeFromWasmTable = map[string]BlockFragmentationType{
	"none": NoneBlockFragmentationType, "page": PageBlockFragmentationType, "column": ColumnBlockFragmentationType, "region": RegionBlockFragmentationType,
}

// JSValue is converting this enum into a java object
func (this *BlockFragmentationType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this BlockFragmentationType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(blockFragmentationTypeToWasmTable) {
		return blockFragmentationTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// BlockFragmentationTypeFromJS is converting a javascript value into
// a BlockFragmentationType enum value.
func BlockFragmentationTypeFromJS(value js.Value) BlockFragmentationType {
	key := value.String()
	conv, ok := blockFragmentationTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: BreakType
type BreakType int

const (
	NoneBreakType BreakType = iota
	LineBreakType
	ColumnBreakType
	PageBreakType
	RegionBreakType
)

var breakTypeToWasmTable = []string{
	"none", "line", "column", "page", "region",
}

var breakTypeFromWasmTable = map[string]BreakType{
	"none": NoneBreakType, "line": LineBreakType, "column": ColumnBreakType, "page": PageBreakType, "region": RegionBreakType,
}

// JSValue is converting this enum into a java object
func (this *BreakType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this BreakType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(breakTypeToWasmTable) {
		return breakTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// BreakTypeFromJS is converting a javascript value into
// a BreakType enum value.
func BreakTypeFromJS(value js.Value) BreakType {
	key := value.String()
	conv, ok := breakTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: BreakTokenOptions
type BreakTokenOptions struct {
	ChildBreakTokens []*ChildBreakToken
	Data             js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BreakTokenOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.ChildBreakTokens))
	for __idx0, __seq_in0 := range _this.ChildBreakTokens {
		__seq_out0 := __seq_in0.JSValue()
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("childBreakTokens", value0)
	value1 := _this.Data
	out.Set("data", value1)
	return out
}

// BreakTokenOptionsFromJS is allocating a new
// BreakTokenOptions object and copy all values from
// input javascript object
func BreakTokenOptionsFromJS(value js.Wrapper) *BreakTokenOptions {
	input := value.JSValue()
	var out BreakTokenOptions
	var (
		value0 []*ChildBreakToken // javascript: sequence<ChildBreakToken> {childBreakTokens ChildBreakTokens childBreakTokens}
		value1 js.Value           // javascript: any {data Data data}
	)
	__length0 := input.Get("childBreakTokens").Length()
	__array0 := make([]*ChildBreakToken, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *ChildBreakToken
		__seq_in0 := input.Get("childBreakTokens").Index(__idx0)
		__seq_out0 = ChildBreakTokenFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.ChildBreakTokens = value0
	value1 = input.Get("data")
	out.Data = value1
	return &out
}

// dictionary: FragmentResultOptions
type FragmentResultOptions struct {
	InlineSize     float64
	BlockSize      float64
	AutoBlockSize  float64
	ChildFragments []*LayoutFragment
	Data           js.Value
	BreakToken     *BreakTokenOptions
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FragmentResultOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.InlineSize
	out.Set("inlineSize", value0)
	value1 := _this.BlockSize
	out.Set("blockSize", value1)
	value2 := _this.AutoBlockSize
	out.Set("autoBlockSize", value2)
	value3 := js.Global().Get("Array").New(len(_this.ChildFragments))
	for __idx3, __seq_in3 := range _this.ChildFragments {
		__seq_out3 := __seq_in3.JSValue()
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("childFragments", value3)
	value4 := _this.Data
	out.Set("data", value4)
	value5 := _this.BreakToken.JSValue()
	out.Set("breakToken", value5)
	return out
}

// FragmentResultOptionsFromJS is allocating a new
// FragmentResultOptions object and copy all values from
// input javascript object
func FragmentResultOptionsFromJS(value js.Wrapper) *FragmentResultOptions {
	input := value.JSValue()
	var out FragmentResultOptions
	var (
		value0 float64            // javascript: double {inlineSize InlineSize inlineSize}
		value1 float64            // javascript: double {blockSize BlockSize blockSize}
		value2 float64            // javascript: double {autoBlockSize AutoBlockSize autoBlockSize}
		value3 []*LayoutFragment  // javascript: sequence<LayoutFragment> {childFragments ChildFragments childFragments}
		value4 js.Value           // javascript: any {data Data data}
		value5 *BreakTokenOptions // javascript: BreakTokenOptions {breakToken BreakToken breakToken}
	)
	value0 = (input.Get("inlineSize")).Float()
	out.InlineSize = value0
	value1 = (input.Get("blockSize")).Float()
	out.BlockSize = value1
	value2 = (input.Get("autoBlockSize")).Float()
	out.AutoBlockSize = value2
	__length3 := input.Get("childFragments").Length()
	__array3 := make([]*LayoutFragment, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 *LayoutFragment
		__seq_in3 := input.Get("childFragments").Index(__idx3)
		__seq_out3 = LayoutFragmentFromJS(__seq_in3)
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.ChildFragments = value3
	value4 = input.Get("data")
	out.Data = value4
	value5 = BreakTokenOptionsFromJS(input.Get("breakToken"))
	out.BreakToken = value5
	return &out
}

// dictionary: LayoutConstraintsOptions
type LayoutConstraintsOptions struct {
	AvailableInlineSize      float64
	AvailableBlockSize       float64
	FixedInlineSize          float64
	FixedBlockSize           float64
	PercentageInlineSize     float64
	PercentageBlockSize      float64
	BlockFragmentationOffset float64
	BlockFragmentationType   BlockFragmentationType
	Data                     js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *LayoutConstraintsOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.AvailableInlineSize
	out.Set("availableInlineSize", value0)
	value1 := _this.AvailableBlockSize
	out.Set("availableBlockSize", value1)
	value2 := _this.FixedInlineSize
	out.Set("fixedInlineSize", value2)
	value3 := _this.FixedBlockSize
	out.Set("fixedBlockSize", value3)
	value4 := _this.PercentageInlineSize
	out.Set("percentageInlineSize", value4)
	value5 := _this.PercentageBlockSize
	out.Set("percentageBlockSize", value5)
	value6 := _this.BlockFragmentationOffset
	out.Set("blockFragmentationOffset", value6)
	value7 := _this.BlockFragmentationType.JSValue()
	out.Set("blockFragmentationType", value7)
	value8 := _this.Data
	out.Set("data", value8)
	return out
}

// LayoutConstraintsOptionsFromJS is allocating a new
// LayoutConstraintsOptions object and copy all values from
// input javascript object
func LayoutConstraintsOptionsFromJS(value js.Wrapper) *LayoutConstraintsOptions {
	input := value.JSValue()
	var out LayoutConstraintsOptions
	var (
		value0 float64                // javascript: double {availableInlineSize AvailableInlineSize availableInlineSize}
		value1 float64                // javascript: double {availableBlockSize AvailableBlockSize availableBlockSize}
		value2 float64                // javascript: double {fixedInlineSize FixedInlineSize fixedInlineSize}
		value3 float64                // javascript: double {fixedBlockSize FixedBlockSize fixedBlockSize}
		value4 float64                // javascript: double {percentageInlineSize PercentageInlineSize percentageInlineSize}
		value5 float64                // javascript: double {percentageBlockSize PercentageBlockSize percentageBlockSize}
		value6 float64                // javascript: double {blockFragmentationOffset BlockFragmentationOffset blockFragmentationOffset}
		value7 BlockFragmentationType // javascript: BlockFragmentationType {blockFragmentationType BlockFragmentationType blockFragmentationType}
		value8 js.Value               // javascript: any {data Data data}
	)
	value0 = (input.Get("availableInlineSize")).Float()
	out.AvailableInlineSize = value0
	value1 = (input.Get("availableBlockSize")).Float()
	out.AvailableBlockSize = value1
	value2 = (input.Get("fixedInlineSize")).Float()
	out.FixedInlineSize = value2
	value3 = (input.Get("fixedBlockSize")).Float()
	out.FixedBlockSize = value3
	value4 = (input.Get("percentageInlineSize")).Float()
	out.PercentageInlineSize = value4
	value5 = (input.Get("percentageBlockSize")).Float()
	out.PercentageBlockSize = value5
	value6 = (input.Get("blockFragmentationOffset")).Float()
	out.BlockFragmentationOffset = value6
	value7 = BlockFragmentationTypeFromJS(input.Get("blockFragmentationType"))
	out.BlockFragmentationType = value7
	value8 = input.Get("data")
	out.Data = value8
	return &out
}

// interface: BreakToken
type BreakToken struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *BreakToken) JSValue() js.Value {
	return _this.Value_JS
}

// BreakTokenFromJS is casting a js.Wrapper into BreakToken.
func BreakTokenFromJS(value js.Wrapper) *BreakToken {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BreakToken{}
	ret.Value_JS = input
	return ret
}

// ChildBreakTokens returning attribute 'childBreakTokens' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *BreakToken) ChildBreakTokens() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("childBreakTokens")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Data returning attribute 'data' with
// type Any (idl: any).
func (_this *BreakToken) Data() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("data")
	ret = value
	return ret
}

// interface: ChildBreakToken
type ChildBreakToken struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ChildBreakToken) JSValue() js.Value {
	return _this.Value_JS
}

// ChildBreakTokenFromJS is casting a js.Wrapper into ChildBreakToken.
func ChildBreakTokenFromJS(value js.Wrapper) *ChildBreakToken {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ChildBreakToken{}
	ret.Value_JS = input
	return ret
}

// BreakType returning attribute 'breakType' with
// type BreakType (idl: BreakType).
func (_this *ChildBreakToken) BreakType() BreakType {
	var ret BreakType
	value := _this.Value_JS.Get("breakType")
	ret = BreakTypeFromJS(value)
	return ret
}

// Child returning attribute 'child' with
// type LayoutChild (idl: LayoutChild).
func (_this *ChildBreakToken) Child() *LayoutChild {
	var ret *LayoutChild
	value := _this.Value_JS.Get("child")
	ret = LayoutChildFromJS(value)
	return ret
}

// interface: FragmentResult
type FragmentResult struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FragmentResult) JSValue() js.Value {
	return _this.Value_JS
}

// FragmentResultFromJS is casting a js.Wrapper into FragmentResult.
func FragmentResultFromJS(value js.Wrapper) *FragmentResult {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FragmentResult{}
	ret.Value_JS = input
	return ret
}

func NewFragmentResult(options *FragmentResultOptions) (_result *FragmentResult) {
	_klass := js.Global().Get("FragmentResult")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := options.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FragmentResult // javascript: FragmentResult _what_return_name
	)
	_converted = FragmentResultFromJS(_returned)
	_result = _converted
	return
}

// InlineSize returning attribute 'inlineSize' with
// type float64 (idl: double).
func (_this *FragmentResult) InlineSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("inlineSize")
	ret = (value).Float()
	return ret
}

// BlockSize returning attribute 'blockSize' with
// type float64 (idl: double).
func (_this *FragmentResult) BlockSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("blockSize")
	ret = (value).Float()
	return ret
}

// interface: IntrinsicSizes
type IntrinsicSizes struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *IntrinsicSizes) JSValue() js.Value {
	return _this.Value_JS
}

// IntrinsicSizesFromJS is casting a js.Wrapper into IntrinsicSizes.
func IntrinsicSizesFromJS(value js.Wrapper) *IntrinsicSizes {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &IntrinsicSizes{}
	ret.Value_JS = input
	return ret
}

// MinContentSize returning attribute 'minContentSize' with
// type float64 (idl: double).
func (_this *IntrinsicSizes) MinContentSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("minContentSize")
	ret = (value).Float()
	return ret
}

// MaxContentSize returning attribute 'maxContentSize' with
// type float64 (idl: double).
func (_this *IntrinsicSizes) MaxContentSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("maxContentSize")
	ret = (value).Float()
	return ret
}

// interface: LayoutChild
type LayoutChild struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *LayoutChild) JSValue() js.Value {
	return _this.Value_JS
}

// LayoutChildFromJS is casting a js.Wrapper into LayoutChild.
func LayoutChildFromJS(value js.Wrapper) *LayoutChild {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &LayoutChild{}
	ret.Value_JS = input
	return ret
}

// StyleMap returning attribute 'styleMap' with
// type typedom.StylePropertyMapReadOnly (idl: StylePropertyMapReadOnly).
func (_this *LayoutChild) StyleMap() *typedom.StylePropertyMapReadOnly {
	var ret *typedom.StylePropertyMapReadOnly
	value := _this.Value_JS.Get("styleMap")
	ret = typedom.StylePropertyMapReadOnlyFromJS(value)
	return ret
}

func (_this *LayoutChild) IntrinsicSizes() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("intrinsicSizes", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *LayoutChild) LayoutNextFragment(constraints *LayoutConstraintsOptions, breakToken *ChildBreakToken) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := constraints.JSValue()
	_args[0] = _p0
	_end++
	_p1 := breakToken.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("layoutNextFragment", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: LayoutConstraints
type LayoutConstraints struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *LayoutConstraints) JSValue() js.Value {
	return _this.Value_JS
}

// LayoutConstraintsFromJS is casting a js.Wrapper into LayoutConstraints.
func LayoutConstraintsFromJS(value js.Wrapper) *LayoutConstraints {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &LayoutConstraints{}
	ret.Value_JS = input
	return ret
}

// AvailableInlineSize returning attribute 'availableInlineSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) AvailableInlineSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("availableInlineSize")
	ret = (value).Float()
	return ret
}

// AvailableBlockSize returning attribute 'availableBlockSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) AvailableBlockSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("availableBlockSize")
	ret = (value).Float()
	return ret
}

// FixedInlineSize returning attribute 'fixedInlineSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) FixedInlineSize() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("fixedInlineSize")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// FixedBlockSize returning attribute 'fixedBlockSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) FixedBlockSize() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("fixedBlockSize")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// PercentageInlineSize returning attribute 'percentageInlineSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) PercentageInlineSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("percentageInlineSize")
	ret = (value).Float()
	return ret
}

// PercentageBlockSize returning attribute 'percentageBlockSize' with
// type float64 (idl: double).
func (_this *LayoutConstraints) PercentageBlockSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("percentageBlockSize")
	ret = (value).Float()
	return ret
}

// BlockFragmentationOffset returning attribute 'blockFragmentationOffset' with
// type float64 (idl: double).
func (_this *LayoutConstraints) BlockFragmentationOffset() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("blockFragmentationOffset")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// BlockFragmentationType returning attribute 'blockFragmentationType' with
// type BlockFragmentationType (idl: BlockFragmentationType).
func (_this *LayoutConstraints) BlockFragmentationType() BlockFragmentationType {
	var ret BlockFragmentationType
	value := _this.Value_JS.Get("blockFragmentationType")
	ret = BlockFragmentationTypeFromJS(value)
	return ret
}

// Data returning attribute 'data' with
// type Any (idl: any).
func (_this *LayoutConstraints) Data() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("data")
	ret = value
	return ret
}

// interface: LayoutEdges
type LayoutEdges struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *LayoutEdges) JSValue() js.Value {
	return _this.Value_JS
}

// LayoutEdgesFromJS is casting a js.Wrapper into LayoutEdges.
func LayoutEdgesFromJS(value js.Wrapper) *LayoutEdges {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &LayoutEdges{}
	ret.Value_JS = input
	return ret
}

// InlineStart returning attribute 'inlineStart' with
// type float64 (idl: double).
func (_this *LayoutEdges) InlineStart() float64 {
	var ret float64
	value := _this.Value_JS.Get("inlineStart")
	ret = (value).Float()
	return ret
}

// InlineEnd returning attribute 'inlineEnd' with
// type float64 (idl: double).
func (_this *LayoutEdges) InlineEnd() float64 {
	var ret float64
	value := _this.Value_JS.Get("inlineEnd")
	ret = (value).Float()
	return ret
}

// BlockStart returning attribute 'blockStart' with
// type float64 (idl: double).
func (_this *LayoutEdges) BlockStart() float64 {
	var ret float64
	value := _this.Value_JS.Get("blockStart")
	ret = (value).Float()
	return ret
}

// BlockEnd returning attribute 'blockEnd' with
// type float64 (idl: double).
func (_this *LayoutEdges) BlockEnd() float64 {
	var ret float64
	value := _this.Value_JS.Get("blockEnd")
	ret = (value).Float()
	return ret
}

// Inline returning attribute 'inline' with
// type float64 (idl: double).
func (_this *LayoutEdges) Inline() float64 {
	var ret float64
	value := _this.Value_JS.Get("inline")
	ret = (value).Float()
	return ret
}

// Block returning attribute 'block' with
// type float64 (idl: double).
func (_this *LayoutEdges) Block() float64 {
	var ret float64
	value := _this.Value_JS.Get("block")
	ret = (value).Float()
	return ret
}

// interface: LayoutFragment
type LayoutFragment struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *LayoutFragment) JSValue() js.Value {
	return _this.Value_JS
}

// LayoutFragmentFromJS is casting a js.Wrapper into LayoutFragment.
func LayoutFragmentFromJS(value js.Wrapper) *LayoutFragment {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &LayoutFragment{}
	ret.Value_JS = input
	return ret
}

// InlineSize returning attribute 'inlineSize' with
// type float64 (idl: double).
func (_this *LayoutFragment) InlineSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("inlineSize")
	ret = (value).Float()
	return ret
}

// BlockSize returning attribute 'blockSize' with
// type float64 (idl: double).
func (_this *LayoutFragment) BlockSize() float64 {
	var ret float64
	value := _this.Value_JS.Get("blockSize")
	ret = (value).Float()
	return ret
}

// InlineOffset returning attribute 'inlineOffset' with
// type float64 (idl: double).
func (_this *LayoutFragment) InlineOffset() float64 {
	var ret float64
	value := _this.Value_JS.Get("inlineOffset")
	ret = (value).Float()
	return ret
}

// SetInlineOffset setting attribute 'inlineOffset' with
// type float64 (idl: double).
func (_this *LayoutFragment) SetInlineOffset(value float64) {
	input := value
	_this.Value_JS.Set("inlineOffset", input)
}

// BlockOffset returning attribute 'blockOffset' with
// type float64 (idl: double).
func (_this *LayoutFragment) BlockOffset() float64 {
	var ret float64
	value := _this.Value_JS.Get("blockOffset")
	ret = (value).Float()
	return ret
}

// SetBlockOffset setting attribute 'blockOffset' with
// type float64 (idl: double).
func (_this *LayoutFragment) SetBlockOffset(value float64) {
	input := value
	_this.Value_JS.Set("blockOffset", input)
}

// Data returning attribute 'data' with
// type Any (idl: any).
func (_this *LayoutFragment) Data() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("data")
	ret = value
	return ret
}

// BreakToken returning attribute 'breakToken' with
// type ChildBreakToken (idl: ChildBreakToken).
func (_this *LayoutFragment) BreakToken() *ChildBreakToken {
	var ret *ChildBreakToken
	value := _this.Value_JS.Get("breakToken")
	if value.Type() != js.TypeNull {
		ret = ChildBreakTokenFromJS(value)
	}
	return ret
}

// interface: LayoutWorkletGlobalScope
type LayoutWorkletGlobalScope struct {
	worklets.WorkletGlobalScope
}

// LayoutWorkletGlobalScopeFromJS is casting a js.Wrapper into LayoutWorkletGlobalScope.
func LayoutWorkletGlobalScopeFromJS(value js.Wrapper) *LayoutWorkletGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &LayoutWorkletGlobalScope{}
	ret.Value_JS = input
	return ret
}

func (_this *LayoutWorkletGlobalScope) RegisterLayout(name string, layoutCtor *webidl.VoidFunction) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++

	var __callback1 js.Value
	if layoutCtor != nil {
		__callback1 = (*layoutCtor).Value
	} else {
		__callback1 = js.Null()
	}
	_p1 := __callback1
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("registerLayout", _args[0:_end]...)
	return
}
