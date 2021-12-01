/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package lifecycle_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -llifecycle_msgs__rosidl_typesupport_c -llifecycle_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/srv/change_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/ChangeState_Response", ChangeState_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewChangeState_Response
// function instead.
type ChangeState_Response struct {
	Success bool `yaml:"success"`// Indicates whether the service was able to initiate the state transition
}

// NewChangeState_Response creates a new ChangeState_Response with default values.
func NewChangeState_Response() *ChangeState_Response {
	self := ChangeState_Response{}
	self.SetDefaults()
	return &self
}

func (t *ChangeState_Response) Clone() *ChangeState_Response {
	c := &ChangeState_Response{}
	c.Success = t.Success
	return c
}

func (t *ChangeState_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ChangeState_Response) SetDefaults() {
	t.Success = false
}

// CloneChangeState_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneChangeState_ResponseSlice(dst, src []ChangeState_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ChangeState_ResponseTypeSupport types.MessageTypeSupport = _ChangeState_ResponseTypeSupport{}

type _ChangeState_ResponseTypeSupport struct{}

func (t _ChangeState_ResponseTypeSupport) New() types.Message {
	return NewChangeState_Response()
}

func (t _ChangeState_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__srv__ChangeState_Response
	return (unsafe.Pointer)(C.lifecycle_msgs__srv__ChangeState_Response__create())
}

func (t _ChangeState_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__srv__ChangeState_Response__destroy((*C.lifecycle_msgs__srv__ChangeState_Response)(pointer_to_free))
}

func (t _ChangeState_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ChangeState_Response)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _ChangeState_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ChangeState_Response)
	mem := (*C.lifecycle_msgs__srv__ChangeState_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _ChangeState_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__srv__ChangeState_Response())
}

type CChangeState_Response = C.lifecycle_msgs__srv__ChangeState_Response
type CChangeState_Response__Sequence = C.lifecycle_msgs__srv__ChangeState_Response__Sequence

func ChangeState_Response__Sequence_to_Go(goSlice *[]ChangeState_Response, cSlice CChangeState_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ChangeState_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.lifecycle_msgs__srv__ChangeState_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Response * uintptr(i)),
		))
		ChangeState_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ChangeState_Response__Sequence_to_C(cSlice *CChangeState_Response__Sequence, goSlice []ChangeState_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.lifecycle_msgs__srv__ChangeState_Response)(C.malloc((C.size_t)(C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.lifecycle_msgs__srv__ChangeState_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_lifecycle_msgs__srv__ChangeState_Response * uintptr(i)),
		))
		ChangeState_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ChangeState_Response__Array_to_Go(goSlice []ChangeState_Response, cSlice []CChangeState_Response) {
	for i := 0; i < len(cSlice); i++ {
		ChangeState_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ChangeState_Response__Array_to_C(cSlice []CChangeState_Response, goSlice []ChangeState_Response) {
	for i := 0; i < len(goSlice); i++ {
		ChangeState_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
