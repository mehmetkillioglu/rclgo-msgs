/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rosbag2_interfaces_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrosbag2_interfaces__rosidl_typesupport_c -lrosbag2_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rosbag2_interfaces/srv/toggle_paused.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("rosbag2_interfaces/TogglePaused", TogglePausedTypeSupport)
}

type _TogglePausedTypeSupport struct {}

func (s _TogglePausedTypeSupport) Request() types.MessageTypeSupport {
	return TogglePaused_RequestTypeSupport
}

func (s _TogglePausedTypeSupport) Response() types.MessageTypeSupport {
	return TogglePaused_ResponseTypeSupport
}

func (s _TogglePausedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rosbag2_interfaces__srv__TogglePaused())
}

// Modifying this variable is undefined behavior.
var TogglePausedTypeSupport types.ServiceTypeSupport = _TogglePausedTypeSupport{}
