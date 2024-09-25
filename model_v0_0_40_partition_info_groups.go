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

// checks if the V0040PartitionInfoGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040PartitionInfoGroups{}

// V0040PartitionInfoGroups struct for V0040PartitionInfoGroups
type V0040PartitionInfoGroups struct {
	// AllowGroups
	Allowed *string `json:"allowed,omitempty"`
}

// NewV0040PartitionInfoGroups instantiates a new V0040PartitionInfoGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040PartitionInfoGroups() *V0040PartitionInfoGroups {
	this := V0040PartitionInfoGroups{}
	return &this
}

// NewV0040PartitionInfoGroupsWithDefaults instantiates a new V0040PartitionInfoGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040PartitionInfoGroupsWithDefaults() *V0040PartitionInfoGroups {
	this := V0040PartitionInfoGroups{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *V0040PartitionInfoGroups) GetAllowed() string {
	if o == nil || IsNil(o.Allowed) {
		var ret string
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040PartitionInfoGroups) GetAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *V0040PartitionInfoGroups) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given string and assigns it to the Allowed field.
func (o *V0040PartitionInfoGroups) SetAllowed(v string) {
	o.Allowed = &v
}

func (o V0040PartitionInfoGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040PartitionInfoGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableV0040PartitionInfoGroups struct {
	value *V0040PartitionInfoGroups
	isSet bool
}

func (v NullableV0040PartitionInfoGroups) Get() *V0040PartitionInfoGroups {
	return v.value
}

func (v *NullableV0040PartitionInfoGroups) Set(val *V0040PartitionInfoGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040PartitionInfoGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040PartitionInfoGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040PartitionInfoGroups(val *V0040PartitionInfoGroups) *NullableV0040PartitionInfoGroups {
	return &NullableV0040PartitionInfoGroups{value: val, isSet: true}
}

func (v NullableV0040PartitionInfoGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040PartitionInfoGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


