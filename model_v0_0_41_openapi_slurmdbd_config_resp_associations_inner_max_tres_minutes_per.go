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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer struct {
	// MaxTRESMinsPerJob
	Job []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"job,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) GetJob() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Job) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) GetJobOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Job field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) SetJob(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Job = v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxTresMinutesPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


