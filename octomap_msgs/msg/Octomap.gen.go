/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package octomap_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -loctomap_msgs__rosidl_typesupport_c -loctomap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <octomap_msgs/msg/octomap.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("octomap_msgs/Octomap", OctomapTypeSupport)
}

// Do not create instances of this type directly. Always use NewOctomap
// function instead.
type Octomap struct {
	Header std_msgs_msg.Header `yaml:"header"`// A 3D map in binary format, as Octree
	Binary bool `yaml:"binary"`// Flag to denote a binary (only free/occupied) or full occupancy octree (.bt/.ot file)
	Id string `yaml:"id"`// Class id of the contained octree
	Resolution float64 `yaml:"resolution"`// Resolution (in m) of the smallest octree nodes
	Data []int8 `yaml:"data"`// binary serialization of octree, use conversions.h to read and write octrees
}

// NewOctomap creates a new Octomap with default values.
func NewOctomap() *Octomap {
	self := Octomap{}
	self.SetDefaults()
	return &self
}

func (t *Octomap) Clone() *Octomap {
	c := &Octomap{}
	c.Header = *t.Header.Clone()
	c.Binary = t.Binary
	c.Id = t.Id
	c.Resolution = t.Resolution
	if t.Data != nil {
		c.Data = make([]int8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *Octomap) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Octomap) SetDefaults() {
	t.Header.SetDefaults()
	t.Binary = false
	t.Id = ""
	t.Resolution = 0
	t.Data = nil
}

// CloneOctomapSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneOctomapSlice(dst, src []Octomap) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var OctomapTypeSupport types.MessageTypeSupport = _OctomapTypeSupport{}

type _OctomapTypeSupport struct{}

func (t _OctomapTypeSupport) New() types.Message {
	return NewOctomap()
}

func (t _OctomapTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.octomap_msgs__msg__Octomap
	return (unsafe.Pointer)(C.octomap_msgs__msg__Octomap__create())
}

func (t _OctomapTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.octomap_msgs__msg__Octomap__destroy((*C.octomap_msgs__msg__Octomap)(pointer_to_free))
}

func (t _OctomapTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Octomap)
	mem := (*C.octomap_msgs__msg__Octomap)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.binary = C.bool(m.Binary)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.id), m.Id)
	mem.resolution = C.double(m.Resolution)
	primitives.Int8__Sequence_to_C((*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _OctomapTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Octomap)
	mem := (*C.octomap_msgs__msg__Octomap)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Binary = bool(mem.binary)
	primitives.StringAsGoStruct(&m.Id, unsafe.Pointer(&mem.id))
	m.Resolution = float64(mem.resolution)
	primitives.Int8__Sequence_to_Go(&m.Data, *(*primitives.CInt8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _OctomapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__octomap_msgs__msg__Octomap())
}

type COctomap = C.octomap_msgs__msg__Octomap
type COctomap__Sequence = C.octomap_msgs__msg__Octomap__Sequence

func Octomap__Sequence_to_Go(goSlice *[]Octomap, cSlice COctomap__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Octomap, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.octomap_msgs__msg__Octomap__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__msg__Octomap * uintptr(i)),
		))
		OctomapTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Octomap__Sequence_to_C(cSlice *COctomap__Sequence, goSlice []Octomap) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.octomap_msgs__msg__Octomap)(C.malloc((C.size_t)(C.sizeof_struct_octomap_msgs__msg__Octomap * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.octomap_msgs__msg__Octomap)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_octomap_msgs__msg__Octomap * uintptr(i)),
		))
		OctomapTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Octomap__Array_to_Go(goSlice []Octomap, cSlice []COctomap) {
	for i := 0; i < len(cSlice); i++ {
		OctomapTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Octomap__Array_to_C(cSlice []COctomap, goSlice []Octomap) {
	for i := 0; i < len(goSlice); i++ {
		OctomapTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
