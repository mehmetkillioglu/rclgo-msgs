/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <diagnostic_msgs/srv/add_diagnostics.h>
*/
import "C"

import (
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("diagnostic_msgs/AddDiagnostics", AddDiagnosticsTypeSupport)
}

type _AddDiagnosticsTypeSupport struct {}

func (s _AddDiagnosticsTypeSupport) Request() types.MessageTypeSupport {
	return AddDiagnostics_RequestTypeSupport
}

func (s _AddDiagnosticsTypeSupport) Response() types.MessageTypeSupport {
	return AddDiagnostics_ResponseTypeSupport
}

func (s _AddDiagnosticsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__diagnostic_msgs__srv__AddDiagnostics())
}

// Modifying this variable is undefined behavior.
var AddDiagnosticsTypeSupport types.ServiceTypeSupport = _AddDiagnosticsTypeSupport{}
