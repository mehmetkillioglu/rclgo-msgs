/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/channel_float32.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/ChannelFloat32", ChannelFloat32TypeSupport)
}

// Do not create instances of this type directly. Always use NewChannelFloat32
// function instead.
type ChannelFloat32 struct {
	Name string `yaml:"name"`// The channel name should give semantics of the channel (e.g."intensity" instead of "value").
	Values []float32 `yaml:"values"`// The values array should be 1-1 with the elements of the associatedPointCloud.
}

// NewChannelFloat32 creates a new ChannelFloat32 with default values.
func NewChannelFloat32() *ChannelFloat32 {
	self := ChannelFloat32{}
	self.SetDefaults()
	return &self
}

func (t *ChannelFloat32) Clone() *ChannelFloat32 {
	c := &ChannelFloat32{}
	c.Name = t.Name
	if t.Values != nil {
		c.Values = make([]float32, len(t.Values))
		copy(c.Values, t.Values)
	}
	return c
}

func (t *ChannelFloat32) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ChannelFloat32) SetDefaults() {
	t.Name = ""
	t.Values = nil
}

// CloneChannelFloat32Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneChannelFloat32Slice(dst, src []ChannelFloat32) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ChannelFloat32TypeSupport types.MessageTypeSupport = _ChannelFloat32TypeSupport{}

type _ChannelFloat32TypeSupport struct{}

func (t _ChannelFloat32TypeSupport) New() types.Message {
	return NewChannelFloat32()
}

func (t _ChannelFloat32TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__ChannelFloat32
	return (unsafe.Pointer)(C.sensor_msgs__msg__ChannelFloat32__create())
}

func (t _ChannelFloat32TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__ChannelFloat32__destroy((*C.sensor_msgs__msg__ChannelFloat32)(pointer_to_free))
}

func (t _ChannelFloat32TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ChannelFloat32)
	mem := (*C.sensor_msgs__msg__ChannelFloat32)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.name), m.Name)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.values)), m.Values)
}

func (t _ChannelFloat32TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ChannelFloat32)
	mem := (*C.sensor_msgs__msg__ChannelFloat32)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Name, unsafe.Pointer(&mem.name))
	primitives.Float32__Sequence_to_Go(&m.Values, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.values)))
}

func (t _ChannelFloat32TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__ChannelFloat32())
}

type CChannelFloat32 = C.sensor_msgs__msg__ChannelFloat32
type CChannelFloat32__Sequence = C.sensor_msgs__msg__ChannelFloat32__Sequence

func ChannelFloat32__Sequence_to_Go(goSlice *[]ChannelFloat32, cSlice CChannelFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ChannelFloat32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__ChannelFloat32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(i)),
		))
		ChannelFloat32TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ChannelFloat32__Sequence_to_C(cSlice *CChannelFloat32__Sequence, goSlice []ChannelFloat32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__ChannelFloat32)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__ChannelFloat32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__ChannelFloat32 * uintptr(i)),
		))
		ChannelFloat32TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ChannelFloat32__Array_to_Go(goSlice []ChannelFloat32, cSlice []CChannelFloat32) {
	for i := 0; i < len(cSlice); i++ {
		ChannelFloat32TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ChannelFloat32__Array_to_C(cSlice []CChannelFloat32, goSlice []ChannelFloat32) {
	for i := 0; i < len(goSlice); i++ {
		ChannelFloat32TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
