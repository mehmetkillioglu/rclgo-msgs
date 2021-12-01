/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rosbag2_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrosbag2_interfaces__rosidl_typesupport_c -lrosbag2_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rosbag2_interfaces/srv/get_rate.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rosbag2_interfaces/GetRate_Response", GetRate_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetRate_Response
// function instead.
type GetRate_Response struct {
	Rate float64 `yaml:"rate"`
}

// NewGetRate_Response creates a new GetRate_Response with default values.
func NewGetRate_Response() *GetRate_Response {
	self := GetRate_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetRate_Response) Clone() *GetRate_Response {
	c := &GetRate_Response{}
	c.Rate = t.Rate
	return c
}

func (t *GetRate_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetRate_Response) SetDefaults() {
	t.Rate = 0
}

// CloneGetRate_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetRate_ResponseSlice(dst, src []GetRate_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetRate_ResponseTypeSupport types.MessageTypeSupport = _GetRate_ResponseTypeSupport{}

type _GetRate_ResponseTypeSupport struct{}

func (t _GetRate_ResponseTypeSupport) New() types.Message {
	return NewGetRate_Response()
}

func (t _GetRate_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosbag2_interfaces__srv__GetRate_Response
	return (unsafe.Pointer)(C.rosbag2_interfaces__srv__GetRate_Response__create())
}

func (t _GetRate_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosbag2_interfaces__srv__GetRate_Response__destroy((*C.rosbag2_interfaces__srv__GetRate_Response)(pointer_to_free))
}

func (t _GetRate_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetRate_Response)
	mem := (*C.rosbag2_interfaces__srv__GetRate_Response)(dst)
	mem.rate = C.double(m.Rate)
}

func (t _GetRate_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetRate_Response)
	mem := (*C.rosbag2_interfaces__srv__GetRate_Response)(ros2_message_buffer)
	m.Rate = float64(mem.rate)
}

func (t _GetRate_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosbag2_interfaces__srv__GetRate_Response())
}

type CGetRate_Response = C.rosbag2_interfaces__srv__GetRate_Response
type CGetRate_Response__Sequence = C.rosbag2_interfaces__srv__GetRate_Response__Sequence

func GetRate_Response__Sequence_to_Go(goSlice *[]GetRate_Response, cSlice CGetRate_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetRate_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rosbag2_interfaces__srv__GetRate_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosbag2_interfaces__srv__GetRate_Response * uintptr(i)),
		))
		GetRate_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetRate_Response__Sequence_to_C(cSlice *CGetRate_Response__Sequence, goSlice []GetRate_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rosbag2_interfaces__srv__GetRate_Response)(C.malloc((C.size_t)(C.sizeof_struct_rosbag2_interfaces__srv__GetRate_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rosbag2_interfaces__srv__GetRate_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosbag2_interfaces__srv__GetRate_Response * uintptr(i)),
		))
		GetRate_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetRate_Response__Array_to_Go(goSlice []GetRate_Response, cSlice []CGetRate_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetRate_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetRate_Response__Array_to_C(cSlice []CGetRate_Response, goSlice []GetRate_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetRate_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}