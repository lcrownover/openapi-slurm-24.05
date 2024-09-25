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
	"bytes"
	"fmt"
)

// checks if the SlurmdbV0041GetDiag200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041GetDiag200Response{}

// SlurmdbV0041GetDiag200Response struct for SlurmdbV0041GetDiag200Response
type SlurmdbV0041GetDiag200Response struct {
	Statistics SlurmdbV0041GetDiag200ResponseStatistics `json:"statistics"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _SlurmdbV0041GetDiag200Response SlurmdbV0041GetDiag200Response

// NewSlurmdbV0041GetDiag200Response instantiates a new SlurmdbV0041GetDiag200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041GetDiag200Response(statistics SlurmdbV0041GetDiag200ResponseStatistics) *SlurmdbV0041GetDiag200Response {
	this := SlurmdbV0041GetDiag200Response{}
	this.Statistics = statistics
	return &this
}

// NewSlurmdbV0041GetDiag200ResponseWithDefaults instantiates a new SlurmdbV0041GetDiag200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041GetDiag200ResponseWithDefaults() *SlurmdbV0041GetDiag200Response {
	this := SlurmdbV0041GetDiag200Response{}
	return &this
}

// GetStatistics returns the Statistics field value
func (o *SlurmdbV0041GetDiag200Response) GetStatistics() SlurmdbV0041GetDiag200ResponseStatistics {
	if o == nil {
		var ret SlurmdbV0041GetDiag200ResponseStatistics
		return ret
	}

	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200Response) GetStatisticsOk() (*SlurmdbV0041GetDiag200ResponseStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// SetStatistics sets field value
func (o *SlurmdbV0041GetDiag200Response) SetStatistics(v SlurmdbV0041GetDiag200ResponseStatistics) {
	o.Statistics = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *SlurmdbV0041GetDiag200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200Response) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *SlurmdbV0041GetDiag200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200Response) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *SlurmdbV0041GetDiag200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o SlurmdbV0041GetDiag200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041GetDiag200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statistics"] = o.Statistics
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

func (o *SlurmdbV0041GetDiag200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statistics",
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

	varSlurmdbV0041GetDiag200Response := _SlurmdbV0041GetDiag200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlurmdbV0041GetDiag200Response)

	if err != nil {
		return err
	}

	*o = SlurmdbV0041GetDiag200Response(varSlurmdbV0041GetDiag200Response)

	return err
}

type NullableSlurmdbV0041GetDiag200Response struct {
	value *SlurmdbV0041GetDiag200Response
	isSet bool
}

func (v NullableSlurmdbV0041GetDiag200Response) Get() *SlurmdbV0041GetDiag200Response {
	return v.value
}

func (v *NullableSlurmdbV0041GetDiag200Response) Set(val *SlurmdbV0041GetDiag200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041GetDiag200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041GetDiag200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041GetDiag200Response(val *SlurmdbV0041GetDiag200Response) *NullableSlurmdbV0041GetDiag200Response {
	return &NullableSlurmdbV0041GetDiag200Response{value: val, isSet: true}
}

func (v NullableSlurmdbV0041GetDiag200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041GetDiag200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


