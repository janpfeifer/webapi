// Code generated by webidl-bind. DO NOT EDIT.

package appmenifestres

import "syscall/js"

// using following types:

// source idl files:
// appmanifest.idl

// transform files:
// appmanifest.go.md

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

// dictionary: ExternalApplicationResource
type ExternalApplicationResource struct {
	Platform     string
	Url          string
	Id           string
	MinVersion   string
	Fingerprints []*Fingerprint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ExternalApplicationResource) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Platform
	out.Set("platform", value0)
	value1 := _this.Url
	out.Set("url", value1)
	value2 := _this.Id
	out.Set("id", value2)
	value3 := _this.MinVersion
	out.Set("min_version", value3)
	value4 := js.Global().Get("Array").New(len(_this.Fingerprints))
	for __idx4, __seq_in4 := range _this.Fingerprints {
		__seq_out4 := __seq_in4.JSValue()
		value4.SetIndex(__idx4, __seq_out4)
	}
	out.Set("fingerprints", value4)
	return out
}

// ExternalApplicationResourceFromJS is allocating a new
// ExternalApplicationResource object and copy all values from
// input javascript object
func ExternalApplicationResourceFromJS(value js.Wrapper) *ExternalApplicationResource {
	input := value.JSValue()
	var out ExternalApplicationResource
	var (
		value0 string         // javascript: USVString {platform Platform platform}
		value1 string         // javascript: USVString {url Url url}
		value2 string         // javascript: DOMString {id Id id}
		value3 string         // javascript: USVString {min_version MinVersion minVersion}
		value4 []*Fingerprint // javascript: sequence<Fingerprint> {fingerprints Fingerprints fingerprints}
	)
	value0 = (input.Get("platform")).String()
	out.Platform = value0
	value1 = (input.Get("url")).String()
	out.Url = value1
	value2 = (input.Get("id")).String()
	out.Id = value2
	value3 = (input.Get("min_version")).String()
	out.MinVersion = value3
	__length4 := input.Get("fingerprints").Length()
	__array4 := make([]*Fingerprint, __length4, __length4)
	for __idx4 := 0; __idx4 < __length4; __idx4++ {
		var __seq_out4 *Fingerprint
		__seq_in4 := input.Get("fingerprints").Index(__idx4)
		__seq_out4 = FingerprintFromJS(__seq_in4)
		__array4[__idx4] = __seq_out4
	}
	value4 = __array4
	out.Fingerprints = value4
	return &out
}

// dictionary: Fingerprint
type Fingerprint struct {
	Type  string
	Value string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Fingerprint) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type
	out.Set("type", value0)
	value1 := _this.Value
	out.Set("value", value1)
	return out
}

// FingerprintFromJS is allocating a new
// Fingerprint object and copy all values from
// input javascript object
func FingerprintFromJS(value js.Wrapper) *Fingerprint {
	input := value.JSValue()
	var out Fingerprint
	var (
		value0 string // javascript: USVString {type Type _type}
		value1 string // javascript: USVString {value Value value}
	)
	value0 = (input.Get("type")).String()
	out.Type = value0
	value1 = (input.Get("value")).String()
	out.Value = value1
	return &out
}

// dictionary: ImageResource
type ImageResource struct {
	Src      string
	Sizes    string
	Type     string
	Purpose  string
	Platform string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ImageResource) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Src
	out.Set("src", value0)
	value1 := _this.Sizes
	out.Set("sizes", value1)
	value2 := _this.Type
	out.Set("type", value2)
	value3 := _this.Purpose
	out.Set("purpose", value3)
	value4 := _this.Platform
	out.Set("platform", value4)
	return out
}

// ImageResourceFromJS is allocating a new
// ImageResource object and copy all values from
// input javascript object
func ImageResourceFromJS(value js.Wrapper) *ImageResource {
	input := value.JSValue()
	var out ImageResource
	var (
		value0 string // javascript: USVString {src Src src}
		value1 string // javascript: DOMString {sizes Sizes sizes}
		value2 string // javascript: USVString {type Type _type}
		value3 string // javascript: USVString {purpose Purpose purpose}
		value4 string // javascript: USVString {platform Platform platform}
	)
	value0 = (input.Get("src")).String()
	out.Src = value0
	value1 = (input.Get("sizes")).String()
	out.Sizes = value1
	value2 = (input.Get("type")).String()
	out.Type = value2
	value3 = (input.Get("purpose")).String()
	out.Purpose = value3
	value4 = (input.Get("platform")).String()
	out.Platform = value4
	return &out
}
