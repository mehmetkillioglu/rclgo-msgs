/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geographic_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geographic_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/geographic_msgs/msg"
	unique_identifier_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/unique_identifier_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeographic_msgs__rosidl_typesupport_c -lgeographic_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeographic_msgs__rosidl_typesupport_c -lgeographic_msgs__rosidl_generator_c
#cgo LDFLAGS: -lunique_identifier_msgs__rosidl_typesupport_c -lunique_identifier_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/srv/get_geo_path.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/GetGeoPath_Response", GetGeoPath_ResponseTypeSupport)
}

// Do not create instances of this type directly. Always use NewGetGeoPath_Response
// function instead.
type GetGeoPath_Response struct {
	Success bool `yaml:"success"`// true if the call succeeded
	Status string `yaml:"status"`// more details
	Plan geographic_msgs_msg.GeoPath `yaml:"plan"`// path to follow
	Network unique_identifier_msgs_msg.UUID `yaml:"network"`// the uuid of the RouteNetwork
	StartSeg unique_identifier_msgs_msg.UUID `yaml:"start_seg"`// the uuid of the starting RouteSegment
	GoalSeg unique_identifier_msgs_msg.UUID `yaml:"goal_seg"`// the uuid of the ending RouteSegment
	Distance float64 `yaml:"distance"`// the length of the plan
}

// NewGetGeoPath_Response creates a new GetGeoPath_Response with default values.
func NewGetGeoPath_Response() *GetGeoPath_Response {
	self := GetGeoPath_Response{}
	self.SetDefaults()
	return &self
}

func (t *GetGeoPath_Response) Clone() *GetGeoPath_Response {
	c := &GetGeoPath_Response{}
	c.Success = t.Success
	c.Status = t.Status
	c.Plan = *t.Plan.Clone()
	c.Network = *t.Network.Clone()
	c.StartSeg = *t.StartSeg.Clone()
	c.GoalSeg = *t.GoalSeg.Clone()
	c.Distance = t.Distance
	return c
}

func (t *GetGeoPath_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GetGeoPath_Response) SetDefaults() {
	t.Success = false
	t.Status = ""
	t.Plan.SetDefaults()
	t.Network.SetDefaults()
	t.StartSeg.SetDefaults()
	t.GoalSeg.SetDefaults()
	t.Distance = 0
}

// CloneGetGeoPath_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGetGeoPath_ResponseSlice(dst, src []GetGeoPath_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GetGeoPath_ResponseTypeSupport types.MessageTypeSupport = _GetGeoPath_ResponseTypeSupport{}

type _GetGeoPath_ResponseTypeSupport struct{}

func (t _GetGeoPath_ResponseTypeSupport) New() types.Message {
	return NewGetGeoPath_Response()
}

func (t _GetGeoPath_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__srv__GetGeoPath_Response
	return (unsafe.Pointer)(C.geographic_msgs__srv__GetGeoPath_Response__create())
}

func (t _GetGeoPath_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__srv__GetGeoPath_Response__destroy((*C.geographic_msgs__srv__GetGeoPath_Response)(pointer_to_free))
}

func (t _GetGeoPath_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GetGeoPath_Response)
	mem := (*C.geographic_msgs__srv__GetGeoPath_Response)(dst)
	mem.success = C.bool(m.Success)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.status), m.Status)
	geographic_msgs_msg.GeoPathTypeSupport.AsCStruct(unsafe.Pointer(&mem.plan), &m.Plan)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.network), &m.Network)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.start_seg), &m.StartSeg)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_seg), &m.GoalSeg)
	mem.distance = C.double(m.Distance)
}

func (t _GetGeoPath_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GetGeoPath_Response)
	mem := (*C.geographic_msgs__srv__GetGeoPath_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	primitives.StringAsGoStruct(&m.Status, unsafe.Pointer(&mem.status))
	geographic_msgs_msg.GeoPathTypeSupport.AsGoStruct(&m.Plan, unsafe.Pointer(&mem.plan))
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.Network, unsafe.Pointer(&mem.network))
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.StartSeg, unsafe.Pointer(&mem.start_seg))
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalSeg, unsafe.Pointer(&mem.goal_seg))
	m.Distance = float64(mem.distance)
}

func (t _GetGeoPath_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__srv__GetGeoPath_Response())
}

type CGetGeoPath_Response = C.geographic_msgs__srv__GetGeoPath_Response
type CGetGeoPath_Response__Sequence = C.geographic_msgs__srv__GetGeoPath_Response__Sequence

func GetGeoPath_Response__Sequence_to_Go(goSlice *[]GetGeoPath_Response, cSlice CGetGeoPath_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetGeoPath_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geographic_msgs__srv__GetGeoPath_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__srv__GetGeoPath_Response * uintptr(i)),
		))
		GetGeoPath_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GetGeoPath_Response__Sequence_to_C(cSlice *CGetGeoPath_Response__Sequence, goSlice []GetGeoPath_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geographic_msgs__srv__GetGeoPath_Response)(C.malloc((C.size_t)(C.sizeof_struct_geographic_msgs__srv__GetGeoPath_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geographic_msgs__srv__GetGeoPath_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__srv__GetGeoPath_Response * uintptr(i)),
		))
		GetGeoPath_ResponseTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GetGeoPath_Response__Array_to_Go(goSlice []GetGeoPath_Response, cSlice []CGetGeoPath_Response) {
	for i := 0; i < len(cSlice); i++ {
		GetGeoPath_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GetGeoPath_Response__Array_to_C(cSlice []CGetGeoPath_Response, goSlice []GetGeoPath_Response) {
	for i := 0; i < len(goSlice); i++ {
		GetGeoPath_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
