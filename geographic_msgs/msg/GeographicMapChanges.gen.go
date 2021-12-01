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
	std_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/std_msgs/msg"
	unique_identifier_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/unique_identifier_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeographic_msgs__rosidl_typesupport_c -lgeographic_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c
#cgo LDFLAGS: -lunique_identifier_msgs__rosidl_typesupport_c -lunique_identifier_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/msg/geographic_map_changes.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/GeographicMapChanges", GeographicMapChangesTypeSupport)
}

// Do not create instances of this type directly. Always use NewGeographicMapChanges
// function instead.
type GeographicMapChanges struct {
	Header std_msgs_msg.Header `yaml:"header"`// stamp specifies time of change
	Diffs GeographicMap `yaml:"diffs"`// new and changed points and features
	Deletes []unique_identifier_msgs_msg.UUID `yaml:"deletes"`// deleted map components
}

// NewGeographicMapChanges creates a new GeographicMapChanges with default values.
func NewGeographicMapChanges() *GeographicMapChanges {
	self := GeographicMapChanges{}
	self.SetDefaults()
	return &self
}

func (t *GeographicMapChanges) Clone() *GeographicMapChanges {
	c := &GeographicMapChanges{}
	c.Header = *t.Header.Clone()
	c.Diffs = *t.Diffs.Clone()
	if t.Deletes != nil {
		c.Deletes = make([]unique_identifier_msgs_msg.UUID, len(t.Deletes))
		unique_identifier_msgs_msg.CloneUUIDSlice(c.Deletes, t.Deletes)
	}
	return c
}

func (t *GeographicMapChanges) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GeographicMapChanges) SetDefaults() {
	t.Header.SetDefaults()
	t.Diffs.SetDefaults()
	t.Deletes = nil
}

// CloneGeographicMapChangesSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGeographicMapChangesSlice(dst, src []GeographicMapChanges) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GeographicMapChangesTypeSupport types.MessageTypeSupport = _GeographicMapChangesTypeSupport{}

type _GeographicMapChangesTypeSupport struct{}

func (t _GeographicMapChangesTypeSupport) New() types.Message {
	return NewGeographicMapChanges()
}

func (t _GeographicMapChangesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__GeographicMapChanges
	return (unsafe.Pointer)(C.geographic_msgs__msg__GeographicMapChanges__create())
}

func (t _GeographicMapChangesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__GeographicMapChanges__destroy((*C.geographic_msgs__msg__GeographicMapChanges)(pointer_to_free))
}

func (t _GeographicMapChangesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GeographicMapChanges)
	mem := (*C.geographic_msgs__msg__GeographicMapChanges)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	GeographicMapTypeSupport.AsCStruct(unsafe.Pointer(&mem.diffs), &m.Diffs)
	unique_identifier_msgs_msg.UUID__Sequence_to_C((*unique_identifier_msgs_msg.CUUID__Sequence)(unsafe.Pointer(&mem.deletes)), m.Deletes)
}

func (t _GeographicMapChangesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GeographicMapChanges)
	mem := (*C.geographic_msgs__msg__GeographicMapChanges)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	GeographicMapTypeSupport.AsGoStruct(&m.Diffs, unsafe.Pointer(&mem.diffs))
	unique_identifier_msgs_msg.UUID__Sequence_to_Go(&m.Deletes, *(*unique_identifier_msgs_msg.CUUID__Sequence)(unsafe.Pointer(&mem.deletes)))
}

func (t _GeographicMapChangesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__GeographicMapChanges())
}

type CGeographicMapChanges = C.geographic_msgs__msg__GeographicMapChanges
type CGeographicMapChanges__Sequence = C.geographic_msgs__msg__GeographicMapChanges__Sequence

func GeographicMapChanges__Sequence_to_Go(goSlice *[]GeographicMapChanges, cSlice CGeographicMapChanges__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GeographicMapChanges, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geographic_msgs__msg__GeographicMapChanges__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__msg__GeographicMapChanges * uintptr(i)),
		))
		GeographicMapChangesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func GeographicMapChanges__Sequence_to_C(cSlice *CGeographicMapChanges__Sequence, goSlice []GeographicMapChanges) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__GeographicMapChanges)(C.malloc((C.size_t)(C.sizeof_struct_geographic_msgs__msg__GeographicMapChanges * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geographic_msgs__msg__GeographicMapChanges)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geographic_msgs__msg__GeographicMapChanges * uintptr(i)),
		))
		GeographicMapChangesTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func GeographicMapChanges__Array_to_Go(goSlice []GeographicMapChanges, cSlice []CGeographicMapChanges) {
	for i := 0; i < len(cSlice); i++ {
		GeographicMapChangesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GeographicMapChanges__Array_to_C(cSlice []CGeographicMapChanges, goSlice []GeographicMapChanges) {
	for i := 0; i < len(goSlice); i++ {
		GeographicMapChangesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
