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

// checks if the V0041OpenapiPartitionRespPartitionsInnerDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiPartitionRespPartitionsInnerDefaults{}

// V0041OpenapiPartitionRespPartitionsInnerDefaults struct for V0041OpenapiPartitionRespPartitionsInnerDefaults
type V0041OpenapiPartitionRespPartitionsInnerDefaults struct {
	// DefMemPerCPU or DefMemPerNode
	MemoryPerCpu *int64 `json:"memory_per_cpu,omitempty"`
	PartitionMemoryPerCpu *V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu `json:"partition_memory_per_cpu,omitempty"`
	PartitionMemoryPerNode *V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode `json:"partition_memory_per_node,omitempty"`
	Time *V0041OpenapiPartitionRespPartitionsInnerDefaultsTime `json:"time,omitempty"`
	// JobDefaults
	Job *string `json:"job,omitempty"`
}

// NewV0041OpenapiPartitionRespPartitionsInnerDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiPartitionRespPartitionsInnerDefaults() *V0041OpenapiPartitionRespPartitionsInnerDefaults {
	this := V0041OpenapiPartitionRespPartitionsInnerDefaults{}
	return &this
}

// NewV0041OpenapiPartitionRespPartitionsInnerDefaultsWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiPartitionRespPartitionsInnerDefaultsWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerDefaults {
	this := V0041OpenapiPartitionRespPartitionsInnerDefaults{}
	return &this
}

// GetMemoryPerCpu returns the MemoryPerCpu field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetMemoryPerCpu() int64 {
	if o == nil || IsNil(o.MemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.MemoryPerCpu
}

// GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryPerCpu) {
		return nil, false
	}
	return o.MemoryPerCpu, true
}

// HasMemoryPerCpu returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MemoryPerCpu) {
		return true
	}

	return false
}

// SetMemoryPerCpu gets a reference to the given int64 and assigns it to the MemoryPerCpu field.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetMemoryPerCpu(v int64) {
	o.MemoryPerCpu = &v
}

// GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerCpu() V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		var ret V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu
		return ret
	}
	return *o.PartitionMemoryPerCpu
}

// GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerCpuOk() (*V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerCpu) {
		return nil, false
	}
	return o.PartitionMemoryPerCpu, true
}

// HasPartitionMemoryPerCpu returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasPartitionMemoryPerCpu() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerCpu) {
		return true
	}

	return false
}

// SetPartitionMemoryPerCpu gets a reference to the given V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu and assigns it to the PartitionMemoryPerCpu field.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetPartitionMemoryPerCpu(v V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerCpu) {
	o.PartitionMemoryPerCpu = &v
}

// GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerNode() V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		var ret V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode
		return ret
	}
	return *o.PartitionMemoryPerNode
}

// GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerNodeOk() (*V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode, bool) {
	if o == nil || IsNil(o.PartitionMemoryPerNode) {
		return nil, false
	}
	return o.PartitionMemoryPerNode, true
}

// HasPartitionMemoryPerNode returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasPartitionMemoryPerNode() bool {
	if o != nil && !IsNil(o.PartitionMemoryPerNode) {
		return true
	}

	return false
}

// SetPartitionMemoryPerNode gets a reference to the given V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode and assigns it to the PartitionMemoryPerNode field.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetPartitionMemoryPerNode(v V0041OpenapiPartitionRespPartitionsInnerDefaultsPartitionMemoryPerNode) {
	o.PartitionMemoryPerNode = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetTime() V0041OpenapiPartitionRespPartitionsInnerDefaultsTime {
	if o == nil || IsNil(o.Time) {
		var ret V0041OpenapiPartitionRespPartitionsInnerDefaultsTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetTimeOk() (*V0041OpenapiPartitionRespPartitionsInnerDefaultsTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0041OpenapiPartitionRespPartitionsInnerDefaultsTime and assigns it to the Time field.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetTime(v V0041OpenapiPartitionRespPartitionsInnerDefaultsTime) {
	o.Time = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetJob(v string) {
	o.Job = &v
}

func (o V0041OpenapiPartitionRespPartitionsInnerDefaults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiPartitionRespPartitionsInnerDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemoryPerCpu) {
		toSerialize["memory_per_cpu"] = o.MemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerCpu) {
		toSerialize["partition_memory_per_cpu"] = o.PartitionMemoryPerCpu
	}
	if !IsNil(o.PartitionMemoryPerNode) {
		toSerialize["partition_memory_per_node"] = o.PartitionMemoryPerNode
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0041OpenapiPartitionRespPartitionsInnerDefaults struct {
	value *V0041OpenapiPartitionRespPartitionsInnerDefaults
	isSet bool
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) Get() *V0041OpenapiPartitionRespPartitionsInnerDefaults {
	return v.value
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) Set(val *V0041OpenapiPartitionRespPartitionsInnerDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiPartitionRespPartitionsInnerDefaults(val *V0041OpenapiPartitionRespPartitionsInnerDefaults) *NullableV0041OpenapiPartitionRespPartitionsInnerDefaults {
	return &NullableV0041OpenapiPartitionRespPartitionsInnerDefaults{value: val, isSet: true}
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


