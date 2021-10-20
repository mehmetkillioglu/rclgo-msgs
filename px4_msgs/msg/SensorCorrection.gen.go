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
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/sensor_correction.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorCorrection", SensorCorrectionTypeSupport)
}

// Do not create instances of this type directly. Always use NewSensorCorrection
// function instead.
type SensorCorrection struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	GyroDeviceIds [4]uint32 `yaml:"gyro_device_ids"`// Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroTemperature [4]float32 `yaml:"gyro_temperature"`// Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset0 [3]float32 `yaml:"gyro_offset_0"`// gyro 0 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset1 [3]float32 `yaml:"gyro_offset_1"`// gyro 1 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset2 [3]float32 `yaml:"gyro_offset_2"`// gyro 2 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	GyroOffset3 [3]float32 `yaml:"gyro_offset_3"`// gyro 3 XYZ offsets in the sensor frame in rad/s. Corrections for gyro angular rate outputs where corrected_rate = raw_rate * gyro_scale + gyro_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelDeviceIds [4]uint32 `yaml:"accel_device_ids"`// Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelTemperature [4]float32 `yaml:"accel_temperature"`// Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset0 [3]float32 `yaml:"accel_offset_0"`// accelerometer 0 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset1 [3]float32 `yaml:"accel_offset_1"`// accelerometer 1 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset2 [3]float32 `yaml:"accel_offset_2"`// accelerometer 2 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	AccelOffset3 [3]float32 `yaml:"accel_offset_3"`// accelerometer 3 offsets in the FRD board frame XYZ-axis in m/s^s. Corrections for acceleromter acceleration outputs where corrected_accel = raw_accel * accel_scale + accel_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroDeviceIds [4]uint32 `yaml:"baro_device_ids"`// Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroTemperature [4]float32 `yaml:"baro_temperature"`// Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset0 float32 `yaml:"baro_offset_0"`// barometric pressure 0 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset1 float32 `yaml:"baro_offset_1"`// barometric pressure 1 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset2 float32 `yaml:"baro_offset_2"`// barometric pressure 2 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
	BaroOffset3 float32 `yaml:"baro_offset_3"`// barometric pressure 3 offsets in the sensor frame in Pascals. Corrections for barometric pressure outputs where corrected_pressure = raw_pressure * pressure_scale + pressure_offsetNote the corrections are in the sensor frame and must be applied before the sensor data is rotated into body frame
}

// NewSensorCorrection creates a new SensorCorrection with default values.
func NewSensorCorrection() *SensorCorrection {
	self := SensorCorrection{}
	self.SetDefaults()
	return &self
}

func (t *SensorCorrection) Clone() *SensorCorrection {
	c := &SensorCorrection{}
	c.Timestamp = t.Timestamp
	c.GyroDeviceIds = t.GyroDeviceIds
	c.GyroTemperature = t.GyroTemperature
	c.GyroOffset0 = t.GyroOffset0
	c.GyroOffset1 = t.GyroOffset1
	c.GyroOffset2 = t.GyroOffset2
	c.GyroOffset3 = t.GyroOffset3
	c.AccelDeviceIds = t.AccelDeviceIds
	c.AccelTemperature = t.AccelTemperature
	c.AccelOffset0 = t.AccelOffset0
	c.AccelOffset1 = t.AccelOffset1
	c.AccelOffset2 = t.AccelOffset2
	c.AccelOffset3 = t.AccelOffset3
	c.BaroDeviceIds = t.BaroDeviceIds
	c.BaroTemperature = t.BaroTemperature
	c.BaroOffset0 = t.BaroOffset0
	c.BaroOffset1 = t.BaroOffset1
	c.BaroOffset2 = t.BaroOffset2
	c.BaroOffset3 = t.BaroOffset3
	return c
}

func (t *SensorCorrection) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SensorCorrection) SetDefaults() {
	t.Timestamp = 0
	t.GyroDeviceIds = [4]uint32{}
	t.GyroTemperature = [4]float32{}
	t.GyroOffset0 = [3]float32{}
	t.GyroOffset1 = [3]float32{}
	t.GyroOffset2 = [3]float32{}
	t.GyroOffset3 = [3]float32{}
	t.AccelDeviceIds = [4]uint32{}
	t.AccelTemperature = [4]float32{}
	t.AccelOffset0 = [3]float32{}
	t.AccelOffset1 = [3]float32{}
	t.AccelOffset2 = [3]float32{}
	t.AccelOffset3 = [3]float32{}
	t.BaroDeviceIds = [4]uint32{}
	t.BaroTemperature = [4]float32{}
	t.BaroOffset0 = 0
	t.BaroOffset1 = 0
	t.BaroOffset2 = 0
	t.BaroOffset3 = 0
}

// CloneSensorCorrectionSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSensorCorrectionSlice(dst, src []SensorCorrection) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SensorCorrectionTypeSupport types.MessageTypeSupport = _SensorCorrectionTypeSupport{}

type _SensorCorrectionTypeSupport struct{}

func (t _SensorCorrectionTypeSupport) New() types.Message {
	return NewSensorCorrection()
}

func (t _SensorCorrectionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorCorrection
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorCorrection__create())
}

func (t _SensorCorrectionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorCorrection__destroy((*C.px4_msgs__msg__SensorCorrection)(pointer_to_free))
}

func (t _SensorCorrectionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorCorrection)
	mem := (*C.px4_msgs__msg__SensorCorrection)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)), m.GyroDeviceIds[:])
	cSlice_gyro_temperature := mem.gyro_temperature[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_temperature)), m.GyroTemperature[:])
	cSlice_gyro_offset_0 := mem.gyro_offset_0[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_0)), m.GyroOffset0[:])
	cSlice_gyro_offset_1 := mem.gyro_offset_1[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_1)), m.GyroOffset1[:])
	cSlice_gyro_offset_2 := mem.gyro_offset_2[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_2)), m.GyroOffset2[:])
	cSlice_gyro_offset_3 := mem.gyro_offset_3[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_3)), m.GyroOffset3[:])
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)), m.AccelDeviceIds[:])
	cSlice_accel_temperature := mem.accel_temperature[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_temperature)), m.AccelTemperature[:])
	cSlice_accel_offset_0 := mem.accel_offset_0[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_0)), m.AccelOffset0[:])
	cSlice_accel_offset_1 := mem.accel_offset_1[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_1)), m.AccelOffset1[:])
	cSlice_accel_offset_2 := mem.accel_offset_2[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_2)), m.AccelOffset2[:])
	cSlice_accel_offset_3 := mem.accel_offset_3[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_3)), m.AccelOffset3[:])
	cSlice_baro_device_ids := mem.baro_device_ids[:]
	primitives.Uint32__Array_to_C(*(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_baro_device_ids)), m.BaroDeviceIds[:])
	cSlice_baro_temperature := mem.baro_temperature[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_baro_temperature)), m.BaroTemperature[:])
	mem.baro_offset_0 = C.float(m.BaroOffset0)
	mem.baro_offset_1 = C.float(m.BaroOffset1)
	mem.baro_offset_2 = C.float(m.BaroOffset2)
	mem.baro_offset_3 = C.float(m.BaroOffset3)
}

func (t _SensorCorrectionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorCorrection)
	mem := (*C.px4_msgs__msg__SensorCorrection)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_gyro_device_ids := mem.gyro_device_ids[:]
	primitives.Uint32__Array_to_Go(m.GyroDeviceIds[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_gyro_device_ids)))
	cSlice_gyro_temperature := mem.gyro_temperature[:]
	primitives.Float32__Array_to_Go(m.GyroTemperature[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_temperature)))
	cSlice_gyro_offset_0 := mem.gyro_offset_0[:]
	primitives.Float32__Array_to_Go(m.GyroOffset0[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_0)))
	cSlice_gyro_offset_1 := mem.gyro_offset_1[:]
	primitives.Float32__Array_to_Go(m.GyroOffset1[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_1)))
	cSlice_gyro_offset_2 := mem.gyro_offset_2[:]
	primitives.Float32__Array_to_Go(m.GyroOffset2[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_2)))
	cSlice_gyro_offset_3 := mem.gyro_offset_3[:]
	primitives.Float32__Array_to_Go(m.GyroOffset3[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_offset_3)))
	cSlice_accel_device_ids := mem.accel_device_ids[:]
	primitives.Uint32__Array_to_Go(m.AccelDeviceIds[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_accel_device_ids)))
	cSlice_accel_temperature := mem.accel_temperature[:]
	primitives.Float32__Array_to_Go(m.AccelTemperature[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_temperature)))
	cSlice_accel_offset_0 := mem.accel_offset_0[:]
	primitives.Float32__Array_to_Go(m.AccelOffset0[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_0)))
	cSlice_accel_offset_1 := mem.accel_offset_1[:]
	primitives.Float32__Array_to_Go(m.AccelOffset1[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_1)))
	cSlice_accel_offset_2 := mem.accel_offset_2[:]
	primitives.Float32__Array_to_Go(m.AccelOffset2[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_2)))
	cSlice_accel_offset_3 := mem.accel_offset_3[:]
	primitives.Float32__Array_to_Go(m.AccelOffset3[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accel_offset_3)))
	cSlice_baro_device_ids := mem.baro_device_ids[:]
	primitives.Uint32__Array_to_Go(m.BaroDeviceIds[:], *(*[]primitives.CUint32)(unsafe.Pointer(&cSlice_baro_device_ids)))
	cSlice_baro_temperature := mem.baro_temperature[:]
	primitives.Float32__Array_to_Go(m.BaroTemperature[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_baro_temperature)))
	m.BaroOffset0 = float32(mem.baro_offset_0)
	m.BaroOffset1 = float32(mem.baro_offset_1)
	m.BaroOffset2 = float32(mem.baro_offset_2)
	m.BaroOffset3 = float32(mem.baro_offset_3)
}

func (t _SensorCorrectionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorCorrection())
}

type CSensorCorrection = C.px4_msgs__msg__SensorCorrection
type CSensorCorrection__Sequence = C.px4_msgs__msg__SensorCorrection__Sequence

func SensorCorrection__Sequence_to_Go(goSlice *[]SensorCorrection, cSlice CSensorCorrection__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorCorrection, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorCorrection__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(i)),
		))
		SensorCorrectionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorCorrection__Sequence_to_C(cSlice *CSensorCorrection__Sequence, goSlice []SensorCorrection) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorCorrection)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorCorrection)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCorrection * uintptr(i)),
		))
		SensorCorrectionTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorCorrection__Array_to_Go(goSlice []SensorCorrection, cSlice []CSensorCorrection) {
	for i := 0; i < len(cSlice); i++ {
		SensorCorrectionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorCorrection__Array_to_C(cSlice []CSensorCorrection, goSlice []SensorCorrection) {
	for i := 0; i < len(goSlice); i++ {
		SensorCorrectionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
