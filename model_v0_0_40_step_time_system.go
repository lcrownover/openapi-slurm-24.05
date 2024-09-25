/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0040StepTimeSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepTimeSystem{}

// V0040StepTimeSystem struct for V0040StepTimeSystem
type V0040StepTimeSystem struct {
	// System CPU time used by the step in seconds
	Seconds *int64 `json:"seconds,omitempty"`
	// System CPU time used by the step in microseconds
	Microseconds *int32 `json:"microseconds,omitempty"`
}

// NewV0040StepTimeSystem instantiates a new V0040StepTimeSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepTimeSystem() *V0040StepTimeSystem {
	this := V0040StepTimeSystem{}
	return &this
}

// NewV0040StepTimeSystemWithDefaults instantiates a new V0040StepTimeSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepTimeSystemWithDefaults() *V0040StepTimeSystem {
	this := V0040StepTimeSystem{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0040StepTimeSystem) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepTimeSystem) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0040StepTimeSystem) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0040StepTimeSystem) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0040StepTimeSystem) GetMicroseconds() int32 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int32
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepTimeSystem) GetMicrosecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0040StepTimeSystem) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int32 and assigns it to the Microseconds field.
func (o *V0040StepTimeSystem) SetMicroseconds(v int32) {
	o.Microseconds = &v
}

func (o V0040StepTimeSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepTimeSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0040StepTimeSystem struct {
	value *V0040StepTimeSystem
	isSet bool
}

func (v NullableV0040StepTimeSystem) Get() *V0040StepTimeSystem {
	return v.value
}

func (v *NullableV0040StepTimeSystem) Set(val *V0040StepTimeSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepTimeSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepTimeSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepTimeSystem(val *V0040StepTimeSystem) *NullableV0040StepTimeSystem {
	return &NullableV0040StepTimeSystem{value: val, isSet: true}
}

func (v NullableV0040StepTimeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepTimeSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


