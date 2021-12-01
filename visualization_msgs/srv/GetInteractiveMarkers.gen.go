/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <visualization_msgs/srv/get_interactive_markers.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("visualization_msgs/GetInteractiveMarkers", GetInteractiveMarkersTypeSupport)
}

type _GetInteractiveMarkersTypeSupport struct {}

func (s _GetInteractiveMarkersTypeSupport) Request() types.MessageTypeSupport {
	return GetInteractiveMarkers_RequestTypeSupport
}

func (s _GetInteractiveMarkersTypeSupport) Response() types.MessageTypeSupport {
	return GetInteractiveMarkers_ResponseTypeSupport
}

func (s _GetInteractiveMarkersTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__visualization_msgs__srv__GetInteractiveMarkers())
}

// Modifying this variable is undefined behavior.
var GetInteractiveMarkersTypeSupport types.ServiceTypeSupport = _GetInteractiveMarkersTypeSupport{}
