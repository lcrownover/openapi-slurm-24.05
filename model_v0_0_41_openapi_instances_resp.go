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

// checks if the V0041OpenapiInstancesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiInstancesResp{}

// V0041OpenapiInstancesResp struct for V0041OpenapiInstancesResp
type V0041OpenapiInstancesResp struct {
	// instances
	Instances []V0041OpenapiSlurmdbdConfigRespInstancesInner `json:"instances"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiInstancesResp V0041OpenapiInstancesResp

// NewV0041OpenapiInstancesResp instantiates a new V0041OpenapiInstancesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiInstancesResp(instances []V0041OpenapiSlurmdbdConfigRespInstancesInner) *V0041OpenapiInstancesResp {
	this := V0041OpenapiInstancesResp{}
	this.Instances = instances
	return &this
}

// NewV0041OpenapiInstancesRespWithDefaults instantiates a new V0041OpenapiInstancesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiInstancesRespWithDefaults() *V0041OpenapiInstancesResp {
	this := V0041OpenapiInstancesResp{}
	return &this
}

// GetInstances returns the Instances field value
func (o *V0041OpenapiInstancesResp) GetInstances() []V0041OpenapiSlurmdbdConfigRespInstancesInner {
	if o == nil {
		var ret []V0041OpenapiSlurmdbdConfigRespInstancesInner
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiInstancesResp) GetInstancesOk() ([]V0041OpenapiSlurmdbdConfigRespInstancesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *V0041OpenapiInstancesResp) SetInstances(v []V0041OpenapiSlurmdbdConfigRespInstancesInner) {
	o.Instances = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiInstancesResp) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiInstancesResp) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiInstancesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *V0041OpenapiInstancesResp) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiInstancesResp) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiInstancesResp) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiInstancesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiInstancesResp) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiInstancesResp) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiInstancesResp) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiInstancesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiInstancesResp) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiInstancesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiInstancesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
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

func (o *V0041OpenapiInstancesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instances",
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

	varV0041OpenapiInstancesResp := _V0041OpenapiInstancesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiInstancesResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiInstancesResp(varV0041OpenapiInstancesResp)

	return err
}

type NullableV0041OpenapiInstancesResp struct {
	value *V0041OpenapiInstancesResp
	isSet bool
}

func (v NullableV0041OpenapiInstancesResp) Get() *V0041OpenapiInstancesResp {
	return v.value
}

func (v *NullableV0041OpenapiInstancesResp) Set(val *V0041OpenapiInstancesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiInstancesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiInstancesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiInstancesResp(val *V0041OpenapiInstancesResp) *NullableV0041OpenapiInstancesResp {
	return &NullableV0041OpenapiInstancesResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiInstancesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiInstancesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


