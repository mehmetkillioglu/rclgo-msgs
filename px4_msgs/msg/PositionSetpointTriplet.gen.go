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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/position_setpoint_triplet.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/PositionSetpointTriplet", PositionSetpointTripletTypeSupport)
}

// Do not create instances of this type directly. Always use NewPositionSetpointTriplet
// function instead.
type PositionSetpointTriplet struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Previous PositionSetpoint `yaml:"previous"`
	Current PositionSetpoint `yaml:"current"`
	Next PositionSetpoint `yaml:"next"`
}

// NewPositionSetpointTriplet creates a new PositionSetpointTriplet with default values.
func NewPositionSetpointTriplet() *PositionSetpointTriplet {
	self := PositionSetpointTriplet{}
	self.SetDefaults()
	return &self
}

func (t *PositionSetpointTriplet) Clone() *PositionSetpointTriplet {
	c := &PositionSetpointTriplet{}
	c.Timestamp = t.Timestamp
	c.Previous = *t.Previous.Clone()
	c.Current = *t.Current.Clone()
	c.Next = *t.Next.Clone()
	return c
}

func (t *PositionSetpointTriplet) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PositionSetpointTriplet) SetDefaults() {
	t.Timestamp = 0
	t.Previous.SetDefaults()
	t.Current.SetDefaults()
	t.Next.SetDefaults()
}

// ClonePositionSetpointTripletSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePositionSetpointTripletSlice(dst, src []PositionSetpointTriplet) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PositionSetpointTripletTypeSupport types.MessageTypeSupport = _PositionSetpointTripletTypeSupport{}

type _PositionSetpointTripletTypeSupport struct{}

func (t _PositionSetpointTripletTypeSupport) New() types.Message {
	return NewPositionSetpointTriplet()
}

func (t _PositionSetpointTripletTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__PositionSetpointTriplet
	return (unsafe.Pointer)(C.px4_msgs__msg__PositionSetpointTriplet__create())
}

func (t _PositionSetpointTripletTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__PositionSetpointTriplet__destroy((*C.px4_msgs__msg__PositionSetpointTriplet)(pointer_to_free))
}

func (t _PositionSetpointTripletTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PositionSetpointTriplet)
	mem := (*C.px4_msgs__msg__PositionSetpointTriplet)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	PositionSetpointTypeSupport.AsCStruct(unsafe.Pointer(&mem.previous), &m.Previous)
	PositionSetpointTypeSupport.AsCStruct(unsafe.Pointer(&mem.current), &m.Current)
	PositionSetpointTypeSupport.AsCStruct(unsafe.Pointer(&mem.next), &m.Next)
}

func (t _PositionSetpointTripletTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PositionSetpointTriplet)
	mem := (*C.px4_msgs__msg__PositionSetpointTriplet)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	PositionSetpointTypeSupport.AsGoStruct(&m.Previous, unsafe.Pointer(&mem.previous))
	PositionSetpointTypeSupport.AsGoStruct(&m.Current, unsafe.Pointer(&mem.current))
	PositionSetpointTypeSupport.AsGoStruct(&m.Next, unsafe.Pointer(&mem.next))
}

func (t _PositionSetpointTripletTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__PositionSetpointTriplet())
}

type CPositionSetpointTriplet = C.px4_msgs__msg__PositionSetpointTriplet
type CPositionSetpointTriplet__Sequence = C.px4_msgs__msg__PositionSetpointTriplet__Sequence

func PositionSetpointTriplet__Sequence_to_Go(goSlice *[]PositionSetpointTriplet, cSlice CPositionSetpointTriplet__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PositionSetpointTriplet, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__PositionSetpointTriplet__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionSetpointTriplet * uintptr(i)),
		))
		PositionSetpointTripletTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PositionSetpointTriplet__Sequence_to_C(cSlice *CPositionSetpointTriplet__Sequence, goSlice []PositionSetpointTriplet) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__PositionSetpointTriplet)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__PositionSetpointTriplet * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__PositionSetpointTriplet)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__PositionSetpointTriplet * uintptr(i)),
		))
		PositionSetpointTripletTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PositionSetpointTriplet__Array_to_Go(goSlice []PositionSetpointTriplet, cSlice []CPositionSetpointTriplet) {
	for i := 0; i < len(cSlice); i++ {
		PositionSetpointTripletTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PositionSetpointTriplet__Array_to_C(cSlice []CPositionSetpointTriplet, goSlice []PositionSetpointTriplet) {
	for i := 0; i < len(goSlice); i++ {
		PositionSetpointTripletTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
