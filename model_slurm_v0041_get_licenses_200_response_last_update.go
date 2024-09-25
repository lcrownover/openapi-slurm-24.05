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

// checks if the SlurmV0041GetLicenses200ResponseLastUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041GetLicenses200ResponseLastUpdate{}

// SlurmV0041GetLicenses200ResponseLastUpdate Time of last licenses change (UNIX timestamp)
type SlurmV0041GetLicenses200ResponseLastUpdate struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int64 `json:"number,omitempty"`
}

// NewSlurmV0041GetLicenses200ResponseLastUpdate instantiates a new SlurmV0041GetLicenses200ResponseLastUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041GetLicenses200ResponseLastUpdate() *SlurmV0041GetLicenses200ResponseLastUpdate {
	this := SlurmV0041GetLicenses200ResponseLastUpdate{}
	return &this
}

// NewSlurmV0041GetLicenses200ResponseLastUpdateWithDefaults instantiates a new SlurmV0041GetLicenses200ResponseLastUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041GetLicenses200ResponseLastUpdateWithDefaults() *SlurmV0041GetLicenses200ResponseLastUpdate {
	this := SlurmV0041GetLicenses200ResponseLastUpdate{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetNumber() int64 {
	if o == nil || IsNil(o.Number) {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) GetNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *SlurmV0041GetLicenses200ResponseLastUpdate) SetNumber(v int64) {
	o.Number = &v
}

func (o SlurmV0041GetLicenses200ResponseLastUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041GetLicenses200ResponseLastUpdate) ToMap() (map[string]interface{}, error) {
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

type NullableSlurmV0041GetLicenses200ResponseLastUpdate struct {
	value *SlurmV0041GetLicenses200ResponseLastUpdate
	isSet bool
}

func (v NullableSlurmV0041GetLicenses200ResponseLastUpdate) Get() *SlurmV0041GetLicenses200ResponseLastUpdate {
	return v.value
}

func (v *NullableSlurmV0041GetLicenses200ResponseLastUpdate) Set(val *SlurmV0041GetLicenses200ResponseLastUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041GetLicenses200ResponseLastUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041GetLicenses200ResponseLastUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041GetLicenses200ResponseLastUpdate(val *SlurmV0041GetLicenses200ResponseLastUpdate) *NullableSlurmV0041GetLicenses200ResponseLastUpdate {
	return &NullableSlurmV0041GetLicenses200ResponseLastUpdate{value: val, isSet: true}
}

func (v NullableSlurmV0041GetLicenses200ResponseLastUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041GetLicenses200ResponseLastUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


