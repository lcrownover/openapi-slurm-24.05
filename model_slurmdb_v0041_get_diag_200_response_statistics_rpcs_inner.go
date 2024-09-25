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

// checks if the SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner{}

// SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner struct for SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner
type SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner struct {
	// RPC type
	Rpc *string `json:"rpc,omitempty"`
	// Number of RPCs processed
	Count *int32 `json:"count,omitempty"`
	Time *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime `json:"time,omitempty"`
}

// NewSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner() *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner {
	this := SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner{}
	return &this
}

// NewSlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerWithDefaults() *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner {
	this := SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner{}
	return &this
}

// GetRpc returns the Rpc field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetRpc() string {
	if o == nil || IsNil(o.Rpc) {
		var ret string
		return ret
	}
	return *o.Rpc
}

// GetRpcOk returns a tuple with the Rpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetRpcOk() (*string, bool) {
	if o == nil || IsNil(o.Rpc) {
		return nil, false
	}
	return o.Rpc, true
}

// HasRpc returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) HasRpc() bool {
	if o != nil && !IsNil(o.Rpc) {
		return true
	}

	return false
}

// SetRpc gets a reference to the given string and assigns it to the Rpc field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) SetRpc(v string) {
	o.Rpc = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetTime() SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime {
	if o == nil || IsNil(o.Time) {
		var ret SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) GetTimeOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime and assigns it to the Time field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) SetTime(v SlurmdbV0041GetDiag200ResponseStatisticsRPCsInnerTime) {
	o.Time = &v
}

func (o SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rpc) {
		toSerialize["rpc"] = o.Rpc
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner struct {
	value *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner
	isSet bool
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) Get() *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner {
	return v.value
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) Set(val *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner(val *SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) *NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner {
	return &NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner{value: val, isSet: true}
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


