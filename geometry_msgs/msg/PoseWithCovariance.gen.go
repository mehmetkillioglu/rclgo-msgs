/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/pose_with_covariance.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PoseWithCovariance", PoseWithCovarianceTypeSupport)
}

// Do not create instances of this type directly. Always use NewPoseWithCovariance
// function instead.
type PoseWithCovariance struct {
	Pose Pose `yaml:"pose"`
	Covariance [36]float64 `yaml:"covariance"`// Row-major representation of the 6x6 covariance matrixThe orientation parameters use a fixed-axis representation.In order, the parameters are:(x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
}

// NewPoseWithCovariance creates a new PoseWithCovariance with default values.
func NewPoseWithCovariance() *PoseWithCovariance {
	self := PoseWithCovariance{}
	self.SetDefaults()
	return &self
}

func (t *PoseWithCovariance) Clone() *PoseWithCovariance {
	c := &PoseWithCovariance{}
	c.Pose = *t.Pose.Clone()
	c.Covariance = t.Covariance
	return c
}

func (t *PoseWithCovariance) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PoseWithCovariance) SetDefaults() {
	t.Pose.SetDefaults()
	t.Covariance = [36]float64{}
}

// ClonePoseWithCovarianceSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePoseWithCovarianceSlice(dst, src []PoseWithCovariance) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PoseWithCovarianceTypeSupport types.MessageTypeSupport = _PoseWithCovarianceTypeSupport{}

type _PoseWithCovarianceTypeSupport struct{}

func (t _PoseWithCovarianceTypeSupport) New() types.Message {
	return NewPoseWithCovariance()
}

func (t _PoseWithCovarianceTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PoseWithCovariance
	return (unsafe.Pointer)(C.geometry_msgs__msg__PoseWithCovariance__create())
}

func (t _PoseWithCovarianceTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PoseWithCovariance__destroy((*C.geometry_msgs__msg__PoseWithCovariance)(pointer_to_free))
}

func (t _PoseWithCovarianceTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PoseWithCovariance)
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(dst)
	PoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)), m.Covariance[:])
}

func (t _PoseWithCovarianceTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PoseWithCovariance)
	mem := (*C.geometry_msgs__msg__PoseWithCovariance)(ros2_message_buffer)
	PoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
	cSlice_covariance := mem.covariance[:]
	primitives.Float64__Array_to_Go(m.Covariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_covariance)))
}

func (t _PoseWithCovarianceTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PoseWithCovariance())
}

type CPoseWithCovariance = C.geometry_msgs__msg__PoseWithCovariance
type CPoseWithCovariance__Sequence = C.geometry_msgs__msg__PoseWithCovariance__Sequence

func PoseWithCovariance__Sequence_to_Go(goSlice *[]PoseWithCovariance, cSlice CPoseWithCovariance__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PoseWithCovariance, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__PoseWithCovariance__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(i)),
		))
		PoseWithCovarianceTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func PoseWithCovariance__Sequence_to_C(cSlice *CPoseWithCovariance__Sequence, goSlice []PoseWithCovariance) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PoseWithCovariance)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__PoseWithCovariance)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__PoseWithCovariance * uintptr(i)),
		))
		PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func PoseWithCovariance__Array_to_Go(goSlice []PoseWithCovariance, cSlice []CPoseWithCovariance) {
	for i := 0; i < len(cSlice); i++ {
		PoseWithCovarianceTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PoseWithCovariance__Array_to_C(cSlice []CPoseWithCovariance, goSlice []PoseWithCovariance) {
	for i := 0; i < len(goSlice); i++ {
		PoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
