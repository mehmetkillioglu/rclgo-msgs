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

#include <gazebo_msgs/srv/get_model_properties.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/GetModelProperties_Request", GetModelProperties_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetModelProperties_Request
// function instead.
type GetModelProperties_Request struct {
	ModelName string `yaml:"model_name"`// name of Gazebo Model
}

// NewGetModelProperties_Request creates a new GetModelProperties_Request with default values.
func NewGetModelProperties_Request() *GetModelProperties_Request {
	self := GetModelProperties_Request{}
	self.SetDefaults()
	return &self
}

func (t *GetModelProperties_Request) Clone() *GetModelProperties_Request {
	c := &GetModelProperties_Request{}
	c.ModelName = t.ModelName
	return c
}

func (t *GetModelProperties_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetModelProperties_Request) SetDefaults() {
	t.ModelName = ""
}

// CloneGetModelProperties_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetModelProperties_RequestSlice(dst, src []GetModelProperties_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetModelProperties_RequestTypeSupport types.MessageTypeSupport = _GetModelProperties_RequestTypeSupport{}

type _GetModelProperties_RequestTypeSupport struct{}

func (t _GetModelProperties_RequestTypeSupport) New() types.Message {
	return NewGetModelProperties_Request()
}

func (t _GetModelProperties_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__GetModelProperties_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__GetModelProperties_Request__create())
}

func (t _GetModelProperties_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__GetModelProperties_Request__destroy((*C.gazebo_msgs__srv__GetModelProperties_Request)(pointer_to_free))
}

func (t _GetModelProperties_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetModelProperties_Request)
	mem := (*C.gazebo_msgs__srv__GetModelProperties_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.model_name), m.ModelName)
}

func (t _GetModelProperties_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetModelProperties_Request)
	mem := (*C.gazebo_msgs__srv__GetModelProperties_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.ModelName, unsafe.Pointer(&mem.model_name))
}

func (t _GetModelProperties_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__GetModelProperties_Request())
}

type CGetModelProperties_Request = C.gazebo_msgs__srv__GetModelProperties_Request
type CGetModelProperties_Request__Sequence = C.gazebo_msgs__srv__GetModelProperties_Request__Sequence

func GetModelProperties_Request__Sequence_to_Go(goSlice *[]GetModelProperties_Request, cSlice CGetModelProperties_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetModelProperties_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__GetModelProperties_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetModelProperties_Request * uintptr(i)),
		))
		GetModelProperties_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetModelProperties_Request__Sequence_to_C(cSlice *CGetModelProperties_Request__Sequence, goSlice []GetModelProperties_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__GetModelProperties_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__GetModelProperties_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__GetModelProperties_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__GetModelProperties_Request * uintptr(i)),
		))
		GetModelProperties_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetModelProperties_Request__Array_to_Go(goSlice []GetModelProperties_Request, cSlice []CGetModelProperties_Request) {
	for i := 0; i < len(cSlice); i++ {
		GetModelProperties_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetModelProperties_Request__Array_to_C(cSlice []CGetModelProperties_Request, goSlice []GetModelProperties_Request) {
	for i := 0; i < len(goSlice); i++ {
		GetModelProperties_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
