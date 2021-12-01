/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geographic_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	unique_identifier_msgs_msg "github.com/tiiuae/rclgo-msgs/unique_identifier_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeographic_msgs__rosidl_typesupport_c -lgeographic_msgs__rosidl_generator_c
#cgo LDFLAGS: -lunique_identifier_msgs__rosidl_typesupport_c -lunique_identifier_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/msg/way_point.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/WayPoint", WayPointTypeSupport)
}

// Do not create instances of this type directly. Always use NewWayPoint
// function instead.
type WayPoint struct {
	Id unique_identifier_msgs_msg.UUID `yaml:"id"`// Unique way-point identifier
	Position GeoPoint `yaml:"position"`// Position relative to WGS 84 ellipsoid
	Props []KeyValue `yaml:"props"`// Key/value properties for this point
}

// NewWayPoint creates a new WayPoint with default values.
func NewWayPoint() *WayPoint {
	self := WayPoint{}
	self.SetDefaults()
	return &self
}

func (t *WayPoint) Clone() *WayPoint {
	c := &WayPoint{}
	c.Id = *t.Id.Clone()
	c.Position = *t.Position.Clone()
	if t.Props != nil {
		c.Props = make([]KeyValue, len(t.Props))
		CloneKeyValueSlice(c.Props, t.Props)
	}
	return c
}

func (t *WayPoint) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WayPoint) SetDefaults() {
	t.Id.SetDefaults()
	t.Position.SetDefaults()
	t.Props = nil
}

// CloneWayPointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWayPointSlice(dst, src []WayPoint) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WayPointTypeSupport types.MessageTypeSupport = _WayPointTypeSupport{}

type _WayPointTypeSupport struct{}

func (t _WayPointTypeSupport) New() types.Message {
	return NewWayPoint()
}

func (t _WayPointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__WayPoint
	return (unsafe.Pointer)(C.geographic_msgs__msg__WayPoint__create())
}

func (t _WayPointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__WayPoint__destroy((*C.geographic_msgs__msg__WayPoint)(pointer_to_free))
}

func (t _WayPointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WayPoint)
	mem := (*C.geographic_msgs__msg__WayPoint)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.id), &m.Id)
	GeoPointTypeSupport.AsCStruct(unsafe.Pointer(&mem.position), &m.Position)
	KeyValue__Sequence_to_C(&mem.props, m.Props)
}

func (t _WayPointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WayPoint)
	mem := (*C.geographic_msgs__msg__WayPoint)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.Id, unsafe.Pointer(&mem.id))
	GeoPointTypeSupport.AsGoStruct(&m.Position, unsafe.Pointer(&mem.position))
	KeyValue__Sequence_to_Go(&m.Props, mem.props)
}

func (t _WayPointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__WayPoint())
}

type CWayPoint = C.geographic_msgs__msg__WayPoint
type CWayPoint__Sequence = C.geographic_msgs__msg__WayPoint__Sequence

func WayPoint__Sequence_to_Go(goSlice *[]WayPoint, cSlice CWayPoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WayPoint, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geographic_msgs__msg__WayPoint__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__msg__WayPoint * uintptr(i)),
		))
		WayPointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func WayPoint__Sequence_to_C(cSlice *CWayPoint__Sequence, goSlice []WayPoint) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__WayPoint)(C.malloc((C.size_t)(C.sizeof_struct_geographic_msgs__msg__WayPoint * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geographic_msgs__msg__WayPoint)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__msg__WayPoint * uintptr(i)),
		))
		WayPointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func WayPoint__Array_to_Go(goSlice []WayPoint, cSlice []CWayPoint) {
	for i := 0; i < len(cSlice); i++ {
		WayPointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WayPoint__Array_to_C(cSlice []CWayPoint, goSlice []WayPoint) {
	for i := 0; i < len(goSlice); i++ {
		WayPointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
