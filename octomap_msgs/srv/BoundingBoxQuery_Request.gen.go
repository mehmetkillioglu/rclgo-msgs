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
	geometry_msgs_msg "github.com/tiiuae/rclgo-msgs/geometry_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -loctomap_msgs__rosidl_typesupport_c -loctomap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <octomap_msgs/srv/bounding_box_query.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("octomap_msgs/BoundingBoxQuery_Request", BoundingBoxQuery_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewBoundingBoxQuery_Request
// function instead.
type BoundingBoxQuery_Request struct {
	Min geometry_msgs_msg.Point `yaml:"min"`// minimum corner point of axis-aligned bounding box in global frame
	Max geometry_msgs_msg.Point `yaml:"max"`// minimum corner point of axis-aligned bounding box in global framemaximum corner point of axis-aligned bounding box in global frame
}

// NewBoundingBoxQuery_Request creates a new BoundingBoxQuery_Request with default values.
func NewBoundingBoxQuery_Request() *BoundingBoxQuery_Request {
	self := BoundingBoxQuery_Request{}
	self.SetDefaults()
	return &self
}

func (t *BoundingBoxQuery_Request) Clone() *BoundingBoxQuery_Request {
	c := &BoundingBoxQuery_Request{}
	c.Min = *t.Min.Clone()
	c.Max = *t.Max.Clone()
	return c
}

func (t *BoundingBoxQuery_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *BoundingBoxQuery_Request) SetDefaults() {
	t.Min.SetDefaults()
	t.Max.SetDefaults()
}

// CloneBoundingBoxQuery_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneBoundingBoxQuery_RequestSlice(dst, src []BoundingBoxQuery_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var BoundingBoxQuery_RequestTypeSupport types.MessageTypeSupport = _BoundingBoxQuery_RequestTypeSupport{}

type _BoundingBoxQuery_RequestTypeSupport struct{}

func (t _BoundingBoxQuery_RequestTypeSupport) New() types.Message {
	return NewBoundingBoxQuery_Request()
}

func (t _BoundingBoxQuery_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.octomap_msgs__srv__BoundingBoxQuery_Request
	return (unsafe.Pointer)(C.octomap_msgs__srv__BoundingBoxQuery_Request__create())
}

func (t _BoundingBoxQuery_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.octomap_msgs__srv__BoundingBoxQuery_Request__destroy((*C.octomap_msgs__srv__BoundingBoxQuery_Request)(pointer_to_free))
}

func (t _BoundingBoxQuery_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*BoundingBoxQuery_Request)
	mem := (*C.octomap_msgs__srv__BoundingBoxQuery_Request)(dst)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.min), &m.Min)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.max), &m.Max)
}

func (t _BoundingBoxQuery_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*BoundingBoxQuery_Request)
	mem := (*C.octomap_msgs__srv__BoundingBoxQuery_Request)(ros2_message_buffer)
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.Min, unsafe.Pointer(&mem.min))
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.Max, unsafe.Pointer(&mem.max))
}

func (t _BoundingBoxQuery_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__octomap_msgs__srv__BoundingBoxQuery_Request())
}

type CBoundingBoxQuery_Request = C.octomap_msgs__srv__BoundingBoxQuery_Request
type CBoundingBoxQuery_Request__Sequence = C.octomap_msgs__srv__BoundingBoxQuery_Request__Sequence

func BoundingBoxQuery_Request__Sequence_to_Go(goSlice *[]BoundingBoxQuery_Request, cSlice CBoundingBoxQuery_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]BoundingBoxQuery_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.octomap_msgs__srv__BoundingBoxQuery_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Request * uintptr(i)),
		))
		BoundingBoxQuery_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func BoundingBoxQuery_Request__Sequence_to_C(cSlice *CBoundingBoxQuery_Request__Sequence, goSlice []BoundingBoxQuery_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.octomap_msgs__srv__BoundingBoxQuery_Request)(C.malloc((C.size_t)(C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.octomap_msgs__srv__BoundingBoxQuery_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__srv__BoundingBoxQuery_Request * uintptr(i)),
		))
		BoundingBoxQuery_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func BoundingBoxQuery_Request__Array_to_Go(goSlice []BoundingBoxQuery_Request, cSlice []CBoundingBoxQuery_Request) {
	for i := 0; i < len(cSlice); i++ {
		BoundingBoxQuery_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func BoundingBoxQuery_Request__Array_to_C(cSlice []CBoundingBoxQuery_Request, goSlice []BoundingBoxQuery_Request) {
	for i := 0; i < len(goSlice); i++ {
		BoundingBoxQuery_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
