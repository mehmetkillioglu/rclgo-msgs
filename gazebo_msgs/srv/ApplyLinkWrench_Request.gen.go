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

#include <gazebo_msgs/srv/apply_link_wrench.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/ApplyLinkWrench_Request", ApplyLinkWrench_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewApplyLinkWrench_Request
// function instead.
type ApplyLinkWrench_Request struct {
	LinkName string `yaml:"link_name"`// Gazebo link to apply wrench (linear force and torque). Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world frame
	ReferenceFrame string `yaml:"reference_frame"`// wrench is defined in the reference frame of this entity. Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_link
	ReferencePoint geometry_msgs_msg.Point `yaml:"reference_point"`// wrench is defined at this location in the reference frame. Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_link
	Wrench geometry_msgs_msg.Wrench `yaml:"wrench"`// wrench applied to the origin of the link. Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_link
	StartTime builtin_interfaces_msg.Time `yaml:"start_time"`// (optional) wrench application start time (seconds). Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_link
	Duration builtin_interfaces_msg.Duration `yaml:"duration"`// optional duration of wrench application time (seconds). Apply Wrench to Gazebo Link.via the callback mechanismall Gazebo operations are made in world framewrench is applied in the gazebo world by defaultlink names are prefixed by model name, e.g. pr2::base_linkuse inertial frame if left emptyframe names are links prefixed by model name, e.g. pr2::base_linkif start_time is not specified, orstart_time < current time, start as soon as possible
}

// NewApplyLinkWrench_Request creates a new ApplyLinkWrench_Request with default values.
func NewApplyLinkWrench_Request() *ApplyLinkWrench_Request {
	self := ApplyLinkWrench_Request{}
	self.SetDefaults()
	return &self
}

func (t *ApplyLinkWrench_Request) Clone() *ApplyLinkWrench_Request {
	c := &ApplyLinkWrench_Request{}
	c.LinkName = t.LinkName
	c.ReferenceFrame = t.ReferenceFrame
	c.ReferencePoint = *t.ReferencePoint.Clone()
	c.Wrench = *t.Wrench.Clone()
	c.StartTime = *t.StartTime.Clone()
	c.Duration = *t.Duration.Clone()
	return c
}

func (t *ApplyLinkWrench_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ApplyLinkWrench_Request) SetDefaults() {
	t.LinkName = ""
	t.ReferenceFrame = ""
	t.ReferencePoint.SetDefaults()
	t.Wrench.SetDefaults()
	t.StartTime.SetDefaults()
	t.Duration.SetDefaults()
}

// CloneApplyLinkWrench_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneApplyLinkWrench_RequestSlice(dst, src []ApplyLinkWrench_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ApplyLinkWrench_RequestTypeSupport types.MessageTypeSupport = _ApplyLinkWrench_RequestTypeSupport{}

type _ApplyLinkWrench_RequestTypeSupport struct{}

func (t _ApplyLinkWrench_RequestTypeSupport) New() types.Message {
	return NewApplyLinkWrench_Request()
}

func (t _ApplyLinkWrench_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__ApplyLinkWrench_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__ApplyLinkWrench_Request__create())
}

func (t _ApplyLinkWrench_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__ApplyLinkWrench_Request__destroy((*C.gazebo_msgs__srv__ApplyLinkWrench_Request)(pointer_to_free))
}

func (t _ApplyLinkWrench_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ApplyLinkWrench_Request)
	mem := (*C.gazebo_msgs__srv__ApplyLinkWrench_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.link_name), m.LinkName)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.reference_frame), m.ReferenceFrame)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.reference_point), &m.ReferencePoint)
	geometry_msgs_msg.WrenchTypeSupport.AsCStruct(unsafe.Pointer(&mem.wrench), &m.Wrench)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.start_time), &m.StartTime)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.duration), &m.Duration)
}

func (t _ApplyLinkWrench_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ApplyLinkWrench_Request)
	mem := (*C.gazebo_msgs__srv__ApplyLinkWrench_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.LinkName, unsafe.Pointer(&mem.link_name))
	primitives.StringAsGoStruct(&m.ReferenceFrame, unsafe.Pointer(&mem.reference_frame))
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.ReferencePoint, unsafe.Pointer(&mem.reference_point))
	geometry_msgs_msg.WrenchTypeSupport.AsGoStruct(&m.Wrench, unsafe.Pointer(&mem.wrench))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.StartTime, unsafe.Pointer(&mem.start_time))
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Duration, unsafe.Pointer(&mem.duration))
}

func (t _ApplyLinkWrench_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__ApplyLinkWrench_Request())
}

type CApplyLinkWrench_Request = C.gazebo_msgs__srv__ApplyLinkWrench_Request
type CApplyLinkWrench_Request__Sequence = C.gazebo_msgs__srv__ApplyLinkWrench_Request__Sequence

func ApplyLinkWrench_Request__Sequence_to_Go(goSlice *[]ApplyLinkWrench_Request, cSlice CApplyLinkWrench_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ApplyLinkWrench_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__ApplyLinkWrench_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Request * uintptr(i)),
		))
		ApplyLinkWrench_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func ApplyLinkWrench_Request__Sequence_to_C(cSlice *CApplyLinkWrench_Request__Sequence, goSlice []ApplyLinkWrench_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__ApplyLinkWrench_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__ApplyLinkWrench_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__ApplyLinkWrench_Request * uintptr(i)),
		))
		ApplyLinkWrench_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func ApplyLinkWrench_Request__Array_to_Go(goSlice []ApplyLinkWrench_Request, cSlice []CApplyLinkWrench_Request) {
	for i := 0; i < len(cSlice); i++ {
		ApplyLinkWrench_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ApplyLinkWrench_Request__Array_to_C(cSlice []CApplyLinkWrench_Request, goSlice []ApplyLinkWrench_Request) {
	for i := 0; i < len(goSlice); i++ {
		ApplyLinkWrench_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
