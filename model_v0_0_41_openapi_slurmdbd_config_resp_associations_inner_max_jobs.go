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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs struct {
	Per *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer `json:"per,omitempty"`
	Active *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive `json:"active,omitempty"`
	Accruing *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing `json:"accruing,omitempty"`
	Total *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal `json:"total,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs{}
	return &this
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetPer() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetPerOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer and assigns it to the Per field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) SetPer(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsPer) {
	o.Per = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetActive() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive {
	if o == nil || IsNil(o.Active) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetActiveOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive and assigns it to the Active field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) SetActive(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsActive) {
	o.Active = &v
}

// GetAccruing returns the Accruing field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetAccruing() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing {
	if o == nil || IsNil(o.Accruing) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing
		return ret
	}
	return *o.Accruing
}

// GetAccruingOk returns a tuple with the Accruing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetAccruingOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing, bool) {
	if o == nil || IsNil(o.Accruing) {
		return nil, false
	}
	return o.Accruing, true
}

// HasAccruing returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) HasAccruing() bool {
	if o != nil && !IsNil(o.Accruing) {
		return true
	}

	return false
}

// SetAccruing gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing and assigns it to the Accruing field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) SetAccruing(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsAccruing) {
	o.Accruing = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetTotal() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal {
	if o == nil || IsNil(o.Total) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) GetTotalOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal and assigns it to the Total field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) SetTotal(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobsTotal) {
	o.Total = &v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Accruing) {
		toSerialize["accruing"] = o.Accruing
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


