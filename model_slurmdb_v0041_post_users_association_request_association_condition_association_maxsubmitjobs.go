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

// checks if the SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs{}

// SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs Maximum number of jobs which can be in a pending or running state at any time in this association
type SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int32 `json:"number,omitempty"`
}

// NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs {
	this := SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs{}
	return &this
}

// NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobsWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobsWithDefaults() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs {
	this := SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) SetNumber(v int32) {
	o.Number = &v
}

func (o SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) ToMap() (map[string]interface{}, error) {
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

type NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs struct {
	value *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs
	isSet bool
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) Get() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs {
	return v.value
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) Set(val *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs(val *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) *NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs {
	return &NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs{value: val, isSet: true}
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


