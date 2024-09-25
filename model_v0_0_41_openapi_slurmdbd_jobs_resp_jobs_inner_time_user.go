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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser struct for V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser
type V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser struct {
	// User CPU time used by the job in seconds
	Seconds *int64 `json:"seconds,omitempty"`
	// User CPU time used by the job in microseconds
	Microseconds *int64 `json:"microseconds,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUserWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUserWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) GetMicroseconds() int64 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int64
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) GetMicrosecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int64 and assigns it to the Microseconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) SetMicroseconds(v int64) {
	o.Microseconds = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


