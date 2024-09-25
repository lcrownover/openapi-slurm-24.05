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

// checks if the SlurmdbV0041DeleteSingleQos200ResponseWarningsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041DeleteSingleQos200ResponseWarningsInner{}

// SlurmdbV0041DeleteSingleQos200ResponseWarningsInner struct for SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
type SlurmdbV0041DeleteSingleQos200ResponseWarningsInner struct {
	// Long form warning description
	Description *string `json:"description,omitempty"`
	// Source of warning or where warning was first detected
	Source *string `json:"source,omitempty"`
}

// NewSlurmdbV0041DeleteSingleQos200ResponseWarningsInner instantiates a new SlurmdbV0041DeleteSingleQos200ResponseWarningsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041DeleteSingleQos200ResponseWarningsInner() *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	this := SlurmdbV0041DeleteSingleQos200ResponseWarningsInner{}
	return &this
}

// NewSlurmdbV0041DeleteSingleQos200ResponseWarningsInnerWithDefaults instantiates a new SlurmdbV0041DeleteSingleQos200ResponseWarningsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041DeleteSingleQos200ResponseWarningsInnerWithDefaults() *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	this := SlurmdbV0041DeleteSingleQos200ResponseWarningsInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) SetDescription(v string) {
	o.Description = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) SetSource(v string) {
	o.Source = &v
}

func (o SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner struct {
	value *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
	isSet bool
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) Get() *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	return v.value
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) Set(val *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner(val *SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) *NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	return &NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner{value: val, isSet: true}
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseWarningsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


