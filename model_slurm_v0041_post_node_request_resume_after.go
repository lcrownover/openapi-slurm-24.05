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

// checks if the SlurmV0041PostNodeRequestResumeAfter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041PostNodeRequestResumeAfter{}

// SlurmV0041PostNodeRequestResumeAfter Number of seconds after which to automatically resume DOWN or DRAINED node
type SlurmV0041PostNodeRequestResumeAfter struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int32 `json:"number,omitempty"`
}

// NewSlurmV0041PostNodeRequestResumeAfter instantiates a new SlurmV0041PostNodeRequestResumeAfter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041PostNodeRequestResumeAfter() *SlurmV0041PostNodeRequestResumeAfter {
	this := SlurmV0041PostNodeRequestResumeAfter{}
	return &this
}

// NewSlurmV0041PostNodeRequestResumeAfterWithDefaults instantiates a new SlurmV0041PostNodeRequestResumeAfter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041PostNodeRequestResumeAfterWithDefaults() *SlurmV0041PostNodeRequestResumeAfter {
	this := SlurmV0041PostNodeRequestResumeAfter{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *SlurmV0041PostNodeRequestResumeAfter) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *SlurmV0041PostNodeRequestResumeAfter) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SlurmV0041PostNodeRequestResumeAfter) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *SlurmV0041PostNodeRequestResumeAfter) SetNumber(v int32) {
	o.Number = &v
}

func (o SlurmV0041PostNodeRequestResumeAfter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041PostNodeRequestResumeAfter) ToMap() (map[string]interface{}, error) {
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

type NullableSlurmV0041PostNodeRequestResumeAfter struct {
	value *SlurmV0041PostNodeRequestResumeAfter
	isSet bool
}

func (v NullableSlurmV0041PostNodeRequestResumeAfter) Get() *SlurmV0041PostNodeRequestResumeAfter {
	return v.value
}

func (v *NullableSlurmV0041PostNodeRequestResumeAfter) Set(val *SlurmV0041PostNodeRequestResumeAfter) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041PostNodeRequestResumeAfter) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041PostNodeRequestResumeAfter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041PostNodeRequestResumeAfter(val *SlurmV0041PostNodeRequestResumeAfter) *NullableSlurmV0041PostNodeRequestResumeAfter {
	return &NullableSlurmV0041PostNodeRequestResumeAfter{value: val, isSet: true}
}

func (v NullableSlurmV0041PostNodeRequestResumeAfter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041PostNodeRequestResumeAfter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


