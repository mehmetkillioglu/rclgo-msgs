/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	map_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/map_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/set_map_projections.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/SetMapProjections_Response", SetMapProjections_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetMapProjections_Response
// function instead.
type SetMapProjections_Response struct {
	ProjectedMapsInfo []map_msgs_msg.ProjectedMapInfo `yaml:"projected_maps_info"`
}

// NewSetMapProjections_Response creates a new SetMapProjections_Response with default values.
func NewSetMapProjections_Response() *SetMapProjections_Response {
	self := SetMapProjections_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetMapProjections_Response) Clone() *SetMapProjections_Response {
	c := &SetMapProjections_Response{}
	if t.ProjectedMapsInfo != nil {
		c.ProjectedMapsInfo = make([]map_msgs_msg.ProjectedMapInfo, len(t.ProjectedMapsInfo))
		map_msgs_msg.CloneProjectedMapInfoSlice(c.ProjectedMapsInfo, t.ProjectedMapsInfo)
	}
	return c
}

func (t *SetMapProjections_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetMapProjections_Response) SetDefaults() {
	t.ProjectedMapsInfo = nil
}

// CloneSetMapProjections_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetMapProjections_ResponseSlice(dst, src []SetMapProjections_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetMapProjections_ResponseTypeSupport types.MessageTypeSupport = _SetMapProjections_ResponseTypeSupport{}

type _SetMapProjections_ResponseTypeSupport struct{}

func (t _SetMapProjections_ResponseTypeSupport) New() types.Message {
	return NewSetMapProjections_Response()
}

func (t _SetMapProjections_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__SetMapProjections_Response
	return (unsafe.Pointer)(C.map_msgs__srv__SetMapProjections_Response__create())
}

func (t _SetMapProjections_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__SetMapProjections_Response__destroy((*C.map_msgs__srv__SetMapProjections_Response)(pointer_to_free))
}

func (t _SetMapProjections_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetMapProjections_Response)
	mem := (*C.map_msgs__srv__SetMapProjections_Response)(dst)
	map_msgs_msg.ProjectedMapInfo__Sequence_to_C((*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)), m.ProjectedMapsInfo)
}

func (t _SetMapProjections_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetMapProjections_Response)
	mem := (*C.map_msgs__srv__SetMapProjections_Response)(ros2_message_buffer)
	map_msgs_msg.ProjectedMapInfo__Sequence_to_Go(&m.ProjectedMapsInfo, *(*map_msgs_msg.CProjectedMapInfo__Sequence)(unsafe.Pointer(&mem.projected_maps_info)))
}

func (t _SetMapProjections_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__SetMapProjections_Response())
}

type CSetMapProjections_Response = C.map_msgs__srv__SetMapProjections_Response
type CSetMapProjections_Response__Sequence = C.map_msgs__srv__SetMapProjections_Response__Sequence

func SetMapProjections_Response__Sequence_to_Go(goSlice *[]SetMapProjections_Response, cSlice CSetMapProjections_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMapProjections_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__SetMapProjections_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(i)),
		))
		SetMapProjections_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetMapProjections_Response__Sequence_to_C(cSlice *CSetMapProjections_Response__Sequence, goSlice []SetMapProjections_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__SetMapProjections_Response)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__SetMapProjections_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__SetMapProjections_Response * uintptr(i)),
		))
		SetMapProjections_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetMapProjections_Response__Array_to_Go(goSlice []SetMapProjections_Response, cSlice []CSetMapProjections_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetMapProjections_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetMapProjections_Response__Array_to_C(cSlice []CSetMapProjections_Response, goSlice []SetMapProjections_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetMapProjections_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
