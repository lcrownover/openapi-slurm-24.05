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

// checks if the SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion{}

// SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion struct for SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion
type SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion struct {
	// Slurm release major version
	Major *string `json:"major,omitempty"`
	// Slurm release micro version
	Micro *string `json:"micro,omitempty"`
	// Slurm release minor version
	Minor *string `json:"minor,omitempty"`
}

// NewSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion instantiates a new SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion() *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion {
	this := SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion{}
	return &this
}

// NewSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersionWithDefaults instantiates a new SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersionWithDefaults() *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion {
	this := SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMajor() string {
	if o == nil || IsNil(o.Major) {
		var ret string
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMajorOk() (*string, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given string and assigns it to the Major field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) SetMajor(v string) {
	o.Major = &v
}

// GetMicro returns the Micro field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMicro() string {
	if o == nil || IsNil(o.Micro) {
		var ret string
		return ret
	}
	return *o.Micro
}

// GetMicroOk returns a tuple with the Micro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMicroOk() (*string, bool) {
	if o == nil || IsNil(o.Micro) {
		return nil, false
	}
	return o.Micro, true
}

// HasMicro returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) HasMicro() bool {
	if o != nil && !IsNil(o.Micro) {
		return true
	}

	return false
}

// SetMicro gets a reference to the given string and assigns it to the Micro field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) SetMicro(v string) {
	o.Micro = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMinor() string {
	if o == nil || IsNil(o.Minor) {
		var ret string
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) GetMinorOk() (*string, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given string and assigns it to the Minor field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) SetMinor(v string) {
	o.Minor = &v
}

func (o SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Micro) {
		toSerialize["micro"] = o.Micro
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion struct {
	value *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion
	isSet bool
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) Get() *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion {
	return v.value
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) Set(val *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion(val *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) *NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion {
	return &NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion{value: val, isSet: true}
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMetaSlurmVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


