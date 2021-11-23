/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package fog_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lfog_msgs__rosidl_typesupport_c -lfog_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <fog_msgs/srv/change_odometry.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("fog_msgs/ChangeOdometry", ChangeOdometryTypeSupport)
}

type _ChangeOdometryTypeSupport struct {}

func (s _ChangeOdometryTypeSupport) Request() types.MessageTypeSupport {
	return ChangeOdometry_RequestTypeSupport
}

func (s _ChangeOdometryTypeSupport) Response() types.MessageTypeSupport {
	return ChangeOdometry_ResponseTypeSupport
}

func (s _ChangeOdometryTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__fog_msgs__srv__ChangeOdometry())
}

// Modifying this variable is undefined behavior.
var ChangeOdometryTypeSupport types.ServiceTypeSupport = _ChangeOdometryTypeSupport{}
