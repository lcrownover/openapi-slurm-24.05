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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU struct for V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU struct {
	RequestedFrequency *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency `json:"requested_frequency,omitempty"`
	// Requested CPU frequency governor in kHz
	Governor *string `json:"governor,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPUWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPUWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU{}
	return &this
}

// GetRequestedFrequency returns the RequestedFrequency field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) GetRequestedFrequency() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency {
	if o == nil || IsNil(o.RequestedFrequency) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency
		return ret
	}
	return *o.RequestedFrequency
}

// GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) GetRequestedFrequencyOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency, bool) {
	if o == nil || IsNil(o.RequestedFrequency) {
		return nil, false
	}
	return o.RequestedFrequency, true
}

// HasRequestedFrequency returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) HasRequestedFrequency() bool {
	if o != nil && !IsNil(o.RequestedFrequency) {
		return true
	}

	return false
}

// SetRequestedFrequency gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency and assigns it to the RequestedFrequency field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) SetRequestedFrequency(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPURequestedFrequency) {
	o.RequestedFrequency = &v
}

// GetGovernor returns the Governor field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) GetGovernor() string {
	if o == nil || IsNil(o.Governor) {
		var ret string
		return ret
	}
	return *o.Governor
}

// GetGovernorOk returns a tuple with the Governor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) GetGovernorOk() (*string, bool) {
	if o == nil || IsNil(o.Governor) {
		return nil, false
	}
	return o.Governor, true
}

// HasGovernor returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) HasGovernor() bool {
	if o != nil && !IsNil(o.Governor) {
		return true
	}

	return false
}

// SetGovernor gets a reference to the given string and assigns it to the Governor field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) SetGovernor(v string) {
	o.Governor = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestedFrequency) {
		toSerialize["requested_frequency"] = o.RequestedFrequency
	}
	if !IsNil(o.Governor) {
		toSerialize["governor"] = o.Governor
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


