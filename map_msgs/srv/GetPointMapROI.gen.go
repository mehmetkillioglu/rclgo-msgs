/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <map_msgs/srv/get_point_map_roi.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("map_msgs/GetPointMapROI", GetPointMapROITypeSupport)
}

type _GetPointMapROITypeSupport struct {}

func (s _GetPointMapROITypeSupport) Request() types.MessageTypeSupport {
	return GetPointMapROI_RequestTypeSupport
}

func (s _GetPointMapROITypeSupport) Response() types.MessageTypeSupport {
	return GetPointMapROI_ResponseTypeSupport
}

func (s _GetPointMapROITypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__map_msgs__srv__GetPointMapROI())
}

// Modifying this variable is undefined behavior.
var GetPointMapROITypeSupport types.ServiceTypeSupport = _GetPointMapROITypeSupport{}
