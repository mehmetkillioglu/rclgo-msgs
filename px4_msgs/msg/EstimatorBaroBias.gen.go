/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/estimator_baro_bias.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/EstimatorBaroBias", EstimatorBaroBiasTypeSupport)
}

// Do not create instances of this type directly. Always use NewEstimatorBaroBias
// function instead.
type EstimatorBaroBias struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	BaroDeviceId uint32 `yaml:"baro_device_id"`// unique device ID for the sensor that does not change between power cycles
	Bias float32 `yaml:"bias"`// estimated barometric altitude bias (m)
	BiasVar float32 `yaml:"bias_var"`// estimated barometric altitude bias variance (m^2)
	Innov float32 `yaml:"innov"`// innovation of the last measurement fusion (m)
	InnovVar float32 `yaml:"innov_var"`// innovation variance of the last measurement fusion (m^2)
	InnovTestRatio float32 `yaml:"innov_test_ratio"`// normalized innovation squared test ratio
}

// NewEstimatorBaroBias creates a new EstimatorBaroBias with default values.
func NewEstimatorBaroBias() *EstimatorBaroBias {
	self := EstimatorBaroBias{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorBaroBias) Clone() *EstimatorBaroBias {
	c := &EstimatorBaroBias{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.BaroDeviceId = t.BaroDeviceId
	c.Bias = t.Bias
	c.BiasVar = t.BiasVar
	c.Innov = t.Innov
	c.InnovVar = t.InnovVar
	c.InnovTestRatio = t.InnovTestRatio
	return c
}

func (t *EstimatorBaroBias) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorBaroBias) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.BaroDeviceId = 0
	t.Bias = 0
	t.BiasVar = 0
	t.Innov = 0
	t.InnovVar = 0
	t.InnovTestRatio = 0
}

// CloneEstimatorBaroBiasSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorBaroBiasSlice(dst, src []EstimatorBaroBias) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorBaroBiasTypeSupport types.MessageTypeSupport = _EstimatorBaroBiasTypeSupport{}

type _EstimatorBaroBiasTypeSupport struct{}

func (t _EstimatorBaroBiasTypeSupport) New() types.Message {
	return NewEstimatorBaroBias()
}

func (t _EstimatorBaroBiasTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__EstimatorBaroBias
	return (unsafe.Pointer)(C.px4_msgs__msg__EstimatorBaroBias__create())
}

func (t _EstimatorBaroBiasTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__EstimatorBaroBias__destroy((*C.px4_msgs__msg__EstimatorBaroBias)(pointer_to_free))
}

func (t _EstimatorBaroBiasTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorBaroBias)
	mem := (*C.px4_msgs__msg__EstimatorBaroBias)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.baro_device_id = C.uint32_t(m.BaroDeviceId)
	mem.bias = C.float(m.Bias)
	mem.bias_var = C.float(m.BiasVar)
	mem.innov = C.float(m.Innov)
	mem.innov_var = C.float(m.InnovVar)
	mem.innov_test_ratio = C.float(m.InnovTestRatio)
}

func (t _EstimatorBaroBiasTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorBaroBias)
	mem := (*C.px4_msgs__msg__EstimatorBaroBias)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.BaroDeviceId = uint32(mem.baro_device_id)
	m.Bias = float32(mem.bias)
	m.BiasVar = float32(mem.bias_var)
	m.Innov = float32(mem.innov)
	m.InnovVar = float32(mem.innov_var)
	m.InnovTestRatio = float32(mem.innov_test_ratio)
}

func (t _EstimatorBaroBiasTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__EstimatorBaroBias())
}

type CEstimatorBaroBias = C.px4_msgs__msg__EstimatorBaroBias
type CEstimatorBaroBias__Sequence = C.px4_msgs__msg__EstimatorBaroBias__Sequence

func EstimatorBaroBias__Sequence_to_Go(goSlice *[]EstimatorBaroBias, cSlice CEstimatorBaroBias__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorBaroBias, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__EstimatorBaroBias__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorBaroBias * uintptr(i)),
		))
		EstimatorBaroBiasTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func EstimatorBaroBias__Sequence_to_C(cSlice *CEstimatorBaroBias__Sequence, goSlice []EstimatorBaroBias) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__EstimatorBaroBias)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__EstimatorBaroBias * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__EstimatorBaroBias)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__EstimatorBaroBias * uintptr(i)),
		))
		EstimatorBaroBiasTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func EstimatorBaroBias__Array_to_Go(goSlice []EstimatorBaroBias, cSlice []CEstimatorBaroBias) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorBaroBiasTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorBaroBias__Array_to_C(cSlice []CEstimatorBaroBias, goSlice []EstimatorBaroBias) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorBaroBiasTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
