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

// checks if the V0040StepStatisticsEnergy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepStatisticsEnergy{}

// V0040StepStatisticsEnergy struct for V0040StepStatisticsEnergy
type V0040StepStatisticsEnergy struct {
	Consumed *V0040Uint64NoVal `json:"consumed,omitempty"`
}

// NewV0040StepStatisticsEnergy instantiates a new V0040StepStatisticsEnergy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepStatisticsEnergy() *V0040StepStatisticsEnergy {
	this := V0040StepStatisticsEnergy{}
	return &this
}

// NewV0040StepStatisticsEnergyWithDefaults instantiates a new V0040StepStatisticsEnergy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepStatisticsEnergyWithDefaults() *V0040StepStatisticsEnergy {
	this := V0040StepStatisticsEnergy{}
	return &this
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *V0040StepStatisticsEnergy) GetConsumed() V0040Uint64NoVal {
	if o == nil || IsNil(o.Consumed) {
		var ret V0040Uint64NoVal
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepStatisticsEnergy) GetConsumedOk() (*V0040Uint64NoVal, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *V0040StepStatisticsEnergy) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given V0040Uint64NoVal and assigns it to the Consumed field.
func (o *V0040StepStatisticsEnergy) SetConsumed(v V0040Uint64NoVal) {
	o.Consumed = &v
}

func (o V0040StepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepStatisticsEnergy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	return toSerialize, nil
}

type NullableV0040StepStatisticsEnergy struct {
	value *V0040StepStatisticsEnergy
	isSet bool
}

func (v NullableV0040StepStatisticsEnergy) Get() *V0040StepStatisticsEnergy {
	return v.value
}

func (v *NullableV0040StepStatisticsEnergy) Set(val *V0040StepStatisticsEnergy) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepStatisticsEnergy) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepStatisticsEnergy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepStatisticsEnergy(val *V0040StepStatisticsEnergy) *NullableV0040StepStatisticsEnergy {
	return &NullableV0040StepStatisticsEnergy{value: val, isSet: true}
}

func (v NullableV0040StepStatisticsEnergy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepStatisticsEnergy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


