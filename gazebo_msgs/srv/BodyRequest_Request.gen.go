/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package gazebo_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <gazebo_msgs/srv/body_request.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/BodyRequest_Request", BodyRequest_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewBodyRequest_Request
// function instead.
type BodyRequest_Request struct {
	BodyName string `yaml:"body_name"`// name of the body requested. body names are prefixed by model name, e.g. pr2::base_link. Deprecated, kept for ROS 1 bridge.Use LinkRequest
}

// NewBodyRequest_Request creates a new BodyRequest_Request with default values.
func NewBodyRequest_Request() *BodyRequest_Request {
	self := BodyRequest_Request{}
	self.SetDefaults()
	return &self
}

func (t *BodyRequest_Request) Clone() *BodyRequest_Request {
	c := &BodyRequest_Request{}
	c.BodyName = t.BodyName
	return c
}

func (t *BodyRequest_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BodyRequest_Request) SetDefaults() {
	t.BodyName = ""
}

// CloneBodyRequest_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBodyRequest_RequestSlice(dst, src []BodyRequest_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BodyRequest_RequestTypeSupport types.MessageTypeSupport = _BodyRequest_RequestTypeSupport{}

type _BodyRequest_RequestTypeSupport struct{}

func (t _BodyRequest_RequestTypeSupport) New() types.Message {
	return NewBodyRequest_Request()
}

func (t _BodyRequest_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__BodyRequest_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__BodyRequest_Request__create())
}

func (t _BodyRequest_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__BodyRequest_Request__destroy((*C.gazebo_msgs__srv__BodyRequest_Request)(pointer_to_free))
}

func (t _BodyRequest_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BodyRequest_Request)
	mem := (*C.gazebo_msgs__srv__BodyRequest_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.body_name), m.BodyName)
}

func (t _BodyRequest_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BodyRequest_Request)
	mem := (*C.gazebo_msgs__srv__BodyRequest_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.BodyName, unsafe.Pointer(&mem.body_name))
}

func (t _BodyRequest_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__BodyRequest_Request())
}

type CBodyRequest_Request = C.gazebo_msgs__srv__BodyRequest_Request
type CBodyRequest_Request__Sequence = C.gazebo_msgs__srv__BodyRequest_Request__Sequence

func BodyRequest_Request__Sequence_to_Go(goSlice *[]BodyRequest_Request, cSlice CBodyRequest_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BodyRequest_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__BodyRequest_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__BodyRequest_Request * uintptr(i)),
		))
		BodyRequest_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func BodyRequest_Request__Sequence_to_C(cSlice *CBodyRequest_Request__Sequence, goSlice []BodyRequest_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__BodyRequest_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__BodyRequest_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__BodyRequest_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__BodyRequest_Request * uintptr(i)),
		))
		BodyRequest_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func BodyRequest_Request__Array_to_Go(goSlice []BodyRequest_Request, cSlice []CBodyRequest_Request) {
	for i := 0; i < len(cSlice); i++ {
		BodyRequest_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BodyRequest_Request__Array_to_C(cSlice []CBodyRequest_Request, goSlice []BodyRequest_Request) {
	for i := 0; i < len(goSlice); i++ {
		BodyRequest_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
