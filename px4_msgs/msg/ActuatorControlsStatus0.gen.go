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
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/actuator_controls_status0.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/ActuatorControlsStatus0", ActuatorControlsStatus0TypeSupport)
}
const (
	ActuatorControlsStatus0_INDEX_ROLL uint8 = 0
	ActuatorControlsStatus0_INDEX_PITCH uint8 = 1
	ActuatorControlsStatus0_INDEX_YAW uint8 = 2
	ActuatorControlsStatus0_INDEX_THROTTLE uint8 = 3
)

// Do not create instances of this type directly. Always use NewActuatorControlsStatus0
// function instead.
type ActuatorControlsStatus0 struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	ControlPower [4]float32 `yaml:"control_power"`
}

// NewActuatorControlsStatus0 creates a new ActuatorControlsStatus0 with default values.
func NewActuatorControlsStatus0() *ActuatorControlsStatus0 {
	self := ActuatorControlsStatus0{}
	self.SetDefaults()
	return &self
}

func (t *ActuatorControlsStatus0) Clone() *ActuatorControlsStatus0 {
	c := &ActuatorControlsStatus0{}
	c.Timestamp = t.Timestamp
	c.ControlPower = t.ControlPower
	return c
}

func (t *ActuatorControlsStatus0) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ActuatorControlsStatus0) SetDefaults() {
	t.Timestamp = 0
	t.ControlPower = [4]float32{}
}

// CloneActuatorControlsStatus0Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneActuatorControlsStatus0Slice(dst, src []ActuatorControlsStatus0) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ActuatorControlsStatus0TypeSupport types.MessageTypeSupport = _ActuatorControlsStatus0TypeSupport{}

type _ActuatorControlsStatus0TypeSupport struct{}

func (t _ActuatorControlsStatus0TypeSupport) New() types.Message {
	return NewActuatorControlsStatus0()
}

func (t _ActuatorControlsStatus0TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__ActuatorControlsStatus0
	return (unsafe.Pointer)(C.px4_msgs__msg__ActuatorControlsStatus0__create())
}

func (t _ActuatorControlsStatus0TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__ActuatorControlsStatus0__destroy((*C.px4_msgs__msg__ActuatorControlsStatus0)(pointer_to_free))
}

func (t _ActuatorControlsStatus0TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ActuatorControlsStatus0)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus0)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)), m.ControlPower[:])
}

func (t _ActuatorControlsStatus0TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ActuatorControlsStatus0)
	mem := (*C.px4_msgs__msg__ActuatorControlsStatus0)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_control_power := mem.control_power[:]
	primitives.Float32__Array_to_Go(m.ControlPower[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_control_power)))
}

func (t _ActuatorControlsStatus0TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__ActuatorControlsStatus0())
}

type CActuatorControlsStatus0 = C.px4_msgs__msg__ActuatorControlsStatus0
type CActuatorControlsStatus0__Sequence = C.px4_msgs__msg__ActuatorControlsStatus0__Sequence

func ActuatorControlsStatus0__Sequence_to_Go(goSlice *[]ActuatorControlsStatus0, cSlice CActuatorControlsStatus0__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ActuatorControlsStatus0, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus0__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus0 * uintptr(i)),
		))
		ActuatorControlsStatus0TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ActuatorControlsStatus0__Sequence_to_C(cSlice *CActuatorControlsStatus0__Sequence, goSlice []ActuatorControlsStatus0) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__ActuatorControlsStatus0)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus0 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__ActuatorControlsStatus0)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__ActuatorControlsStatus0 * uintptr(i)),
		))
		ActuatorControlsStatus0TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ActuatorControlsStatus0__Array_to_Go(goSlice []ActuatorControlsStatus0, cSlice []CActuatorControlsStatus0) {
	for i := 0; i < len(cSlice); i++ {
		ActuatorControlsStatus0TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ActuatorControlsStatus0__Array_to_C(cSlice []CActuatorControlsStatus0, goSlice []ActuatorControlsStatus0) {
	for i := 0; i < len(goSlice); i++ {
		ActuatorControlsStatus0TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
