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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerReservation{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerReservation struct for V0041OpenapiSlurmdbdJobsRespJobsInnerReservation
type V0041OpenapiSlurmdbdJobsRespJobsInnerReservation struct {
	// Unique identifier of requested reservation
	Id *int32 `json:"id,omitempty"`
	// Name of reservation to use
	Name *string `json:"name,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerReservation instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerReservation() *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerReservation{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerReservationWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerReservationWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerReservation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) SetName(v string) {
	o.Name = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation(val *V0041OpenapiSlurmdbdJobsRespJobsInnerReservation) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


