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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerPreempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerPreempt{}

// V0041OpenapiSlurmdbdConfigRespQosInnerPreempt struct for V0041OpenapiSlurmdbdConfigRespQosInnerPreempt
type V0041OpenapiSlurmdbdConfigRespQosInnerPreempt struct {
	// Other QOS's this QOS can preempt
	List []string `json:"list,omitempty"`
	// PreemptMode
	Mode []string `json:"mode,omitempty"`
	ExemptTime *V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime `json:"exempt_time,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerPreempt instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerPreempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerPreempt() *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerPreempt{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerPreemptWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerPreempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerPreemptWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerPreempt{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) SetList(v []string) {
	o.List = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetMode() []string {
	if o == nil || IsNil(o.Mode) {
		var ret []string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetModeOk() ([]string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given []string and assigns it to the Mode field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) SetMode(v []string) {
	o.Mode = v
}

// GetExemptTime returns the ExemptTime field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetExemptTime() V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime {
	if o == nil || IsNil(o.ExemptTime) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime
		return ret
	}
	return *o.ExemptTime
}

// GetExemptTimeOk returns a tuple with the ExemptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) GetExemptTimeOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime, bool) {
	if o == nil || IsNil(o.ExemptTime) {
		return nil, false
	}
	return o.ExemptTime, true
}

// HasExemptTime returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) HasExemptTime() bool {
	if o != nil && !IsNil(o.ExemptTime) {
		return true
	}

	return false
}

// SetExemptTime gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime and assigns it to the ExemptTime field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) SetExemptTime(v V0041OpenapiSlurmdbdConfigRespQosInnerPreemptExemptTime) {
	o.ExemptTime = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ExemptTime) {
		toSerialize["exempt_time"] = o.ExemptTime
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt(val *V0041OpenapiSlurmdbdConfigRespQosInnerPreempt) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerPreempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


