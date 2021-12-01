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

#include <gazebo_msgs/srv/get_world_properties.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/GetWorldProperties_Response", GetWorldProperties_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetWorldProperties_Response
// function instead.
type GetWorldProperties_Response struct {
	SimTime float64 `yaml:"sim_time"`// current sim time. Deprecated, kept for ROS 1 bridge.Use GetModelList
	ModelNames []string `yaml:"model_names"`// list of models in the world. Deprecated, kept for ROS 1 bridge.Use GetModelList
	RenderingEnabled bool `yaml:"rendering_enabled"`// if X is used. Deprecated, kept for ROS 1 bridge.Use GetModelList
	Success bool `yaml:"success"`// return true if get successful. Deprecated, kept for ROS 1 bridge.Use GetModelList
	StatusMessage string `yaml:"status_message"`// comments if available. Deprecated, kept for ROS 1 bridge.Use GetModelList
}

// NewGetWorldProperties_Response creates a new GetWorldProperties_Response with default values.
func NewGetWorldProperties_Response() *GetWorldProperties_Response {
	self := GetWorldProperties_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetWorldProperties_Response) Clone() *GetWorldProperties_Response {
	c := &GetWorldProperties_Response{}
	c.SimTime = t.SimTime
	if t.ModelNames != nil {
		c.ModelNames = make([]string, len(t.ModelNames))
		copy(c.ModelNames, t.ModelNames)
	}
	c.RenderingEnabled = t.RenderingEnabled
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *GetWorldProperties_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetWorldProperties_Response) SetDefaults() {
	t.SimTime = 0
	t.ModelNames = nil
	t.RenderingEnabled = false
	t.Success = false
	t.StatusMessage = ""
}

// CloneGetWorldProperties_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetWorldProperties_ResponseSlice(dst, src []GetWorldProperties_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetWorldProperties_ResponseTypeSupport types.MessageTypeSupport = _GetWorldProperties_ResponseTypeSupport{}

type _GetWorldProperties_ResponseTypeSupport struct{}

func (t _GetWorldProperties_ResponseTypeSupport) New() types.Message {
	return NewGetWorldProperties_Response()
}

func (t _GetWorldProperties_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__GetWorldProperties_Response
	return (unsafe.Pointer)(C.gazebo_msgs__srv__GetWorldProperties_Response__create())
}

func (t _GetWorldProperties_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__GetWorldProperties_Response__destroy((*C.gazebo_msgs__srv__GetWorldProperties_Response)(pointer_to_free))
}

func (t _GetWorldProperties_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetWorldProperties_Response)
	mem := (*C.gazebo_msgs__srv__GetWorldProperties_Response)(dst)
	mem.sim_time = C.double(m.SimTime)
	primitives.String__Sequence_to_C((*primitives.CString__Sequence)(unsafe.Pointer(&mem.model_names)), m.ModelNames)
	mem.rendering_enabled = C.bool(m.RenderingEnabled)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _GetWorldProperties_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetWorldProperties_Response)
	mem := (*C.gazebo_msgs__srv__GetWorldProperties_Response)(ros2_message_buffer)
	m.SimTime = float64(mem.sim_time)
	primitives.String__Sequence_to_Go(&m.ModelNames, *(*primitives.CString__Sequence)(unsafe.Pointer(&mem.model_names)))
	m.RenderingEnabled = bool(mem.rendering_enabled)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _GetWorldProperties_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__GetWorldProperties_Response())
}

type CGetWorldProperties_Response = C.gazebo_msgs__srv__GetWorldProperties_Response
type CGetWorldProperties_Response__Sequence = C.gazebo_msgs__srv__GetWorldProperties_Response__Sequence

func GetWorldProperties_Response__Sequence_to_Go(goSlice *[]GetWorldProperties_Response, cSlice CGetWorldProperties_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetWorldProperties_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__GetWorldProperties_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetWorldProperties_Response * uintptr(i)),
		))
		GetWorldProperties_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetWorldProperties_Response__Sequence_to_C(cSlice *CGetWorldProperties_Response__Sequence, goSlice []GetWorldProperties_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__GetWorldProperties_Response)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__GetWorldProperties_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__GetWorldProperties_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetWorldProperties_Response * uintptr(i)),
		))
		GetWorldProperties_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetWorldProperties_Response__Array_to_Go(goSlice []GetWorldProperties_Response, cSlice []CGetWorldProperties_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetWorldProperties_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetWorldProperties_Response__Array_to_C(cSlice []CGetWorldProperties_Response, goSlice []GetWorldProperties_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetWorldProperties_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}