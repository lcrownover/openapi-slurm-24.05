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

// checks if the V0040StatsMsgRpcsByTypeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StatsMsgRpcsByTypeInner{}

// V0040StatsMsgRpcsByTypeInner RPC
type V0040StatsMsgRpcsByTypeInner struct {
	// Message type as string
	MessageType *string `json:"message_type,omitempty"`
	// Message type as integer
	TypeId *int32 `json:"type_id,omitempty"`
	// Number of RPCs received
	Count *int64 `json:"count,omitempty"`
	// Average time spent processing RPC in seconds
	AverageTime *int64 `json:"average_time,omitempty"`
	// Total time spent processing RPC in seconds
	TotalTime *int64 `json:"total_time,omitempty"`
}

// NewV0040StatsMsgRpcsByTypeInner instantiates a new V0040StatsMsgRpcsByTypeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StatsMsgRpcsByTypeInner() *V0040StatsMsgRpcsByTypeInner {
	this := V0040StatsMsgRpcsByTypeInner{}
	return &this
}

// NewV0040StatsMsgRpcsByTypeInnerWithDefaults instantiates a new V0040StatsMsgRpcsByTypeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StatsMsgRpcsByTypeInnerWithDefaults() *V0040StatsMsgRpcsByTypeInner {
	this := V0040StatsMsgRpcsByTypeInner{}
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *V0040StatsMsgRpcsByTypeInner) GetMessageType() string {
	if o == nil || IsNil(o.MessageType) {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsMsgRpcsByTypeInner) GetMessageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *V0040StatsMsgRpcsByTypeInner) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *V0040StatsMsgRpcsByTypeInner) SetMessageType(v string) {
	o.MessageType = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *V0040StatsMsgRpcsByTypeInner) GetTypeId() int32 {
	if o == nil || IsNil(o.TypeId) {
		var ret int32
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsMsgRpcsByTypeInner) GetTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *V0040StatsMsgRpcsByTypeInner) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.
func (o *V0040StatsMsgRpcsByTypeInner) SetTypeId(v int32) {
	o.TypeId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0040StatsMsgRpcsByTypeInner) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsMsgRpcsByTypeInner) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0040StatsMsgRpcsByTypeInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *V0040StatsMsgRpcsByTypeInner) SetCount(v int64) {
	o.Count = &v
}

// GetAverageTime returns the AverageTime field value if set, zero value otherwise.
func (o *V0040StatsMsgRpcsByTypeInner) GetAverageTime() int64 {
	if o == nil || IsNil(o.AverageTime) {
		var ret int64
		return ret
	}
	return *o.AverageTime
}

// GetAverageTimeOk returns a tuple with the AverageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsMsgRpcsByTypeInner) GetAverageTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AverageTime) {
		return nil, false
	}
	return o.AverageTime, true
}

// HasAverageTime returns a boolean if a field has been set.
func (o *V0040StatsMsgRpcsByTypeInner) HasAverageTime() bool {
	if o != nil && !IsNil(o.AverageTime) {
		return true
	}

	return false
}

// SetAverageTime gets a reference to the given int64 and assigns it to the AverageTime field.
func (o *V0040StatsMsgRpcsByTypeInner) SetAverageTime(v int64) {
	o.AverageTime = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *V0040StatsMsgRpcsByTypeInner) GetTotalTime() int64 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int64
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StatsMsgRpcsByTypeInner) GetTotalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *V0040StatsMsgRpcsByTypeInner) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int64 and assigns it to the TotalTime field.
func (o *V0040StatsMsgRpcsByTypeInner) SetTotalTime(v int64) {
	o.TotalTime = &v
}

func (o V0040StatsMsgRpcsByTypeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StatsMsgRpcsByTypeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageType) {
		toSerialize["message_type"] = o.MessageType
	}
	if !IsNil(o.TypeId) {
		toSerialize["type_id"] = o.TypeId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageTime) {
		toSerialize["average_time"] = o.AverageTime
	}
	if !IsNil(o.TotalTime) {
		toSerialize["total_time"] = o.TotalTime
	}
	return toSerialize, nil
}

type NullableV0040StatsMsgRpcsByTypeInner struct {
	value *V0040StatsMsgRpcsByTypeInner
	isSet bool
}

func (v NullableV0040StatsMsgRpcsByTypeInner) Get() *V0040StatsMsgRpcsByTypeInner {
	return v.value
}

func (v *NullableV0040StatsMsgRpcsByTypeInner) Set(val *V0040StatsMsgRpcsByTypeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StatsMsgRpcsByTypeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StatsMsgRpcsByTypeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StatsMsgRpcsByTypeInner(val *V0040StatsMsgRpcsByTypeInner) *NullableV0040StatsMsgRpcsByTypeInner {
	return &NullableV0040StatsMsgRpcsByTypeInner{value: val, isSet: true}
}

func (v NullableV0040StatsMsgRpcsByTypeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StatsMsgRpcsByTypeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


