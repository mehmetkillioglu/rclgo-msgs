/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package octomap_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -loctomap_msgs__rosidl_typesupport_c -loctomap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <octomap_msgs/srv/bounding_box_query.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("octomap_msgs/BoundingBoxQuery_Response", BoundingBoxQuery_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewBoundingBoxQuery_Response
// function instead.
type BoundingBoxQuery_Response struct {
}

// NewBoundingBoxQuery_Response creates a new BoundingBoxQuery_Response with default values.
func NewBoundingBoxQuery_Response() *BoundingBoxQuery_Response {
	self := BoundingBoxQuery_Response{}
	self.SetDefaults()
	return &self
}

func (t *BoundingBoxQuery_Response) Clone() *BoundingBoxQuery_Response {
	c := &BoundingBoxQuery_Response{}
	return c
}

func (t *BoundingBoxQuery_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BoundingBoxQuery_Response) SetDefaults() {
}

// CloneBoundingBoxQuery_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBoundingBoxQuery_ResponseSlice(dst, src []BoundingBoxQuery_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BoundingBoxQuery_ResponseTypeSupport types.MessageTypeSupport = _BoundingBoxQuery_ResponseTypeSupport{}

type _BoundingBoxQuery_ResponseTypeSupport struct{}

func (t _BoundingBoxQuery_ResponseTypeSupport) New() types.Message {
	return NewBoundingBoxQuery_Response()
}

func (t _BoundingBoxQuery_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.octomap_msgs__srv__BoundingBoxQuery_Response
	return (unsafe.Pointer)(C.octomap_msgs__srv__BoundingBoxQuery_Response__create())
}

func (t _BoundingBoxQuery_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.octomap_msgs__srv__BoundingBoxQuery_Response__destroy((*C.octomap_msgs__srv__BoundingBoxQuery_Response)(pointer_to_free))
}

func (t _BoundingBoxQuery_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _BoundingBoxQuery_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _BoundingBoxQuery_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__octomap_msgs__srv__BoundingBoxQuery_Response())
}

type CBoundingBoxQuery_Response = C.octomap_msgs__srv__BoundingBoxQuery_Response
type CBoundingBoxQuery_Response__Sequence = C.octomap_msgs__srv__BoundingBoxQuery_Response__Sequence

func BoundingBoxQuery_Response__Sequence_to_Go(goSlice *[]BoundingBoxQuery_Response, cSlice CBoundingBoxQuery_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BoundingBoxQuery_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.octomap_msgs__srv__BoundingBoxQuery_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Response * uintptr(i)),
		))
		BoundingBoxQuery_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func BoundingBoxQuery_Response__Sequence_to_C(cSlice *CBoundingBoxQuery_Response__Sequence, goSlice []BoundingBoxQuery_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.octomap_msgs__srv__BoundingBoxQuery_Response)(C.malloc((C.size_t)(C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.octomap_msgs__srv__BoundingBoxQuery_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Response * uintptr(i)),
		))
		BoundingBoxQuery_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func BoundingBoxQuery_Response__Array_to_Go(goSlice []BoundingBoxQuery_Response, cSlice []CBoundingBoxQuery_Response) {
	for i := 0; i < len(cSlice); i++ {
		BoundingBoxQuery_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BoundingBoxQuery_Response__Array_to_C(cSlice []CBoundingBoxQuery_Response, goSlice []BoundingBoxQuery_Response) {
	for i := 0; i < len(goSlice); i++ {
		BoundingBoxQuery_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
