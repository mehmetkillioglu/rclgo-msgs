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

#include <gazebo_msgs/srv/apply_link_wrench.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/ApplyLinkWrench_Response", ApplyLinkWrench_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewApplyLinkWrench_Response
// function instead.
type ApplyLinkWrench_Response struct {
	Success bool `yaml:"success"`// return true if set wrench successful. Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_linkif start_time is not specified, orstart_time < current time, start as soon as possibleif duration < 0, apply wrench continuously without endif duration = 0, do nothingif duration < step size, apply wrenchfor one step size
	StatusMessage string `yaml:"status_message"`// comments if available. Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_linkif start_time is not specified, orstart_time < current time, start as soon as possibleif duration < 0, apply wrench continuously without endif duration = 0, do nothingif duration < step size, apply wrenchfor one step size
}

// NewApplyLinkWrench_Response creates a new ApplyLinkWrench_Response with default values.
func NewApplyLinkWrench_Response() *ApplyLinkWrench_Response {
	self := ApplyLinkWrench_Response{}
	self.SetDefaults()
	return &self
}

func (t *ApplyLinkWrench_Response) Clone() *ApplyLinkWrench_Response {
	c := &ApplyLinkWrench_Response{}
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *ApplyLinkWrench_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ApplyLinkWrench_Response) SetDefaults() {
	t.Success = false
	t.StatusMessage = ""
}

// CloneApplyLinkWrench_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneApplyLinkWrench_ResponseSlice(dst, src []ApplyLinkWrench_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ApplyLinkWrench_ResponseTypeSupport types.MessageTypeSupport = _ApplyLinkWrench_ResponseTypeSupport{}

type _ApplyLinkWrench_ResponseTypeSupport struct{}

func (t _ApplyLinkWrench_ResponseTypeSupport) New() types.Message {
	return NewApplyLinkWrench_Response()
}

func (t _ApplyLinkWrench_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__ApplyLinkWrench_Response
	return (unsafe.Pointer)(C.gazebo_msgs__srv__ApplyLinkWrench_Response__create())
}

func (t _ApplyLinkWrench_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__ApplyLinkWrench_Response__destroy((*C.gazebo_msgs__srv__ApplyLinkWrench_Response)(pointer_to_free))
}

func (t _ApplyLinkWrench_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ApplyLinkWrench_Response)
	mem := (*C.gazebo_msgs__srv__ApplyLinkWrench_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _ApplyLinkWrench_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ApplyLinkWrench_Response)
	mem := (*C.gazebo_msgs__srv__ApplyLinkWrench_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _ApplyLinkWrench_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__ApplyLinkWrench_Response())
}

type CApplyLinkWrench_Response = C.gazebo_msgs__srv__ApplyLinkWrench_Response
type CApplyLinkWrench_Response__Sequence = C.gazebo_msgs__srv__ApplyLinkWrench_Response__Sequence

func ApplyLinkWrench_Response__Sequence_to_Go(goSlice *[]ApplyLinkWrench_Response, cSlice CApplyLinkWrench_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ApplyLinkWrench_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__ApplyLinkWrench_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Response * uintptr(i)),
		))
		ApplyLinkWrench_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ApplyLinkWrench_Response__Sequence_to_C(cSlice *CApplyLinkWrench_Response__Sequence, goSlice []ApplyLinkWrench_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__ApplyLinkWrench_Response)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__ApplyLinkWrench_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Response * uintptr(i)),
		))
		ApplyLinkWrench_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ApplyLinkWrench_Response__Array_to_Go(goSlice []ApplyLinkWrench_Response, cSlice []CApplyLinkWrench_Response) {
	for i := 0; i < len(cSlice); i++ {
		ApplyLinkWrench_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ApplyLinkWrench_Response__Array_to_C(cSlice []CApplyLinkWrench_Response, goSlice []ApplyLinkWrench_Response) {
	for i := 0; i < len(goSlice); i++ {
		ApplyLinkWrench_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}