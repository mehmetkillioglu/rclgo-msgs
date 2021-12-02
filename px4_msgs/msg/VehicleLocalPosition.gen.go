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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/vehicle_local_position.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/VehicleLocalPosition", VehicleLocalPositionTypeSupport)
}
const (
	VehicleLocalPosition_DIST_BOTTOM_SENSOR_NONE uint8 = 0// Distance to surface
	VehicleLocalPosition_DIST_BOTTOM_SENSOR_RANGE uint8 = 1// (1 << 0) a range sensor is used to estimate dist_bottom field. Distance to surface
	VehicleLocalPosition_DIST_BOTTOM_SENSOR_FLOW uint8 = 2// (1 << 1) a flow sensor is used to estimate dist_bottom field (mostly fixed-wing use case). Distance to surface
)

// Do not create instances of this type directly. Always use NewVehicleLocalPosition
// function instead.
type VehicleLocalPosition struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	TimestampSample uint64 `yaml:"timestamp_sample"`// the timestamp of the raw data (microseconds)
	XyValid bool `yaml:"xy_valid"`// true if x and y are valid
	ZValid bool `yaml:"z_valid"`// true if z is valid
	VXyValid bool `yaml:"v_xy_valid"`// true if vy and vy are valid
	VZValid bool `yaml:"v_z_valid"`// true if vz is valid
	X float32 `yaml:"x"`// North position in NED earth-fixed frame, (metres). Position in local NED frame
	Y float32 `yaml:"y"`// East position in NED earth-fixed frame, (metres). Position in local NED frame
	Z float32 `yaml:"z"`// Down position (negative altitude) in NED earth-fixed frame, (metres). Position in local NED frame
	DeltaXy [2]float32 `yaml:"delta_xy"`// Position reset delta
	XyResetCounter uint8 `yaml:"xy_reset_counter"`// Position reset delta
	DeltaZ float32 `yaml:"delta_z"`
	ZResetCounter uint8 `yaml:"z_reset_counter"`
	Vx float32 `yaml:"vx"`// North velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	Vy float32 `yaml:"vy"`// East velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	Vz float32 `yaml:"vz"`// Down velocity in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	ZDeriv float32 `yaml:"z_deriv"`// Down position time derivative in NED earth-fixed frame, (metres/sec). Velocity in NED frame
	DeltaVxy [2]float32 `yaml:"delta_vxy"`// Velocity reset delta
	VxyResetCounter uint8 `yaml:"vxy_reset_counter"`// Velocity reset delta
	DeltaVz float32 `yaml:"delta_vz"`
	VzResetCounter uint8 `yaml:"vz_reset_counter"`
	Ax float32 `yaml:"ax"`// North velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Ay float32 `yaml:"ay"`// East velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Az float32 `yaml:"az"`// Down velocity derivative in NED earth-fixed frame, (metres/sec^2). Acceleration in NED frame
	Heading float32 `yaml:"heading"`// Euler yaw angle transforming the tangent plane relative to NED earth-fixed frame, -PI..+PI,  (radians)
	DeltaHeading float32 `yaml:"delta_heading"`
	HeadingResetCounter uint8 `yaml:"heading_reset_counter"`
	XyGlobal bool `yaml:"xy_global"`// true if position (x, y) has a valid global reference (ref_lat, ref_lon). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	ZGlobal bool `yaml:"z_global"`// true if z has a valid global reference (ref_alt). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefTimestamp uint64 `yaml:"ref_timestamp"`// Time when reference position was set since system start, (microseconds). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefLat float64 `yaml:"ref_lat"`// Reference point latitude, (degrees). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefLon float64 `yaml:"ref_lon"`// Reference point longitude, (degrees). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	RefAlt float32 `yaml:"ref_alt"`// Reference altitude AMSL, (metres). Position of reference point (local NED frame origin) in global (GPS / WGS84) frame
	DistBottom float32 `yaml:"dist_bottom"`// Distance from from bottom surface to ground, (metres). Distance to surface
	DistBottomValid bool `yaml:"dist_bottom_valid"`// true if distance to bottom surface is valid. Distance to surface
	DistBottomSensorBitfield uint8 `yaml:"dist_bottom_sensor_bitfield"`// bitfield indicating what type of sensor is used to estimate dist_bottom. Distance to surface
	Eph float32 `yaml:"eph"`// Standard deviation of horizontal position error, (metres)
	Epv float32 `yaml:"epv"`// Standard deviation of vertical position error, (metres)
	Evh float32 `yaml:"evh"`// Standard deviation of horizontal velocity error, (metres/sec)
	Evv float32 `yaml:"evv"`// Standard deviation of horizontal velocity error, (metres/sec)
	VxyMax float32 `yaml:"vxy_max"`// maximum horizontal speed - set to 0 when limiting not required (meters/sec). estimator specified vehicle limits
	VzMax float32 `yaml:"vz_max"`// maximum vertical speed - set to 0 when limiting not required (meters/sec). estimator specified vehicle limits
	HaglMin float32 `yaml:"hagl_min"`// minimum height above ground level - set to 0 when limiting not required (meters). estimator specified vehicle limits
	HaglMax float32 `yaml:"hagl_max"`// maximum height above ground level - set to 0 when limiting not required (meters). estimator specified vehicle limits
}

// NewVehicleLocalPosition creates a new VehicleLocalPosition with default values.
func NewVehicleLocalPosition() *VehicleLocalPosition {
	self := VehicleLocalPosition{}
	self.SetDefaults()
	return &self
}

func (t *VehicleLocalPosition) Clone() *VehicleLocalPosition {
	c := &VehicleLocalPosition{}
	c.Timestamp = t.Timestamp
	c.TimestampSample = t.TimestampSample
	c.XyValid = t.XyValid
	c.ZValid = t.ZValid
	c.VXyValid = t.VXyValid
	c.VZValid = t.VZValid
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.DeltaXy = t.DeltaXy
	c.XyResetCounter = t.XyResetCounter
	c.DeltaZ = t.DeltaZ
	c.ZResetCounter = t.ZResetCounter
	c.Vx = t.Vx
	c.Vy = t.Vy
	c.Vz = t.Vz
	c.ZDeriv = t.ZDeriv
	c.DeltaVxy = t.DeltaVxy
	c.VxyResetCounter = t.VxyResetCounter
	c.DeltaVz = t.DeltaVz
	c.VzResetCounter = t.VzResetCounter
	c.Ax = t.Ax
	c.Ay = t.Ay
	c.Az = t.Az
	c.Heading = t.Heading
	c.DeltaHeading = t.DeltaHeading
	c.HeadingResetCounter = t.HeadingResetCounter
	c.XyGlobal = t.XyGlobal
	c.ZGlobal = t.ZGlobal
	c.RefTimestamp = t.RefTimestamp
	c.RefLat = t.RefLat
	c.RefLon = t.RefLon
	c.RefAlt = t.RefAlt
	c.DistBottom = t.DistBottom
	c.DistBottomValid = t.DistBottomValid
	c.DistBottomSensorBitfield = t.DistBottomSensorBitfield
	c.Eph = t.Eph
	c.Epv = t.Epv
	c.Evh = t.Evh
	c.Evv = t.Evv
	c.VxyMax = t.VxyMax
	c.VzMax = t.VzMax
	c.HaglMin = t.HaglMin
	c.HaglMax = t.HaglMax
	return c
}

func (t *VehicleLocalPosition) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VehicleLocalPosition) SetDefaults() {
	t.Timestamp = 0
	t.TimestampSample = 0
	t.XyValid = false
	t.ZValid = false
	t.VXyValid = false
	t.VZValid = false
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.DeltaXy = [2]float32{}
	t.XyResetCounter = 0
	t.DeltaZ = 0
	t.ZResetCounter = 0
	t.Vx = 0
	t.Vy = 0
	t.Vz = 0
	t.ZDeriv = 0
	t.DeltaVxy = [2]float32{}
	t.VxyResetCounter = 0
	t.DeltaVz = 0
	t.VzResetCounter = 0
	t.Ax = 0
	t.Ay = 0
	t.Az = 0
	t.Heading = 0
	t.DeltaHeading = 0
	t.HeadingResetCounter = 0
	t.XyGlobal = false
	t.ZGlobal = false
	t.RefTimestamp = 0
	t.RefLat = 0
	t.RefLon = 0
	t.RefAlt = 0
	t.DistBottom = 0
	t.DistBottomValid = false
	t.DistBottomSensorBitfield = 0
	t.Eph = 0
	t.Epv = 0
	t.Evh = 0
	t.Evv = 0
	t.VxyMax = 0
	t.VzMax = 0
	t.HaglMin = 0
	t.HaglMax = 0
}

// CloneVehicleLocalPositionSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVehicleLocalPositionSlice(dst, src []VehicleLocalPosition) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VehicleLocalPositionTypeSupport types.MessageTypeSupport = _VehicleLocalPositionTypeSupport{}

type _VehicleLocalPositionTypeSupport struct{}

func (t _VehicleLocalPositionTypeSupport) New() types.Message {
	return NewVehicleLocalPosition()
}

func (t _VehicleLocalPositionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__VehicleLocalPosition
	return (unsafe.Pointer)(C.px4_msgs__msg__VehicleLocalPosition__create())
}

func (t _VehicleLocalPositionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__VehicleLocalPosition__destroy((*C.px4_msgs__msg__VehicleLocalPosition)(pointer_to_free))
}

func (t _VehicleLocalPositionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VehicleLocalPosition)
	mem := (*C.px4_msgs__msg__VehicleLocalPosition)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.timestamp_sample = C.uint64_t(m.TimestampSample)
	mem.xy_valid = C.bool(m.XyValid)
	mem.z_valid = C.bool(m.ZValid)
	mem.v_xy_valid = C.bool(m.VXyValid)
	mem.v_z_valid = C.bool(m.VZValid)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	cSlice_delta_xy := mem.delta_xy[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_xy)), m.DeltaXy[:])
	mem.xy_reset_counter = C.uint8_t(m.XyResetCounter)
	mem.delta_z = C.float(m.DeltaZ)
	mem.z_reset_counter = C.uint8_t(m.ZResetCounter)
	mem.vx = C.float(m.Vx)
	mem.vy = C.float(m.Vy)
	mem.vz = C.float(m.Vz)
	mem.z_deriv = C.float(m.ZDeriv)
	cSlice_delta_vxy := mem.delta_vxy[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_vxy)), m.DeltaVxy[:])
	mem.vxy_reset_counter = C.uint8_t(m.VxyResetCounter)
	mem.delta_vz = C.float(m.DeltaVz)
	mem.vz_reset_counter = C.uint8_t(m.VzResetCounter)
	mem.ax = C.float(m.Ax)
	mem.ay = C.float(m.Ay)
	mem.az = C.float(m.Az)
	mem.heading = C.float(m.Heading)
	mem.delta_heading = C.float(m.DeltaHeading)
	mem.heading_reset_counter = C.uint8_t(m.HeadingResetCounter)
	mem.xy_global = C.bool(m.XyGlobal)
	mem.z_global = C.bool(m.ZGlobal)
	mem.ref_timestamp = C.uint64_t(m.RefTimestamp)
	mem.ref_lat = C.double(m.RefLat)
	mem.ref_lon = C.double(m.RefLon)
	mem.ref_alt = C.float(m.RefAlt)
	mem.dist_bottom = C.float(m.DistBottom)
	mem.dist_bottom_valid = C.bool(m.DistBottomValid)
	mem.dist_bottom_sensor_bitfield = C.uint8_t(m.DistBottomSensorBitfield)
	mem.eph = C.float(m.Eph)
	mem.epv = C.float(m.Epv)
	mem.evh = C.float(m.Evh)
	mem.evv = C.float(m.Evv)
	mem.vxy_max = C.float(m.VxyMax)
	mem.vz_max = C.float(m.VzMax)
	mem.hagl_min = C.float(m.HaglMin)
	mem.hagl_max = C.float(m.HaglMax)
}

func (t _VehicleLocalPositionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VehicleLocalPosition)
	mem := (*C.px4_msgs__msg__VehicleLocalPosition)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.TimestampSample = uint64(mem.timestamp_sample)
	m.XyValid = bool(mem.xy_valid)
	m.ZValid = bool(mem.z_valid)
	m.VXyValid = bool(mem.v_xy_valid)
	m.VZValid = bool(mem.v_z_valid)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	cSlice_delta_xy := mem.delta_xy[:]
	primitives.Float32__Array_to_Go(m.DeltaXy[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_xy)))
	m.XyResetCounter = uint8(mem.xy_reset_counter)
	m.DeltaZ = float32(mem.delta_z)
	m.ZResetCounter = uint8(mem.z_reset_counter)
	m.Vx = float32(mem.vx)
	m.Vy = float32(mem.vy)
	m.Vz = float32(mem.vz)
	m.ZDeriv = float32(mem.z_deriv)
	cSlice_delta_vxy := mem.delta_vxy[:]
	primitives.Float32__Array_to_Go(m.DeltaVxy[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_delta_vxy)))
	m.VxyResetCounter = uint8(mem.vxy_reset_counter)
	m.DeltaVz = float32(mem.delta_vz)
	m.VzResetCounter = uint8(mem.vz_reset_counter)
	m.Ax = float32(mem.ax)
	m.Ay = float32(mem.ay)
	m.Az = float32(mem.az)
	m.Heading = float32(mem.heading)
	m.DeltaHeading = float32(mem.delta_heading)
	m.HeadingResetCounter = uint8(mem.heading_reset_counter)
	m.XyGlobal = bool(mem.xy_global)
	m.ZGlobal = bool(mem.z_global)
	m.RefTimestamp = uint64(mem.ref_timestamp)
	m.RefLat = float64(mem.ref_lat)
	m.RefLon = float64(mem.ref_lon)
	m.RefAlt = float32(mem.ref_alt)
	m.DistBottom = float32(mem.dist_bottom)
	m.DistBottomValid = bool(mem.dist_bottom_valid)
	m.DistBottomSensorBitfield = uint8(mem.dist_bottom_sensor_bitfield)
	m.Eph = float32(mem.eph)
	m.Epv = float32(mem.epv)
	m.Evh = float32(mem.evh)
	m.Evv = float32(mem.evv)
	m.VxyMax = float32(mem.vxy_max)
	m.VzMax = float32(mem.vz_max)
	m.HaglMin = float32(mem.hagl_min)
	m.HaglMax = float32(mem.hagl_max)
}

func (t _VehicleLocalPositionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__VehicleLocalPosition())
}

type CVehicleLocalPosition = C.px4_msgs__msg__VehicleLocalPosition
type CVehicleLocalPosition__Sequence = C.px4_msgs__msg__VehicleLocalPosition__Sequence

func VehicleLocalPosition__Sequence_to_Go(goSlice *[]VehicleLocalPosition, cSlice CVehicleLocalPosition__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VehicleLocalPosition, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__VehicleLocalPosition__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLocalPosition * uintptr(i)),
		))
		VehicleLocalPositionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func VehicleLocalPosition__Sequence_to_C(cSlice *CVehicleLocalPosition__Sequence, goSlice []VehicleLocalPosition) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__VehicleLocalPosition)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__VehicleLocalPosition * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__VehicleLocalPosition)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__VehicleLocalPosition * uintptr(i)),
		))
		VehicleLocalPositionTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func VehicleLocalPosition__Array_to_Go(goSlice []VehicleLocalPosition, cSlice []CVehicleLocalPosition) {
	for i := 0; i < len(cSlice); i++ {
		VehicleLocalPositionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VehicleLocalPosition__Array_to_C(cSlice []CVehicleLocalPosition, goSlice []VehicleLocalPosition) {
	for i := 0; i < len(goSlice); i++ {
		VehicleLocalPositionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
