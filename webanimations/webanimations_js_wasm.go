// Code generated by webidl-bind. DO NOT EDIT.

package webanimations

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.Object
// javascript.Promise

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

// enum: AnimationPlayState
type AnimationPlayState int

const (
	IdleAnimationPlayState AnimationPlayState = iota
	RunningAnimationPlayState
	PausedAnimationPlayState
	FinishedAnimationPlayState
)

var animationPlayStateToWasmTable = []string{
	"idle", "running", "paused", "finished",
}

var animationPlayStateFromWasmTable = map[string]AnimationPlayState{
	"idle": IdleAnimationPlayState, "running": RunningAnimationPlayState, "paused": PausedAnimationPlayState, "finished": FinishedAnimationPlayState,
}

// JSValue is converting this enum into a java object
func (this *AnimationPlayState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AnimationPlayState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(animationPlayStateToWasmTable) {
		return animationPlayStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// AnimationPlayStateFromJS is converting a javascript value into
// a AnimationPlayState enum value.
func AnimationPlayStateFromJS(value js.Value) AnimationPlayState {
	key := value.String()
	conv, ok := animationPlayStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: CompositeOperation
type CompositeOperation int

const (
	ReplaceCompositeOperation CompositeOperation = iota
	AddCompositeOperation
	AccumulateCompositeOperation
)

var compositeOperationToWasmTable = []string{
	"replace", "add", "accumulate",
}

var compositeOperationFromWasmTable = map[string]CompositeOperation{
	"replace": ReplaceCompositeOperation, "add": AddCompositeOperation, "accumulate": AccumulateCompositeOperation,
}

// JSValue is converting this enum into a java object
func (this *CompositeOperation) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this CompositeOperation) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(compositeOperationToWasmTable) {
		return compositeOperationToWasmTable[idx]
	}
	panic("unknown input value")
}

// CompositeOperationFromJS is converting a javascript value into
// a CompositeOperation enum value.
func CompositeOperationFromJS(value js.Value) CompositeOperation {
	key := value.String()
	conv, ok := compositeOperationFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: FillMode
type FillMode int

const (
	NoneFillMode FillMode = iota
	ForwardsFillMode
	BackwardsFillMode
	BothFillMode
	AutoFillMode
)

var fillModeToWasmTable = []string{
	"none", "forwards", "backwards", "both", "auto",
}

var fillModeFromWasmTable = map[string]FillMode{
	"none": NoneFillMode, "forwards": ForwardsFillMode, "backwards": BackwardsFillMode, "both": BothFillMode, "auto": AutoFillMode,
}

// JSValue is converting this enum into a java object
func (this *FillMode) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this FillMode) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(fillModeToWasmTable) {
		return fillModeToWasmTable[idx]
	}
	panic("unknown input value")
}

// FillModeFromJS is converting a javascript value into
// a FillMode enum value.
func FillModeFromJS(value js.Value) FillMode {
	key := value.String()
	conv, ok := fillModeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: IterationCompositeOperation
type IterationCompositeOperation int

const (
	ReplaceIterationCompositeOperation IterationCompositeOperation = iota
	AccumulateIterationCompositeOperation
)

var iterationCompositeOperationToWasmTable = []string{
	"replace", "accumulate",
}

var iterationCompositeOperationFromWasmTable = map[string]IterationCompositeOperation{
	"replace": ReplaceIterationCompositeOperation, "accumulate": AccumulateIterationCompositeOperation,
}

// JSValue is converting this enum into a java object
func (this *IterationCompositeOperation) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this IterationCompositeOperation) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(iterationCompositeOperationToWasmTable) {
		return iterationCompositeOperationToWasmTable[idx]
	}
	panic("unknown input value")
}

// IterationCompositeOperationFromJS is converting a javascript value into
// a IterationCompositeOperation enum value.
func IterationCompositeOperationFromJS(value js.Value) IterationCompositeOperation {
	key := value.String()
	conv, ok := iterationCompositeOperationFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: PlaybackDirection
type PlaybackDirection int

const (
	NormalPlaybackDirection PlaybackDirection = iota
	ReversePlaybackDirection
	AlternatePlaybackDirection
	AlternateReversePlaybackDirection
)

var playbackDirectionToWasmTable = []string{
	"normal", "reverse", "alternate", "alternate-reverse",
}

var playbackDirectionFromWasmTable = map[string]PlaybackDirection{
	"normal": NormalPlaybackDirection, "reverse": ReversePlaybackDirection, "alternate": AlternatePlaybackDirection, "alternate-reverse": AlternateReversePlaybackDirection,
}

// JSValue is converting this enum into a java object
func (this *PlaybackDirection) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PlaybackDirection) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(playbackDirectionToWasmTable) {
		return playbackDirectionToWasmTable[idx]
	}
	panic("unknown input value")
}

// PlaybackDirectionFromJS is converting a javascript value into
// a PlaybackDirection enum value.
func PlaybackDirectionFromJS(value js.Value) PlaybackDirection {
	key := value.String()
	conv, ok := playbackDirectionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: AnimationPlaybackEventInit
type AnimationPlaybackEventInit struct {
	Bubbles      bool
	Cancelable   bool
	Composed     bool
	CurrentTime  *float64
	TimelineTime *float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *AnimationPlaybackEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.CurrentTime
	out.Set("currentTime", value3)
	value4 := _this.TimelineTime
	out.Set("timelineTime", value4)
	return out
}

// AnimationPlaybackEventInitFromJS is allocating a new
// AnimationPlaybackEventInit object and copy all values from
// input javascript object
func AnimationPlaybackEventInitFromJS(value js.Wrapper) *AnimationPlaybackEventInit {
	input := value.JSValue()
	var out AnimationPlaybackEventInit
	var (
		value0 bool     // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool     // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool     // javascript: boolean {composed Composed composed}
		value3 *float64 // javascript: double {currentTime CurrentTime currentTime}
		value4 *float64 // javascript: double {timelineTime TimelineTime timelineTime}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	if input.Get("currentTime").Type() != js.TypeNull {
		__tmp := (input.Get("currentTime")).Float()
		value3 = &__tmp
	}
	out.CurrentTime = value3
	if input.Get("timelineTime").Type() != js.TypeNull {
		__tmp := (input.Get("timelineTime")).Float()
		value4 = &__tmp
	}
	out.TimelineTime = value4
	return &out
}

// dictionary: ComputedEffectTiming
type ComputedEffectTiming struct {
	Delay            float64
	EndDelay         float64
	Fill             FillMode
	IterationStart   float64
	Iterations       float64
	Duration         *Union
	Direction        PlaybackDirection
	Easing           string
	EndTime          float64
	ActiveDuration   float64
	LocalTime        *float64
	Progress         *float64
	CurrentIteration *float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ComputedEffectTiming) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Delay
	out.Set("delay", value0)
	value1 := _this.EndDelay
	out.Set("endDelay", value1)
	value2 := _this.Fill.JSValue()
	out.Set("fill", value2)
	value3 := _this.IterationStart
	out.Set("iterationStart", value3)
	value4 := _this.Iterations
	out.Set("iterations", value4)
	value5 := _this.Duration.JSValue()
	out.Set("duration", value5)
	value6 := _this.Direction.JSValue()
	out.Set("direction", value6)
	value7 := _this.Easing
	out.Set("easing", value7)
	value8 := _this.EndTime
	out.Set("endTime", value8)
	value9 := _this.ActiveDuration
	out.Set("activeDuration", value9)
	value10 := _this.LocalTime
	out.Set("localTime", value10)
	value11 := _this.Progress
	out.Set("progress", value11)
	value12 := _this.CurrentIteration
	out.Set("currentIteration", value12)
	return out
}

// ComputedEffectTimingFromJS is allocating a new
// ComputedEffectTiming object and copy all values from
// input javascript object
func ComputedEffectTimingFromJS(value js.Wrapper) *ComputedEffectTiming {
	input := value.JSValue()
	var out ComputedEffectTiming
	var (
		value0  float64           // javascript: double {delay Delay delay}
		value1  float64           // javascript: double {endDelay EndDelay endDelay}
		value2  FillMode          // javascript: FillMode {fill Fill fill}
		value3  float64           // javascript: double {iterationStart IterationStart iterationStart}
		value4  float64           // javascript: unrestricted double {iterations Iterations iterations}
		value5  *Union            // javascript: Union {duration Duration duration}
		value6  PlaybackDirection // javascript: PlaybackDirection {direction Direction direction}
		value7  string            // javascript: DOMString {easing Easing easing}
		value8  float64           // javascript: unrestricted double {endTime EndTime endTime}
		value9  float64           // javascript: unrestricted double {activeDuration ActiveDuration activeDuration}
		value10 *float64          // javascript: double {localTime LocalTime localTime}
		value11 *float64          // javascript: double {progress Progress progress}
		value12 *float64          // javascript: unrestricted double {currentIteration CurrentIteration currentIteration}
	)
	value0 = (input.Get("delay")).Float()
	out.Delay = value0
	value1 = (input.Get("endDelay")).Float()
	out.EndDelay = value1
	value2 = FillModeFromJS(input.Get("fill"))
	out.Fill = value2
	value3 = (input.Get("iterationStart")).Float()
	out.IterationStart = value3
	value4 = (input.Get("iterations")).Float()
	out.Iterations = value4
	value5 = UnionFromJS(input.Get("duration"))
	out.Duration = value5
	value6 = PlaybackDirectionFromJS(input.Get("direction"))
	out.Direction = value6
	value7 = (input.Get("easing")).String()
	out.Easing = value7
	value8 = (input.Get("endTime")).Float()
	out.EndTime = value8
	value9 = (input.Get("activeDuration")).Float()
	out.ActiveDuration = value9
	if input.Get("localTime").Type() != js.TypeNull {
		__tmp := (input.Get("localTime")).Float()
		value10 = &__tmp
	}
	out.LocalTime = value10
	if input.Get("progress").Type() != js.TypeNull {
		__tmp := (input.Get("progress")).Float()
		value11 = &__tmp
	}
	out.Progress = value11
	if input.Get("currentIteration").Type() != js.TypeNull {
		__tmp := (input.Get("currentIteration")).Float()
		value12 = &__tmp
	}
	out.CurrentIteration = value12
	return &out
}

// dictionary: DocumentTimelineOptions
type DocumentTimelineOptions struct {
	OriginTime float64
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *DocumentTimelineOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.OriginTime
	out.Set("originTime", value0)
	return out
}

// DocumentTimelineOptionsFromJS is allocating a new
// DocumentTimelineOptions object and copy all values from
// input javascript object
func DocumentTimelineOptionsFromJS(value js.Wrapper) *DocumentTimelineOptions {
	input := value.JSValue()
	var out DocumentTimelineOptions
	var (
		value0 float64 // javascript: double {originTime OriginTime originTime}
	)
	value0 = (input.Get("originTime")).Float()
	out.OriginTime = value0
	return &out
}

// dictionary: EffectTiming
type EffectTiming struct {
	Delay          float64
	EndDelay       float64
	Fill           FillMode
	IterationStart float64
	Iterations     float64
	Duration       *Union
	Direction      PlaybackDirection
	Easing         string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *EffectTiming) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Delay
	out.Set("delay", value0)
	value1 := _this.EndDelay
	out.Set("endDelay", value1)
	value2 := _this.Fill.JSValue()
	out.Set("fill", value2)
	value3 := _this.IterationStart
	out.Set("iterationStart", value3)
	value4 := _this.Iterations
	out.Set("iterations", value4)
	value5 := _this.Duration.JSValue()
	out.Set("duration", value5)
	value6 := _this.Direction.JSValue()
	out.Set("direction", value6)
	value7 := _this.Easing
	out.Set("easing", value7)
	return out
}

// EffectTimingFromJS is allocating a new
// EffectTiming object and copy all values from
// input javascript object
func EffectTimingFromJS(value js.Wrapper) *EffectTiming {
	input := value.JSValue()
	var out EffectTiming
	var (
		value0 float64           // javascript: double {delay Delay delay}
		value1 float64           // javascript: double {endDelay EndDelay endDelay}
		value2 FillMode          // javascript: FillMode {fill Fill fill}
		value3 float64           // javascript: double {iterationStart IterationStart iterationStart}
		value4 float64           // javascript: unrestricted double {iterations Iterations iterations}
		value5 *Union            // javascript: Union {duration Duration duration}
		value6 PlaybackDirection // javascript: PlaybackDirection {direction Direction direction}
		value7 string            // javascript: DOMString {easing Easing easing}
	)
	value0 = (input.Get("delay")).Float()
	out.Delay = value0
	value1 = (input.Get("endDelay")).Float()
	out.EndDelay = value1
	value2 = FillModeFromJS(input.Get("fill"))
	out.Fill = value2
	value3 = (input.Get("iterationStart")).Float()
	out.IterationStart = value3
	value4 = (input.Get("iterations")).Float()
	out.Iterations = value4
	value5 = UnionFromJS(input.Get("duration"))
	out.Duration = value5
	value6 = PlaybackDirectionFromJS(input.Get("direction"))
	out.Direction = value6
	value7 = (input.Get("easing")).String()
	out.Easing = value7
	return &out
}

// dictionary: KeyframeAnimationOptions
type KeyframeAnimationOptions struct {
	Delay              float64
	EndDelay           float64
	Fill               FillMode
	IterationStart     float64
	Iterations         float64
	Duration           *Union
	Direction          PlaybackDirection
	Easing             string
	IterationComposite IterationCompositeOperation
	Composite          CompositeOperation
	Id                 string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *KeyframeAnimationOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Delay
	out.Set("delay", value0)
	value1 := _this.EndDelay
	out.Set("endDelay", value1)
	value2 := _this.Fill.JSValue()
	out.Set("fill", value2)
	value3 := _this.IterationStart
	out.Set("iterationStart", value3)
	value4 := _this.Iterations
	out.Set("iterations", value4)
	value5 := _this.Duration.JSValue()
	out.Set("duration", value5)
	value6 := _this.Direction.JSValue()
	out.Set("direction", value6)
	value7 := _this.Easing
	out.Set("easing", value7)
	value8 := _this.IterationComposite.JSValue()
	out.Set("iterationComposite", value8)
	value9 := _this.Composite.JSValue()
	out.Set("composite", value9)
	value10 := _this.Id
	out.Set("id", value10)
	return out
}

// KeyframeAnimationOptionsFromJS is allocating a new
// KeyframeAnimationOptions object and copy all values from
// input javascript object
func KeyframeAnimationOptionsFromJS(value js.Wrapper) *KeyframeAnimationOptions {
	input := value.JSValue()
	var out KeyframeAnimationOptions
	var (
		value0  float64                     // javascript: double {delay Delay delay}
		value1  float64                     // javascript: double {endDelay EndDelay endDelay}
		value2  FillMode                    // javascript: FillMode {fill Fill fill}
		value3  float64                     // javascript: double {iterationStart IterationStart iterationStart}
		value4  float64                     // javascript: unrestricted double {iterations Iterations iterations}
		value5  *Union                      // javascript: Union {duration Duration duration}
		value6  PlaybackDirection           // javascript: PlaybackDirection {direction Direction direction}
		value7  string                      // javascript: DOMString {easing Easing easing}
		value8  IterationCompositeOperation // javascript: IterationCompositeOperation {iterationComposite IterationComposite iterationComposite}
		value9  CompositeOperation          // javascript: CompositeOperation {composite Composite composite}
		value10 string                      // javascript: DOMString {id Id id}
	)
	value0 = (input.Get("delay")).Float()
	out.Delay = value0
	value1 = (input.Get("endDelay")).Float()
	out.EndDelay = value1
	value2 = FillModeFromJS(input.Get("fill"))
	out.Fill = value2
	value3 = (input.Get("iterationStart")).Float()
	out.IterationStart = value3
	value4 = (input.Get("iterations")).Float()
	out.Iterations = value4
	value5 = UnionFromJS(input.Get("duration"))
	out.Duration = value5
	value6 = PlaybackDirectionFromJS(input.Get("direction"))
	out.Direction = value6
	value7 = (input.Get("easing")).String()
	out.Easing = value7
	value8 = IterationCompositeOperationFromJS(input.Get("iterationComposite"))
	out.IterationComposite = value8
	value9 = CompositeOperationFromJS(input.Get("composite"))
	out.Composite = value9
	value10 = (input.Get("id")).String()
	out.Id = value10
	return &out
}

// dictionary: KeyframeEffectOptions
type KeyframeEffectOptions struct {
	Delay              float64
	EndDelay           float64
	Fill               FillMode
	IterationStart     float64
	Iterations         float64
	Duration           *Union
	Direction          PlaybackDirection
	Easing             string
	IterationComposite IterationCompositeOperation
	Composite          CompositeOperation
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *KeyframeEffectOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Delay
	out.Set("delay", value0)
	value1 := _this.EndDelay
	out.Set("endDelay", value1)
	value2 := _this.Fill.JSValue()
	out.Set("fill", value2)
	value3 := _this.IterationStart
	out.Set("iterationStart", value3)
	value4 := _this.Iterations
	out.Set("iterations", value4)
	value5 := _this.Duration.JSValue()
	out.Set("duration", value5)
	value6 := _this.Direction.JSValue()
	out.Set("direction", value6)
	value7 := _this.Easing
	out.Set("easing", value7)
	value8 := _this.IterationComposite.JSValue()
	out.Set("iterationComposite", value8)
	value9 := _this.Composite.JSValue()
	out.Set("composite", value9)
	return out
}

// KeyframeEffectOptionsFromJS is allocating a new
// KeyframeEffectOptions object and copy all values from
// input javascript object
func KeyframeEffectOptionsFromJS(value js.Wrapper) *KeyframeEffectOptions {
	input := value.JSValue()
	var out KeyframeEffectOptions
	var (
		value0 float64                     // javascript: double {delay Delay delay}
		value1 float64                     // javascript: double {endDelay EndDelay endDelay}
		value2 FillMode                    // javascript: FillMode {fill Fill fill}
		value3 float64                     // javascript: double {iterationStart IterationStart iterationStart}
		value4 float64                     // javascript: unrestricted double {iterations Iterations iterations}
		value5 *Union                      // javascript: Union {duration Duration duration}
		value6 PlaybackDirection           // javascript: PlaybackDirection {direction Direction direction}
		value7 string                      // javascript: DOMString {easing Easing easing}
		value8 IterationCompositeOperation // javascript: IterationCompositeOperation {iterationComposite IterationComposite iterationComposite}
		value9 CompositeOperation          // javascript: CompositeOperation {composite Composite composite}
	)
	value0 = (input.Get("delay")).Float()
	out.Delay = value0
	value1 = (input.Get("endDelay")).Float()
	out.EndDelay = value1
	value2 = FillModeFromJS(input.Get("fill"))
	out.Fill = value2
	value3 = (input.Get("iterationStart")).Float()
	out.IterationStart = value3
	value4 = (input.Get("iterations")).Float()
	out.Iterations = value4
	value5 = UnionFromJS(input.Get("duration"))
	out.Duration = value5
	value6 = PlaybackDirectionFromJS(input.Get("direction"))
	out.Direction = value6
	value7 = (input.Get("easing")).String()
	out.Easing = value7
	value8 = IterationCompositeOperationFromJS(input.Get("iterationComposite"))
	out.IterationComposite = value8
	value9 = CompositeOperationFromJS(input.Get("composite"))
	out.Composite = value9
	return &out
}

// dictionary: OptionalEffectTiming
type OptionalEffectTiming struct {
	Delay          float64
	EndDelay       float64
	Fill           FillMode
	IterationStart float64
	Iterations     float64
	Duration       *Union
	Direction      PlaybackDirection
	Easing         string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *OptionalEffectTiming) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Delay
	out.Set("delay", value0)
	value1 := _this.EndDelay
	out.Set("endDelay", value1)
	value2 := _this.Fill.JSValue()
	out.Set("fill", value2)
	value3 := _this.IterationStart
	out.Set("iterationStart", value3)
	value4 := _this.Iterations
	out.Set("iterations", value4)
	value5 := _this.Duration.JSValue()
	out.Set("duration", value5)
	value6 := _this.Direction.JSValue()
	out.Set("direction", value6)
	value7 := _this.Easing
	out.Set("easing", value7)
	return out
}

// OptionalEffectTimingFromJS is allocating a new
// OptionalEffectTiming object and copy all values from
// input javascript object
func OptionalEffectTimingFromJS(value js.Wrapper) *OptionalEffectTiming {
	input := value.JSValue()
	var out OptionalEffectTiming
	var (
		value0 float64           // javascript: double {delay Delay delay}
		value1 float64           // javascript: double {endDelay EndDelay endDelay}
		value2 FillMode          // javascript: FillMode {fill Fill fill}
		value3 float64           // javascript: double {iterationStart IterationStart iterationStart}
		value4 float64           // javascript: unrestricted double {iterations Iterations iterations}
		value5 *Union            // javascript: Union {duration Duration duration}
		value6 PlaybackDirection // javascript: PlaybackDirection {direction Direction direction}
		value7 string            // javascript: DOMString {easing Easing easing}
	)
	value0 = (input.Get("delay")).Float()
	out.Delay = value0
	value1 = (input.Get("endDelay")).Float()
	out.EndDelay = value1
	value2 = FillModeFromJS(input.Get("fill"))
	out.Fill = value2
	value3 = (input.Get("iterationStart")).Float()
	out.IterationStart = value3
	value4 = (input.Get("iterations")).Float()
	out.Iterations = value4
	value5 = UnionFromJS(input.Get("duration"))
	out.Duration = value5
	value6 = PlaybackDirectionFromJS(input.Get("direction"))
	out.Direction = value6
	value7 = (input.Get("easing")).String()
	out.Easing = value7
	return &out
}

// interface: Animation
type Animation struct {
	domcore.EventTarget
}

// AnimationFromJS is casting a js.Wrapper into Animation.
func AnimationFromJS(value js.Wrapper) *Animation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Animation{}
	ret.Value_JS = input
	return ret
}

func NewAnimation(effect *AnimationEffect, timeline *AnimationTimeline) (_result *Animation) {
	_klass := js.Global().Get("Animation")
	var (
		_args [2]interface{}
		_end  int
	)
	if effect != nil {
		_p0 := effect.JSValue()
		_args[0] = _p0
		_end++
	}
	if timeline != nil {
		_p1 := timeline.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Animation // javascript: Animation _what_return_name
	)
	_converted = AnimationFromJS(_returned)
	_result = _converted
	return
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *Animation) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// SetId setting attribute 'id' with
// type string (idl: DOMString).
func (_this *Animation) SetId(value string) {
	input := value
	_this.Value_JS.Set("id", input)
}

// Effect returning attribute 'effect' with
// type AnimationEffect (idl: AnimationEffect).
func (_this *Animation) Effect() *AnimationEffect {
	var ret *AnimationEffect
	value := _this.Value_JS.Get("effect")
	if value.Type() != js.TypeNull {
		ret = AnimationEffectFromJS(value)
	}
	return ret
}

// SetEffect setting attribute 'effect' with
// type AnimationEffect (idl: AnimationEffect).
func (_this *Animation) SetEffect(value *AnimationEffect) {
	input := value.JSValue()
	_this.Value_JS.Set("effect", input)
}

// Timeline returning attribute 'timeline' with
// type AnimationTimeline (idl: AnimationTimeline).
func (_this *Animation) Timeline() *AnimationTimeline {
	var ret *AnimationTimeline
	value := _this.Value_JS.Get("timeline")
	if value.Type() != js.TypeNull {
		ret = AnimationTimelineFromJS(value)
	}
	return ret
}

// SetTimeline setting attribute 'timeline' with
// type AnimationTimeline (idl: AnimationTimeline).
func (_this *Animation) SetTimeline(value *AnimationTimeline) {
	input := value.JSValue()
	_this.Value_JS.Set("timeline", input)
}

// StartTime returning attribute 'startTime' with
// type float64 (idl: double).
func (_this *Animation) StartTime() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("startTime")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// SetStartTime setting attribute 'startTime' with
// type float64 (idl: double).
func (_this *Animation) SetStartTime(value *float64) {
	input := value
	_this.Value_JS.Set("startTime", input)
}

// CurrentTime returning attribute 'currentTime' with
// type float64 (idl: double).
func (_this *Animation) CurrentTime() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("currentTime")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// SetCurrentTime setting attribute 'currentTime' with
// type float64 (idl: double).
func (_this *Animation) SetCurrentTime(value *float64) {
	input := value
	_this.Value_JS.Set("currentTime", input)
}

// PlaybackRate returning attribute 'playbackRate' with
// type float64 (idl: double).
func (_this *Animation) PlaybackRate() float64 {
	var ret float64
	value := _this.Value_JS.Get("playbackRate")
	ret = (value).Float()
	return ret
}

// SetPlaybackRate setting attribute 'playbackRate' with
// type float64 (idl: double).
func (_this *Animation) SetPlaybackRate(value float64) {
	input := value
	_this.Value_JS.Set("playbackRate", input)
}

// PlayState returning attribute 'playState' with
// type AnimationPlayState (idl: AnimationPlayState).
func (_this *Animation) PlayState() AnimationPlayState {
	var ret AnimationPlayState
	value := _this.Value_JS.Get("playState")
	ret = AnimationPlayStateFromJS(value)
	return ret
}

// Pending returning attribute 'pending' with
// type bool (idl: boolean).
func (_this *Animation) Pending() bool {
	var ret bool
	value := _this.Value_JS.Get("pending")
	ret = (value).Bool()
	return ret
}

// Ready returning attribute 'ready' with
// type javascript.Promise (idl: Promise).
func (_this *Animation) Ready() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("ready")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// Finished returning attribute 'finished' with
// type javascript.Promise (idl: Promise).
func (_this *Animation) Finished() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("finished")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// Onfinish returning attribute 'onfinish' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Animation) Onfinish() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onfinish")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnfinish setting attribute 'onfinish' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Animation) SetOnfinish(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onfinish", input)
}

// Oncancel returning attribute 'oncancel' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Animation) Oncancel() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("oncancel")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOncancel setting attribute 'oncancel' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Animation) SetOncancel(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("oncancel", input)
}

func (_this *Animation) Cancel() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("cancel", _args[0:_end]...)
	return
}

func (_this *Animation) Finish() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("finish", _args[0:_end]...)
	return
}

func (_this *Animation) Play() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("play", _args[0:_end]...)
	return
}

func (_this *Animation) Pause() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("pause", _args[0:_end]...)
	return
}

func (_this *Animation) UpdatePlaybackRate(playbackRate float64) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := playbackRate
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("updatePlaybackRate", _args[0:_end]...)
	return
}

func (_this *Animation) Reverse() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("reverse", _args[0:_end]...)
	return
}

// interface: AnimationEffect
type AnimationEffect struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *AnimationEffect) JSValue() js.Value {
	return _this.Value_JS
}

// AnimationEffectFromJS is casting a js.Wrapper into AnimationEffect.
func AnimationEffectFromJS(value js.Wrapper) *AnimationEffect {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AnimationEffect{}
	ret.Value_JS = input
	return ret
}

func (_this *AnimationEffect) GetTiming() (_result *EffectTiming) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getTiming", _args[0:_end]...)
	var (
		_converted *EffectTiming // javascript: EffectTiming _what_return_name
	)
	_converted = EffectTimingFromJS(_returned)
	_result = _converted
	return
}

func (_this *AnimationEffect) GetComputedTiming() (_result *ComputedEffectTiming) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getComputedTiming", _args[0:_end]...)
	var (
		_converted *ComputedEffectTiming // javascript: ComputedEffectTiming _what_return_name
	)
	_converted = ComputedEffectTimingFromJS(_returned)
	_result = _converted
	return
}

func (_this *AnimationEffect) UpdateTiming(timing *OptionalEffectTiming) {
	var (
		_args [1]interface{}
		_end  int
	)
	if timing != nil {
		_p0 := timing.JSValue()
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("updateTiming", _args[0:_end]...)
	return
}

// interface: AnimationPlaybackEvent
type AnimationPlaybackEvent struct {
	domcore.Event
}

// AnimationPlaybackEventFromJS is casting a js.Wrapper into AnimationPlaybackEvent.
func AnimationPlaybackEventFromJS(value js.Wrapper) *AnimationPlaybackEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AnimationPlaybackEvent{}
	ret.Value_JS = input
	return ret
}

func NewAnimationPlaybackEvent(_type string, eventInitDict *AnimationPlaybackEventInit) (_result *AnimationPlaybackEvent) {
	_klass := js.Global().Get("AnimationPlaybackEvent")
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
		_converted *AnimationPlaybackEvent // javascript: AnimationPlaybackEvent _what_return_name
	)
	_converted = AnimationPlaybackEventFromJS(_returned)
	_result = _converted
	return
}

// CurrentTime returning attribute 'currentTime' with
// type float64 (idl: double).
func (_this *AnimationPlaybackEvent) CurrentTime() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("currentTime")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// TimelineTime returning attribute 'timelineTime' with
// type float64 (idl: double).
func (_this *AnimationPlaybackEvent) TimelineTime() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("timelineTime")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// interface: AnimationTimeline
type AnimationTimeline struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *AnimationTimeline) JSValue() js.Value {
	return _this.Value_JS
}

// AnimationTimelineFromJS is casting a js.Wrapper into AnimationTimeline.
func AnimationTimelineFromJS(value js.Wrapper) *AnimationTimeline {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &AnimationTimeline{}
	ret.Value_JS = input
	return ret
}

// CurrentTime returning attribute 'currentTime' with
// type float64 (idl: double).
func (_this *AnimationTimeline) CurrentTime() *float64 {
	var ret *float64
	value := _this.Value_JS.Get("currentTime")
	if value.Type() != js.TypeNull {
		__tmp := (value).Float()
		ret = &__tmp
	}
	return ret
}

// interface: DocumentTimeline
type DocumentTimeline struct {
	AnimationTimeline
}

// DocumentTimelineFromJS is casting a js.Wrapper into DocumentTimeline.
func DocumentTimelineFromJS(value js.Wrapper) *DocumentTimeline {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DocumentTimeline{}
	ret.Value_JS = input
	return ret
}

func NewDocumentTimeline(options *DocumentTimelineOptions) (_result *DocumentTimeline) {
	_klass := js.Global().Get("DocumentTimeline")
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DocumentTimeline // javascript: DocumentTimeline _what_return_name
	)
	_converted = DocumentTimelineFromJS(_returned)
	_result = _converted
	return
}

// interface: KeyframeEffect
type KeyframeEffect struct {
	AnimationEffect
}

// KeyframeEffectFromJS is casting a js.Wrapper into KeyframeEffect.
func KeyframeEffectFromJS(value js.Wrapper) *KeyframeEffect {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &KeyframeEffect{}
	ret.Value_JS = input
	return ret
}

func NewKeyframeEffect(source *KeyframeEffect) (_result *KeyframeEffect) {
	_klass := js.Global().Get("KeyframeEffect")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := source.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *KeyframeEffect // javascript: KeyframeEffect _what_return_name
	)
	_converted = KeyframeEffectFromJS(_returned)
	_result = _converted
	return
}

// Target returning attribute 'target' with
// type Union (idl: Union).
func (_this *KeyframeEffect) Target() *Union {
	var ret *Union
	value := _this.Value_JS.Get("target")
	if value.Type() != js.TypeNull {
		ret = UnionFromJS(value)
	}
	return ret
}

// SetTarget setting attribute 'target' with
// type Union (idl: Union).
func (_this *KeyframeEffect) SetTarget(value *Union) {
	input := value.JSValue()
	_this.Value_JS.Set("target", input)
}

// IterationComposite returning attribute 'iterationComposite' with
// type IterationCompositeOperation (idl: IterationCompositeOperation).
func (_this *KeyframeEffect) IterationComposite() IterationCompositeOperation {
	var ret IterationCompositeOperation
	value := _this.Value_JS.Get("iterationComposite")
	ret = IterationCompositeOperationFromJS(value)
	return ret
}

// SetIterationComposite setting attribute 'iterationComposite' with
// type IterationCompositeOperation (idl: IterationCompositeOperation).
func (_this *KeyframeEffect) SetIterationComposite(value IterationCompositeOperation) {
	input := value.JSValue()
	_this.Value_JS.Set("iterationComposite", input)
}

// Composite returning attribute 'composite' with
// type CompositeOperation (idl: CompositeOperation).
func (_this *KeyframeEffect) Composite() CompositeOperation {
	var ret CompositeOperation
	value := _this.Value_JS.Get("composite")
	ret = CompositeOperationFromJS(value)
	return ret
}

// SetComposite setting attribute 'composite' with
// type CompositeOperation (idl: CompositeOperation).
func (_this *KeyframeEffect) SetComposite(value CompositeOperation) {
	input := value.JSValue()
	_this.Value_JS.Set("composite", input)
}

func (_this *KeyframeEffect) GetKeyframes() (_result []*javascript.Object) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getKeyframes", _args[0:_end]...)
	var (
		_converted []*javascript.Object // javascript: sequence<object> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*javascript.Object, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *javascript.Object
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = javascript.ObjectFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *KeyframeEffect) SetKeyframes(keyframes *javascript.Object) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := keyframes.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("setKeyframes", _args[0:_end]...)
	return
}
