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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs struct {
	ActiveJobs *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs `json:"active_jobs,omitempty"`
	Per *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer `json:"per,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs{}
	return &this
}

// GetActiveJobs returns the ActiveJobs field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) GetActiveJobs() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs {
	if o == nil || IsNil(o.ActiveJobs) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs
		return ret
	}
	return *o.ActiveJobs
}

// GetActiveJobsOk returns a tuple with the ActiveJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) GetActiveJobsOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs, bool) {
	if o == nil || IsNil(o.ActiveJobs) {
		return nil, false
	}
	return o.ActiveJobs, true
}

// HasActiveJobs returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) HasActiveJobs() bool {
	if o != nil && !IsNil(o.ActiveJobs) {
		return true
	}

	return false
}

// SetActiveJobs gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs and assigns it to the ActiveJobs field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) SetActiveJobs(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsActiveJobs) {
	o.ActiveJobs = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) GetPer() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) GetPerOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer and assigns it to the Per field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) SetPer(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobsPer) {
	o.Per = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveJobs) {
		toSerialize["active_jobs"] = o.ActiveJobs
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


