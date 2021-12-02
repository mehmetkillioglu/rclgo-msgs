/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	sensor_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/sensor_msgs/msg"
	std_msgs_msg "github.com/mehmetkillioglu/rclgo-msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/msg/point_cloud2_update.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/PointCloud2Update", PointCloud2UpdateTypeSupport)
}
const (
	PointCloud2Update_ADD uint32 = 0
	PointCloud2Update_DELETE uint32 = 1
)

// Do not create instances of this type directly. Always use NewPointCloud2Update
// function instead.
type PointCloud2Update struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Type uint32 `yaml:"type"`// type of update, one of ADD or DELETE
	Points sensor_msgs_msg.PointCloud2 `yaml:"points"`
}

// NewPointCloud2Update creates a new PointCloud2Update with default values.
func NewPointCloud2Update() *PointCloud2Update {
	self := PointCloud2Update{}
	self.SetDefaults()
	return &self
}

func (t *PointCloud2Update) Clone() *PointCloud2Update {
	c := &PointCloud2Update{}
	c.Header = *t.Header.Clone()
	c.Type = t.Type
	c.Points = *t.Points.Clone()
	return c
}

func (t *PointCloud2Update) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PointCloud2Update) SetDefaults() {
	t.Header.SetDefaults()
	t.Type = 0
	t.Points.SetDefaults()
}

// ClonePointCloud2UpdateSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointCloud2UpdateSlice(dst, src []PointCloud2Update) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointCloud2UpdateTypeSupport types.MessageTypeSupport = _PointCloud2UpdateTypeSupport{}

type _PointCloud2UpdateTypeSupport struct{}

func (t _PointCloud2UpdateTypeSupport) New() types.Message {
	return NewPointCloud2Update()
}

func (t _PointCloud2UpdateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__msg__PointCloud2Update
	return (unsafe.Pointer)(C.map_msgs__msg__PointCloud2Update__create())
}

func (t _PointCloud2UpdateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__msg__PointCloud2Update__destroy((*C.map_msgs__msg__PointCloud2Update)(pointer_to_free))
}

func (t _PointCloud2UpdateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointCloud2Update)
	mem := (*C.map_msgs__msg__PointCloud2Update)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem._type = C.uint32_t(m.Type)
	sensor_msgs_msg.PointCloud2TypeSupport.AsCStruct(unsafe.Pointer(&mem.points), &m.Points)
}

func (t _PointCloud2UpdateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointCloud2Update)
	mem := (*C.map_msgs__msg__PointCloud2Update)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Type = uint32(mem._type)
	sensor_msgs_msg.PointCloud2TypeSupport.AsGoStruct(&m.Points, unsafe.Pointer(&mem.points))
}

func (t _PointCloud2UpdateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__msg__PointCloud2Update())
}

type CPointCloud2Update = C.map_msgs__msg__PointCloud2Update
type CPointCloud2Update__Sequence = C.map_msgs__msg__PointCloud2Update__Sequence

func PointCloud2Update__Sequence_to_Go(goSlice *[]PointCloud2Update, cSlice CPointCloud2Update__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointCloud2Update, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__msg__PointCloud2Update__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__msg__PointCloud2Update * uintptr(i)),
		))
		PointCloud2UpdateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PointCloud2Update__Sequence_to_C(cSlice *CPointCloud2Update__Sequence, goSlice []PointCloud2Update) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__msg__PointCloud2Update)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__msg__PointCloud2Update * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__msg__PointCloud2Update)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__msg__PointCloud2Update * uintptr(i)),
		))
		PointCloud2UpdateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PointCloud2Update__Array_to_Go(goSlice []PointCloud2Update, cSlice []CPointCloud2Update) {
	for i := 0; i < len(cSlice); i++ {
		PointCloud2UpdateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointCloud2Update__Array_to_C(cSlice []CPointCloud2Update, goSlice []PointCloud2Update) {
	for i := 0; i < len(goSlice); i++ {
		PointCloud2UpdateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
