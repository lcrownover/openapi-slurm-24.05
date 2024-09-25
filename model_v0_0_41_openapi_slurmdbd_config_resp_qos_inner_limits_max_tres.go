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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres struct {
	// GrpTRES
	Total []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"total,omitempty"`
	Minutes *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes `json:"minutes,omitempty"`
	Per *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer `json:"per,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetTotal() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Total) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetTotalOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Total field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) SetTotal(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Total = v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetMinutes() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes {
	if o == nil || IsNil(o.Minutes) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetMinutesOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes and assigns it to the Minutes field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) SetMinutes(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutes) {
	o.Minutes = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetPer() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer {
	if o == nil || IsNil(o.Per) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) GetPerOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer and assigns it to the Per field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) SetPer(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) {
	o.Per = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


