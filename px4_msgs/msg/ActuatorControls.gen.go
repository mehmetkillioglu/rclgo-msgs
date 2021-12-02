/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/actuator_controls.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControls", ActuatorControlsTypeSupport)
}
const (
	ActuatorControls_NUM_ACTUATOR_CONTROLS uint8 = 8
	ActuatorControls_NUM_ACTUATOR_CONTROL_GROUPS uint8 = 6
	ActuatorControls_INDEX_ROLL uint8 = 0
	ActuatorControls_INDEX_PITCH uint8 = 1
	ActuatorControls_INDEX_YAW uint8 = 2
	ActuatorControls_INDEX_THROTTLE uint8 = 3
	ActuatorControls_INDEX_FLAPS uint8 = 4
	ActuatorControls_INDEX_SPOILERS uint8 = 5
	ActuatorControls_INDEX_AIRBRAKES uint8 = 6
	ActuatorControls_INDEX_LANDING_GEAR uint8 = 7
	ActuatorControls_INDEX_GIMBAL_SHUTTER uint8 = 3
	ActuatorControls_INDEX_CAMERA_ZOOM uint8 = 4
	ActuatorControls_GROUP_INDEX_ATTITUDE uint8 = 0
	ActuatorControls_GROUP_INDEX_ATTITUDE_ALTERNATE uint8 = 1
	ActuatorControls_GROUP_INDEX_GIMBAL uint8 = 2
	ActuatorControls_GROUP_INDEX_MANUAL_PASSTHROUGH uint8 = 3
	ActuatorControls_GROUP_INDEX_ALLOCATED_PART1 uint8 = 4
	ActuatorControls_GROUP_INDEX_ALLOCATED_PART2 uint8 = 5
	ActuatorControls_GROUP_INDEX_PAYLOAD uint8 = 6
)

// Do not create instances of this type directly. Always use NewActuatorControls
// function instead.
type ActuatorControls struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp the data this control response is based on was sampled
	Control [8]float32 `yaml:"control"`
}

// NewActuatorControls creates a new ActuatorControls with default values.
func NewActuatorControls() *ActuatorControls {
	self := ActuatorControls{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControls) Clone() *ActuatorControls {
	c := &ActuatorControls{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.Control = t.Control
	return c
}

func (t *ActuatorControls) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControls) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.Control = [8]float32{}
}

// CloneActuatorControlsSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControlsSlice(dst, src []ActuatorControls) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControlsTypeSupport types.MessageTypeSupport = _ActuatorControlsTypeSupport{}

type _ActuatorControlsTypeSupport struct{}

func (t _ActuatorControlsTypeSupport) New() types.Message {
	return NewActuatorControls()
}

func (t _ActuatorControlsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControls
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControls__create())
}

func (t _ActuatorControlsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControls__destroy((*C.px4_msgs__msg__ActuatorControls)(pointer_to_free))
}

func (t _ActuatorControlsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControls)
	mem := (*C.px4_msgs__msg__ActuatorControls)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)), m.Control[:])
}

func (t _ActuatorControlsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControls)
	mem := (*C.px4_msgs__msg__ActuatorControls)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	cSlice_control := mem.control[:]
	primitives.Float32__Array_to_Go(m.Control[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control)))
}

func (t _ActuatorControlsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControls())
}

type CActuatorControls = C.px4_msgs__msg__ActuatorControls
type CActuatorControls__Sequence = C.px4_msgs__msg__ActuatorControls__Sequence

func ActuatorControls__Sequence_to_Go(goSlice *[]ActuatorControls, cSlice CActuatorControls__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControls, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControls__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls * uintptr(i)),
		))
		ActuatorControlsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControls__Sequence_to_C(cSlice *CActuatorControls__Sequence, goSlice []ActuatorControls) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControls)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControls * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControls)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControls * uintptr(i)),
		))
		ActuatorControlsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControls__Array_to_Go(goSlice []ActuatorControls, cSlice []CActuatorControls) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControlsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControls__Array_to_C(cSlice []CActuatorControls, goSlice []ActuatorControls) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControlsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
