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

#include <gazebo_msgs/srv/get_link_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/GetLinkState_Request", GetLinkState_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetLinkState_Request
// function instead.
type GetLinkState_Request struct {
	LinkName string `yaml:"link_name"`// name of link. Deprecated, kept for ROS 1 bridge.Use GetEntityState
	ReferenceFrame string `yaml:"reference_frame"`// reference frame of returned information, must be a valid link. Deprecated, kept for ROS 1 bridge.Use GetEntityStatelink names are prefixed by model name, e.g. pr2::base_link
}

// NewGetLinkState_Request creates a new GetLinkState_Request with default values.
func NewGetLinkState_Request() *GetLinkState_Request {
	self := GetLinkState_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetLinkState_Request) Clone() *GetLinkState_Request {
	c := &GetLinkState_Request{}
	c.LinkName = t.LinkName
	c.ReferenceFrame = t.ReferenceFrame
	return c
}

func (t *GetLinkState_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetLinkState_Request) SetDefaults() {
	t.LinkName = ""
	t.ReferenceFrame = ""
}

// CloneGetLinkState_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetLinkState_RequestSlice(dst, src []GetLinkState_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetLinkState_RequestTypeSupport types.MessageTypeSupport = _GetLinkState_RequestTypeSupport{}

type _GetLinkState_RequestTypeSupport struct{}

func (t _GetLinkState_RequestTypeSupport) New() types.Message {
	return NewGetLinkState_Request()
}

func (t _GetLinkState_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__GetLinkState_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__GetLinkState_Request__create())
}

func (t _GetLinkState_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__GetLinkState_Request__destroy((*C.gazebo_msgs__srv__GetLinkState_Request)(pointer_to_free))
}

func (t _GetLinkState_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetLinkState_Request)
	mem := (*C.gazebo_msgs__srv__GetLinkState_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.link_name), m.LinkName)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.reference_frame), m.ReferenceFrame)
}

func (t _GetLinkState_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetLinkState_Request)
	mem := (*C.gazebo_msgs__srv__GetLinkState_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.LinkName, unsafe.Pointer(&mem.link_name))
	primitives.StringAsGoStruct(&m.ReferenceFrame, unsafe.Pointer(&mem.reference_frame))
}

func (t _GetLinkState_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__GetLinkState_Request())
}

type CGetLinkState_Request = C.gazebo_msgs__srv__GetLinkState_Request
type CGetLinkState_Request__Sequence = C.gazebo_msgs__srv__GetLinkState_Request__Sequence

func GetLinkState_Request__Sequence_to_Go(goSlice *[]GetLinkState_Request, cSlice CGetLinkState_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetLinkState_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__GetLinkState_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Request * uintptr(i)),
		))
		GetLinkState_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetLinkState_Request__Sequence_to_C(cSlice *CGetLinkState_Request__Sequence, goSlice []GetLinkState_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__GetLinkState_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__GetLinkState_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Request * uintptr(i)),
		))
		GetLinkState_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetLinkState_Request__Array_to_Go(goSlice []GetLinkState_Request, cSlice []CGetLinkState_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetLinkState_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetLinkState_Request__Array_to_C(cSlice []CGetLinkState_Request, goSlice []GetLinkState_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetLinkState_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
