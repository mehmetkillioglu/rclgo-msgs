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

#include <composition_interfaces/srv/unload_node.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("composition_interfaces/UnloadNode_Request", UnloadNode_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewUnloadNode_Request
// function instead.
type UnloadNode_Request struct {
	UniqueId uint64 `yaml:"unique_id"`// Container specific unique id of a loaded node.
}

// NewUnloadNode_Request creates a new UnloadNode_Request with default values.
func NewUnloadNode_Request() *UnloadNode_Request {
	self := UnloadNode_Request{}
	self.SetDefaults()
	return &self
}

func (t *UnloadNode_Request) Clone() *UnloadNode_Request {
	c := &UnloadNode_Request{}
	c.UniqueId = t.UniqueId
	return c
}

func (t *UnloadNode_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UnloadNode_Request) SetDefaults() {
	t.UniqueId = 0
}

// CloneUnloadNode_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUnloadNode_RequestSlice(dst, src []UnloadNode_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UnloadNode_RequestTypeSupport types.MessageTypeSupport = _UnloadNode_RequestTypeSupport{}

type _UnloadNode_RequestTypeSupport struct{}

func (t _UnloadNode_RequestTypeSupport) New() types.Message {
	return NewUnloadNode_Request()
}

func (t _UnloadNode_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.composition_interfaces__srv__UnloadNode_Request
	return (unsafe.Pointer)(C.composition_interfaces__srv__UnloadNode_Request__create())
}

func (t _UnloadNode_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.composition_interfaces__srv__UnloadNode_Request__destroy((*C.composition_interfaces__srv__UnloadNode_Request)(pointer_to_free))
}

func (t _UnloadNode_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UnloadNode_Request)
	mem := (*C.composition_interfaces__srv__UnloadNode_Request)(dst)
	mem.unique_id = C.uint64_t(m.UniqueId)
}

func (t _UnloadNode_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UnloadNode_Request)
	mem := (*C.composition_interfaces__srv__UnloadNode_Request)(ros2_message_buffer)
	m.UniqueId = uint64(mem.unique_id)
}

func (t _UnloadNode_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__composition_interfaces__srv__UnloadNode_Request())
}

type CUnloadNode_Request = C.composition_interfaces__srv__UnloadNode_Request
type CUnloadNode_Request__Sequence = C.composition_interfaces__srv__UnloadNode_Request__Sequence

func UnloadNode_Request__Sequence_to_Go(goSlice *[]UnloadNode_Request, cSlice CUnloadNode_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UnloadNode_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.composition_interfaces__srv__UnloadNode_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__UnloadNode_Request * uintptr(i)),
		))
		UnloadNode_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UnloadNode_Request__Sequence_to_C(cSlice *CUnloadNode_Request__Sequence, goSlice []UnloadNode_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.composition_interfaces__srv__UnloadNode_Request)(C.malloc((C.size_t)(C.sizeof_struct_composition_interfaces__srv__UnloadNode_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.composition_interfaces__srv__UnloadNode_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_composition_interfaces__srv__UnloadNode_Request * uintptr(i)),
		))
		UnloadNode_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UnloadNode_Request__Array_to_Go(goSlice []UnloadNode_Request, cSlice []CUnloadNode_Request) {
	for i := 0; i < len(cSlice); i++ {
		UnloadNode_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UnloadNode_Request__Array_to_C(cSlice []CUnloadNode_Request, goSlice []UnloadNode_Request) {
	for i := 0; i < len(goSlice); i++ {
		UnloadNode_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
