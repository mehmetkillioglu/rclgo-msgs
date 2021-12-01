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

#include <rosbag2_interfaces/srv/seek.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rosbag2_interfaces/Seek_Response", Seek_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSeek_Response
// function instead.
type Seek_Response struct {
	Success bool `yaml:"success"`// return true if valid time in bag duration, and successful seek
}

// NewSeek_Response creates a new Seek_Response with default values.
func NewSeek_Response() *Seek_Response {
	self := Seek_Response{}
	self.SetDefaults()
	return &self
}

func (t *Seek_Response) Clone() *Seek_Response {
	c := &Seek_Response{}
	c.Success = t.Success
	return c
}

func (t *Seek_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Seek_Response) SetDefaults() {
	t.Success = false
}

// CloneSeek_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSeek_ResponseSlice(dst, src []Seek_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Seek_ResponseTypeSupport types.MessageTypeSupport = _Seek_ResponseTypeSupport{}

type _Seek_ResponseTypeSupport struct{}

func (t _Seek_ResponseTypeSupport) New() types.Message {
	return NewSeek_Response()
}

func (t _Seek_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosbag2_interfaces__srv__Seek_Response
	return (unsafe.Pointer)(C.rosbag2_interfaces__srv__Seek_Response__create())
}

func (t _Seek_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosbag2_interfaces__srv__Seek_Response__destroy((*C.rosbag2_interfaces__srv__Seek_Response)(pointer_to_free))
}

func (t _Seek_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Seek_Response)
	mem := (*C.rosbag2_interfaces__srv__Seek_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _Seek_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Seek_Response)
	mem := (*C.rosbag2_interfaces__srv__Seek_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _Seek_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosbag2_interfaces__srv__Seek_Response())
}

type CSeek_Response = C.rosbag2_interfaces__srv__Seek_Response
type CSeek_Response__Sequence = C.rosbag2_interfaces__srv__Seek_Response__Sequence

func Seek_Response__Sequence_to_Go(goSlice *[]Seek_Response, cSlice CSeek_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Seek_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rosbag2_interfaces__srv__Seek_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosbag2_interfaces__srv__Seek_Response * uintptr(i)),
		))
		Seek_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Seek_Response__Sequence_to_C(cSlice *CSeek_Response__Sequence, goSlice []Seek_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rosbag2_interfaces__srv__Seek_Response)(C.malloc((C.size_t)(C.sizeof_struct_rosbag2_interfaces__srv__Seek_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rosbag2_interfaces__srv__Seek_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosbag2_interfaces__srv__Seek_Response * uintptr(i)),
		))
		Seek_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Seek_Response__Array_to_Go(goSlice []Seek_Response, cSlice []CSeek_Response) {
	for i := 0; i < len(cSlice); i++ {
		Seek_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Seek_Response__Array_to_C(cSlice []CSeek_Response, goSlice []Seek_Response) {
	for i := 0; i < len(goSlice); i++ {
		Seek_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
