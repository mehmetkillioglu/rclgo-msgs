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
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/get_bool.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/GetBool_Response", GetBool_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetBool_Response
// function instead.
type GetBool_Response struct {
	Value bool `yaml:"value"`
}

// NewGetBool_Response creates a new GetBool_Response with default values.
func NewGetBool_Response() *GetBool_Response {
	self := GetBool_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetBool_Response) Clone() *GetBool_Response {
	c := &GetBool_Response{}
	c.Value = t.Value
	return c
}

func (t *GetBool_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetBool_Response) SetDefaults() {
	t.Value = false
}

// CloneGetBool_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetBool_ResponseSlice(dst, src []GetBool_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetBool_ResponseTypeSupport types.MessageTypeSupport = _GetBool_ResponseTypeSupport{}

type _GetBool_ResponseTypeSupport struct{}

func (t _GetBool_ResponseTypeSupport) New() types.Message {
	return NewGetBool_Response()
}

func (t _GetBool_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__GetBool_Response
	return (unsafe.Pointer)(C.fog_msgs__srv__GetBool_Response__create())
}

func (t _GetBool_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__GetBool_Response__destroy((*C.fog_msgs__srv__GetBool_Response)(pointer_to_free))
}

func (t _GetBool_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetBool_Response)
	mem := (*C.fog_msgs__srv__GetBool_Response)(dst)
	mem.value = C.bool(m.Value)
}

func (t _GetBool_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetBool_Response)
	mem := (*C.fog_msgs__srv__GetBool_Response)(ros2_message_buffer)
	m.Value = bool(mem.value)
}

func (t _GetBool_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__GetBool_Response())
}

type CGetBool_Response = C.fog_msgs__srv__GetBool_Response
type CGetBool_Response__Sequence = C.fog_msgs__srv__GetBool_Response__Sequence

func GetBool_Response__Sequence_to_Go(goSlice *[]GetBool_Response, cSlice CGetBool_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetBool_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__GetBool_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__GetBool_Response * uintptr(i)),
		))
		GetBool_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetBool_Response__Sequence_to_C(cSlice *CGetBool_Response__Sequence, goSlice []GetBool_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__GetBool_Response)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__GetBool_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__GetBool_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__GetBool_Response * uintptr(i)),
		))
		GetBool_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetBool_Response__Array_to_Go(goSlice []GetBool_Response, cSlice []CGetBool_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetBool_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetBool_Response__Array_to_C(cSlice []CGetBool_Response, goSlice []GetBool_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetBool_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
