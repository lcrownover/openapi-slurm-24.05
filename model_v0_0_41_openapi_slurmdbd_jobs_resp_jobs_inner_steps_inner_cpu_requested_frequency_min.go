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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin Minimum requested CPU frequency in kHz
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int32 `json:"number,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMinWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMinWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) SetNumber(v int32) {
	o.Number = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.Infinite) {
		toSerialize["infinite"] = o.Infinite
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequencyMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


