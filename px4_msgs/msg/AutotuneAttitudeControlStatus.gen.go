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

#include <px4_msgs/msg/autotune_attitude_control_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/AutotuneAttitudeControlStatus", AutotuneAttitudeControlStatusTypeSupport)
}
const (
	AutotuneAttitudeControlStatus_STATE_IDLE uint8 = 0
	AutotuneAttitudeControlStatus_STATE_INIT uint8 = 1
	AutotuneAttitudeControlStatus_STATE_ROLL uint8 = 2
	AutotuneAttitudeControlStatus_STATE_ROLL_PAUSE uint8 = 3
	AutotuneAttitudeControlStatus_STATE_PITCH uint8 = 4
	AutotuneAttitudeControlStatus_STATE_PITCH_PAUSE uint8 = 5
	AutotuneAttitudeControlStatus_STATE_YAW uint8 = 6
	AutotuneAttitudeControlStatus_STATE_YAW_PAUSE uint8 = 7
	AutotuneAttitudeControlStatus_STATE_VERIFICATION uint8 = 8
	AutotuneAttitudeControlStatus_STATE_APPLY uint8 = 9
	AutotuneAttitudeControlStatus_STATE_TEST uint8 = 10
	AutotuneAttitudeControlStatus_STATE_COMPLETE uint8 = 11
	AutotuneAttitudeControlStatus_STATE_FAIL uint8 = 12
	AutotuneAttitudeControlStatus_STATE_WAIT_FOR_DISARM uint8 = 13
)

// Do not create instances of this type directly. Always use NewAutotuneAttitudeControlStatus
// function instead.
type AutotuneAttitudeControlStatus struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	Coeff [5]float32 `yaml:"coeff"`// coefficients of the identified discrete-time model
	CoeffVar [5]float32 `yaml:"coeff_var"`// coefficients' variance of the identified discrete-time model
	Fitness float32 `yaml:"fitness"`// fitness of the parameter estimate
	Innov float32 `yaml:"innov"`
	DtModel float32 `yaml:"dt_model"`
	Kc float32 `yaml:"kc"`
	Ki float32 `yaml:"ki"`
	Kd float32 `yaml:"kd"`
	Kff float32 `yaml:"kff"`
	AttP float32 `yaml:"att_p"`
	RateSp [3]float32 `yaml:"rate_sp"`
	UFilt float32 `yaml:"u_filt"`
	YFilt float32 `yaml:"y_filt"`
	State uint8 `yaml:"state"`
}

// NewAutotuneAttitudeControlStatus creates a new AutotuneAttitudeControlStatus with default values.
func NewAutotuneAttitudeControlStatus() *AutotuneAttitudeControlStatus {
	self := AutotuneAttitudeControlStatus{}
	self.SetDefaults()
	return &self
}

func (t *AutotuneAttitudeControlStatus) Clone() *AutotuneAttitudeControlStatus {
	c := &AutotuneAttitudeControlStatus{}
	c.Timestamp = t.Timestamp
	c.Coeff = t.Coeff
	c.CoeffVar = t.CoeffVar
	c.Fitness = t.Fitness
	c.Innov = t.Innov
	c.DtModel = t.DtModel
	c.Kc = t.Kc
	c.Ki = t.Ki
	c.Kd = t.Kd
	c.Kff = t.Kff
	c.AttP = t.AttP
	c.RateSp = t.RateSp
	c.UFilt = t.UFilt
	c.YFilt = t.YFilt
	c.State = t.State
	return c
}

func (t *AutotuneAttitudeControlStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AutotuneAttitudeControlStatus) SetDefaults() {
	t.Timestamp = 0
	t.Coeff = [5]float32{}
	t.CoeffVar = [5]float32{}
	t.Fitness = 0
	t.Innov = 0
	t.DtModel = 0
	t.Kc = 0
	t.Ki = 0
	t.Kd = 0
	t.Kff = 0
	t.AttP = 0
	t.RateSp = [3]float32{}
	t.UFilt = 0
	t.YFilt = 0
	t.State = 0
}

// CloneAutotuneAttitudeControlStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAutotuneAttitudeControlStatusSlice(dst, src []AutotuneAttitudeControlStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AutotuneAttitudeControlStatusTypeSupport types.MessageTypeSupport = _AutotuneAttitudeControlStatusTypeSupport{}

type _AutotuneAttitudeControlStatusTypeSupport struct{}

func (t _AutotuneAttitudeControlStatusTypeSupport) New() types.Message {
	return NewAutotuneAttitudeControlStatus()
}

func (t _AutotuneAttitudeControlStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__AutotuneAttitudeControlStatus
	return (unsafe.Pointer)(C.px4_msgs__msg__AutotuneAttitudeControlStatus__create())
}

func (t _AutotuneAttitudeControlStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__AutotuneAttitudeControlStatus__destroy((*C.px4_msgs__msg__AutotuneAttitudeControlStatus)(pointer_to_free))
}

func (t _AutotuneAttitudeControlStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AutotuneAttitudeControlStatus)
	mem := (*C.px4_msgs__msg__AutotuneAttitudeControlStatus)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_coeff := mem.coeff[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_coeff)), m.Coeff[:])
	cSlice_coeff_var := mem.coeff_var[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_coeff_var)), m.CoeffVar[:])
	mem.fitness = C.float(m.Fitness)
	mem.innov = C.float(m.Innov)
	mem.dt_model = C.float(m.DtModel)
	mem.kc = C.float(m.Kc)
	mem.ki = C.float(m.Ki)
	mem.kd = C.float(m.Kd)
	mem.kff = C.float(m.Kff)
	mem.att_p = C.float(m.AttP)
	cSlice_rate_sp := mem.rate_sp[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_rate_sp)), m.RateSp[:])
	mem.u_filt = C.float(m.UFilt)
	mem.y_filt = C.float(m.YFilt)
	mem.state = C.uint8_t(m.State)
}

func (t _AutotuneAttitudeControlStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AutotuneAttitudeControlStatus)
	mem := (*C.px4_msgs__msg__AutotuneAttitudeControlStatus)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_coeff := mem.coeff[:]
	primitives.Float32__Array_to_Go(m.Coeff[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_coeff)))
	cSlice_coeff_var := mem.coeff_var[:]
	primitives.Float32__Array_to_Go(m.CoeffVar[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_coeff_var)))
	m.Fitness = float32(mem.fitness)
	m.Innov = float32(mem.innov)
	m.DtModel = float32(mem.dt_model)
	m.Kc = float32(mem.kc)
	m.Ki = float32(mem.ki)
	m.Kd = float32(mem.kd)
	m.Kff = float32(mem.kff)
	m.AttP = float32(mem.att_p)
	cSlice_rate_sp := mem.rate_sp[:]
	primitives.Float32__Array_to_Go(m.RateSp[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_rate_sp)))
	m.UFilt = float32(mem.u_filt)
	m.YFilt = float32(mem.y_filt)
	m.State = uint8(mem.state)
}

func (t _AutotuneAttitudeControlStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__AutotuneAttitudeControlStatus())
}

type CAutotuneAttitudeControlStatus = C.px4_msgs__msg__AutotuneAttitudeControlStatus
type CAutotuneAttitudeControlStatus__Sequence = C.px4_msgs__msg__AutotuneAttitudeControlStatus__Sequence

func AutotuneAttitudeControlStatus__Sequence_to_Go(goSlice *[]AutotuneAttitudeControlStatus, cSlice CAutotuneAttitudeControlStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AutotuneAttitudeControlStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__AutotuneAttitudeControlStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__AutotuneAttitudeControlStatus * uintptr(i)),
		))
		AutotuneAttitudeControlStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func AutotuneAttitudeControlStatus__Sequence_to_C(cSlice *CAutotuneAttitudeControlStatus__Sequence, goSlice []AutotuneAttitudeControlStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__AutotuneAttitudeControlStatus)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__AutotuneAttitudeControlStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__AutotuneAttitudeControlStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__AutotuneAttitudeControlStatus * uintptr(i)),
		))
		AutotuneAttitudeControlStatusTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func AutotuneAttitudeControlStatus__Array_to_Go(goSlice []AutotuneAttitudeControlStatus, cSlice []CAutotuneAttitudeControlStatus) {
	for i := 0; i < len(cSlice); i++ {
		AutotuneAttitudeControlStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AutotuneAttitudeControlStatus__Array_to_C(cSlice []CAutotuneAttitudeControlStatus, goSlice []AutotuneAttitudeControlStatus) {
	for i := 0; i < len(goSlice); i++ {
		AutotuneAttitudeControlStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
