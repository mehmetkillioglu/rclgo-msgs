/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package gazebo_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo-msgs/builtin_interfaces/msg"
	geometry_msgs_msg "github.com/tiiuae/rclgo-msgs/geometry_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c
#cgo LDFLAGS: -lbuiltin_interfaces__rosidl_typesupport_c -lbuiltin_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <gazebo_msgs/srv/apply_body_wrench.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/ApplyBodyWrench_Request", ApplyBodyWrench_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewApplyBodyWrench_Request
// function instead.
type ApplyBodyWrench_Request struct {
	BodyName string `yaml:"body_name"`// Gazebo body to apply wrench (linear force and torque). Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world frame
	ReferenceFrame string `yaml:"reference_frame"`// wrench is defined in the reference frame of this entity. Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultbody names are prefixed by model name, e.g. pr2::base_link
	ReferencePoint geometry_msgs_msg.Point `yaml:"reference_point"`// wrench is defined at this location in the reference frame. Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultbody names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are bodies prefixed by model name, e.g. pr2::base_link
	Wrench geometry_msgs_msg.Wrench `yaml:"wrench"`// wrench applied to the origin of the body. Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultbody names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are bodies prefixed by model name, e.g. pr2::base_link
	StartTime builtin_interfaces_msg.Time `yaml:"start_time"`// (optional) wrench application start time (seconds). Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultbody names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are bodies prefixed by model name, e.g. pr2::base_link
	Duration builtin_interfaces_msg.Duration `yaml:"duration"`// optional duration of wrench application time (seconds). Deprecated, kept for ROS 1 bridge.Use ApplyLinkWrenchApply Wrench to Gazebo Body.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultbody names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are bodies prefixed by model name, e.g. pr2::base_linkif start_time is not specified, orstart_time < current time, start as soon as possible
}

// NewApplyBodyWrench_Request creates a new ApplyBodyWrench_Request with default values.
func NewApplyBodyWrench_Request() *ApplyBodyWrench_Request {
	self := ApplyBodyWrench_Request{}
	self.SetDefaults()
	return &self
}

func (t *ApplyBodyWrench_Request) Clone() *ApplyBodyWrench_Request {
	c := &ApplyBodyWrench_Request{}
	c.BodyName = t.BodyName
	c.ReferenceFrame = t.ReferenceFrame
	c.ReferencePoint = *t.ReferencePoint.Clone()
	c.Wrench = *t.Wrench.Clone()
	c.StartTime = *t.StartTime.Clone()
	c.Duration = *t.Duration.Clone()
	return c
}

func (t *ApplyBodyWrench_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ApplyBodyWrench_Request) SetDefaults() {
	t.BodyName = ""
	t.ReferenceFrame = ""
	t.ReferencePoint.SetDefaults()
	t.Wrench.SetDefaults()
	t.StartTime.SetDefaults()
	t.Duration.SetDefaults()
}

// CloneApplyBodyWrench_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneApplyBodyWrench_RequestSlice(dst, src []ApplyBodyWrench_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ApplyBodyWrench_RequestTypeSupport types.MessageTypeSupport = _ApplyBodyWrench_RequestTypeSupport{}

type _ApplyBodyWrench_RequestTypeSupport struct{}

func (t _ApplyBodyWrench_RequestTypeSupport) New() types.Message {
	return NewApplyBodyWrench_Request()
}

func (t _ApplyBodyWrench_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__ApplyBodyWrench_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__ApplyBodyWrench_Request__create())
}

func (t _ApplyBodyWrench_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__ApplyBodyWrench_Request__destroy((*C.gazebo_msgs__srv__ApplyBodyWrench_Request)(pointer_to_free))
}

func (t _ApplyBodyWrench_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ApplyBodyWrench_Request)
	mem := (*C.gazebo_msgs__srv__ApplyBodyWrench_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.body_name), m.BodyName)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.reference_frame), m.ReferenceFrame)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.reference_point), &m.ReferencePoint)
	geometry_msgs_msg.WrenchTypeSupport.AsCStruct(unsafe.Pointer(&mem.wrench), &m.Wrench)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.start_time), &m.StartTime)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.duration), &m.Duration)
}

func (t _ApplyBodyWrench_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ApplyBodyWrench_Request)
	mem := (*C.gazebo_msgs__srv__ApplyBodyWrench_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.BodyName, unsafe.Pointer(&mem.body_name))
	primitives.StringAsGoStruct(&m.ReferenceFrame, unsafe.Pointer(&mem.reference_frame))
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.ReferencePoint, unsafe.Pointer(&mem.reference_point))
	geometry_msgs_msg.WrenchTypeSupport.AsGoStruct(&m.Wrench, unsafe.Pointer(&mem.wrench))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.StartTime, unsafe.Pointer(&mem.start_time))
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Duration, unsafe.Pointer(&mem.duration))
}

func (t _ApplyBodyWrench_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__ApplyBodyWrench_Request())
}

type CApplyBodyWrench_Request = C.gazebo_msgs__srv__ApplyBodyWrench_Request
type CApplyBodyWrench_Request__Sequence = C.gazebo_msgs__srv__ApplyBodyWrench_Request__Sequence

func ApplyBodyWrench_Request__Sequence_to_Go(goSlice *[]ApplyBodyWrench_Request, cSlice CApplyBodyWrench_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ApplyBodyWrench_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__ApplyBodyWrench_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyBodyWrench_Request * uintptr(i)),
		))
		ApplyBodyWrench_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ApplyBodyWrench_Request__Sequence_to_C(cSlice *CApplyBodyWrench_Request__Sequence, goSlice []ApplyBodyWrench_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__ApplyBodyWrench_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__ApplyBodyWrench_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__ApplyBodyWrench_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyBodyWrench_Request * uintptr(i)),
		))
		ApplyBodyWrench_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ApplyBodyWrench_Request__Array_to_Go(goSlice []ApplyBodyWrench_Request, cSlice []CApplyBodyWrench_Request) {
	for i := 0; i < len(cSlice); i++ {
		ApplyBodyWrench_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ApplyBodyWrench_Request__Array_to_C(cSlice []CApplyBodyWrench_Request, goSlice []ApplyBodyWrench_Request) {
	for i := 0; i < len(goSlice); i++ {
		ApplyBodyWrench_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
