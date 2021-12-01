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

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgazebo_msgs__rosidl_typesupport_c -lgazebo_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <gazebo_msgs/srv/get_physics_properties.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("gazebo_msgs/GetPhysicsProperties", GetPhysicsPropertiesTypeSupport)
}

type _GetPhysicsPropertiesTypeSupport struct {}

func (s _GetPhysicsPropertiesTypeSupport) Request() types.MessageTypeSupport {
	return GetPhysicsProperties_RequestTypeSupport
}

func (s _GetPhysicsPropertiesTypeSupport) Response() types.MessageTypeSupport {
	return GetPhysicsProperties_ResponseTypeSupport
}

func (s _GetPhysicsPropertiesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__gazebo_msgs__srv__GetPhysicsProperties())
}

// Modifying this variable is undefined behavior.
var GetPhysicsPropertiesTypeSupport types.ServiceTypeSupport = _GetPhysicsPropertiesTypeSupport{}
