/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <fog_msgs/srv/waypoint_to_local.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("fog_msgs/WaypointToLocal_Request", WaypointToLocal_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewWaypointToLocal_Request
// function instead.
type WaypointToLocal_Request struct {
	LatitudeDeg float64 `yaml:"latitude_deg"`
	LongitudeDeg float64 `yaml:"longitude_deg"`
	RelativeAltitudeM float64 `yaml:"relative_altitude_m"`
}

// NewWaypointToLocal_Request creates a new WaypointToLocal_Request with default values.
func NewWaypointToLocal_Request() *WaypointToLocal_Request {
	self := WaypointToLocal_Request{}
	self.SetDefaults()
	return &self
}

func (t *WaypointToLocal_Request) Clone() *WaypointToLocal_Request {
	c := &WaypointToLocal_Request{}
	c.LatitudeDeg = t.LatitudeDeg
	c.LongitudeDeg = t.LongitudeDeg
	c.RelativeAltitudeM = t.RelativeAltitudeM
	return c
}

func (t *WaypointToLocal_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointToLocal_Request) SetDefaults() {
	t.LatitudeDeg = 0
	t.LongitudeDeg = 0
	t.RelativeAltitudeM = 0
}

// CloneWaypointToLocal_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointToLocal_RequestSlice(dst, src []WaypointToLocal_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointToLocal_RequestTypeSupport types.MessageTypeSupport = _WaypointToLocal_RequestTypeSupport{}

type _WaypointToLocal_RequestTypeSupport struct{}

func (t _WaypointToLocal_RequestTypeSupport) New() types.Message {
	return NewWaypointToLocal_Request()
}

func (t _WaypointToLocal_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.fog_msgs__srv__WaypointToLocal_Request
	return (unsafe.Pointer)(C.fog_msgs__srv__WaypointToLocal_Request__create())
}

func (t _WaypointToLocal_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.fog_msgs__srv__WaypointToLocal_Request__destroy((*C.fog_msgs__srv__WaypointToLocal_Request)(pointer_to_free))
}

func (t _WaypointToLocal_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WaypointToLocal_Request)
	mem := (*C.fog_msgs__srv__WaypointToLocal_Request)(dst)
	mem.latitude_deg = C.double(m.LatitudeDeg)
	mem.longitude_deg = C.double(m.LongitudeDeg)
	mem.relative_altitude_m = C.double(m.RelativeAltitudeM)
}

func (t _WaypointToLocal_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WaypointToLocal_Request)
	mem := (*C.fog_msgs__srv__WaypointToLocal_Request)(ros2_message_buffer)
	m.LatitudeDeg = float64(mem.latitude_deg)
	m.LongitudeDeg = float64(mem.longitude_deg)
	m.RelativeAltitudeM = float64(mem.relative_altitude_m)
}

func (t _WaypointToLocal_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__fog_msgs__srv__WaypointToLocal_Request())
}

type CWaypointToLocal_Request = C.fog_msgs__srv__WaypointToLocal_Request
type CWaypointToLocal_Request__Sequence = C.fog_msgs__srv__WaypointToLocal_Request__Sequence

func WaypointToLocal_Request__Sequence_to_Go(goSlice *[]WaypointToLocal_Request, cSlice CWaypointToLocal_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointToLocal_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.fog_msgs__srv__WaypointToLocal_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__WaypointToLocal_Request * uintptr(i)),
		))
		WaypointToLocal_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func WaypointToLocal_Request__Sequence_to_C(cSlice *CWaypointToLocal_Request__Sequence, goSlice []WaypointToLocal_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.fog_msgs__srv__WaypointToLocal_Request)(C.malloc((C.size_t)(C.sizeof_struct_fog_msgs__srv__WaypointToLocal_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.fog_msgs__srv__WaypointToLocal_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_fog_msgs__srv__WaypointToLocal_Request * uintptr(i)),
		))
		WaypointToLocal_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func WaypointToLocal_Request__Array_to_Go(goSlice []WaypointToLocal_Request, cSlice []CWaypointToLocal_Request) {
	for i := 0; i < len(cSlice); i++ {
		WaypointToLocal_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointToLocal_Request__Array_to_C(cSlice []CWaypointToLocal_Request, goSlice []WaypointToLocal_Request) {
	for i := 0; i < len(goSlice); i++ {
		WaypointToLocal_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
