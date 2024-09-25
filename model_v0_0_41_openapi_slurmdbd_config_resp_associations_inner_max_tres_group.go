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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup struct {
	// GrpTRESMins
	Minutes []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"minutes,omitempty"`
	// GrpTRESRunMins
	Active []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"active,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroupWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroupWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup{}
	return &this
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) GetMinutes() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Minutes) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) GetMinutesOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Minutes field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) SetMinutes(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Minutes = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) GetActive() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Active) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) GetActiveOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Active field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) SetActive(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Active = v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


