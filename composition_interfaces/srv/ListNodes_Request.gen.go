/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package composition_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lcomposition_interfaces__rosidl_typesupport_c -lcomposition_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <composition_interfaces/srv/list_nodes.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("composition_interfaces/ListNodes_Request", ListNodes_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewListNodes_Request
// function instead.
type ListNodes_Request struct {
}

// NewListNodes_Request creates a new ListNodes_Request with default values.
func NewListNodes_Request() *ListNodes_Request {
	self := ListNodes_Request{}
	self.SetDefaults()
	return &self
}

func (t *ListNodes_Request) Clone() *ListNodes_Request {
	c := &ListNodes_Request{}
	return c
}

func (t *ListNodes_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ListNodes_Request) SetDefaults() {
}

// CloneListNodes_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneListNodes_RequestSlice(dst, src []ListNodes_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ListNodes_RequestTypeSupport types.MessageTypeSupport = _ListNodes_RequestTypeSupport{}

type _ListNodes_RequestTypeSupport struct{}

func (t _ListNodes_RequestTypeSupport) New() types.Message {
	return NewListNodes_Request()
}

func (t _ListNodes_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.composition_interfaces__srv__ListNodes_Request
	return (unsafe.Pointer)(C.composition_interfaces__srv__ListNodes_Request__create())
}

func (t _ListNodes_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.composition_interfaces__srv__ListNodes_Request__destroy((*C.composition_interfaces__srv__ListNodes_Request)(pointer_to_free))
}

func (t _ListNodes_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _ListNodes_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _ListNodes_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__composition_interfaces__srv__ListNodes_Request())
}

type CListNodes_Request = C.composition_interfaces__srv__ListNodes_Request
type CListNodes_Request__Sequence = C.composition_interfaces__srv__ListNodes_Request__Sequence

func ListNodes_Request__Sequence_to_Go(goSlice *[]ListNodes_Request, cSlice CListNodes_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ListNodes_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.composition_interfaces__srv__ListNodes_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__ListNodes_Request * uintptr(i)),
		))
		ListNodes_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ListNodes_Request__Sequence_to_C(cSlice *CListNodes_Request__Sequence, goSlice []ListNodes_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.composition_interfaces__srv__ListNodes_Request)(C.malloc((C.size_t)(C.sizeof_struct_composition_interfaces__srv__ListNodes_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.composition_interfaces__srv__ListNodes_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__ListNodes_Request * uintptr(i)),
		))
		ListNodes_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ListNodes_Request__Array_to_Go(goSlice []ListNodes_Request, cSlice []CListNodes_Request) {
	for i := 0; i < len(cSlice); i++ {
		ListNodes_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ListNodes_Request__Array_to_C(cSlice []CListNodes_Request, goSlice []ListNodes_Request) {
	for i := 0; i < len(goSlice); i++ {
		ListNodes_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
