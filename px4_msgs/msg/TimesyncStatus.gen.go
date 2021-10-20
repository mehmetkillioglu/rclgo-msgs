/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/timesync_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/TimesyncStatus", TimesyncStatusTypeSupport)
}
const (
	TimesyncStatus_SOURCE_PROTOCOL_MAVLINK uint8 = 0
	TimesyncStatus_SOURCE_PROTOCOL_RTPS uint8 = 1
)

// Do not create instances of this type directly. Always use NewTimesyncStatus
// function instead.
type TimesyncStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	SourceProtocol uint8 `yaml:"source_protocol"`// timesync source. Source can be MAVLink or the microRTPS bridge
	RemoteTimestamp uint64 `yaml:"remote_timestamp"`// remote system timestamp (microseconds)
	ObservedOffset int64 `yaml:"observed_offset"`// raw time offset directly observed from this timesync packet (microseconds)
	EstimatedOffset int64 `yaml:"estimated_offset"`// smoothed time offset between companion system and PX4 (microseconds)
	RoundTripTime uint32 `yaml:"round_trip_time"`// round trip time of this timesync packet (microseconds)
}

// NewTimesyncStatus creates a new TimesyncStatus with default values.
func NewTimesyncStatus() *TimesyncStatus {
	self := TimesyncStatus{}
	self.SetDefaults()
	return &self
}

func (t *TimesyncStatus) Clone() *TimesyncStatus {
	c := &TimesyncStatus{}
	c.Timestamp = t.Timestamp
	c.SourceProtocol = t.SourceProtocol
	c.RemoteTimestamp = t.RemoteTimestamp
	c.ObservedOffset = t.ObservedOffset
	c.EstimatedOffset = t.EstimatedOffset
	c.RoundTripTime = t.RoundTripTime
	return c
}

func (t *TimesyncStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TimesyncStatus) SetDefaults() {
	t.Timestamp = 0
	t.SourceProtocol = 0
	t.RemoteTimestamp = 0
	t.ObservedOffset = 0
	t.EstimatedOffset = 0
	t.RoundTripTime = 0
}

// CloneTimesyncStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTimesyncStatusSlice(dst, src []TimesyncStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TimesyncStatusTypeSupport types.MessageTypeSupport = _TimesyncStatusTypeSupport{}

type _TimesyncStatusTypeSupport struct{}

func (t _TimesyncStatusTypeSupport) New() types.Message {
	return NewTimesyncStatus()
}

func (t _TimesyncStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__TimesyncStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__TimesyncStatus__create())
}

func (t _TimesyncStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__TimesyncStatus__destroy((*C.px4_msgs__msg__TimesyncStatus)(pointer_to_free))
}

func (t _TimesyncStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TimesyncStatus)
	mem := (*C.px4_msgs__msg__TimesyncStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.source_protocol = C.uint8_t(m.SourceProtocol)
	mem.remote_timestamp = C.uint64_t(m.RemoteTimestamp)
	mem.observed_offset = C.int64_t(m.ObservedOffset)
	mem.estimated_offset = C.int64_t(m.EstimatedOffset)
	mem.round_trip_time = C.uint32_t(m.RoundTripTime)
}

func (t _TimesyncStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TimesyncStatus)
	mem := (*C.px4_msgs__msg__TimesyncStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.SourceProtocol = uint8(mem.source_protocol)
	m.RemoteTimestamp = uint64(mem.remote_timestamp)
	m.ObservedOffset = int64(mem.observed_offset)
	m.EstimatedOffset = int64(mem.estimated_offset)
	m.RoundTripTime = uint32(mem.round_trip_time)
}

func (t _TimesyncStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__TimesyncStatus())
}

type CTimesyncStatus = C.px4_msgs__msg__TimesyncStatus
type CTimesyncStatus__Sequence = C.px4_msgs__msg__TimesyncStatus__Sequence

func TimesyncStatus__Sequence_to_Go(goSlice *[]TimesyncStatus, cSlice CTimesyncStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TimesyncStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__TimesyncStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TimesyncStatus * uintptr(i)),
		))
		TimesyncStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func TimesyncStatus__Sequence_to_C(cSlice *CTimesyncStatus__Sequence, goSlice []TimesyncStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__TimesyncStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__TimesyncStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__TimesyncStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__TimesyncStatus * uintptr(i)),
		))
		TimesyncStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func TimesyncStatus__Array_to_Go(goSlice []TimesyncStatus, cSlice []CTimesyncStatus) {
	for i := 0; i < len(cSlice); i++ {
		TimesyncStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TimesyncStatus__Array_to_C(cSlice []CTimesyncStatus, goSlice []TimesyncStatus) {
	for i := 0; i < len(goSlice); i++ {
		TimesyncStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
