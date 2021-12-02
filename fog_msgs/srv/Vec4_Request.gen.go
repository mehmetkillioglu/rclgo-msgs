/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/vec4.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/Vec4_Request", Vec4_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewVec4_Request
// function instead.
type Vec4_Request struct {
	Goal []float64 `yaml:"goal"`
}

// NewVec4_Request creates a new Vec4_Request with default values.
func NewVec4_Request() *Vec4_Request {
	self := Vec4_Request{}
	self.SetDefaults()
	return &self
}

func (t *Vec4_Request) Clone() *Vec4_Request {
	c := &Vec4_Request{}
	if t.Goal != nil {
		c.Goal = make([]float64, len(t.Goal))
		copy(c.Goal, t.Goal)
	}
	return c
}

func (t *Vec4_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Vec4_Request) SetDefaults() {
	t.Goal = nil
}

// CloneVec4_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVec4_RequestSlice(dst, src []Vec4_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Vec4_RequestTypeSupport types.MessageTypeSupport = _Vec4_RequestTypeSupport{}

type _Vec4_RequestTypeSupport struct{}

func (t _Vec4_RequestTypeSupport) New() types.Message {
	return NewVec4_Request()
}

func (t _Vec4_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__Vec4_Request
	return (unsafe.Pointer)(C.fog_msgs__srv__Vec4_Request__create())
}

func (t _Vec4_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__Vec4_Request__destroy((*C.fog_msgs__srv__Vec4_Request)(pointer_to_free))
}

func (t _Vec4_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Vec4_Request)
	mem := (*C.fog_msgs__srv__Vec4_Request)(dst)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.goal)), m.Goal)
}

func (t _Vec4_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Vec4_Request)
	mem := (*C.fog_msgs__srv__Vec4_Request)(ros2_message_buffer)
	primitives.Float64__Sequence_to_Go(&m.Goal, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.goal)))
}

func (t _Vec4_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__Vec4_Request())
}

type CVec4_Request = C.fog_msgs__srv__Vec4_Request
type CVec4_Request__Sequence = C.fog_msgs__srv__Vec4_Request__Sequence

func Vec4_Request__Sequence_to_Go(goSlice *[]Vec4_Request, cSlice CVec4_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Vec4_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__Vec4_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Vec4_Request * uintptr(i)),
		))
		Vec4_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Vec4_Request__Sequence_to_C(cSlice *CVec4_Request__Sequence, goSlice []Vec4_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__Vec4_Request)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__Vec4_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__Vec4_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__Vec4_Request * uintptr(i)),
		))
		Vec4_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Vec4_Request__Array_to_Go(goSlice []Vec4_Request, cSlice []CVec4_Request) {
	for i := 0; i < len(cSlice); i++ {
		Vec4_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Vec4_Request__Array_to_C(cSlice []CVec4_Request, goSlice []Vec4_Request) {
	for i := 0; i < len(goSlice); i++ {
		Vec4_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
