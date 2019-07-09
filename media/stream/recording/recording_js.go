// Code generated by webidl-bind. DO NOT EDIT.

package recording

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/media/capture/local"
)

// using following types:
// domcore.DOMException
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// file.Blob
// local.MediaStream

// source idl files:
// mediastream-recording.idl

// transform files:
// mediastream-recording.go.md

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

// enum: RecordingState
type RecordingState int

const (
	InactiveRecordingState RecordingState = iota
	RecordingRecordingState
	PausedRecordingState
)

var recordingStateToWasmTable = []string{
	"inactive", "recording", "paused",
}

var recordingStateFromWasmTable = map[string]RecordingState{
	"inactive": InactiveRecordingState, "recording": RecordingRecordingState, "paused": PausedRecordingState,
}

// JSValue is converting this enum into a java object
func (this *RecordingState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RecordingState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(recordingStateToWasmTable) {
		return recordingStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// RecordingStateFromJS is converting a javascript value into
// a RecordingState enum value.
func RecordingStateFromJS(value js.Value) RecordingState {
	key := value.String()
	conv, ok := recordingStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: BlobEventInit
type BlobEventInit struct {
	Data     *file.Blob
	Timecode float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BlobEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Data.JSValue()
	out.Set("data", value0)
	value1 := _this.Timecode
	out.Set("timecode", value1)
	return out
}

// BlobEventInitFromJS is allocating a new
// BlobEventInit object and copy all values from
// input javascript object
func BlobEventInitFromJS(value js.Wrapper) *BlobEventInit {
	input := value.JSValue()
	var out BlobEventInit
	var (
		value0 *file.Blob // javascript: Blob {data Data data}
		value1 float64    // javascript: double {timecode Timecode timecode}
	)
	value0 = file.BlobFromJS(input.Get("data"))
	out.Data = value0
	value1 = (input.Get("timecode")).Float()
	out.Timecode = value1
	return &out
}

// dictionary: MediaRecorderErrorEventInit
type MediaRecorderErrorEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Error      *domcore.DOMException
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaRecorderErrorEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Error.JSValue()
	out.Set("error", value3)
	return out
}

// MediaRecorderErrorEventInitFromJS is allocating a new
// MediaRecorderErrorEventInit object and copy all values from
// input javascript object
func MediaRecorderErrorEventInitFromJS(value js.Wrapper) *MediaRecorderErrorEventInit {
	input := value.JSValue()
	var out MediaRecorderErrorEventInit
	var (
		value0 bool                  // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                  // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                  // javascript: boolean {composed Composed composed}
		value3 *domcore.DOMException // javascript: DOMException {error Error _error}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = domcore.DOMExceptionFromJS(input.Get("error"))
	out.Error = value3
	return &out
}

// dictionary: MediaRecorderOptions
type MediaRecorderOptions struct {
	MimeType           string
	AudioBitsPerSecond uint
	VideoBitsPerSecond uint
	BitsPerSecond      uint
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MediaRecorderOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.MimeType
	out.Set("mimeType", value0)
	value1 := _this.AudioBitsPerSecond
	out.Set("audioBitsPerSecond", value1)
	value2 := _this.VideoBitsPerSecond
	out.Set("videoBitsPerSecond", value2)
	value3 := _this.BitsPerSecond
	out.Set("bitsPerSecond", value3)
	return out
}

// MediaRecorderOptionsFromJS is allocating a new
// MediaRecorderOptions object and copy all values from
// input javascript object
func MediaRecorderOptionsFromJS(value js.Wrapper) *MediaRecorderOptions {
	input := value.JSValue()
	var out MediaRecorderOptions
	var (
		value0 string // javascript: DOMString {mimeType MimeType mimeType}
		value1 uint   // javascript: unsigned long {audioBitsPerSecond AudioBitsPerSecond audioBitsPerSecond}
		value2 uint   // javascript: unsigned long {videoBitsPerSecond VideoBitsPerSecond videoBitsPerSecond}
		value3 uint   // javascript: unsigned long {bitsPerSecond BitsPerSecond bitsPerSecond}
	)
	value0 = (input.Get("mimeType")).String()
	out.MimeType = value0
	value1 = (uint)((input.Get("audioBitsPerSecond")).Int())
	out.AudioBitsPerSecond = value1
	value2 = (uint)((input.Get("videoBitsPerSecond")).Int())
	out.VideoBitsPerSecond = value2
	value3 = (uint)((input.Get("bitsPerSecond")).Int())
	out.BitsPerSecond = value3
	return &out
}

// interface: BlobEvent
type BlobEvent struct {
	domcore.Event
}

// BlobEventFromJS is casting a js.Wrapper into BlobEvent.
func BlobEventFromJS(value js.Wrapper) *BlobEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BlobEvent{}
	ret.Value_JS = input
	return ret
}

func NewBlobEvent(_type string, eventInitDict *BlobEventInit) (_result *BlobEvent) {
	_klass := js.Global().Get("BlobEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *BlobEvent // javascript: BlobEvent _what_return_name
	)
	_converted = BlobEventFromJS(_returned)
	_result = _converted
	return
}

// Data returning attribute 'data' with
// type file.Blob (idl: Blob).
func (_this *BlobEvent) Data() *file.Blob {
	var ret *file.Blob
	value := _this.Value_JS.Get("data")
	ret = file.BlobFromJS(value)
	return ret
}

// Timecode returning attribute 'timecode' with
// type float64 (idl: double).
func (_this *BlobEvent) Timecode() float64 {
	var ret float64
	value := _this.Value_JS.Get("timecode")
	ret = (value).Float()
	return ret
}

// interface: MediaRecorder
type MediaRecorder struct {
	domcore.EventTarget
}

// MediaRecorderFromJS is casting a js.Wrapper into MediaRecorder.
func MediaRecorderFromJS(value js.Wrapper) *MediaRecorder {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaRecorder{}
	ret.Value_JS = input
	return ret
}

func IsTypeSupported(_type string) (_result bool) {
	_klass := js.Global().Get("MediaRecorder")
	_method := _klass.Get("isTypeSupported")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func NewMediaRecorder(stream *local.MediaStream, options *MediaRecorderOptions) (_result *MediaRecorder) {
	_klass := js.Global().Get("MediaRecorder")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := stream.JSValue()
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *MediaRecorder // javascript: MediaRecorder _what_return_name
	)
	_converted = MediaRecorderFromJS(_returned)
	_result = _converted
	return
}

// Stream returning attribute 'stream' with
// type local.MediaStream (idl: MediaStream).
func (_this *MediaRecorder) Stream() *local.MediaStream {
	var ret *local.MediaStream
	value := _this.Value_JS.Get("stream")
	ret = local.MediaStreamFromJS(value)
	return ret
}

// MimeType returning attribute 'mimeType' with
// type string (idl: DOMString).
func (_this *MediaRecorder) MimeType() string {
	var ret string
	value := _this.Value_JS.Get("mimeType")
	ret = (value).String()
	return ret
}

// State returning attribute 'state' with
// type RecordingState (idl: RecordingState).
func (_this *MediaRecorder) State() RecordingState {
	var ret RecordingState
	value := _this.Value_JS.Get("state")
	ret = RecordingStateFromJS(value)
	return ret
}

// Onstart returning attribute 'onstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Onstart() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onstart")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnstart setting attribute 'onstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOnstart(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onstart", input)
}

// Onstop returning attribute 'onstop' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Onstop() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onstop")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnstop setting attribute 'onstop' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOnstop(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onstop", input)
}

// Ondataavailable returning attribute 'ondataavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Ondataavailable() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ondataavailable")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOndataavailable setting attribute 'ondataavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOndataavailable(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("ondataavailable", input)
}

// Onpause returning attribute 'onpause' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Onpause() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onpause")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnpause setting attribute 'onpause' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOnpause(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onpause", input)
}

// Onresume returning attribute 'onresume' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Onresume() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onresume")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnresume setting attribute 'onresume' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOnresume(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onresume", input)
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MediaRecorder) SetOnerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onerror", input)
}

// VideoBitsPerSecond returning attribute 'videoBitsPerSecond' with
// type uint (idl: unsigned long).
func (_this *MediaRecorder) VideoBitsPerSecond() uint {
	var ret uint
	value := _this.Value_JS.Get("videoBitsPerSecond")
	ret = (uint)((value).Int())
	return ret
}

// AudioBitsPerSecond returning attribute 'audioBitsPerSecond' with
// type uint (idl: unsigned long).
func (_this *MediaRecorder) AudioBitsPerSecond() uint {
	var ret uint
	value := _this.Value_JS.Get("audioBitsPerSecond")
	ret = (uint)((value).Int())
	return ret
}

func (_this *MediaRecorder) Start(timeslice *uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	if timeslice != nil {
		_p0 := timeslice
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("start", _args[0:_end]...)
	return
}

func (_this *MediaRecorder) Stop() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("stop", _args[0:_end]...)
	return
}

func (_this *MediaRecorder) Pause() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("pause", _args[0:_end]...)
	return
}

func (_this *MediaRecorder) Resume() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("resume", _args[0:_end]...)
	return
}

func (_this *MediaRecorder) RequestData() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("requestData", _args[0:_end]...)
	return
}

// interface: MediaRecorderErrorEvent
type MediaRecorderErrorEvent struct {
	domcore.Event
}

// MediaRecorderErrorEventFromJS is casting a js.Wrapper into MediaRecorderErrorEvent.
func MediaRecorderErrorEventFromJS(value js.Wrapper) *MediaRecorderErrorEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaRecorderErrorEvent{}
	ret.Value_JS = input
	return ret
}

func NewMediaRecorderErrorEvent(_type string, eventInitDict *MediaRecorderErrorEventInit) (_result *MediaRecorderErrorEvent) {
	_klass := js.Global().Get("MediaRecorderErrorEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *MediaRecorderErrorEvent // javascript: MediaRecorderErrorEvent _what_return_name
	)
	_converted = MediaRecorderErrorEventFromJS(_returned)
	_result = _converted
	return
}

// Error returning attribute 'error' with
// type domcore.DOMException (idl: DOMException).
func (_this *MediaRecorderErrorEvent) Error() *domcore.DOMException {
	var ret *domcore.DOMException
	value := _this.Value_JS.Get("error")
	ret = domcore.DOMExceptionFromJS(value)
	return ret
}
