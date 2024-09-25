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
	"bytes"
	"fmt"
)

// checks if the V0041OpenapiTresResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiTresResp{}

// V0041OpenapiTresResp struct for V0041OpenapiTresResp
type V0041OpenapiTresResp struct {
	// TRES
	TRES []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"TRES"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiTresResp V0041OpenapiTresResp

// NewV0041OpenapiTresResp instantiates a new V0041OpenapiTresResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiTresResp(tRES []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) *V0041OpenapiTresResp {
	this := V0041OpenapiTresResp{}
	this.TRES = tRES
	return &this
}

// NewV0041OpenapiTresRespWithDefaults instantiates a new V0041OpenapiTresResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiTresRespWithDefaults() *V0041OpenapiTresResp {
	this := V0041OpenapiTresResp{}
	return &this
}

// GetTRES returns the TRES field value
func (o *V0041OpenapiTresResp) GetTRES() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}

	return o.TRES
}

// GetTRESOk returns a tuple with the TRES field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiTresResp) GetTRESOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TRES, true
}

// SetTRES sets field value
func (o *V0041OpenapiTresResp) SetTRES(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.TRES = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiTresResp) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiTresResp) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiTresResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *V0041OpenapiTresResp) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiTresResp) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiTresResp) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiTresResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiTresResp) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiTresResp) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiTresResp) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiTresResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiTresResp) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiTresResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiTresResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TRES"] = o.TRES
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0041OpenapiTresResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"TRES",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0041OpenapiTresResp := _V0041OpenapiTresResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiTresResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiTresResp(varV0041OpenapiTresResp)

	return err
}

type NullableV0041OpenapiTresResp struct {
	value *V0041OpenapiTresResp
	isSet bool
}

func (v NullableV0041OpenapiTresResp) Get() *V0041OpenapiTresResp {
	return v.value
}

func (v *NullableV0041OpenapiTresResp) Set(val *V0041OpenapiTresResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiTresResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiTresResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiTresResp(val *V0041OpenapiTresResp) *NullableV0041OpenapiTresResp {
	return &NullableV0041OpenapiTresResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiTresResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiTresResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


