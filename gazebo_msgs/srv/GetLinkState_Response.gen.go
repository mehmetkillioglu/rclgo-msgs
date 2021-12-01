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
	gazebo_msgs_msg "github.com/tiiuae/rclgo-msgs/gazebo_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <gazebo_msgs/srv/get_link_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/GetLinkState_Response", GetLinkState_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetLinkState_Response
// function instead.
type GetLinkState_Response struct {
	LinkState gazebo_msgs_msg.LinkState `yaml:"link_state"`// Deprecated, kept for ROS 1 bridge.Use GetEntityStatelink names are prefixed by model name, e.g. pr2::base_linkif empty, use inertial (gazebo world) framereference_frame names are prefixed by model name, e.g. pr2::base_link
	Success bool `yaml:"success"`// return true if get info is successful. Deprecated, kept for ROS 1 bridge.Use GetEntityStatelink names are prefixed by model name, e.g. pr2::base_linkif empty, use inertial (gazebo world) framereference_frame names are prefixed by model name, e.g. pr2::base_link
	StatusMessage string `yaml:"status_message"`// comments if available. Deprecated, kept for ROS 1 bridge.Use GetEntityStatelink names are prefixed by model name, e.g. pr2::base_linkif empty, use inertial (gazebo world) framereference_frame names are prefixed by model name, e.g. pr2::base_link
}

// NewGetLinkState_Response creates a new GetLinkState_Response with default values.
func NewGetLinkState_Response() *GetLinkState_Response {
	self := GetLinkState_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetLinkState_Response) Clone() *GetLinkState_Response {
	c := &GetLinkState_Response{}
	c.LinkState = *t.LinkState.Clone()
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *GetLinkState_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetLinkState_Response) SetDefaults() {
	t.LinkState.SetDefaults()
	t.Success = false
	t.StatusMessage = ""
}

// CloneGetLinkState_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetLinkState_ResponseSlice(dst, src []GetLinkState_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetLinkState_ResponseTypeSupport types.MessageTypeSupport = _GetLinkState_ResponseTypeSupport{}

type _GetLinkState_ResponseTypeSupport struct{}

func (t _GetLinkState_ResponseTypeSupport) New() types.Message {
	return NewGetLinkState_Response()
}

func (t _GetLinkState_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__GetLinkState_Response
	return (unsafe.Pointer)(C.gazebo_msgs__srv__GetLinkState_Response__create())
}

func (t _GetLinkState_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__GetLinkState_Response__destroy((*C.gazebo_msgs__srv__GetLinkState_Response)(pointer_to_free))
}

func (t _GetLinkState_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetLinkState_Response)
	mem := (*C.gazebo_msgs__srv__GetLinkState_Response)(dst)
	gazebo_msgs_msg.LinkStateTypeSupport.AsCStruct(unsafe.Pointer(&mem.link_state), &m.LinkState)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _GetLinkState_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetLinkState_Response)
	mem := (*C.gazebo_msgs__srv__GetLinkState_Response)(ros2_message_buffer)
	gazebo_msgs_msg.LinkStateTypeSupport.AsGoStruct(&m.LinkState, unsafe.Pointer(&mem.link_state))
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _GetLinkState_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__GetLinkState_Response())
}

type CGetLinkState_Response = C.gazebo_msgs__srv__GetLinkState_Response
type CGetLinkState_Response__Sequence = C.gazebo_msgs__srv__GetLinkState_Response__Sequence

func GetLinkState_Response__Sequence_to_Go(goSlice *[]GetLinkState_Response, cSlice CGetLinkState_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetLinkState_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__GetLinkState_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Response * uintptr(i)),
		))
		GetLinkState_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetLinkState_Response__Sequence_to_C(cSlice *CGetLinkState_Response__Sequence, goSlice []GetLinkState_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__GetLinkState_Response)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__GetLinkState_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetLinkState_Response * uintptr(i)),
		))
		GetLinkState_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetLinkState_Response__Array_to_Go(goSlice []GetLinkState_Response, cSlice []CGetLinkState_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetLinkState_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetLinkState_Response__Array_to_C(cSlice []CGetLinkState_Response, goSlice []GetLinkState_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetLinkState_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}