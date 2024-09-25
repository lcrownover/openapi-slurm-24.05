/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_slurm_24_05

import (
	"encoding/json"
)

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal struct for V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal
type V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal struct {
	// Sum of System and User CPU time used by the job in seconds
	Seconds *int64 `json:"seconds,omitempty"`
	// Sum of System and User CPU time used by the job in microseconds
	Microseconds *int64 `json:"microseconds,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotalWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotalWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) GetMicroseconds() int64 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int64
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) GetMicrosecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int64 and assigns it to the Microseconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) SetMicroseconds(v int64) {
	o.Microseconds = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


