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

#include <gazebo_msgs/srv/delete_model.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/DeleteModel_Response", DeleteModel_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewDeleteModel_Response
// function instead.
type DeleteModel_Response struct {
	Success bool `yaml:"success"`// return true if deletion is successful. Deprecated, kept for ROS 1 bridge.Use DeleteEntity
	StatusMessage string `yaml:"status_message"`// comments if available. Deprecated, kept for ROS 1 bridge.Use DeleteEntity
}

// NewDeleteModel_Response creates a new DeleteModel_Response with default values.
func NewDeleteModel_Response() *DeleteModel_Response {
	self := DeleteModel_Response{}
	self.SetDefaults()
	return &self
}

func (t *DeleteModel_Response) Clone() *DeleteModel_Response {
	c := &DeleteModel_Response{}
	c.Success = t.Success
	c.StatusMessage = t.StatusMessage
	return c
}

func (t *DeleteModel_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DeleteModel_Response) SetDefaults() {
	t.Success = false
	t.StatusMessage = ""
}

// CloneDeleteModel_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDeleteModel_ResponseSlice(dst, src []DeleteModel_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DeleteModel_ResponseTypeSupport types.MessageTypeSupport = _DeleteModel_ResponseTypeSupport{}

type _DeleteModel_ResponseTypeSupport struct{}

func (t _DeleteModel_ResponseTypeSupport) New() types.Message {
	return NewDeleteModel_Response()
}

func (t _DeleteModel_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__DeleteModel_Response
	return (unsafe.Pointer)(C.gazebo_msgs__srv__DeleteModel_Response__create())
}

func (t _DeleteModel_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__DeleteModel_Response__destroy((*C.gazebo_msgs__srv__DeleteModel_Response)(pointer_to_free))
}

func (t _DeleteModel_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DeleteModel_Response)
	mem := (*C.gazebo_msgs__srv__DeleteModel_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status_message), m.StatusMessage)
}

func (t _DeleteModel_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DeleteModel_Response)
	mem := (*C.gazebo_msgs__srv__DeleteModel_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.StatusMessage, unsafe.Pointer(&mem.status_message))
}

func (t _DeleteModel_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__DeleteModel_Response())
}

type CDeleteModel_Response = C.gazebo_msgs__srv__DeleteModel_Response
type CDeleteModel_Response__Sequence = C.gazebo_msgs__srv__DeleteModel_Response__Sequence

func DeleteModel_Response__Sequence_to_Go(goSlice *[]DeleteModel_Response, cSlice CDeleteModel_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DeleteModel_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__DeleteModel_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__DeleteModel_Response * uintptr(i)),
		))
		DeleteModel_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func DeleteModel_Response__Sequence_to_C(cSlice *CDeleteModel_Response__Sequence, goSlice []DeleteModel_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__DeleteModel_Response)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__DeleteModel_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__DeleteModel_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__DeleteModel_Response * uintptr(i)),
		))
		DeleteModel_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func DeleteModel_Response__Array_to_Go(goSlice []DeleteModel_Response, cSlice []CDeleteModel_Response) {
	for i := 0; i < len(cSlice); i++ {
		DeleteModel_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DeleteModel_Response__Array_to_C(cSlice []CDeleteModel_Response, goSlice []DeleteModel_Response) {
	for i := 0; i < len(goSlice); i++ {
		DeleteModel_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
