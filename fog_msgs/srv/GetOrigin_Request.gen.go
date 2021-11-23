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

#include <fog_msgs/srv/get_origin.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/GetOrigin_Request", GetOrigin_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetOrigin_Request
// function instead.
type GetOrigin_Request struct {
}

// NewGetOrigin_Request creates a new GetOrigin_Request with default values.
func NewGetOrigin_Request() *GetOrigin_Request {
	self := GetOrigin_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetOrigin_Request) Clone() *GetOrigin_Request {
	c := &GetOrigin_Request{}
	return c
}

func (t *GetOrigin_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetOrigin_Request) SetDefaults() {
}

// CloneGetOrigin_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetOrigin_RequestSlice(dst, src []GetOrigin_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetOrigin_RequestTypeSupport types.MessageTypeSupport = _GetOrigin_RequestTypeSupport{}

type _GetOrigin_RequestTypeSupport struct{}

func (t _GetOrigin_RequestTypeSupport) New() types.Message {
	return NewGetOrigin_Request()
}

func (t _GetOrigin_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__GetOrigin_Request
	return (unsafe.Pointer)(C.fog_msgs__srv__GetOrigin_Request__create())
}

func (t _GetOrigin_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__GetOrigin_Request__destroy((*C.fog_msgs__srv__GetOrigin_Request)(pointer_to_free))
}

func (t _GetOrigin_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _GetOrigin_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _GetOrigin_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__GetOrigin_Request())
}

type CGetOrigin_Request = C.fog_msgs__srv__GetOrigin_Request
type CGetOrigin_Request__Sequence = C.fog_msgs__srv__GetOrigin_Request__Sequence

func GetOrigin_Request__Sequence_to_Go(goSlice *[]GetOrigin_Request, cSlice CGetOrigin_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetOrigin_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__GetOrigin_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__GetOrigin_Request * uintptr(i)),
		))
		GetOrigin_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetOrigin_Request__Sequence_to_C(cSlice *CGetOrigin_Request__Sequence, goSlice []GetOrigin_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__GetOrigin_Request)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__GetOrigin_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__GetOrigin_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__GetOrigin_Request * uintptr(i)),
		))
		GetOrigin_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetOrigin_Request__Array_to_Go(goSlice []GetOrigin_Request, cSlice []CGetOrigin_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetOrigin_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetOrigin_Request__Array_to_C(cSlice []CGetOrigin_Request, goSlice []GetOrigin_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetOrigin_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
