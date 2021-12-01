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
	gazebo_msgs_msg "github.com/tiiuae/rclgo-msgs/gazebo_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <gazebo_msgs/srv/set_joint_properties.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("gazebo_msgs/SetJointProperties_Request", SetJointProperties_RequestTypeSupport)
}

// Do not create instances of this type directly. Always use NewSetJointProperties_Request
// function instead.
type SetJointProperties_Request struct {
	JointName string `yaml:"joint_name"`// name of joint
	OdeJointConfig gazebo_msgs_msg.ODEJointProperties `yaml:"ode_joint_config"`// access to ODE joint dynamics properties
}

// NewSetJointProperties_Request creates a new SetJointProperties_Request with default values.
func NewSetJointProperties_Request() *SetJointProperties_Request {
	self := SetJointProperties_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetJointProperties_Request) Clone() *SetJointProperties_Request {
	c := &SetJointProperties_Request{}
	c.JointName = t.JointName
	c.OdeJointConfig = *t.OdeJointConfig.Clone()
	return c
}

func (t *SetJointProperties_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetJointProperties_Request) SetDefaults() {
	t.JointName = ""
	t.OdeJointConfig.SetDefaults()
}

// CloneSetJointProperties_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetJointProperties_RequestSlice(dst, src []SetJointProperties_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetJointProperties_RequestTypeSupport types.MessageTypeSupport = _SetJointProperties_RequestTypeSupport{}

type _SetJointProperties_RequestTypeSupport struct{}

func (t _SetJointProperties_RequestTypeSupport) New() types.Message {
	return NewSetJointProperties_Request()
}

func (t _SetJointProperties_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.gazebo_msgs__srv__SetJointProperties_Request
	return (unsafe.Pointer)(C.gazebo_msgs__srv__SetJointProperties_Request__create())
}

func (t _SetJointProperties_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.gazebo_msgs__srv__SetJointProperties_Request__destroy((*C.gazebo_msgs__srv__SetJointProperties_Request)(pointer_to_free))
}

func (t _SetJointProperties_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetJointProperties_Request)
	mem := (*C.gazebo_msgs__srv__SetJointProperties_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.joint_name), m.JointName)
	gazebo_msgs_msg.ODEJointPropertiesTypeSupport.AsCStruct(unsafe.Pointer(&mem.ode_joint_config), &m.OdeJointConfig)
}

func (t _SetJointProperties_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetJointProperties_Request)
	mem := (*C.gazebo_msgs__srv__SetJointProperties_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.JointName, unsafe.Pointer(&mem.joint_name))
	gazebo_msgs_msg.ODEJointPropertiesTypeSupport.AsGoStruct(&m.OdeJointConfig, unsafe.Pointer(&mem.ode_joint_config))
}

func (t _SetJointProperties_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__gazebo_msgs__srv__SetJointProperties_Request())
}

type CSetJointProperties_Request = C.gazebo_msgs__srv__SetJointProperties_Request
type CSetJointProperties_Request__Sequence = C.gazebo_msgs__srv__SetJointProperties_Request__Sequence

func SetJointProperties_Request__Sequence_to_Go(goSlice *[]SetJointProperties_Request, cSlice CSetJointProperties_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetJointProperties_Request, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.gazebo_msgs__srv__SetJointProperties_Request__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__SetJointProperties_Request * uintptr(i)),
		))
		SetJointProperties_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SetJointProperties_Request__Sequence_to_C(cSlice *CSetJointProperties_Request__Sequence, goSlice []SetJointProperties_Request) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.gazebo_msgs__srv__SetJointProperties_Request)(C.malloc((C.size_t)(C.sizeof_struct_gazebo_msgs__srv__SetJointProperties_Request * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.gazebo_msgs__srv__SetJointProperties_Request)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_gazebo_msgs__srv__SetJointProperties_Request * uintptr(i)),
		))
		SetJointProperties_RequestTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SetJointProperties_Request__Array_to_Go(goSlice []SetJointProperties_Request, cSlice []CSetJointProperties_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetJointProperties_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetJointProperties_Request__Array_to_C(cSlice []CSetJointProperties_Request, goSlice []SetJointProperties_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetJointProperties_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
