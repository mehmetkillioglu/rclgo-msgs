/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/msg/future_trajectory.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/FutureTrajectory", FutureTrajectoryTypeSupport)
}

// Do not create instances of this type directly. Always use NewFutureTrajectory
// function instead.
type FutureTrajectory struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Poses []Vector4Stamped `yaml:"poses"`
}

// NewFutureTrajectory creates a new FutureTrajectory with default values.
func NewFutureTrajectory() *FutureTrajectory {
	self := FutureTrajectory{}
	self.SetDefaults()
	return &self
}

func (t *FutureTrajectory) Clone() *FutureTrajectory {
	c := &FutureTrajectory{}
	c.Header = *t.Header.Clone()
	if t.Poses != nil {
		c.Poses = make([]Vector4Stamped, len(t.Poses))
		CloneVector4StampedSlice(c.Poses, t.Poses)
	}
	return c
}

func (t *FutureTrajectory) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FutureTrajectory) SetDefaults() {
	t.Header.SetDefaults()
	t.Poses = nil
}

// CloneFutureTrajectorySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFutureTrajectorySlice(dst, src []FutureTrajectory) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FutureTrajectoryTypeSupport types.MessageTypeSupport = _FutureTrajectoryTypeSupport{}

type _FutureTrajectoryTypeSupport struct{}

func (t _FutureTrajectoryTypeSupport) New() types.Message {
	return NewFutureTrajectory()
}

func (t _FutureTrajectoryTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__msg__FutureTrajectory
	return (unsafe.Pointer)(C.fog_msgs__msg__FutureTrajectory__create())
}

func (t _FutureTrajectoryTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__msg__FutureTrajectory__destroy((*C.fog_msgs__msg__FutureTrajectory)(pointer_to_free))
}

func (t _FutureTrajectoryTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FutureTrajectory)
	mem := (*C.fog_msgs__msg__FutureTrajectory)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	Vector4Stamped__Sequence_to_C(&mem.poses, m.Poses)
}

func (t _FutureTrajectoryTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FutureTrajectory)
	mem := (*C.fog_msgs__msg__FutureTrajectory)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	Vector4Stamped__Sequence_to_Go(&m.Poses, mem.poses)
}

func (t _FutureTrajectoryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__msg__FutureTrajectory())
}

type CFutureTrajectory = C.fog_msgs__msg__FutureTrajectory
type CFutureTrajectory__Sequence = C.fog_msgs__msg__FutureTrajectory__Sequence

func FutureTrajectory__Sequence_to_Go(goSlice *[]FutureTrajectory, cSlice CFutureTrajectory__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FutureTrajectory, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__msg__FutureTrajectory__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__FutureTrajectory * uintptr(i)),
		))
		FutureTrajectoryTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func FutureTrajectory__Sequence_to_C(cSlice *CFutureTrajectory__Sequence, goSlice []FutureTrajectory) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__msg__FutureTrajectory)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__msg__FutureTrajectory * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__msg__FutureTrajectory)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__msg__FutureTrajectory * uintptr(i)),
		))
		FutureTrajectoryTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func FutureTrajectory__Array_to_Go(goSlice []FutureTrajectory, cSlice []CFutureTrajectory) {
	for i := 0; i < len(cSlice); i++ {
		FutureTrajectoryTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FutureTrajectory__Array_to_C(cSlice []CFutureTrajectory, goSlice []FutureTrajectory) {
	for i := 0; i < len(goSlice); i++ {
		FutureTrajectoryTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
