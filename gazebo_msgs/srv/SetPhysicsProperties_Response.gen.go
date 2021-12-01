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

#include <gazebo_msgs/srv/set_physics_properties.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/SetPhysicsProperties_Response", SetPhysicsProperties_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetPhysicsProperties_Response
// function instead.
type SetPhysicsProperties_Response struct {
	Success bool `yaml:"success"`// return true if set wrench successful. sets pose and twist of a link.  All children link poses/twists of the URDF tree will be updated accordingly
	StatusMessage string `yaml:"status_message"`// comments if available. sets pose and twist of a link.  All children link poses/twists of the URDF tree will be updated accordingly
}

// NewSetPhysicsProperties_Response creates a new SetPhysicsProperties_Response with default values.
func NewSetPhysicsProperties_Response() *SetPhysicsProperties_Response {
	self := SetPhysicsProperties_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetPhysicsProperties_Response) Clone() *SetPhysicsProperties_Response {
	c := &SetPhysicsProperties_Response{}
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *SetPhysicsProperties_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetPhysicsProperties_Response) SetDefaults() {
	t.Success = false
	t.StatusMessage = ""
}

// CloneSetPhysicsProperties_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetPhysicsProperties_ResponseSlice(dst, src []SetPhysicsProperties_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetPhysicsProperties_ResponseTypeSupport types.MessageTypeSupport = _SetPhysicsProperties_ResponseTypeSupport{}

type _SetPhysicsProperties_ResponseTypeSupport struct{}

func (t _SetPhysicsProperties_ResponseTypeSupport) New() types.Message {
	return NewSetPhysicsProperties_Response()
}

func (t _SetPhysicsProperties_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__SetPhysicsProperties_Response
	return (unsafe.Pointer)(C.gazebo_msgs__srv__SetPhysicsProperties_Response__create())
}

func (t _SetPhysicsProperties_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__SetPhysicsProperties_Response__destroy((*C.gazebo_msgs__srv__SetPhysicsProperties_Response)(pointer_to_free))
}

func (t _SetPhysicsProperties_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetPhysicsProperties_Response)
	mem := (*C.gazebo_msgs__srv__SetPhysicsProperties_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _SetPhysicsProperties_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetPhysicsProperties_Response)
	mem := (*C.gazebo_msgs__srv__SetPhysicsProperties_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _SetPhysicsProperties_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__SetPhysicsProperties_Response())
}

type CSetPhysicsProperties_Response = C.gazebo_msgs__srv__SetPhysicsProperties_Response
type CSetPhysicsProperties_Response__Sequence = C.gazebo_msgs__srv__SetPhysicsProperties_Response__Sequence

func SetPhysicsProperties_Response__Sequence_to_Go(goSlice *[]SetPhysicsProperties_Response, cSlice CSetPhysicsProperties_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetPhysicsProperties_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__SetPhysicsProperties_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__SetPhysicsProperties_Response * uintptr(i)),
		))
		SetPhysicsProperties_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetPhysicsProperties_Response__Sequence_to_C(cSlice *CSetPhysicsProperties_Response__Sequence, goSlice []SetPhysicsProperties_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__SetPhysicsProperties_Response)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__SetPhysicsProperties_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__SetPhysicsProperties_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__SetPhysicsProperties_Response * uintptr(i)),
		))
		SetPhysicsProperties_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetPhysicsProperties_Response__Array_to_Go(goSlice []SetPhysicsProperties_Response, cSlice []CSetPhysicsProperties_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetPhysicsProperties_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetPhysicsProperties_Response__Array_to_C(cSlice []CSetPhysicsProperties_Response, goSlice []SetPhysicsProperties_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetPhysicsProperties_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
