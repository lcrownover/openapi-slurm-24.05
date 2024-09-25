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

// checks if the V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer{}

// V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer struct for V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer
type V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer struct {
	Account *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount `json:"account,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer {
	this := V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) GetAccount() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount {
	if o == nil || IsNil(o.Account) {
		var ret V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) GetAccountOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount and assigns it to the Account field.
func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) SetAccount(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPerAccount) {
	o.Account = &v
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer struct {
	value *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) Get() *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) Set(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer(val *V0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespAssociationsInnerMaxPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


